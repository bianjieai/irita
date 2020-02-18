package types

import (
	"regexp"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// types for msgs
	TypeMsgSvcDef           = "define_service"
	TypeMsgSvcBind          = "bind_service"
	TypeMsgSvcBindingUpdate = "update_service_binding"
	TypeMsgSvcDisable       = "disable_service"
	TypeMsgSvcEnable        = "enable_service"
	TypeMsgSvcRefundDeposit = "refund_service_deposit"
	TypeMsgSvcRequest       = "call_service"
	TypeMsgSvcResponse      = "respond_service"
	TypeMsgSvcRefundFees    = "refund_service_fees"
	TypeMsgSvcWithdrawFees  = "withdraw_service_fees"
	TypeMsgSvcWithdrawTax   = "withdraw_service_tax"

	MaxNameLength        = 70  // max length of the service name
	MaxChainIDLength     = 50  // max length of the chain ID
	MaxDescriptionLength = 280 // max length of the service and author description
	MaxTagCount          = 10  // max total number of the tags
	MaxTagLength         = 70  // max length of the tag
)

var (
	reSvcName = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`) // the service name only accepts alphanumeric characters, _ and -
)

var _, _, _, _, _, _, _, _, _, _, _ sdk.Msg = MsgSvcDef{}, MsgSvcBind{}, MsgSvcBindingUpdate{}, MsgSvcDisable{}, MsgSvcEnable{}, MsgSvcRefundDeposit{}, MsgSvcRequest{}, MsgSvcResponse{}, MsgSvcRefundFees{}, MsgSvcWithdrawFees{}, MsgSvcWithdrawTax{}

//______________________________________________________________________

// MsgSvcDef - struct for define a service
type MsgSvcDef struct {
	SvcDef
}

func NewMsgSvcDef(name, chainId, description string, tags []string, author sdk.AccAddress, authorDescription, idlContent string) MsgSvcDef {
	return MsgSvcDef{
		SvcDef{
			Name:              name,
			ChainId:           chainId,
			Description:       description,
			Tags:              tags,
			Author:            author,
			AuthorDescription: authorDescription,
			IDLContent:        idlContent,
		},
	}
}

func (msg MsgSvcDef) Route() string { return RouterKey }
func (msg MsgSvcDef) Type() string  { return TypeMsgSvcDef }

func (msg MsgSvcDef) GetSignBytes() []byte {
	if len(msg.Tags) == 0 {
		msg.Tags = nil
	}
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcDef) ValidateBasic() error {
	if len(msg.ChainId) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "chain-id missing")
	}

	if !validServiceName(msg.Name) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.Name)
	}

	if len(msg.Author) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "author missing")
	}

	if len(msg.IDLContent) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "IDL content missing")
	}

	if err := msg.EnsureLength(); err != nil {
		return err
	}

	if err := validateMethods(msg.IDLContent); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid IDL content, %s", err.Error())
	}

	return nil
}

func (msg MsgSvcDef) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Author}
}

//______________________________________________________________________

// MsgSvcBinding - struct for bind a service
type MsgSvcBind struct {
	DefName     string         `json:"def_name" yaml:"def_name"`
	DefChainID  string         `json:"def_chain_id" yaml:"def_chain_id"`
	BindChainID string         `json:"bind_chain_id" yaml:"bind_chain_id"`
	Provider    sdk.AccAddress `json:"provider" yaml:"provider"`
	BindingType BindingType    `json:"binding_type" yaml:"binding_type"`
	Deposit     sdk.Coins      `json:"deposit" yaml:"deposit"`
	Prices      []sdk.Coin     `json:"price" yaml:"price"`
	Level       Level          `json:"level" yaml:"level"`
}

func NewMsgSvcBind(defChainID, defName, bindChainID string, provider sdk.AccAddress, bindingType BindingType, deposit sdk.Coins, prices []sdk.Coin, level Level) MsgSvcBind {
	return MsgSvcBind{
		DefChainID:  defChainID,
		DefName:     defName,
		BindChainID: bindChainID,
		Provider:    provider,
		BindingType: bindingType,
		Deposit:     deposit,
		Prices:      prices,
		Level:       level,
	}
}

func (msg MsgSvcBind) Route() string { return RouterKey }
func (msg MsgSvcBind) Type() string  { return TypeMsgSvcBind }

func (msg MsgSvcBind) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcBind) ValidateBasic() error {
	if len(msg.DefChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "define chain-id missing")
	}

	if len(msg.BindChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bind chain-id missing")
	}

	if err := ensureChainIdLength(msg.DefChainID, "def_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.BindChainID, "bind_chain_id"); err != nil {
		return err
	}

	if !validServiceName(msg.DefName) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.DefName)
	}

	if err := ensureNameLength(msg.DefName); err != nil {
		return err
	}

	if !validBindingType(msg.BindingType) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid binding type: %s", msg.BindingType)
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	if !validServiceCoins(msg.Deposit) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid service deposit: %s", msg.Deposit.String())
	}

	for _, price := range msg.Prices {
		if !validServiceCoins(sdk.Coins{price}) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid service price: %s", price.String())
		}
	}

	if !validLevel(msg.Level) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid level; AvgRspTime: %d, UsableTime: %d", msg.Level.AvgRspTime, msg.Level.UsableTime)
	}

	return nil
}

func (msg MsgSvcBind) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcBindingUpdate - struct for update a service binding
type MsgSvcBindingUpdate struct {
	DefName     string         `json:"def_name" yaml:"def_name"`
	DefChainID  string         `json:"def_chain_id" yaml:"def_chain_id"`
	BindChainID string         `json:"bind_chain_id" yaml:"bind_chain_id"`
	Provider    sdk.AccAddress `json:"provider" yaml:"provider"`
	BindingType BindingType    `json:"binding_type" yaml:"binding_type"`
	Deposit     sdk.Coins      `json:"deposit" yaml:"deposit"`
	Prices      []sdk.Coin     `json:"price" yaml:"price"`
	Level       Level          `json:"level" yaml:"level"`
}

func NewMsgSvcBindingUpdate(defChainID, defName, bindChainID string, provider sdk.AccAddress, bindingType BindingType, deposit sdk.Coins, prices []sdk.Coin, level Level) MsgSvcBindingUpdate {
	return MsgSvcBindingUpdate{
		DefChainID:  defChainID,
		DefName:     defName,
		BindChainID: bindChainID,
		Provider:    provider,
		BindingType: bindingType,
		Deposit:     deposit,
		Prices:      prices,
		Level:       level,
	}
}
func (msg MsgSvcBindingUpdate) Route() string { return RouterKey }
func (msg MsgSvcBindingUpdate) Type() string  { return TypeMsgSvcBindingUpdate }

func (msg MsgSvcBindingUpdate) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcBindingUpdate) ValidateBasic() error {
	if len(msg.DefChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "define chain-id missing")
	}

	if len(msg.BindChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bind chain-id missing")
	}

	if err := ensureChainIdLength(msg.DefChainID, "def_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.BindChainID, "bind_chain_id"); err != nil {
		return err
	}

	if !validServiceName(msg.DefName) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.DefName)
	}

	if err := ensureNameLength(msg.DefName); err != nil {
		return err
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	if msg.BindingType != 0x00 && !validBindingType(msg.BindingType) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid binding type: %s", msg.BindingType.String())
	}

	if !validServiceCoins(msg.Deposit) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid service deposit: %s", msg.Deposit.String())
	}

	for _, price := range msg.Prices {
		if !validServiceCoins(sdk.Coins{price}) {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid service price: %s", price.String())
		}
	}

	if !validUpdateLevel(msg.Level) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid level; AvgRspTime: %d, UsableTime: %d", msg.Level.AvgRspTime, msg.Level.UsableTime)
	}

	return nil
}

func (msg MsgSvcBindingUpdate) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcDisable - struct for disable a service binding
type MsgSvcDisable struct {
	DefName     string         `json:"def_name" yaml:"def_name"`
	DefChainID  string         `json:"def_chain_id" yaml:"def_chain_id"`
	BindChainID string         `json:"bind_chain_id" yaml:"bind_chain_id"`
	Provider    sdk.AccAddress `json:"provider" yaml:"provider"`
}

func NewMsgSvcDisable(defChainID, defName, bindChainID string, provider sdk.AccAddress) MsgSvcDisable {
	return MsgSvcDisable{
		DefChainID:  defChainID,
		DefName:     defName,
		BindChainID: bindChainID,
		Provider:    provider,
	}
}

func (msg MsgSvcDisable) Route() string { return RouterKey }
func (msg MsgSvcDisable) Type() string  { return TypeMsgSvcDisable }

func (msg MsgSvcDisable) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcDisable) ValidateBasic() error {
	if len(msg.DefChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "define chain-id missing")
	}

	if len(msg.BindChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bind chain-id missing")
	}

	if err := ensureChainIdLength(msg.DefChainID, "def_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.BindChainID, "bind_chain_id"); err != nil {
		return err
	}

	if !validServiceName(msg.DefName) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.DefName)
	}

	if err := ensureNameLength(msg.DefName); err != nil {
		return err
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	return nil
}

func (msg MsgSvcDisable) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcEnable - struct for enable a service binding
type MsgSvcEnable struct {
	DefName     string         `json:"def_name" yaml:"def_name"`
	DefChainID  string         `json:"def_chain_id" yaml:"def_chain_id"`
	BindChainID string         `json:"bind_chain_id" yaml:"bind_chain_id"`
	Provider    sdk.AccAddress `json:"provider" yaml:"provider"`
	Deposit     sdk.Coins      `json:"deposit" yaml:"deposit"`
}

func NewMsgSvcEnable(defChainID, defName, bindChainID string, provider sdk.AccAddress, deposit sdk.Coins) MsgSvcEnable {
	return MsgSvcEnable{
		DefChainID:  defChainID,
		DefName:     defName,
		BindChainID: bindChainID,
		Provider:    provider,
		Deposit:     deposit,
	}
}

func (msg MsgSvcEnable) Route() string { return RouterKey }
func (msg MsgSvcEnable) Type() string  { return TypeMsgSvcEnable }

func (msg MsgSvcEnable) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcEnable) ValidateBasic() error {
	if len(msg.DefChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "define chain-id missing")
	}

	if len(msg.BindChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bind chain-id missing")
	}

	if err := ensureChainIdLength(msg.DefChainID, "def_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.BindChainID, "bind_chain_id"); err != nil {
		return err
	}

	if !validServiceName(msg.DefName) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.DefName)
	}

	if err := ensureNameLength(msg.DefName); err != nil {
		return err
	}

	if !validServiceCoins(msg.Deposit) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid service deposit: %s", msg.Deposit.String())
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	return nil
}

func (msg MsgSvcEnable) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcRefundDeposit - struct for refund deposit from a service binding
type MsgSvcRefundDeposit struct {
	DefName     string         `json:"def_name" yaml:"def_name"`
	DefChainID  string         `json:"def_chain_id" yaml:"def_chain_id"`
	BindChainID string         `json:"bind_chain_id" yaml:"bind_chain_id"`
	Provider    sdk.AccAddress `json:"provider" yaml:"provider"`
}

func NewMsgSvcRefundDeposit(defChainID, defName, bindChainID string, provider sdk.AccAddress) MsgSvcRefundDeposit {
	return MsgSvcRefundDeposit{
		DefChainID:  defChainID,
		DefName:     defName,
		BindChainID: bindChainID,
		Provider:    provider,
	}
}

func (msg MsgSvcRefundDeposit) Route() string { return RouterKey }
func (msg MsgSvcRefundDeposit) Type() string  { return TypeMsgSvcRefundDeposit }

func (msg MsgSvcRefundDeposit) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcRefundDeposit) ValidateBasic() error {
	if len(msg.DefChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "define chain-id missing")
	}

	if len(msg.BindChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bind chain-id missing")
	}

	if err := ensureChainIdLength(msg.DefChainID, "def_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.BindChainID, "bind_chain_id"); err != nil {
		return err
	}

	if !validServiceName(msg.DefName) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.DefName)
	}

	if err := ensureNameLength(msg.DefName); err != nil {
		return err
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	return nil
}

func (msg MsgSvcRefundDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcRequest - struct for call a service
type MsgSvcRequest struct {
	DefChainID  string         `json:"def_chain_id" yaml:"def_chain_id"`
	DefName     string         `json:"def_name" yaml:"def_name"`
	BindChainID string         `json:"bind_chain_id" yaml:"bind_chain_id"`
	ReqChainID  string         `json:"req_chain_id" yaml:"req_chain_id"`
	MethodID    int16          `json:"method_id" yaml:"method_id"`
	Provider    sdk.AccAddress `json:"provider" yaml:"provider"`
	Consumer    sdk.AccAddress `json:"consumer" yaml:"consumer"`
	Input       []byte         `json:"input" yaml:"input"`
	ServiceFee  sdk.Coins      `json:"service_fee" yaml:"service_fee"`
	Profiling   bool           `json:"profiling" yaml:"profiling"`
}

func NewMsgSvcRequest(defChainID, defName, bindChainID, reqChainID string, consumer, provider sdk.AccAddress, methodID int16, input []byte, serviceFee sdk.Coins, profiling bool) MsgSvcRequest {
	return MsgSvcRequest{
		DefChainID:  defChainID,
		DefName:     defName,
		BindChainID: bindChainID,
		ReqChainID:  reqChainID,
		Consumer:    consumer,
		Provider:    provider,
		MethodID:    methodID,
		Input:       input,
		ServiceFee:  serviceFee,
		Profiling:   profiling,
	}
}

func (msg MsgSvcRequest) Route() string { return RouterKey }
func (msg MsgSvcRequest) Type() string  { return TypeMsgSvcRequest }

func (msg MsgSvcRequest) GetSignBytes() []byte {
	if len(msg.Input) == 0 {
		msg.Input = nil
	}
	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcRequest) ValidateBasic() error {
	if len(msg.DefChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "define chain-id missing")
	}

	if len(msg.BindChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bind chain-id missing")
	}

	if len(msg.ReqChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "request chain-id missing")
	}

	if err := ensureChainIdLength(msg.DefChainID, "def_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.BindChainID, "bind_chain_id"); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.ReqChainID, "req_chain_id"); err != nil {
		return err
	}

	if !validServiceName(msg.DefName) {
		return sdkerrors.Wrap(ErrInvalidServiceName, msg.DefName)
	}

	if err := ensureNameLength(msg.DefName); err != nil {
		return err
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	if len(msg.Consumer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "consumer missing")
	}

	if !validServiceCoins(msg.ServiceFee) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "invalid service fee: %s", msg.ServiceFee.String())
	}

	return nil
}

func (msg MsgSvcRequest) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

//______________________________________________________________________

// MsgSvcResponse - struct for respond a service call
type MsgSvcResponse struct {
	ReqChainID string         `json:"req_chain_id" yaml:"req_chain_id"`
	RequestID  string         `json:"request_id" yaml:"request_id"`
	Provider   sdk.AccAddress `json:"provider" yaml:"provider"`
	Output     []byte         `json:"output" yaml:"output"`
	ErrorMsg   []byte         `json:"error_msg" yaml:"error_msg"`
}

func NewMsgSvcResponse(reqChainID string, requestId string, provider sdk.AccAddress, output, errorMsg []byte) MsgSvcResponse {
	return MsgSvcResponse{
		ReqChainID: reqChainID,
		RequestID:  requestId,
		Provider:   provider,
		Output:     output,
		ErrorMsg:   errorMsg,
	}
}

func (msg MsgSvcResponse) Route() string { return RouterKey }
func (msg MsgSvcResponse) Type() string  { return TypeMsgSvcResponse }

func (msg MsgSvcResponse) GetSignBytes() []byte {
	if len(msg.Output) == 0 {
		msg.Output = nil
	}

	if len(msg.ErrorMsg) == 0 {
		msg.ErrorMsg = nil
	}

	b, err := ModuleCdc.MarshalJSON(msg)
	if err != nil {
		panic(err)
	}

	return sdk.MustSortJSON(b)
}

func (msg MsgSvcResponse) ValidateBasic() error {
	if len(msg.ReqChainID) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "request chain-id missing")
	}

	if err := ensureChainIdLength(msg.ReqChainID, "req_chain_id"); err != nil {
		return err
	}

	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider missing")
	}

	_, _, _, err := ConvertRequestID(msg.RequestID)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid request id: %s", msg.RequestID)
	}

	return nil
}

func (msg MsgSvcResponse) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcRefundFees - struct for refund fees
type MsgSvcRefundFees struct {
	Consumer sdk.AccAddress `json:"consumer" yaml:"consumer"`
}

func NewMsgSvcRefundFees(consumer sdk.AccAddress) MsgSvcRefundFees {
	return MsgSvcRefundFees{
		Consumer: consumer,
	}
}

func (msg MsgSvcRefundFees) Route() string { return RouterKey }
func (msg MsgSvcRefundFees) Type() string  { return TypeMsgSvcRefundFees }

func (msg MsgSvcRefundFees) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcRefundFees) ValidateBasic() error {
	if len(msg.Consumer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "consumer address missing")
	}

	return nil
}

func (msg MsgSvcRefundFees) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Consumer}
}

//______________________________________________________________________

// MsgSvcWithdrawFees - struct for withdraw fees
type MsgSvcWithdrawFees struct {
	Provider sdk.AccAddress `json:"provider" yaml:"provider"`
}

func NewMsgSvcWithdrawFees(provider sdk.AccAddress) MsgSvcWithdrawFees {
	return MsgSvcWithdrawFees{
		Provider: provider,
	}
}

func (msg MsgSvcWithdrawFees) Route() string { return RouterKey }
func (msg MsgSvcWithdrawFees) Type() string  { return TypeMsgSvcWithdrawFees }

func (msg MsgSvcWithdrawFees) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcWithdrawFees) ValidateBasic() error {
	if len(msg.Provider) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "provider address missing")
	}

	return nil
}

func (msg MsgSvcWithdrawFees) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Provider}
}

//______________________________________________________________________

// MsgSvcWithdrawTax - struct for withdraw tax
type MsgSvcWithdrawTax struct {
	Trustee     sdk.AccAddress `json:"trustee" yaml:"trustee"`
	DestAddress sdk.AccAddress `json:"dest_address" yaml:"dest_address"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
}

