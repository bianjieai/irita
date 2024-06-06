package app

import (
	"encoding/json"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/irisnet/irismod/modules/service"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/bianjieai/iritamod/modules/node"
)

// ExportAppStateAndValidators export the state of irita for a genesis file
func (app *IritaApp) ExportAppStateAndValidators(forZeroHeight bool, jailAllowedAddrs []string) (servertypes.ExportedApp, error) {
	// as if they could withdraw from the start of the next block
	ctx := app.NewContext(true, tmproto.Header{Height: app.LastBlockHeight()})

	height := app.LastBlockHeight() + 1
	if forZeroHeight {
		height = 0
		app.prepForZeroHeightGenesis(ctx, jailAllowedAddrs)
	}

	genState := app.mm.ExportGenesis(ctx, app.appCodec)
	appState, err := json.MarshalIndent(genState, "", "  ")
	if err != nil {
		return servertypes.ExportedApp{}, err
	}

	validators := node.WriteValidators(ctx, app.nodeKeeper)
	return servertypes.ExportedApp{
		AppState:        appState,
		Validators:      validators,
		Height:          height,
		ConsensusParams: app.BaseApp.GetConsensusParams(ctx),
	}, nil
}

// prepare for fresh start at zero height
// NOTE zero height genesis is a temporary feature which will be deprecated
//
//	in favour of export at a block height
func (app *IritaApp) prepForZeroHeightGenesis(ctx sdk.Context, jailAllowedAddrs []string) {

	/* Just to be safe, assert the invariants on current state. */
	app.crisisKeeper.AssertInvariants(ctx)

	service.PrepForZeroHeightGenesis(ctx, app.serviceKeeper)
}
