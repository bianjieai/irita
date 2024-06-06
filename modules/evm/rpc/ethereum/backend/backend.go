package backend

import (
	"context"
	"fmt"
	"math/big"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
	"github.com/tharsis/ethermint/rpc/ethereum/backend"
	"github.com/tharsis/ethermint/rpc/ethereum/types"
	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/bianjieai/irita/modules/evm/crypto"
)

type EVMWBackend struct {
	*backend.EVMBackend
	ctx         *client.Context
	queryClient *types.QueryClient
	logger      log.Logger
}

func NewEVMWBackend(ctx *server.Context, logger log.Logger, clientCtx client.Context) *EVMWBackend {
	evmBackend := backend.NewEVMBackend(ctx, logger, clientCtx)

	return &EVMWBackend{evmBackend, &clientCtx, types.NewQueryClient(clientCtx), logger}
}

func (e *EVMWBackend) SendTransaction(args evmtypes.TransactionArgs) (common.Hash, error) {
	// Look up the wallet containing the requested signer

	info, err := e.ctx.Keyring.KeyByAddress(sdk.AccAddress(args.From.Bytes()))
	if err != nil {

		e.logger.Error("failed to find key in keyring", "address", args.From, "error", err.Error())
		return common.Hash{}, fmt.Errorf("%s; %s", keystore.ErrNoMatch, err.Error())
	}

	args, err = e.SetTxDefaults(args)
	if err != nil {
		return common.Hash{}, err
	}

	msg := args.ToTransaction()
	if err := msg.ValidateBasic(); err != nil {
		e.logger.Debug("tx failed basic validation", "error", err.Error())
		return common.Hash{}, err
	}

	signer := crypto.NewSm2Signer(e.ChainConfig().ChainID)
	if info.GetAlgo() == ethsecp256k1.KeyType {
		// eth
		fmt.Println("SendTransaction", info.GetAlgo())
		bn, err := e.BlockNumber()
		if err != nil {
			e.logger.Debug("failed to fetch latest block number", "error", err.Error())
			return common.Hash{}, err
		}
		signer = ethtypes.MakeSigner(e.ChainConfig(), new(big.Int).SetUint64(uint64(bn)))
	}

	// Sign transaction
	if err := msg.Sign(signer, e.ctx.Keyring); err != nil {
		e.logger.Debug("failed to sign tx", "error", err.Error())
		return common.Hash{}, err
	}

	// Assemble transaction from fields
	builder, ok := e.ctx.TxConfig.NewTxBuilder().(authtx.ExtensionOptionsTxBuilder)
	if !ok {
		e.logger.Error("clientCtx.TxConfig.NewTxBuilder returns unsupported builder", "error", err.Error())
	}

	option, err := codectypes.NewAnyWithValue(&evmtypes.ExtensionOptionsEthereumTx{})
	if err != nil {
		e.logger.Error("codectypes.NewAnyWithValue failed to pack an obvious value", "error", err.Error())
		return common.Hash{}, err
	}

	builder.SetExtensionOptions(option)
	if err = builder.SetMsgs(msg); err != nil {
		e.logger.Error("builder.SetMsgs failed", "error", err.Error())
	}

	// Query params to use the EVM denomination
	res, err := e.queryClient.QueryClient.Params(context.Background(), &evmtypes.QueryParamsRequest{})
	if err != nil {
		e.logger.Error("failed to query evm params", "error", err.Error())
		return common.Hash{}, err
	}

	txData, err := evmtypes.UnpackTxData(msg.Data)
	if err != nil {
		e.logger.Error("failed to unpack tx data", "error", err.Error())
		return common.Hash{}, err
	}

	fees := sdk.Coins{sdk.NewCoin(res.Params.EvmDenom, sdk.NewIntFromBigInt(txData.Fee()))}
	builder.SetFeeAmount(fees)
	builder.SetGasLimit(msg.GetGas())

	// Encode transaction by default Tx encoder
	txEncoder := e.ctx.TxConfig.TxEncoder()
	txBytes, err := txEncoder(builder.GetTx())
	if err != nil {
		e.logger.Error("failed to encode eth tx using default encoder", "error", err.Error())
		return common.Hash{}, err
	}

	txHash := msg.AsTransaction().Hash()

	// Broadcast transaction in sync mode (default)
	// NOTE: If error is encountered on the node, the broadcast will not return an error
	syncCtx := e.ctx.WithBroadcastMode(flags.BroadcastSync)
	rsp, err := syncCtx.BroadcastTx(txBytes)
	if rsp != nil && rsp.Code != 0 {
		err = sdkerrors.ABCIError(rsp.Codespace, rsp.Code, rsp.RawLog)
	}
	if err != nil {
		e.logger.Error("failed to broadcast tx", "error", err.Error())
		return txHash, err
	}

	// Return transaction hash
	return txHash, nil
}
