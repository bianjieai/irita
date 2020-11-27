package app

import (
	"os"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/bianjieai/iritamod/modules/admin"
	"github.com/bianjieai/iritamod/modules/node"
)

var (
	rootAdmin = sdk.AccAddress(tmhash.SumTruncated([]byte("rootAdmin")))
	rootStr   = `-----BEGIN CERTIFICATE-----
MIIBxTCCAXegAwIBAgIUHMPutrm+7FT7fIFf2fEgyQnIg8kwBQYDK2VwMFgxCzAJ
BgNVBAYTAkNOMQ0wCwYDVQQIDARyb290MQ0wCwYDVQQHDARyb290MQ0wCwYDVQQK
DARyb290MQ0wCwYDVQQLDARyb290MQ0wCwYDVQQDDARyb290MB4XDTIwMDYxOTA3
MDExMVoXDTIxMDYxOTA3MDExMVowWDELMAkGA1UEBhMCQ04xDTALBgNVBAgMBHJv
b3QxDTALBgNVBAcMBHJvb3QxDTALBgNVBAoMBHJvb3QxDTALBgNVBAsMBHJvb3Qx
DTALBgNVBAMMBHJvb3QwKjAFBgMrZXADIQDdzGFcck4I7Wa1vRj4JsdQ9RjVgH92
7iOhXJ8mFLwQKaNTMFEwHQYDVR0OBBYEFPrjTGR+/g4RUduZ9E8JSXNyI4mzMB8G
A1UdIwQYMBaAFPrjTGR+/g4RUduZ9E8JSXNyI4mzMA8GA1UdEwEB/wQFMAMBAf8w
BQYDK2VwA0EAT8EG5nGxwCAP4ZlfQvAhrnJI+SojlsOoE3rZ8W6/knZsrnVb6RI8
QAVleeE0pMY+MtENXcQ2wH0QRXs+wO0XCw==
-----END CERTIFICATE-----`
)

func TestIritaExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewIritaApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, simapp.FlagPeriodValue, MakeEncodingConfig(), simapp.EmptyAppOptions{}, interBlockCacheOpt())

	_ = setGenesis(app)

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewIritaApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, simapp.FlagPeriodValue, MakeEncodingConfig(), simapp.EmptyAppOptions{}, interBlockCacheOpt())
	_, err := app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

// ensure that black listed addresses are properly set in bank keeper
func TestBlackListedAddrs(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewIritaApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{}, DefaultNodeHome, simapp.FlagPeriodValue, MakeEncodingConfig(), simapp.EmptyAppOptions{}, interBlockCacheOpt())

	for acc := range maccPerms {
		require.Equal(t, !allowedReceivingModAcc[acc], app.bankKeeper.BlockedAddr(app.accountKeeper.GetModuleAddress(acc)))
	}
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}

func setGenesis(iapp *IritaApp) error {
	genesisState := NewDefaultGenesisState()

	// add root cert
	validatorGenState := node.GetGenesisStateFromAppState(iapp.appCodec, genesisState)
	validatorGenState.RootCert = rootStr
	validatorGenStateBz := iapp.cdc.MustMarshalJSON(validatorGenState)
	genesisState[node.ModuleName] = validatorGenStateBz

	// add root admin
	adminGenState := admin.GetGenesisStateFromAppState(iapp.appCodec, genesisState)
	adminGenState.RoleAccounts = append(
		adminGenState.RoleAccounts,
		admin.RoleAccount{
			Address: rootAdmin.String(),
			Roles:   []admin.Role{admin.RoleRootAdmin},
		},
	)
	adminGenStateBz := iapp.cdc.MustMarshalJSON(adminGenState)
	genesisState[admin.ModuleName] = adminGenStateBz

	stateBytes, err := codec.MarshalJSONIndent(iapp.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	iapp.InitChain(abci.RequestInitChain{
		Validators:    []abci.ValidatorUpdate{},
		AppStateBytes: stateBytes,
	})

	iapp.Commit()
	return nil
}
