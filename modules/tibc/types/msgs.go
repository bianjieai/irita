package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	tibctypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	"github.com/bianjieai/tibc-go/modules/tibc/core/exported"
)

const (
	TypeMsgCreateClient    = "create"  // type for MsgMint
	TypeMsgUpgradeClient   = "upgrade" // type for MsgReclaim
	TypeMsgRegisterRelayer = "registe"
)

var (
	_ sdk.Msg = &MsgCreateClient{}
	_ sdk.Msg = &MsgUpgradeClient{}
	_ sdk.Msg = &MsgRegisterRelayer{}
)

func NewMsgCreateClient(chainName string, clientState exported.ClientState, consensusState exported.ConsensusState, signer sdk.AccAddress) (*MsgCreateClient, error) {
	clientStateAny, err := tibctypes.PackClientState(clientState)
	if err != nil {
		return nil, err
	}

	consensusStateAny, err := tibctypes.PackConsensusState(consensusState)
	if err != nil {
		return nil, err
	}

	return &MsgCreateClient{
		ChainName:      chainName,
		ClientState:    clientStateAny,
		ConsensusState: consensusStateAny,
		Signer:         signer.String(),
	}, nil
}

// Route implements Msg.
func (m MsgCreateClient) Route() string {
	return RouterKey
}

// Type implements Msg.
func (m MsgCreateClient) Type() string {
	return TypeMsgCreateClient
}

// ValidateBasic implements Msg.
func (m MsgCreateClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	// todo
	return nil
}

// GetSigners implements Msg.
func (m MsgCreateClient) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Signer)
	return []sdk.AccAddress{addr}
}

func NewMsgUpgradeClient(chainName string, clientState exported.ClientState, consensusState exported.ConsensusState, signer sdk.AccAddress) (*MsgUpgradeClient, error) {
	clientStateAny, err := tibctypes.PackClientState(clientState)
	if err != nil {
		return nil, err
	}

	consensusStateAny, err := tibctypes.PackConsensusState(consensusState)
	if err != nil {
		return nil, err
	}

	return &MsgUpgradeClient{
		ChainName:      chainName,
		ClientState:    clientStateAny,
		ConsensusState: consensusStateAny,
		Signer:         signer.String(),
	}, nil
}

// Route implements Msg.
func (m MsgUpgradeClient) Route() string {
	return RouterKey
}

// Type implements Msg.
func (m MsgUpgradeClient) Type() string {
	return TypeMsgUpgradeClient
}

// ValidateBasic implements Msg.
func (m MsgUpgradeClient) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	// todo
	return nil
}

// GetSigners implements Msg.
func (m MsgUpgradeClient) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Signer)
	return []sdk.AccAddress{addr}
}

func NewMsgRegisterRelayer(chainName string, relayers []string, signer sdk.AccAddress) (*MsgRegisterRelayer, error) {
	return &MsgRegisterRelayer{
		ChainName: chainName,
		Relayers:  relayers,
		Signer:    signer.String(),
	}, nil
}

// Route implements Msg.
func (m MsgRegisterRelayer) Route() string {
	return RouterKey
}

// Type implements Msg.
func (m MsgRegisterRelayer) Type() string {
	return TypeMsgRegisterRelayer
}

// ValidateBasic implements Msg.
func (m MsgRegisterRelayer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	// todo
	return nil
}

// GetSigners implements Msg.
func (m MsgRegisterRelayer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Signer)
	return []sdk.AccAddress{addr}
}
