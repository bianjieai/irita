package cmd

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/iavl"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/consensus"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	tmmath "github.com/tendermint/tendermint/libs/math"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"
	"github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/store"
	"github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/bianjieai/irita/app"
)

const (
	flagTmpDir               = "tmp-dir"
	flagMaxPruningVersions   = "max-pruning-versions"
	pathSeparator            = string(os.PathSeparator)
	defaultTmpDir            = "data.bak"
	dataDir                  = "data"
	blockStoreDir            = "blockstore"
	stateStoreDir            = "state"
	applicationDBDir         = "application"
	applicationDb            = "application.db"
	evidenceDb               = "evidence.db"
	csWalFile                = "cs.wal"
	privValidatorStateFile   = "priv_validator_state.json"
	upgradeInfoFile          = "upgrade-info.json"
	valSetCheckpointInterval = 100000
	DefaultCacheSize         = 10000
	moduleKeyFmt             = "s/k:%s/"
)

var privValidatorState = `{
  "height": "%d",
  "round": 0,
  "step": 0
}`

var storeKeys = app.GetStoreKeys()

func NewSnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "snapshot the latest information and drop the others, prune version",
	}
	cmd.AddCommand(
		SnapshotCmd(),
		PruneCmd(),
	)
	return cmd
}

// SnapshotCmd delete historical block data and index data
func SnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "snapshot the latest information and drop the others",
		Example: fmt.Sprintf(
			"$ %s snapshot create --home=/root/.%s --tmp-dir=/home/data.bak",
			version.AppName, version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()

			home := viper.GetString(tmcli.HomeFlag)
			targetDir := viper.GetString(flagTmpDir)

			if len(targetDir) == 0 {
				targetDir = filepath.Join(home, defaultTmpDir)
			}

			// if targetDir exist then tips
			if b, _ := pathExists(filepath.Join(targetDir, applicationDb)); b {
				fmt.Printf("target  dir: (%s) existed! \n", targetDir)
				return nil
			}
			dataDir := filepath.Join(home, dataDir)

			if err := snapshot(dataDir, targetDir); err != nil {
				fmt.Errorf("snapshot err: %s", err.Error())
				_ = os.RemoveAll(targetDir)
				return err
			}
			fmt.Println("snapshot completed!")
			return nil
		},
	}
	cmd.Flags().String(flagTmpDir, "", "Snapshot file storage directory")
	return cmd
}

// PruneCmd delete historical version data of application
func PruneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prune",
		Short: "prune snapshot's historical versions of application",
		Example: fmt.Sprintf(
			"$ %s snapshot prune --tmp-dir=/home/data.bak --max-pruning-versions=2000",
			version.AppName,
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()

			home := viper.GetString(tmcli.HomeFlag)
			targetDir := viper.GetString(flagTmpDir)
			batchNum := viper.GetInt64(flagMaxPruningVersions)

			if len(targetDir) == 0 {
				targetDir = filepath.Join(home, defaultTmpDir)
			}

			// if targetDir not exist then tips
			if b, _ := pathExists(targetDir); !b {
				fmt.Printf("target dir: (%s) not existed! \n", targetDir)
				return nil
			}

			if err := pruningVersions(targetDir, batchNum); err != nil {
				fmt.Errorf("prune err: %s", err.Error())
				return err
			}
			fmt.Println("prune completed!")
			return nil
		},
	}
	cmd.Flags().String(flagTmpDir, "", "Snapshot file storage directory")
	cmd.Flags().Int64(flagMaxPruningVersions, 2000, "the number that delete pruning versions in one time, affect memory usage")
	return cmd
}

func loadDb(name, path string) *dbm.GoLevelDB {
	db, err := dbm.NewGoLevelDB(name, path)
	if err != nil {
		panic(err)
	}
	return db
}

