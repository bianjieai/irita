package cmd

import (
	"errors"
	"fmt"
	gogotypes "github.com/gogo/protobuf/types"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"
	"github.com/tendermint/tendermint/proto/tendermint/version"
	"github.com/tendermint/tendermint/state"
	"github.com/tendermint/tendermint/store"
	tmversion "github.com/tendermint/tendermint/version"
	dbm "github.com/tendermint/tm-db"
)

const (
	latestVersionKey = "s/latest"
	commitInfoKeyFmt = "s/%d"
)

// Recover overwrites the current Tendermint state (height n) with the most
// recent previous state (height n - 1).
// Note that this function does not affect application state.
func Recover(bs store.BlockStore, ss state.Store, as dbm.GoLevelDB) (int64, error) {
	defer as.Close()
	invalidState, err := ss.Load()
	if err != nil {
		return -1, err
	}
	if invalidState.IsEmpty() {
		return -1, errors.New("no state found")
	}

	height := bs.Height()
	fmt.Println("Recover", "  storeHeight=", height, "  stateHeight=", invalidState.LastBlockHeight)

	// NOTE: persistence of state and blocks don't happen atomically. Therefore it is possible that
	// when the user stopped the node the state wasn't updated but the blockstore was. In this situation
	// we don't need to rollback any state and can just return early
	if height == invalidState.LastBlockHeight+1 {
		rollBackApplication(as, invalidState.LastBlockHeight)
		return invalidState.LastBlockHeight, nil
	}

	// If the state store isn't one below nor equal to the blockstore height than this violates the
	// invariant
	if height != invalidState.LastBlockHeight {
		return -1, fmt.Errorf("statestore height (%d) is not one below or equal to blockstore height (%d)",
			invalidState.LastBlockHeight, height)
	}

	// state store height is equal to blockstore height. We're good to proceed with rolling back state
	rollbackHeight := invalidState.LastBlockHeight - 1
	rollbackBlock := bs.LoadBlockMeta(rollbackHeight)
	if rollbackBlock == nil {
		return -1, fmt.Errorf("block at height %d not found", rollbackHeight)
	}
	// we also need to retrieve the latest block because the app hash and last results hash is only agreed upon in the following block
	latestBlock := bs.LoadBlockMeta(invalidState.LastBlockHeight)
	if latestBlock == nil {
		return -1, fmt.Errorf("block at height %d not found", invalidState.LastBlockHeight)
	}

	previousLastValidatorSet, err := ss.LoadValidators(rollbackHeight)
	if err != nil {
		return -1, err
	}

	previousParams, err := ss.LoadConsensusParams(rollbackHeight + 1)
	if err != nil {
		return -1, err
	}

	valChangeHeight := invalidState.LastHeightValidatorsChanged
	// this can only happen if the validator set changed since the last block
	if valChangeHeight > rollbackHeight {
		valChangeHeight = rollbackHeight + 1
	}

	paramsChangeHeight := invalidState.LastHeightConsensusParamsChanged
	// this can only happen if params changed from the last block
	if paramsChangeHeight > rollbackHeight {
		paramsChangeHeight = rollbackHeight + 1
	}

	// build the new state from the old state and the prior block
	rolledBackState := state.State{
		Version: tmstate.Version{
			Consensus: version.Consensus{
				Block: tmversion.BlockProtocol,
				App:   previousParams.Version.AppVersion,
			},
			Software: tmversion.TMVersionDefault,
		},
		// immutable fields
		ChainID:       invalidState.ChainID,
		InitialHeight: invalidState.InitialHeight,

		LastBlockHeight: rollbackBlock.Header.Height,
		LastBlockID:     rollbackBlock.BlockID,
		LastBlockTime:   rollbackBlock.Header.Time,

		NextValidators:              invalidState.Validators,
		Validators:                  invalidState.LastValidators,
		LastValidators:              previousLastValidatorSet,
		LastHeightValidatorsChanged: valChangeHeight,

		ConsensusParams:                  previousParams,
		LastHeightConsensusParamsChanged: paramsChangeHeight,

		LastResultsHash: latestBlock.Header.LastResultsHash,
		AppHash:         latestBlock.Header.AppHash,
	}

	// persist the new state. This overrides the invalid one. NOTE: this will also
	// persist the validator set and consensus params over the existing structures,
	// but both should be the same
	if err := ss.Save(rolledBackState); err != nil {
		return -1, fmt.Errorf("failed to save rolled back state: %w", err)
	}

	// rollback application the latest version
	rollBackApplication(as, rolledBackState.LastBlockHeight)
	return rolledBackState.LastBlockHeight, nil
}

func rollBackApplication(as dbm.GoLevelDB, height int64) {
	batch := as.NewBatch()
	defer batch.Close()
	setLatestVersion(batch, height)
	deleteCommitInfo(batch, height+1)
	batch.WriteSync()
}

func setLatestVersion(batch dbm.Batch, version int64) {
	bz, err := gogotypes.StdInt64Marshal(version)
	if err != nil {
		panic(err)
	}

	if err := batch.Set([]byte(latestVersionKey), bz); err != nil {
		panic(err)
	}
}

func deleteCommitInfo(batch dbm.Batch, version int64) {
	cInfoKey := fmt.Sprintf(commitInfoKeyFmt, version)
	if err := batch.Delete([]byte(cInfoKey)); err != nil {
		panic(err)
	}
}
