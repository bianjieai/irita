package keeper

import (
	"github.com/bianjieai/irita/modules/opb/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// ValidateTokenTransferDecorator checks if the token transfer satisfies the underlying constraint
type ValidateTokenTransferDecorator struct {
	keeper      Keeper
	tokenKeeper types.TokenKeeper
}

// NewValidateTokenTransferDecorator constructs a new ValidateTokenTransferDecorator instance
func NewValidateTokenTransferDecorator(
	keeper Keeper,
	tokenKeeper types.TokenKeeper,
) ValidateTokenTransferDecorator {
	return ValidateTokenTransferDecorator{
		keeper:      keeper,
		tokenKeeper: tokenKeeper,
	}
}

// AnteHandle implements AnteHandler
func (vtd ValidateTokenTransferDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	restrictionEnabled := !vtd.keeper.UnrestrictedTokenTransfer(ctx)

	// check only if the transfer restriction is enabled
	if restrictionEnabled {
		for _, msg := range tx.GetMsgs() {
			switch msg := msg.(type) {
			case *banktypes.MsgSend:
				err := vtd.validateMsgSend(ctx, msg)
				if err != nil {
					return ctx, err
				}
			case *banktypes.MsgMultiSend:
				err := vtd.validateMsgMultiSend(ctx, msg)
				if err != nil {
					return ctx, err
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}

// validateMsgSend validates the MsgSend msg
func (vtd ValidateTokenTransferDecorator) validateMsgSend(ctx sdk.Context, msg *banktypes.MsgSend) error {
	for _, coin := range msg.Amount {
		owner, err := vtd.getOwner(ctx, coin.Denom)
		if err != nil {
			continue
		}

		if msg.FromAddress != owner && msg.ToAddress != owner {
			return sdkerrors.Wrapf(
				types.ErrUnauthorized,
				"either the sender or recipient must be the owner %s for token %s",
				owner, coin.Denom,
			)
		}
	}

	return nil
}

// validateMsgMultiSend validates the MsgMultiSend msg
func (vtd ValidateTokenTransferDecorator) validateMsgMultiSend(ctx sdk.Context, msg *banktypes.MsgMultiSend) error {
	inputMap := getInputMap(msg.Inputs)
	outputMap := getOutputMap(msg.Outputs)

	for denom, addresses := range inputMap {
		owner, err := vtd.getOwner(ctx, denom)
		if err != nil {
			continue
		}

		if !owned(owner, addresses) && !owned(owner, outputMap[denom]) {
			return sdkerrors.Wrapf(
				types.ErrUnauthorized,
				"either the sender or recipient must be the owner %s for token %s",
				owner, denom,
			)
		}
	}

	return nil
}

// getOwner gets the owner of the specified denom
func (vtd ValidateTokenTransferDecorator) getOwner(ctx sdk.Context, denom string) (owner string, err error) {
	baseTokenDenom := vtd.keeper.BaseTokenDenom(ctx)

	if denom == baseTokenDenom {
		owner = vtd.keeper.BaseTokenManager(ctx)
	} else {
		ownerAddr, err := vtd.tokenKeeper.GetOwner(ctx, denom)
		if err == nil {
			owner = ownerAddr.String()
		}
	}

	return
}

// validateContractFunds validates the funds in the contract transactions
func (vtd ValidateTokenTransferDecorator) validateContractFunds(ctx sdk.Context, coins sdk.Coins) error {
	baseTokenDenom := vtd.keeper.BaseTokenDenom(ctx)

	for _, coin := range coins {
		if coin.Denom == baseTokenDenom {
			return sdkerrors.Wrapf(
				types.ErrUnauthorized,
				"%s not allowed for contract transactions",
				coin.Denom,
			)
		}
	}

	return nil
}

// owned returns false if any address is not the owner of the denom among the given non-empty addresses
// True otherwise
func owned(owner string, addresses []string) bool {
	for _, addr := range addresses {
		if addr != owner {
			return false
		}
	}

	return true
}

// getInputMap maps input denoms to addresses
func getInputMap(inputs []banktypes.Input) map[string][]string {
	inputMap := make(map[string][]string)

	for _, input := range inputs {
		for _, coin := range input.Coins {
			inputMap[coin.Denom] = append(inputMap[coin.Denom], input.Address)
		}
	}

	return inputMap
}

// getOutputMap maps output denoms to addresses
func getOutputMap(outputs []banktypes.Output) map[string][]string {
	outputMap := make(map[string][]string)

	for _, output := range outputs {
		for _, coin := range output.Coins {
			outputMap[coin.Denom] = append(outputMap[coin.Denom], output.Address)
		}
	}

	return outputMap
}
