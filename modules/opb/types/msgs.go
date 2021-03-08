package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
