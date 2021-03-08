package types

import (
	"errors"
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultBaseTokenDenom is the denom of the base native token by default
	DefaultBaseTokenDenom = "uirita"

	// DefaultPointTokenDenom is the denom of the native point token by default
	DefaultPointTokenDenom = "upoint"

	// DefaultUnrestrictedTokenTransfer is set to false, which
	// means that the token transfer is under certain constraint
	DefaultUnrestrictedTokenTransfer = true
)

// Parameter store keys
var (
	KeyBaseTokenDenom            = []byte("BaseTokenDenom")
	KeyPointTokenDenom           = []byte("PointTokenDenom")
	KeyBaseTokenManager          = []byte("BaseTokenManager")
	KeyUnrestrictedTokenTransfer = []byte("UnrestrictedTokenTransfer")
)

// NewParams creates a new Params instance
func NewParams(
	baseTokenDenom string,
	pointTokenDenom string,
	baseTokenManager string,
	unrestrictedTokenTransfer bool,
) Params {
	return Params{
		BaseTokenDenom:            baseTokenDenom,
		PointTokenDenom:           pointTokenDenom,
		BaseTokenManager:          baseTokenManager,
		UnrestrictedTokenTransfer: unrestrictedTokenTransfer,
	}
}

// DefaultParams returns the module default parameters
func DefaultParams() Params {
	return NewParams(
		DefaultBaseTokenDenom,
		DefaultPointTokenDenom,
		"",
		DefaultUnrestrictedTokenTransfer,
	)
}

// Validate validates the params
func (p Params) Validate() error {
	if err := validateBaseTokenDenom(p.BaseTokenDenom); err != nil {
		return err
	}

	if err := validatePointTokenDenom(p.PointTokenDenom); err != nil {
		return err
	}

	if err := validateBaseTokenManager(p.BaseTokenManager); err != nil {
		return err
	}

	if err := validateUnrestrictedTokenTransfer(p.UnrestrictedTokenTransfer); err != nil {
		return err
	}

	return nil
}

// String implements Stringer.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyBaseTokenDenom, &p.BaseTokenDenom, validateBaseTokenDenom),
		paramtypes.NewParamSetPair(KeyPointTokenDenom, &p.PointTokenDenom, validatePointTokenDenom),
		paramtypes.NewParamSetPair(KeyBaseTokenManager, &p.BaseTokenManager, validateBaseTokenManager),
		paramtypes.NewParamSetPair(KeyUnrestrictedTokenTransfer, &p.UnrestrictedTokenTransfer, validateUnrestrictedTokenTransfer),
	}
}

func validateBaseTokenDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("base token denom can not be empty")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validatePointTokenDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("point token denom can not be empty")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}

func validateBaseTokenManager(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if len(v) > 0 {
		if _, err := sdk.AccAddressFromBech32(v); err != nil {
			return fmt.Errorf("invalid bech32 address %s: %s", v, err)
		}
	}

	return nil
}

func validateUnrestrictedTokenTransfer(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
