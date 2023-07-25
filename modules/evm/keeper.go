package evm

import (
	"math/big"

	opbkeeper "github.com/bianjieai/irita/modules/opb/keeper"
	permkeeper "github.com/bianjieai/iritamod/modules/perm/keeper"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	tokenkeeper "github.com/irisnet/irismod/modules/token/keeper"

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

	opbKeeper   opbkeeper.Keeper
	tokenKeeper tokenkeeper.Keeper
	permKeeper  permkeeper.Keeper

	// custom stateless precompiled smart contracts
	customPrecompiles evm.PrecompiledContracts
}

type Option func(k *Keeper)

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
	opbKeeper opbkeeper.Keeper,
	tokenKeeper tokenkeeper.Keeper,
	permKeeper permkeeper.Keeper,
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
		opbKeeper:         opbKeeper,
		tokenKeeper:       tokenKeeper,
		permKeeper:        permKeeper,
		customPrecompiles: customPrecompiles,
	}
}

func (k *Keeper) NewEVM(
	ctx sdk.Context,
	msg core.Message,
	cfg *statedb.EVMConfig,
	tracer vm.EVMLogger,
	stateDB vm.StateDB,
) evm.EVM {
	validator := k.NewEthOpbValidator(ctx)
	blockCtx := vm.BlockContext{
		CanTransfer: validator.CanTransfer,
		Transfer:    validator.Transfer,
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

func (k *Keeper) NewEthOpbValidator(ctx sdk.Context) *EthOpbValidator {
	return NewEthOpbValidator(ctx,
		k.opbKeeper,
		k.tokenKeeper,
		k,
		k.permKeeper,
	)
}
