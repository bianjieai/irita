package wservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/wservice/keeper"
	"github.com/bianjieai/irita/modules/wservice/types"
)

func InitGenesis(ctx sdk.Context, k keeper.IKeeper, data types.GenesisState) {
	if err := types.ValidateGenesis(data); err != nil {
		panic(err.Error())
	}
	for _, reqSeq := range data.ReqSequence {
		key := types.ModuleName + "_" + reqSeq.Key
		k.SetReqSequence(ctx, []byte(key), []byte(reqSeq.Value))
	}
}

func ExportGenesis(ctx sdk.Context, k keeper.IKeeper) *types.GenesisState {
	var reqSequences []types.RequestSequence

	iterator := k.QueryReqSequencesFromPrefix(ctx, []byte(types.ModuleName))
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		reqSequence := types.RequestSequence{
			Key:   string(iterator.Key()),
			Value: string(iterator.Value()),
		}
		reqSequences = append(reqSequences, reqSequence)
	}
	return types.NewGenesisState(reqSequences)
}

// DefaultGenesisState returns a default genesis state
func DefaultGenesisState() *types.GenesisState {
	return types.NewGenesisState([]types.RequestSequence{})
}