func NewMsgSvcWithdrawTax(trustee, destAddress sdk.AccAddress, amount sdk.Coins) MsgSvcWithdrawTax {
	return MsgSvcWithdrawTax{
		Trustee:     trustee,
		DestAddress: destAddress,
		Amount:      amount,
	}
}

func (msg MsgSvcWithdrawTax) Route() string { return RouterKey }
func (msg MsgSvcWithdrawTax) Type() string  { return TypeMsgSvcWithdrawTax }

func (msg MsgSvcWithdrawTax) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(b)
}

func (msg MsgSvcWithdrawTax) ValidateBasic() error {
	if len(msg.Trustee) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "trustee address missing")
	}

	if len(msg.DestAddress) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "destination address missing")
	}

	if !validServiceCoins(msg.Amount) {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}

	return nil
}

func (msg MsgSvcWithdrawTax) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Trustee}
}

//______________________________________________________________________

func validServiceName(name string) bool {
	return reSvcName.MatchString(name)
}

func (msg MsgSvcDef) EnsureLength() error {
	if err := ensureNameLength(msg.Name); err != nil {
		return err
	}

	if err := ensureChainIdLength(msg.ChainId, "chain_id"); err != nil {
		return err
	}

	if len(msg.Description) > MaxDescriptionLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid description length; got: %d, max: %d", len(msg.Description), MaxDescriptionLength)
	}

	if len(msg.Tags) > MaxTagCount {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tags count; got: %d, max: %d", len(msg.Tags), MaxTagCount)
	} else {
		for i, tag := range msg.Tags {
			if len(tag) > MaxTagLength {
				return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid tag[%d] length; got: %d, max: %d", i, len(tag), MaxTagLength)
			}
		}
	}

	if len(msg.AuthorDescription) > MaxDescriptionLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid author description length; got: %d, max: %d", len(msg.AuthorDescription), MaxDescriptionLength)
	}

	return nil
}

func ensureNameLength(name string) error {
	if len(name) > MaxNameLength {
		return sdkerrors.Wrapf(ErrInvalidServiceName, "invalid service name length; got: %d, max: %d", len(name), MaxNameLength)
	}

	return nil
}

func ensureChainIdLength(chainId, fieldNm string) error {
	if len(chainId) > MaxChainIDLength {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid chain-id length; got: %d, max: %d", len(chainId), MaxChainIDLength)
	}

	return nil
}

func validServiceCoins(coins sdk.Coins) bool {
	if coins == nil || len(coins) != 1 {
		return false
	}

	return coins[0].IsPositive()
}
