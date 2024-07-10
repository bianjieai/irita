// Package rpc contains RPC handler methods and utilities to start
// Ethermint's Web3-compatibly JSON-RPC server.
package rpc

import (
	rpcclient "github.com/cometbft/cometbft/rpc/jsonrpc/client"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/debug"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/eth"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/eth/filters"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/miner"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/net"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/personal"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/txpool"
	"github.com/evmos/ethermint/rpc/namespaces/ethereum/web3"
	ethermint "github.com/evmos/ethermint/types"

	"github.com/bianjieai/irita/modules/evm/rpc/ethereum/backend"
)

// RPC namespaces and API version
const (
	Web3Namespace     = "web3"
	EthNamespace      = "eth"
	PersonalNamespace = "personal"
	NetNamespace      = "net"
	TxPoolNamespace   = "txpool"
	DebugNamespace    = "debug"
	MinerNamespace    = "miner"

	apiVersion = "1.0"
)

// GetRPCAPIs returns the list of all APIs
func GetRPCAPIs(ctx *server.Context,
	clientCtx client.Context,
	tmWSClient *rpcclient.WSClient,
	allowUnprotectedTxs bool,
	indexer ethermint.EVMTxIndexer,
	selectedAPIs []string,
) []rpc.API {
	evmBackend := backend.NewEVMWBackend(ctx, ctx.Logger, clientCtx, allowUnprotectedTxs, indexer)

	var apis []rpc.API
	// remove duplicates
	selectedAPIs = unique(selectedAPIs)

	for index := range selectedAPIs {
		switch selectedAPIs[index] {
		case EthNamespace:
			apis = append(apis,
				rpc.API{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   eth.NewPublicAPI(ctx.Logger, evmBackend),
					Public:    true,
				},
				rpc.API{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   filters.NewPublicAPI(ctx.Logger, clientCtx,tmWSClient, evmBackend),
					Public:    true,
				},
			)
		case Web3Namespace:
			apis = append(apis,
				rpc.API{
					Namespace: Web3Namespace,
					Version:   apiVersion,
					Service:   web3.NewPublicAPI(),
					Public:    true,
				},
			)
		case NetNamespace:
			apis = append(apis,
				rpc.API{
					Namespace: NetNamespace,
					Version:   apiVersion,
					Service:   net.NewPublicAPI(clientCtx),
					Public:    true,
				},
			)
		case PersonalNamespace:
			apis = append(apis,
				rpc.API{
					Namespace: PersonalNamespace,
					Version:   apiVersion,
					Service:   personal.NewAPI(ctx.Logger, evmBackend),
					Public:    false,
				},
			)
		case TxPoolNamespace:
			apis = append(apis,
				rpc.API{
					Namespace: TxPoolNamespace,
					Version:   apiVersion,
					Service:   txpool.NewPublicAPI(ctx.Logger),
					Public:    true,
				},
			)
		case DebugNamespace:
			apis = append(apis,
				rpc.API{
					Namespace: DebugNamespace,
					Version:   apiVersion,
					Service:   debug.NewAPI(ctx, evmBackend),
					Public:    true,
				},
			)
		case MinerNamespace:
			apis = append(apis,
				rpc.API{
					Namespace: MinerNamespace,
					Version:   apiVersion,
					Service:   miner.NewPrivateAPI(ctx,evmBackend),
					Public:    false,
				},
			)
		default:
			ctx.Logger.Error("invalid namespace value", "namespace", selectedAPIs[index])
		}
	}

	return apis
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