func snapshot(dataDir, targetDir string) error {
	blockDB := loadDb(blockStoreDir, dataDir)
	blockStore := store.NewBlockStore(blockDB)

	stateDB := loadDb(stateStoreDir, dataDir)
	ss := state.NewStore(stateDB)
	appDB := loadDb(applicationDBDir, dataDir)

	defer func() {
		blockDB.Close()
		stateDB.Close()
		appDB.Close()
	}()

	var (
		height int64
		err    error
	)
	if height, err = Recover(blockStore, ss, *appDB); err != nil {
		return err
	}

	//save local current block and flush disk
	snapshotBlock(blockStore, targetDir, height)
	//save local current block height state
	snapshotState(stateDB, targetDir, height)
	//save local current block height consensus data
	snapshotCsWAL(dataDir, targetDir, height)
	// create private validator state
	createPrivValidatorState(targetDir, height)
	// copy application
	copyAppDB(dataDir, targetDir)
	// copy evidence.db
	copyEvidenceDB(dataDir, targetDir)
	// copy upgrade info
	copyUpgradeInfo(dataDir, targetDir)
	return nil
}

func copyUpgradeInfo(dataDir string, targetDir string) {
	upgradeInfoFrom := filepath.Join(dataDir, upgradeInfoFile)
	b, err := pathExists(filepath.Join(targetDir, applicationDb))
	if err != nil {
		panic(fmt.Sprintf("read file %s err: %s", upgradeInfoFrom, err))
	}
	if b {
		fmt.Printf("target  dir: (%s) not found! skip\n", targetDir)
		return
	}

	upgradeInfoTo := filepath.Join(targetDir, upgradeInfoFile)

	_, err = copyFile(upgradeInfoFrom, upgradeInfoTo)
	if err != nil {
		panic(fmt.Sprintf("snapshot %s err: %s", upgradeInfoFrom, err))
	}
}

func copyEvidenceDB(dataDir string, targetDir string) {
	evidenceDir := filepath.Join(dataDir, evidenceDb)
	evidenceTargetDir := filepath.Join(targetDir, evidenceDb)
	err := copyDir(evidenceDir, evidenceTargetDir)
	if err != nil {
		panic(fmt.Sprintf("snapshot %s err: %s", evidenceDir, err))
	}
}

func copyAppDB(dataDir string, targetDir string) {
	appDir := filepath.Join(dataDir, applicationDb)
	appTargetDir := filepath.Join(targetDir, applicationDb)
	if err := copyDir(appDir, appTargetDir); err != nil {
		panic(fmt.Sprintf("snapshot %s err: %s", appDir, err))
	}
}

func snapshotState(tmDB *dbm.GoLevelDB, targetDir string, height int64) {
	targetDb := loadDb(stateStoreDir, targetDir)
	defer targetDb.Close()

	newStore := state.NewStore(targetDb)
	oldStore := state.NewStore(tmDB)

	state, err := oldStore.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("snapshotState",
		"lastBlockHeight=", state.LastBlockHeight,
		"lastHeightValidatorsChanged=", state.LastHeightValidatorsChanged,
		"nextValidators=", state.NextValidators,
	)

	abciResponse, err := oldStore.LoadABCIResponses(height)
	if err != nil {
		panic(err)
	}
	if err = newStore.Save(state); err != nil {
		panic(err)
	}
	if err = newStore.SaveABCIResponses(height, abciResponse); err != nil {
		panic(err)
	}

	saveValidatorsInfo(tmDB, targetDb, height, state.LastHeightValidatorsChanged)
	saveConsensusParamsInfo(tmDB, targetDb, height)
}

func snapshotBlock(originStore *store.BlockStore, targetDir string, height int64) {
	targetDb := loadDb(blockStoreDir, targetDir)
	defer targetDb.Close()

	fmt.Println("snapshotBlock", "  storeHeight=", height)

	targetStore := store.NewBlockStore(targetDb)

	block := originStore.LoadBlock(height)
	seenCommit := originStore.LoadSeenCommit(height)
	partSet := block.MakePartSet(types.BlockPartSizeBytes)
	targetStore.SaveBlock(block, partSet, seenCommit)
}

