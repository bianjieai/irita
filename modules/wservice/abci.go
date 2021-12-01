package wservice

import (
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/bianjieai/irita/modules/wservice/keeper"
	"github.com/bianjieai/irita/modules/wservice/types"
)

// EndBlocker handles block ending logic for service
func EndBlocker(ctx sdk.Context, k keeper.IKeeper) {
	ctx = ctx.WithLogger(ctx.Logger().With("handler", "endBlock").With("module", "wservice"))
	prefix := strconv.Itoa(int(ctx.BlockHeight())) + "-" + types.ModuleName
	iterator := k.QueryReqSequencesFromPrefix(ctx, []byte(prefix))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		valueBytes := iterator.Key()
		keyList := strings.Split(string(valueBytes), "-")
		reqSequence := keyList[len(keyList)-1]
		k.DeleteReqSequence(ctx, []byte(reqSequence))
		k.DeleteReqSequence(ctx, valueBytes)
	}
}
