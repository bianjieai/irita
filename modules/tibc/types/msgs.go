package types

import (
	tibctypes "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibchost "github.com/bianjieai/tibc-go/modules/tibc/core/24-host"
	"github.com/bianjieai/tibc-go/modules/tibc/core/exported"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateClient    = "create"
	TypeMsgUpgradeClient   = "upgrade"
	TypeMsgRegisterRelayer = "register"
)

var (
	_ sdk.Msg = &MsgCreateClient{}
	_ sdk.Msg = &MsgUpgradeClient{}
	_ sdk.Msg = &MsgRegisterRelayer{}

	_ codectypes.UnpackInterfacesMessage = &MsgCreateClient{}
	_ codectypes.UnpackInterfacesMessage = &MsgUpgradeClient{}
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
	if err := tibchost.ClientIdentifierValidator(m.ChainName); err != nil {
		return err
	}

	clientState, err := tibctypes.UnpackClientState(m.ClientState)
	if err != nil {
		return err
	}
	return clientState.Validate()
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
	if err := tibchost.ClientIdentifierValidator(m.ChainName); err != nil {
		return err
	}

	clientState, err := tibctypes.UnpackClientState(m.ClientState)
	if err != nil {
		return err
	}
	return clientState.Validate()
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
	if err := tibchost.ClientIdentifierValidator(m.ChainName); err != nil {
		return err
	}

	if len(m.Relayers) == 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrNotFound, "string could not be parsed as address")
	}

	for _, relayer := range m.Relayers {
		if _, err := sdk.AccAddressFromBech32(relayer); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
		}
	}
	return nil
}

// GetSigners implements Msg.
func (m MsgRegisterRelayer) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Signer)
	return []sdk.AccAddress{addr}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (cup MsgCreateClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if err := unpacker.UnpackAny(cup.ClientState, new(exported.ClientState)); err != nil {
		return err
	}

	if err := unpacker.UnpackAny(cup.ConsensusState, new(exported.ConsensusState)); err != nil {
		return err
	}
	return nil
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (cup MsgUpgradeClient) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	if err := unpacker.UnpackAny(cup.ClientState, new(exported.ClientState)); err != nil {
		return err
	}

	if err := unpacker.UnpackAny(cup.ConsensusState, new(exported.ConsensusState)); err != nil {
		return err
	}
	return nil
}
