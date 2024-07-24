package wrapper

import (
	"cosmossdk.io/math"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	nodekeeper "iritamod.bianjie.ai/modules/node/keeper"
	slashingkeeper "iritamod.bianjie.ai/modules/slashing/keeper"

	evmtypes "github.com/evmos/ethermint/x/evm/types"
)

var (
	_ slashingtypes.StakingKeeper = (*StakingKeeper)(nil)
	_ evidencetypes.StakingKeeper = (*StakingKeeper)(nil)
	_ evmtypes.StakingKeeper      = (*StakingKeeper)(nil)
)

type StakingKeeperI interface {
	slashingtypes.StakingKeeper
	evidencetypes.StakingKeeper
	evmtypes.StakingKeeper
}

func ProvideSlashingStakingKeeper(nodeKeeper *nodekeeper.Keeper) slashingtypes.StakingKeeper {
	return NewStakingKeeper(nodeKeeper)
}

func ProvideEvidenceStakingKeeper(nodeKeeper *nodekeeper.Keeper) evidencetypes.StakingKeeper {
	return NewStakingKeeper(nodeKeeper)
}

func ProvideEvmStakingKeeper(nodeKeeper *nodekeeper.Keeper) evmtypes.StakingKeeper {
	return NewStakingKeeper(nodeKeeper)
}

func ProvideStakingHooks(slashingKeeper slashingkeeper.Keeper) []stakingtypes.StakingHooks {
	return []stakingtypes.StakingHooks{slashingKeeper.Hooks()}
}

// StakingKeeper implements the staking keeper interface.
type StakingKeeper struct {
	nk *nodekeeper.Keeper
}

// NewStakingKeeper creates a new instance of the stakingKeeper struct.
//
// It takes a nodekeeper.Keeper as a parameter and returns a pointer to the stakingKeeper struct.
func NewStakingKeeper(nk *nodekeeper.Keeper) StakingKeeperI {
	return &StakingKeeper{
		nk: nk,
	}
}

// Delegation implements types.StakingKeeper.
func (s *StakingKeeper) Delegation(ctx sdk.Context, delegator sdk.AccAddress, validator sdk.ValAddress) types.DelegationI {
	return s.nk.Delegation(ctx, delegator, validator)
}

// GetAllValidators implements types.StakingKeeper.
func (s *StakingKeeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	vs := s.nk.GetAllValidators(ctx)

	for _, v := range vs {
		pubKey, err := v.ConsPubKey()
		if err != nil {
			panic(err)
		}

		pkAny, err := codectypes.NewAnyWithValue(pubKey)
		if err != nil {
			panic(err)
		}

		validators = append(validators, types.Validator{
			OperatorAddress: v.GetOperator().String(),
			ConsensusPubkey: pkAny,
			Jailed:          v.IsJailed(),
			Status:          v.GetStatus(),
			Tokens:          v.GetTokens(),
			DelegatorShares: v.GetDelegatorShares(),
			Description: types.Description{
				Moniker:  v.GetMoniker(),
				Details:  v.Description,
				Identity: v.Certificate,
			},
			MinSelfDelegation: math.ZeroInt(),
		})
	}
	return validators
}

// IsValidatorJailed implements types.StakingKeeper.
func (s *StakingKeeper) IsValidatorJailed(ctx sdk.Context, addr sdk.ConsAddress) bool {
	v, ok := s.nk.GetValidatorByConsAddr(ctx, addr)
	if !ok {
		return false
	}
	return v.Jailed
}

// IterateValidators implements types.StakingKeeper.
func (s *StakingKeeper) IterateValidators(ctx sdk.Context, fn func(index int64, validator types.ValidatorI) (stop bool)) {
	s.nk.IterateValidators(ctx, fn)
}

// Jail implements types.StakingKeeper.
func (s *StakingKeeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	s.nk.Jail(ctx, consAddr)
}

// MaxValidators implements types.StakingKeeper.
func (s *StakingKeeper) MaxValidators(ctx sdk.Context) uint32 {
	return s.nk.MaxValidators(ctx)
}

// Slash implements types.StakingKeeper.
func (s *StakingKeeper) Slash(ctx sdk.Context, consAddr sdk.ConsAddress, i int64, i2 int64, dec sdk.Dec) math.Int {
	s.nk.Slash(ctx, consAddr, i, i2, dec)
	return math.NewInt(0)
}

// SlashWithInfractionReason implements types.StakingKeeper.
func (s *StakingKeeper) SlashWithInfractionReason(ctx sdk.Context, consAddr sdk.ConsAddress, infractionHeight int64, power int64, slashFactor sdk.Dec, _ types.Infraction) math.Int {
	return s.Slash(ctx, consAddr, infractionHeight, power, slashFactor)
}

// Unjail implements types.StakingKeeper.
func (s *StakingKeeper) Unjail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	s.nk.Unjail(ctx, consAddr)
}

// Validator implements types.StakingKeeper.
func (s *StakingKeeper) Validator(ctx sdk.Context, valAddr sdk.ValAddress) types.ValidatorI {
	return s.nk.Validator(ctx, valAddr)
}

// ValidatorByConsAddr implements types.StakingKeeper.
func (s *StakingKeeper) ValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) types.ValidatorI {
	return s.nk.ValidatorByConsAddr(ctx, consAddr)
}

// GetParams implements types.StakingKeeper.
func (s *StakingKeeper) GetParams(ctx sdk.Context) types.Params {
	params := s.nk.GetParams(ctx)
	return types.Params{
		MaxEntries:        10,
		HistoricalEntries: params.HistoricalEntries,
		MinCommissionRate: math.LegacyZeroDec(),
	}
}

// GetHistoricalInfo implements types.StakingKeeper.
func (s *StakingKeeper) GetHistoricalInfo(ctx sdk.Context, height int64) (types.HistoricalInfo, bool) {
	return s.nk.GetHistoricalInfo(ctx, height)
}

// GetValidatorByConsAddr implements types.StakingKeeper.
func (s *StakingKeeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	addr, found := s.nk.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		return types.Validator{}, false
	}
	validator.Jailed = addr.Jailed

	_, i, err := bech32.DecodeAndConvert(addr.Operator)
	if err != nil {
		return types.Validator{}, false
	}
	validator.OperatorAddress = sdk.ValAddress(i).String()

	return validator, found
}
