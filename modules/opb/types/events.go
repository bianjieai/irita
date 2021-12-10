package types

// OPB module event types
const (
	EventTypeMint           = "mint"
	EventTypeReclaim        = "reclaim"
	EventTypeContractAdd    = "contract_add"
	EventTypeContractRemove = "contract_remove"
	EventTypeAccountAdd     = "account_add"
	EventTypeAccountRemove  = "account_remove"

	AttributeKeyAmount     = "amount"
	AttributeKeyDenom      = "denom"
	AttributeKeyRecipient  = "recipient"
	AttributeValueCategory = ModuleName
)
