package evm

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"

	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	"github.com/evmos/ethermint/x/evm/types"
	evm "github.com/evmos/ethermint/x/evm/vm"
	"github.com/evmos/ethermint/x/evm/vm/geth"
)

var (
	_ EVMKeeper = &Keeper{}
)

type Keeper struct {
	*evmkeeper.Keeper

	// custom stateless precompiled smart contracts
	customPrecompiles evm.PrecompiledContracts
	creator           Creator
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey, transientKey storetypes.StoreKey,
	authority sdk.AccAddress,
	ak types.AccountKeeper,
	bankKeeper types.BankKeeper,
	sk types.StakingKeeper,
	fmk types.FeeMarketKeeper,
	customPrecompiles evm.PrecompiledContracts,
	_ evm.Constructor,
	tracer string,
	ss paramstypes.Subspace,
) *Keeper {
	k := &Keeper{
		customPrecompiles: customPrecompiles,
		creator:           DefaultCreator{},
	}
	k.Keeper = evmkeeper.NewKeeper(
		cdc,
		storeKey,
		transientKey,
		authority,
		ak,
		bankKeeper,
		sk,
		fmk,
		customPrecompiles,
		k.evmCreator(),
		tracer,
		ss,
	)
	return k
}

// NewEVM override the evmkeeper.NewEVM method
func (k *Keeper) evmCreator() evm.Constructor {
	return func(
		ctx sdk.Context,
		blockCtx vm.BlockContext,
		txCtx vm.TxContext,
		stateDB vm.StateDB,
		chainConfig *params.ChainConfig,
		config vm.Config,
		_ evm.PrecompiledContracts,
	) evm.EVM {
		blockCtx.CanTransfer = k.creator.CanTransfer(ctx)
		blockCtx.Transfer = k.creator.Transfer(ctx)
		return &geth.EVM{
			EVM: vm.NewEVM(blockCtx, txCtx, stateDB, chainConfig, config),
		}
	}
}

func (k *Keeper) SetCreator(creator Creator) *Keeper {
	if creator != nil {
		k.creator = creator
	}
	return k
}

func (k *Keeper) GetBaseDenom(ctx sdk.Context) string {
	return k.GetParams(ctx).EvmDenom
}

type Creator interface {
	CanTransfer(ctx sdk.Context) vm.CanTransferFunc
	Transfer(ctx sdk.Context) vm.TransferFunc
}

type DefaultCreator struct{}

func (DefaultCreator) CanTransfer(_ sdk.Context) vm.CanTransferFunc {
	return core.CanTransfer
}

func (DefaultCreator) Transfer(_ sdk.Context) vm.TransferFunc {
	return core.Transfer
}
