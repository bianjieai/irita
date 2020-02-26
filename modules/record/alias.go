package record

import (
	"github.com/bianjieai/irita/modules/record/internal/keeper"
	"github.com/bianjieai/irita/modules/record/internal/types"
)

const (
	ModuleName   = types.ModuleName
	StoreKey     = types.StoreKey
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute

	EventTypeCreateRecord  = types.EventTypeCreateRecord
	AttributeKeyCreator    = types.AttributeKeyCreator
	AttributeKeyRecordID   = types.AttributeKeyRecordID
	AttributeValueCategory = types.AttributeValueCategory
)

var (
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	RegisterCodec       = types.RegisterCodec
	NewRecord           = types.NewRecord
	NewQuerier          = keeper.NewQuerier
	NewKeeper           = keeper.NewKeeper

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper          = keeper.Keeper
	GenesisState    = types.GenesisState
	Record          = types.Record
	Content         = types.Content
	MsgCreateRecord = types.MsgCreateRecord
)
