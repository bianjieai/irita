package wservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/wservice/keeper"
)

func NewHandler(keeper keeper.IKeeper) sdk.Handler {
	return nil
}
