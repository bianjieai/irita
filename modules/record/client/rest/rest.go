package rest

import (
	"github.com/bianjieai/irita/modules/record/internal/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
)

// Rest variable names
// nolint
const (
	RestRecordID = "record-id"
)

// RegisterRoutes defines routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}

type RecordCreateReq struct {
	BaseReq  rest.BaseReq    `json:"base_req" yaml:"base_req"` // base req
	Contents []types.Content `json:"contents" yaml:"contents"`
	Creator  sdk.AccAddress  `json:"creator" yaml:"creator"`
}
