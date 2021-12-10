package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
)

const (
	// ModuleName is the name of the OPB module
	ModuleName = "opb"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// QuerierRoute is the querier route for the OPB module
	QuerierRoute = ModuleName

	// RouterKey is the msg router key for the OPB module
	RouterKey = ModuleName

	// PointTokenFeeCollectorName is the root string for the fee collector account address for the point token
	PointTokenFeeCollectorName = "opb_point_token_fee_collector"

	ContractDenyListName = "contract-deny-list"
	AccountDenyListName  = "account-deny-list"
)

const (
	KeyPrefixContractDenyList = "ContractDenyList"
	KeyPrefixAccountDenyList  = "AccountDenyList"
)

func ContractDenyListPath(contractAddress common.Address) string {
	return fmt.Sprintf("%s/%s", KeyPrefixContractDenyList, contractAddress)
}

// ContractDenyListKey defines the full key under which a contract deny list is stored.
func ContractDenyListKey(contractAddress common.Address) []byte {
	return []byte(ContractDenyListPath(contractAddress))
}

func AccountDenyListPath(accountAddress sdk.Address) string {
	return fmt.Sprintf("%s/%s", KeyPrefixAccountDenyList, accountAddress)
}

// AccountDenyListKey defines the full key under which an account deny list is stored.
func AccountDenyListKey(accountAddress sdk.Address) []byte {
	return []byte(AccountDenyListPath(accountAddress))
}
