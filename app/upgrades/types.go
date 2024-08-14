package types

import (
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	consensuskeeper "github.com/cosmos/cosmos-sdk/x/consensus/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	feemarketkeeper "github.com/evmos/ethermint/x/feemarket/keeper"
	identitykeeper "iritamod.bianjie.ai/modules/identity/keeper"
	nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	paramskeeper "iritamod.bianjie.ai/modules/params/keeper"
	slashingkeeper "iritamod.bianjie.ai/modules/slashing/keeper"
	upgradekeeper "iritamod.bianjie.ai/modules/upgrade/keeper"
	mtkeeper "mods.irisnet.org/modules/mt/keeper"
	nftkeeper "mods.irisnet.org/modules/nft/keeper"
	oraclekeeper "mods.irisnet.org/modules/oracle/keeper"
	randomkeeper "mods.irisnet.org/modules/random/keeper"
	recordkeeper "mods.irisnet.org/modules/record/keeper"
	servicekeeper "mods.irisnet.org/modules/service/keeper"
	tokenkeeper "mods.irisnet.org/modules/token/keeper"
)

type Upgrade struct {
	// Upgrade version name, for the upgrade handler, e.g. `v7`
	UpgradeName string

	// UpgradeHandlerConstructor defines the function that creates an upgrade handler
	UpgradeHandlerConstructor func(*module.Manager, module.Configurator, AppKeepers) upgradetypes.UpgradeHandler

	// Store upgrades, should be used for any new modules introduced, new modules deleted, or store names renamed.
	StoreUpgrades *store.StoreUpgrades
}

type ConsensusParamsReaderWriter interface {
	StoreConsensusParams(ctx sdk.Context, cp *tmproto.ConsensusParams)
	GetConsensusParams(ctx sdk.Context) *tmproto.ConsensusParams
}

type AppKeepers struct {
	AppCodec              codec.Codec
	AccountKeeper         authkeeper.AccountKeeper
	BankKeeper            bankkeeper.Keeper
	SlashingKeeper        slashingkeeper.Keeper
	CrisisKeeper          *crisiskeeper.Keeper
	UpgradeKeeper         *upgradekeeper.Keeper
	ParamsKeeper          paramskeeper.Keeper
	EvidenceKeeper        evidencekeeper.Keeper
	RecordKeeper          recordkeeper.Keeper
	TokenKeeper           tokenkeeper.Keeper
	NftKeeper             nftkeeper.Keeper
	MtKeeper              mtkeeper.Keeper
	ServiceKeeper         servicekeeper.Keeper
	OracleKeeper          oraclekeeper.Keeper
	RandomKeeper          randomkeeper.Keeper
	IdentityKeeper        identitykeeper.Keeper
	NodeKeeper            *nodekeeper.Keeper
	FeeGrantKeeper        feegrantkeeper.Keeper
	CapabilityKeeper      *capabilitykeeper.Keeper
	ConsensusParamsKeeper consensuskeeper.Keeper
	EvmKeeper             *evmkeeper.Keeper
	FeeMarketKeeper       feemarketkeeper.Keeper
}

type upgradeRouter struct {
	mu map[string]Upgrade
}

func NewUpgradeRouter() *upgradeRouter {
	return &upgradeRouter{make(map[string]Upgrade)}
}

func (r *upgradeRouter) Register(u Upgrade) *upgradeRouter {
	if _, has := r.mu[u.UpgradeName]; has {
		panic(u.UpgradeName + " already registered")
	}
	r.mu[u.UpgradeName] = u
	return r
}

func (r *upgradeRouter) Routers() map[string]Upgrade {
	return r.mu
}

func (r *upgradeRouter) UpgradeInfo(planName string) Upgrade {
	return r.mu[planName]
}
