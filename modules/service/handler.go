package service

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/bianjieai/irita/modules/service/internal/types"
)

// NewHandler returns a handler for all the "service" type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case MsgSvcDef:
			return handleMsgSvcDef(ctx, k, msg)
		case MsgSvcBind:
			return handleMsgSvcBind(ctx, k, msg)
		case MsgSvcBindingUpdate:
			return handleMsgSvcBindUpdate(ctx, k, msg)
		case MsgSvcDisable:
			return handleMsgSvcDisable(ctx, k, msg)
		case MsgSvcEnable:
			return handleMsgSvcEnable(ctx, k, msg)
		case MsgSvcRefundDeposit:
			return handleMsgSvcRefundDeposit(ctx, k, msg)
		case MsgSvcRequest:
			return handleMsgSvcRequest(ctx, k, msg)
		case MsgSvcResponse:
			return handleMsgSvcResponse(ctx, k, msg)
		case MsgSvcRefundFees:
			return handleMsgSvcRefundFees(ctx, k, msg)
		case MsgSvcWithdrawFees:
			return handleMsgSvcWithdrawFees(ctx, k, msg)
		case MsgSvcWithdrawTax:
			return handleMsgSvcWithdrawTax(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}

func handleMsgSvcDef(ctx sdk.Context, k Keeper, msg MsgSvcDef) (*sdk.Result, error) {
	err := k.AddServiceDefinition(ctx, msg.Name, msg.ChainId, msg.Description, msg.Tags,
		msg.Author, msg.AuthorDescription, msg.IDLContent)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("Create service definition", "name", msg.Name, "author", msg.Author.String())

	return &sdk.Result{}, nil
}

func handleMsgSvcBind(ctx sdk.Context, k Keeper, msg MsgSvcBind) (*sdk.Result, error) {
	err := k.AddServiceBinding(ctx, msg.DefChainID, msg.DefName, msg.BindChainID,
		msg.Provider, msg.BindingType, msg.Deposit, msg.Prices, msg.Level)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("Add service binding", "def_name", msg.DefName, "def_chain_id", msg.DefChainID,
		"provider", msg.Provider.String(), "binding_type", msg.BindingType.String())

	return &sdk.Result{}, nil
}

func handleMsgSvcBindUpdate(ctx sdk.Context, k Keeper, msg MsgSvcBindingUpdate) (*sdk.Result, error) {
	svcBinding, err := k.UpdateServiceBinding(ctx, msg.DefChainID, msg.DefName, msg.BindChainID,
		msg.Provider, msg.BindingType, msg.Deposit, msg.Prices, msg.Level)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("Update service binding", "def_name", msg.DefName, "def_chain_id", msg.DefChainID,
		"provider", msg.Provider.String(), "binding_type", svcBinding.BindingType.String())

	return &sdk.Result{}, nil
}

func handleMsgSvcDisable(ctx sdk.Context, k Keeper, msg MsgSvcDisable) (*sdk.Result, error) {
	err := k.Disable(ctx, msg.DefChainID, msg.DefName, msg.BindChainID, msg.Provider)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("Disable service binding", "def_name", msg.DefName, "def_chain_id", msg.DefChainID,
		"provider", msg.Provider.String())

	return &sdk.Result{}, nil
}

func handleMsgSvcEnable(ctx sdk.Context, k Keeper, msg MsgSvcEnable) (*sdk.Result, error) {
	err := k.Enable(ctx, msg.DefChainID, msg.DefName, msg.BindChainID, msg.Provider, msg.Deposit)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("Enable service binding", "def_name", msg.DefName, "def_chain_id", msg.DefChainID,
		"provider", msg.Provider.String())

	return &sdk.Result{}, nil
}

func handleMsgSvcRefundDeposit(ctx sdk.Context, k Keeper, msg MsgSvcRefundDeposit) (*sdk.Result, error) {
	err := k.RefundDeposit(ctx, msg.DefChainID, msg.DefName, msg.BindChainID, msg.Provider)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Info("Refund deposit", "def_name", msg.DefName, "def_chain_id", msg.DefChainID,
		"provider", msg.Provider.String())

	return &sdk.Result{}, nil
}

func handleMsgSvcRequest(ctx sdk.Context, k Keeper, msg MsgSvcRequest) (*sdk.Result, error) {
	req, err := k.AddRequest(ctx, msg.DefChainID, msg.DefName, msg.BindChainID, msg.ReqChainID,
		msg.Consumer, msg.Provider, msg.MethodID, msg.Input, msg.ServiceFee, msg.Profiling)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debug("Service request", "def_name", req.DefName, "def_chain_id", req.DefChainID,
		"provider", req.Provider.String(), "consumer", req.Consumer.String(), "method_id", req.MethodID,
		"service_fee", req.ServiceFee, "request_id", req.RequestID())

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeySender, req.Consumer.String()),
			),
			sdk.NewEvent(
				EventTypeRequestSvc,
				sdk.NewAttribute(AttributeKeyRequestID, req.RequestID()),
				sdk.NewAttribute(AttributeKeyProvider, req.Provider.String()),
				sdk.NewAttribute(AttributeKeyServiceFee, req.ServiceFee.String()),
			),
		},
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSvcResponse(ctx sdk.Context, k Keeper, msg MsgSvcResponse) (*sdk.Result, error) {
	resp, err := k.AddResponse(ctx, msg.ReqChainID, msg.RequestID, msg.Provider, msg.Output, msg.ErrorMsg)
	if err != nil {
		return nil, err
	}

	ctx.Logger().Debug("Service response", "request_id", msg.RequestID,
		"provider", resp.Provider.String(), "consumer", resp.Consumer.String())

	ctx.EventManager().EmitEvents(
		sdk.Events{
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeySender, resp.Provider.String()),
			),
			sdk.NewEvent(
				types.EventTypeRespondSvc,
				sdk.NewAttribute(AttributeKeyRequestID, msg.RequestID),
				sdk.NewAttribute(AttributeKeyConsumer, resp.Consumer.String()),
			),
		},
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleMsgSvcRefundFees(ctx sdk.Context, k Keeper, msg MsgSvcRefundFees) (*sdk.Result, error) {
	err := k.RefundFee(ctx, msg.Consumer)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}

func handleMsgSvcWithdrawFees(ctx sdk.Context, k Keeper, msg MsgSvcWithdrawFees) (*sdk.Result, error) {
	err := k.WithdrawFee(ctx, msg.Provider)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}

func handleMsgSvcWithdrawTax(ctx sdk.Context, k Keeper, msg MsgSvcWithdrawTax) (*sdk.Result, error) {
	err := k.WithdrawTax(ctx, msg.Trustee, msg.DestAddress, msg.Amount)
	if err != nil {
		return nil, err
	}

	return &sdk.Result{}, nil
}
