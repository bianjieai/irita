package keeper

import (
	"encoding/hex"
	"encoding/json"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	irismodtypes "github.com/irisnet/irismod/modules/service/types"

	"github.com/bianjieai/irita/modules/wservice/types"
)

type DeduplicationTxDecorator struct {
	keeper IKeeper
}

func NewDeduplicationTxDecorator(keeper IKeeper) DeduplicationTxDecorator {
	return DeduplicationTxDecorator{keeper: keeper}
}

func (decorator DeduplicationTxDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *irismodtypes.MsgCallService:
			err = decorator.validateCallServiceReqSequence(ctx, msg)
			if err != nil {
				return ctx, err
			}

		case *irismodtypes.MsgRespondService:
			err = decorator.validateRespondServiceReqSequence(ctx, msg)
			if err != nil {
				return ctx, err
			}
		}
	}
	return next(ctx, tx, simulate)
}

func (decorator *DeduplicationTxDecorator) validateCallServiceReqSequence(ctx sdk.Context, msg *irismodtypes.MsgCallService) error {

	serviceKeeper := decorator.keeper.GetServiceKeeper()

	_, err := sdk.AccAddressFromBech32(msg.Consumer)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrInvalidConsumerAddr,
			"invalid consumer addr: %s ",
			msg.Consumer,
		)
	}

	input := &types.Input{}
	if err := json.Unmarshal([]byte(msg.Input), input); err != nil {
		return sdkerrors.Wrapf(
			types.ErrInvalidInput,
			"input unmarshal error ",
		)
	}
	if len(input.Header.ReqSequence) == 0 || len(input.Header.ID) == 0 {
		return nil
	}

	// 当是BSN特定需求的消息时 判断Providers 是否等于1
	if len(msg.Providers) != 1 {
		return sdkerrors.Wrapf(
			types.ErrInvalidProviderLength,
			"number of providers can only be 1 in de-duplication mode",
		)
	}

	// 获取全局timeout
	timeout := serviceKeeper.MaxRequestTimeout(ctx)

	suffix := input.Header.ID + "_" + input.Header.ReqSequence
	key := types.ReqSequenceInputPrefix + suffix
	// example: wservice_input_test01_123

	if decorator.keeper.ExistReqSequence(ctx, []byte(key)) {
		return sdkerrors.Wrapf(
			types.ErrDuplicatedRequestSequence,
			"duplicated request sequence: %s",
			input.Header.ReqSequence,
		)
	}

	decorator.keeper.SetReqSequence(ctx, []byte(key), []byte{1})

	// height key
	heightKey := strconv.Itoa(int(ctx.BlockHeight()+timeout)) + "-" + key
	// example: 51-wservice_input_test01_123
	decorator.keeper.SetReqSequence(ctx, []byte(heightKey), []byte{1})
	return nil
}

func (decorator *DeduplicationTxDecorator) validateRespondServiceReqSequence(ctx sdk.Context, msg *irismodtypes.MsgRespondService) error {

	// de
	serviceKeeper := decorator.keeper.GetServiceKeeper()
	requestID, _ := hex.DecodeString(msg.RequestId)
	requestResp, isFound := serviceKeeper.GetRequest(ctx, requestID)
	if !isFound {
		return sdkerrors.Wrapf(
			types.ErrRequestNotFound,
			"request not found: %s",
			requestID,
		)
	}

	// 解析callService的参数
	input := &types.Input{}
	if err := json.Unmarshal([]byte(requestResp.Input), input); err != nil {
		return sdkerrors.Wrapf(
			types.ErrInvalidInput,
			"input unmarshal error ",
		)
	}

	// 如果input没有包含 去重的参数，则直接返回
	if len(input.Header.ReqSequence) == 0 || len(input.Header.ID) == 0 {
		return nil
	}

	output := &types.Output{}
	if err := json.Unmarshal([]byte(msg.Output), output); err != nil {
		return sdkerrors.Wrapf(
			types.ErrInvalidOutput,
			"output unmarshal error ",
		)
	}

	// 如果input中有去重参数，但output没有，返回错误
	if len(output.Header.ReqSequence) == 0 || len(output.Header.ID) == 0 {
		return sdkerrors.Wrapf(
			types.ErrInvalidOutput,
			"output params error ",
		)
	}

	providerAddr, err := sdk.AccAddressFromBech32(msg.Provider)
	if err != nil {
		return sdkerrors.Wrapf(
			types.ErrInvalidProviderAddr,
			"invalid provider addr: %s",
			msg.Provider,
		)
	}

	// 断言err 肯定为空
	requestProvider, _ := sdk.AccAddressFromBech32(requestResp.Provider)
	if !providerAddr.Equals(requestProvider) {
		return sdkerrors.Wrapf(
			types.ErrInvalidProvider,
			"invalid provider: %s",
			msg.Provider,
		)
	}

	// get global timeout
	timeout := serviceKeeper.MaxRequestTimeout(ctx)

	outputSuffix := output.Header.ID + "_" + output.Header.ReqSequence
	outputKey := types.ReqSequenceOutputPrefix + outputSuffix
	// example: wservice_output_test01_123

	// if output key exits, completed!
	if decorator.keeper.ExistReqSequence(ctx, []byte(outputKey)) {
		return sdkerrors.Wrapf(
			types.ErrDuplicatedRequestSequence,
			"duplicated request sequence: %s",
			output.Header.ReqSequence,
		)
	}

	decorator.keeper.SetReqSequence(ctx, []byte(outputKey), []byte{1})

	// height key
	heightKey := strconv.Itoa(int(ctx.BlockHeight()+timeout)) + "-" + outputKey
	// example: 52-wservice_output_test01_123
	decorator.keeper.SetReqSequence(ctx, []byte(heightKey), []byte{1})
	return nil
}
