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

type CanTransferCreator func(ctx sdk.Context) vm.CanTransferFunc
type TransferCreator func(ctx sdk.Context) vm.TransferFunc

type Keeper struct {
	*evmkeeper.Keeper

	// custom stateless precompiled smart contracts
	customPrecompiles evm.PrecompiledContracts
	createCanTransfer CanTransferCreator
	createTransfer    TransferCreator
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
		createCanTransfer: func(ctx sdk.Context) vm.CanTransferFunc {
			return core.CanTransfer
		},
		createTransfer: func(ctx sdk.Context) vm.TransferFunc {
			return core.Transfer
		},
	}
}

func (k *Keeper) NewEVM(
	ctx sdk.Context,
	msg core.Message,
	cfg *statedb.EVMConfig,
	tracer vm.EVMLogger,
	stateDB vm.StateDB,
) evm.EVM {
	blockCtx := vm.BlockContext{
		CanTransfer: k.createCanTransfer(ctx),
		Transfer:    k.createTransfer(ctx),
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

func (k *Keeper) SetValidator(
	canTransferCreator CanTransferCreator,
	transferCreator TransferCreator,
) *Keeper {
	if canTransferCreator != nil {
		k.createCanTransfer = canTransferCreator
	}
	if transferCreator != nil {
		k.createTransfer = transferCreator
	}

	return k
}
