package cmd

import (
	"fmt"
	opbtypes "github.com/bianjieai/irita/modules/opb/types"
	identitytypes "github.com/bianjieai/iritamod/modules/identity/types"
	nodetypes "github.com/bianjieai/iritamod/modules/node/types"
	permtypes "github.com/bianjieai/iritamod/modules/perm/types"
	upgradetypes "github.com/bianjieai/iritamod/modules/upgrade/types"
	tibcmttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/mt_transfer/types"
	tibcnfttypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/iavl"
	mttypes "github.com/irisnet/irismod/modules/mt/types"
	nfttypes "github.com/irisnet/irismod/modules/nft/types"
	oracletypes "github.com/irisnet/irismod/modules/oracle/types"
	randomtypes "github.com/irisnet/irismod/modules/random/types"
	recordtypes "github.com/irisnet/irismod/modules/record/types"
	servicetypes "github.com/irisnet/irismod/modules/service/types"
	tokentypes "github.com/irisnet/irismod/modules/token/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"
	feemarkettypes "github.com/tharsis/ethermint/x/feemarket/types"
	"io"
	"os"
	"path/filepath"
	"strings"

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
)

const (
	flagTmpDir               = "tmp-dir"
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
	valSetCheckpointInterval = 100000
	DefaultCacheSize         = 10000
	moduleKeyFmt             = "s/k:%s/"
)

var privValidatorState = `{
  "height": "%d",
  "round": 0,
  "step": 0
}`

var storeKeys = []string{
	authtypes.StoreKey,
	banktypes.StoreKey,
	slashingtypes.StoreKey,
	paramstypes.StoreKey,
	upgradetypes.StoreKey,
	feegrant.StoreKey,
	evidencetypes.StoreKey,
	recordtypes.StoreKey,
	tokentypes.StoreKey,
	nfttypes.StoreKey,
	mttypes.StoreKey,
	servicetypes.StoreKey,
	oracletypes.StoreKey,
	randomtypes.StoreKey,
	permtypes.StoreKey,
	identitytypes.StoreKey,
	nodetypes.StoreKey,
	opbtypes.StoreKey,
	tibchost.StoreKey,
	tibcnfttypes.StoreKey,
	//wasm.StoreKey,
	tibcmttypes.StoreKey,

	// evm
	evmtypes.StoreKey, feemarkettypes.StoreKey,
}

// SnapshotCmd delete historical block data and index data
func SnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "snapshot the latest information and drop the others",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()

			home := viper.GetString(tmcli.HomeFlag)
			targetDir := viper.GetString(flagTmpDir)

			// if targetDir exist then tips
			if b, _ := pathExists(targetDir); b {
				fmt.Printf("target  dir: (%s) existed! \n", targetDir)
				return nil
			}

			if len(targetDir) == 0 {
				targetDir = filepath.Join(home, defaultTmpDir)
			}
			dataDir := filepath.Join(home, dataDir)

			if err := Snapshot(dataDir, targetDir); err != nil {
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

func loadDb(name, path string) *dbm.GoLevelDB {
	db, err := dbm.NewGoLevelDB(name, path)
	if err != nil {
		panic(err)
	}
	return db
}

func Snapshot(dataDir, targetDir string) error {
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
	if height, err = Recover(*blockStore, ss, *appDB); err != nil {
		return err
	}

	//save local current block and flush disk
	snapshotBlock(blockStore, targetDir, height)
	//save local current block height state
	snapshotState(stateDB, targetDir, height)
	//save local current block height consensus data
	snapshotCsWAL(dataDir, targetDir, height)

	privState := filepath.Join(targetDir, privValidatorStateFile)
	createPrivValidatorState(privState, height)

	//copy application
	appDir := filepath.Join(dataDir, applicationDb)
	appTargetDir := filepath.Join(targetDir, applicationDb)
	if err := copyDir(appDir, appTargetDir); err != nil {
		return err
	}

	//copy evidence.db
	evidenceDir := filepath.Join(dataDir, evidenceDb)
	evidenceTargetDir := filepath.Join(targetDir, evidenceDb)
	if err := copyDir(evidenceDir, evidenceTargetDir); err != nil {
		return err
	}

	// pruning application versions
	targetAppDB := loadDb(applicationDBDir, targetDir)
	defer targetAppDB.Close()
	return pruningVersions(targetAppDB, height)
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
	if err := newStore.Save(state); err != nil {
		panic(err)
	}
	if err := newStore.SaveABCIResponses(height, abciResponse); err != nil {
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

func createPrivValidatorState(path string, height int64) {
	valState := fmt.Sprintf(privValidatorState, height)
	tmos.MustWriteFile(path, []byte(valState), 0644)
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

func pruningVersions(db dbm.DB, height int64) error {
	if height <= 0 {
		return nil
	}
	for _, store := range storeKeys {
		fmt.Printf("pruning store: %s start \n", store)
		// read tree
		tree, err := readTree(db, store)
		if err != nil {
			fmt.Printf("read Tree %s err: %s \n", store, err.Error())
			return err
		}

		start := int64(0)
		length := int64(10000)
		for {
			if start >= height {
				break
			}
			end := start + length
			if end > height {
				end = height
			}
			if err := tree.DeleteVersionsRange(start, end); err != nil {
				fmt.Printf("delete version from 0 to %d err: %s \n", height, err.Error())
				return err
			}
			start = end
		}

	}
	return nil
}

func readTree(db dbm.DB, store string) (*iavl.MutableTree, error) {
	prefix := fmt.Sprintf(moduleKeyFmt, store)
	prefixDB := dbm.NewPrefixDB(db, []byte(prefix))

	tree, err := iavl.NewMutableTree(prefixDB, DefaultCacheSize)
	if err != nil {
		return nil, err
	}
	return tree, err
}
