package types

// OPB module event types
const (
	EventTypeMint           = "mint"
	EventTypeReclaim        = "reclaim"
	EventTypeContractAdd    = "contract_add"
	EventTypeContractRemove = "contract_remove"

	AttributeKeyAmount     = "amount"
	AttributeKeyDenom      = "denom"
	AttributeKeyRecipient  = "recipient"
	AttributeValueCategory = ModuleName
)