func snapshotCsWAL(home, targetDir string, height int64) {
	walTargetDir := filepath.Join(targetDir, csWalFile, "wal")
	targetWAL, err := consensus.NewWAL(walTargetDir)

	walSourceDir := filepath.Join(home, csWalFile, "wal")
	sourceWAL, err := consensus.NewWAL(walSourceDir)
	if err != nil {
		return
	}

	gr, found, err := sourceWAL.SearchForEndHeight(height, &consensus.WALSearchOptions{IgnoreDataCorruptionErrors: true})

	if err != nil || !found {

		return
	}

	defer func() {
		if err = gr.Close(); err != nil {
			return
		}
	}()

	var msg *consensus.TimedWALMessage
	dec := consensus.NewWALDecoder(gr)
	for {
		msg, err = dec.Decode()
		if err == io.EOF {
			break
		} else if consensus.IsDataCorruptionError(err) {
			return
		} else if err != nil {
			return
		}
		if err := targetWAL.Write(msg.Msg); err != nil {
			return
		}
	}
	err = targetWAL.WriteSync(consensus.EndHeightMessage{Height: height})
	if err != nil {
		return
	}
}

func copyDir(srcPath string, destPath string) error {
	if _, err := os.Stat(srcPath); err != nil {
		return err
	}

	return filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}
		path = strings.Replace(path, fmt.Sprintf("\\%s", pathSeparator), pathSeparator, -1)
		destNewPath := strings.Replace(path, srcPath, destPath, -1)
		_, err = copyFile(path, destNewPath)
		return err
	})
}

func copyFile(src, dest string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}

	defer srcFile.Close()

	destSplitPathDirs := strings.Split(dest, pathSeparator)

	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + pathSeparator
			if b, _ := pathExists(destSplitPath); b == false {
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					return 0, err
				}
			}
		}
	}

	dstFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer dstFile.Close()

	buf := make([]byte, 1024*1024)
	for {
		n, err := srcFile.Read(buf)
		if err != nil && err != io.EOF {
			return int64(n), err
		}
		if n == 0 {
			break
		}

		tmp := buf[:n]
		dstFile.Write(tmp)
	}
	return 0, nil
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func loadValidatorsInfo(db dbm.DB, height int64) *tmstate.ValidatorsInfo {
	buf, err := db.Get(calcValidatorsKey(height))
	if err != nil {
		panic(err)
	}

	if len(buf) == 0 {
		panic("validatorsInfo is empty")
	}

	v := new(tmstate.ValidatorsInfo)
	if err := v.Unmarshal(buf); err != nil {
		panic(err)
	}
	return v
}

func saveValidatorsInfo(originDb, targetDb dbm.DB, height, lastHeightChanged int64) {
	if lastHeightChanged > height {
		panic("lastHeightChanged cannot be greater than ValidatorsInfo height")
	}

	saveValidators := func(height int64, valInfo *tmstate.ValidatorsInfo) {
		bz, err := valInfo.Marshal()
		if err != nil {
			panic(err)
		}
		targetDb.Set(calcValidatorsKey(height), bz)
	}

	saveCurrentValidator := func() {
		valInfo := loadValidatorsInfo(originDb, height)
		saveValidators(height, valInfo)
	}

	saveNextValidator := func() {
		nextHeight := height + 1
		valInfo := loadValidatorsInfo(originDb, nextHeight)
		saveValidators(nextHeight, valInfo)
	}

	saveLastChangedValidators := func() {
		lastStoredHeight := lastStoredHeightFor(height, lastHeightChanged)
		valInfo := loadValidatorsInfo(originDb, lastStoredHeight)
		saveValidators(lastStoredHeight, valInfo)
	}

	saveCurrentValidator()
	saveNextValidator()
	saveLastChangedValidators()
}

