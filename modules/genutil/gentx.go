package genutil

// DONTCOVER

import (
	"encoding/json"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/genutil/types"
)

// AddGenTxsInAppGenesisState - add the genesis transactions in the app genesis state
func AddGenTxsInAppGenesisState(
	cdc codec.JSONMarshaler, txJSONEncoder sdk.TxEncoder,
	appGenesisState map[string]json.RawMessage, genTxs []sdk.Tx,
) (map[string]json.RawMessage, error) {

	genesisState := GetGenesisStateFromAppState(cdc, appGenesisState)
	genTxsBz := genesisState.GenTxs

	for _, genTx := range genTxs {
		txBz, err := txJSONEncoder(genTx)
		if err != nil {
			return appGenesisState, err
		}

		genTxsBz = append(genTxsBz, txBz)
	}

	genesisState.GenTxs = genTxsBz
	return SetGenesisStateInAppState(cdc, appGenesisState, genesisState), nil
}

type deliverTxfn func(abci.RequestDeliverTx) abci.ResponseDeliverTx

// DeliverGenTxs iterates over all genesis txs, decodes each into a StdTx and
// invokes the provided deliverTxfn with the decoded StdTx. It returns the result
// of the validator module's ApplyAndReturnValidatorSetUpdates.
func DeliverGenTxs(
	ctx sdk.Context, genTxs []json.RawMessage,
	validatorKeeper types.ValidatorKeeper, deliverTx deliverTxfn,
	txEncodingConfig client.TxEncodingConfig,
) []abci.ValidatorUpdate {
	for _, genTx := range genTxs {
		tx, err := txEncodingConfig.TxJSONDecoder()(genTx)
		if err != nil {
			panic(err)
		}

		bz, err := txEncodingConfig.TxEncoder()(tx)
		if err != nil {
			panic(err)
		}

		res := deliverTx(abci.RequestDeliverTx{Tx: bz})
		if !res.IsOK() {
			panic(res.Log)
		}
	}
	validators, err := validatorKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		panic(err)
	}

	return validators
}
