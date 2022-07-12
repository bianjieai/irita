package cmd

import (
	"fmt"
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
	applicationDb            = "application.db"
	evidenceDb               = "evidence.db"
	csWalFile                = "cs.wal"
	privValidatorStateFile   = "priv_validator_state.json"
	valSetCheckpointInterval = 100000
)

var privValidatorState = `{
  "height": "%d",
  "round": 0,
  "step": 0
}`

// SnapshotCmd delete historical block data and index data
func SnapshotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "snapshot",
		Short: "snapshot the latest information and drop the others",
		RunE: func(cmd *cobra.Command, args []string) error {
			defer func() {
				if r := recover(); r != nil {
				}
			}()

			home := viper.GetString(tmcli.HomeFlag)
			targetDir := viper.GetString(flagTmpDir)
			if len(targetDir) == 0 {
				targetDir = filepath.Join(home, defaultTmpDir)
			}
			dataDir := filepath.Join(home, dataDir)

			if err := Snapshot(dataDir, targetDir); err != nil {
				_ = os.RemoveAll(targetDir)
				return err
			}
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

	defer func() {
		blockDB.Close()
		stateDB.Close()
	}()

	var (
		height int64
		err    error
	)
	if height, err = Recover(*blockStore, ss); err != nil {
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
	// TODO
	// appDir := filepath.Join(dataDir, applicationDb)
	// appTargetDir := filepath.Join(targetDir, applicationDb)
	// if err := copyDir(appDir, appTargetDir); err != nil {
	// 	return err
	// }

	//copy evidence.db
	evidenceDir := filepath.Join(dataDir, evidenceDb)
	evidenceTargetDir := filepath.Join(targetDir, evidenceDb)
	return copyDir(evidenceDir, evidenceTargetDir)
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
	defer srcFile.Close()
	if err != nil {
		return
	}

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
	return io.Copy(dstFile, srcFile)
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
