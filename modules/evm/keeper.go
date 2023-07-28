package evm

import (
	"math/big"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/vm"

	ethermint "github.com/evmos/ethermint/types"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
	"github.com/evmos/ethermint/x/evm/statedb"
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
	evmConstructor evm.Constructor,
	tracer string,
	ss paramstypes.Subspace,
) *Keeper {
	evmKeeper := evmkeeper.NewKeeper(
		cdc,
		storeKey,
		transientKey,
		authority,
		ak,
		bankKeeper,
		sk,
		fmk,
		customPrecompiles,
		evmConstructor,
		tracer,
		ss,
	)
	return &Keeper{
		Keeper:            evmKeeper,
		customPrecompiles: customPrecompiles,
		creator:           DefaultCreator{},
	}
}

// NewEVM override the evmkeeper.NewEVM method
func (k *Keeper) NewEVM(
	ctx sdk.Context,
	msg core.Message,
	cfg *statedb.EVMConfig,
	tracer vm.EVMLogger,
	stateDB vm.StateDB,
) evm.EVM {
	blockCtx := vm.BlockContext{
		CanTransfer: k.creator.CanTransfer(ctx),
		Transfer:    k.creator.Transfer(ctx),
		GetHash:     k.GetHashFn(ctx),
		Coinbase:    cfg.CoinBase,
		GasLimit:    ethermint.BlockGasLimit(ctx),
		BlockNumber: big.NewInt(ctx.BlockHeight()),
		Time:        big.NewInt(ctx.BlockHeader().Time.Unix()),
		Difficulty:  big.NewInt(0), // unused. Only required in PoW context
		BaseFee:     cfg.BaseFee,
		Random:      nil, // not supported
	}

	txCtx := core.NewEVMTxContext(msg)
	if tracer == nil {
		tracer = k.Tracer(ctx, msg, cfg.ChainConfig)
	}
	vmConfig := k.VMConfig(ctx, msg, cfg, tracer)
	return &geth.EVM{
		EVM: vm.NewEVM(blockCtx, txCtx, stateDB, cfg.ChainConfig, vmConfig),
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