func loadConsensusParamsInfo(db dbm.DB, height int64) *tmstate.ConsensusParamsInfo {
	buf, err := db.Get(calcConsensusParamsKey(height))
	if err != nil {
		panic(err)
	}

	if len(buf) == 0 {
		panic("validatorsInfo is empty")
	}

	paramsInfo := new(tmstate.ConsensusParamsInfo)
	if err := paramsInfo.Unmarshal(buf); err != nil {
		panic(err)
	}
	return paramsInfo
}

func saveConsensusParamsInfo(originDb, targetDb dbm.DB, height int64) {
	consensusParamsInfo := loadConsensusParamsInfo(originDb, height)
	paramsInfo := &tmstate.ConsensusParamsInfo{
		LastHeightChanged: consensusParamsInfo.LastHeightChanged,
	}
	bz, err := paramsInfo.Marshal()
	if err != nil {
		panic(err)
	}
	targetDb.Set(calcConsensusParamsKey(consensusParamsInfo.LastHeightChanged), bz)
}

func createPrivValidatorState(targetDir string, height int64) {
	privState := filepath.Join(targetDir, privValidatorStateFile)
	valState := fmt.Sprintf(privValidatorState, height)
	tmos.MustWriteFile(privState, []byte(valState), 0644)
}

func calcValidatorsKey(height int64) []byte {
	return []byte(fmt.Sprintf("validatorsKey:%v", height))
}

func calcConsensusParamsKey(height int64) []byte {
	return []byte(fmt.Sprintf("consensusParamsKey:%v", height))
}

func lastStoredHeightFor(height, lastHeightChanged int64) int64 {
	checkpointHeight := height - height%valSetCheckpointInterval
	return tmmath.MaxInt64(checkpointHeight, lastHeightChanged)
}

func pruningVersions(targetDir string, batchNum int64) error {
	// pruning application versions
	targetAppDB := loadDb(applicationDBDir, targetDir)
	defer targetAppDB.Close()

	height := getLatestVersion(targetAppDB)
	fmt.Printf("application latest version: %d \n", height)
	if height <= 0 {
		return nil
	}
	w := &sync.WaitGroup{}
	for _, store := range storeKeys {
		w.Add(1)
		go pruneStore(targetAppDB, store, height, batchNum, w)
	}
	w.Wait()
	return nil
}

func pruneStore(targetAppDB *dbm.GoLevelDB, store string, height int64, batchNum int64, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Printf("pruning store: %s start  - %s \n", store, time.Now().Format("2006-01-02 15:04:05"))
	// read tree
	tree, err := readTree(targetAppDB, height, store)
	if err != nil {
		fmt.Printf("read Tree %s err: %s \n", store, err.Error())
		return
	}

	// if module not exist then return
	if !tree.VersionExists(height) {
		fmt.Printf("store: %s not exists \n", store)
		return
	}

	list := tree.AvailableVersions()
	start := int64(list[0])
	fmt.Printf("pruning store %s from %d to %d \n", store, start, height)
	for {
		if start >= height {
			break
		}
		end := start + batchNum
		if end > height {
			end = height
		}
		if err := tree.DeleteVersionsRange(start, end); err != nil {
			fmt.Printf("delete version from %d to %d err: %s \n", start, end, err.Error())
			return
		}
		start = end
	}
	fmt.Printf("pruning store: %s end  - %s \n", store, time.Now().Format("2006-01-02 15:04:05"))
}

func readTree(db dbm.DB, latestVersion int64, store string) (*iavl.MutableTree, error) {
	prefix := fmt.Sprintf(moduleKeyFmt, store)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	tree, err := iavl.NewMutableTree(prefixDB, DefaultCacheSize)
	if err != nil {
		return nil, err
	}
	tree.LoadVersion(latestVersion)
	return tree, err
}
