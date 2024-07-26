package wrapper

import (
	nfttransfertypes "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"
	nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	slashingkeeper "iritamod.bianjie.ai/modules/slashing/keeper"
	nftkeeper "mods.irisnet.org/modules/nft/keeper"
)

// ProvideSlashingStakingKeeper provides a slashing staking keeper
func ProvideSlashingStakingKeeper(nodeKeeper *nodekeeper.Keeper) slashingtypes.StakingKeeper {
	return NewStakingKeeper(nodeKeeper)
}

// ProvideEvidenceStakingKeeper provides a evidence staking keeper
func ProvideEvidenceStakingKeeper(nodeKeeper *nodekeeper.Keeper) evidencetypes.StakingKeeper {
	return NewStakingKeeper(nodeKeeper)
}

// ProvideEvmStakingKeeper provides a evm staking keeper
func ProvideEvmStakingKeeper(nodeKeeper *nodekeeper.Keeper) evmtypes.StakingKeeper {
	return NewStakingKeeper(nodeKeeper)
}

// ProvideStakingHooks provides a staking hooks
func ProvideStakingHooks(slashingKeeper slashingkeeper.Keeper) []stakingtypes.StakingHooks {
	return []stakingtypes.StakingHooks{slashingKeeper.Hooks()}
}

// ProvideTibcNftKeeper provides a tibc nft keeper
func ProvideTibcNftKeeper(k nftkeeper.Keeper) nfttransfertypes.NftKeeper {
	return nftkeeper.NewLegacyKeeper(k)
}