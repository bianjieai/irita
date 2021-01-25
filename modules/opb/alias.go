package opb

import (
	"github.com/bianjieai/irita/modules/opb/keeper"
	"github.com/bianjieai/irita/modules/opb/types"
)

const (
	ModuleName             = types.ModuleName
	StoreKey               = types.StoreKey
	QuerierRoute           = types.QuerierRoute
	RouterKey              = types.RouterKey
	EventTypeMint          = types.EventTypeMint
	EventTypeReclaim       = types.EventTypeReclaim
	AttributeKeyRecipient  = types.AttributeKeyRecipient
	AttributeValueCategory = types.AttributeValueCategory
)

var (
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewKeeper           = keeper.NewKeeper
)

type (
	MsgMint      = types.MsgMint
	MsgReclaim   = types.MsgReclaim
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
)
