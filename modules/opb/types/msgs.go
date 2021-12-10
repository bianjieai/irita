package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"
)

const (
	TypeMsgMint    = "mint"    // type for MsgMint
	TypeMsgReclaim = "reclaim" // type for MsgReclaim
)

var (
	_ sdk.Msg = &MsgMint{}
	_ sdk.Msg = &MsgReclaim{}
)

// NewMsgMint creates a new MsgMint instance.
func NewMsgMint(amount uint64, recipient sdk.AccAddress, operator sdk.AccAddress) *MsgMint {
	return &MsgMint{
		Amount:    amount,
		Recipient: recipient.String(),
		Operator:  operator.String(),
	}
}

// Route implements Msg.
func (m MsgMint) Route() string {
	return RouterKey
}

// Type implements Msg.
func (m MsgMint) Type() string {
	return TypeMsgMint
}

// ValidateBasic implements Msg.
func (m MsgMint) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid operator")
	}

	if m.Amount == 0 {
		return sdkerrors.Wrap(ErrInvalidAmount, "amount should be greater than 0")
	}

	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient %s: %s", m.Recipient, err)
	}

	return nil
}

// GetSignBytes implements Msg.
func (m MsgMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg.
func (m MsgMint) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

// NewMsgReclaim creates a new MsgReclaim instance.
func NewMsgReclaim(denom string, recipient sdk.AccAddress, operator sdk.AccAddress) *MsgReclaim {
	return &MsgReclaim{
		Denom:     denom,
		Recipient: recipient.String(),
		Operator:  operator.String(),
	}
}

// Route implements Msg.
func (m MsgReclaim) Route() string {
	return RouterKey
}

// Type implements Msg.
func (m MsgReclaim) Type() string {
	return TypeMsgReclaim
}

// ValidateBasic implements Msg.
func (m MsgReclaim) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Operator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid operator %s: %s", m.Operator, err)
	}

	if len(m.Denom) == 0 {
		return sdkerrors.Wrapf(ErrInvalidDenom, "denom missing")
	}

	if err := sdk.ValidateDenom(m.Denom); err != nil {
		return sdkerrors.Wrapf(ErrInvalidDenom, "invalid denom %s: %s", m.Denom, err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient %s: %s", m.Recipient, err)
	}

	return nil
}

// GetSignBytes implements Msg.
func (m MsgReclaim) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg.
func (m MsgReclaim) GetSigners() []sdk.AccAddress {
	addr, _ := sdk.AccAddressFromBech32(m.Operator)
	return []sdk.AccAddress{addr}
}

var (
	_ sdk.Msg = &MsgAddToContractDenyList{}
	_ sdk.Msg = &MsgRemoveFromContractDenyList{}
)

func NewMsgAddToContractDenyList(contractAddr, from string) *MsgAddToContractDenyList {
	return &MsgAddToContractDenyList{
		contractAddr,
		from,
	}
}

func (m MsgAddToContractDenyList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	if !common.IsHexAddress(m.ContractAddress) {
		return sdkerrors.Wrap(ErrInvalidContractAddress, "invalid from address")
	}
	return nil
}

func (m *MsgAddToContractDenyList) GetSigners() []sdk.AccAddress {
	if len(m.From) == 0 {
		panic("do not have signer")
	}
	accAddr, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func NewMsgRemoveFromContractDenyList(contractAddr, from string) *MsgRemoveFromContractDenyList {
	return &MsgRemoveFromContractDenyList{
		contractAddr,
		from,
	}
}

func (m MsgRemoveFromContractDenyList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	if !common.IsHexAddress(m.ContractAddress) {
		return sdkerrors.Wrap(ErrInvalidContractAddress, "invalid from address")
	}
	return nil
}

func (m *MsgRemoveFromContractDenyList) GetSigners() []sdk.AccAddress {
	if len(m.From) == 0 {
		panic("do not have signer")
	}
	accAddr, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

var (
	_ sdk.Msg = &MsgAddToAccountDenyList{}
	_ sdk.Msg = &MsgRemoveFromAccountDenyList{}
)

func NewMsgAddToAccountDenyList(contractAddr, from string) *MsgAddToAccountDenyList {
	return &MsgAddToAccountDenyList{
		contractAddr,
		from,
	}
}

func (m MsgAddToAccountDenyList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}

	_, err = sdk.AccAddressFromBech32(m.AccountAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	return nil
}

func (m *MsgAddToAccountDenyList) GetSigners() []sdk.AccAddress {
	if len(m.From) == 0 {
		panic("do not have signer")
	}
	accAddr, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func NewMsgRemoveFromAccountDenyList(contractAddr, from string) *MsgRemoveFromAccountDenyList {
	return &MsgRemoveFromAccountDenyList{
		contractAddr,
		from,
	}
}

func (m MsgRemoveFromAccountDenyList) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	_, err = sdk.AccAddressFromBech32(m.AccountAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "string could not be parsed as address: %v", err)
	}
	return nil
}

func (m *MsgRemoveFromAccountDenyList) GetSigners() []sdk.AccAddress {
	if len(m.From) == 0 {
		panic("do not have signer")
	}
	accAddr, err := sdk.AccAddressFromBech32(m.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
