// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VRFCoordinatorRequestCommitment is an auto generated low-level Go binding around an user-defined struct.
type VRFCoordinatorRequestCommitment struct {
	BlockNum         uint64
	SubId            uint64
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
}

// VRFProof is an auto generated low-level Go binding around an user-defined struct.
type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

// VrfMetaData contains all meta data concerning the Vrf contract.
var VrfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"internalType\":\"structVRFCoordinator.RequestCommitment\",\"name\":\"rc\",\"type\":\"tuple\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"}],\"name\":\"getBlockHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"getCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"storeBlockHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VrfABI is the input ABI used to generate the binding from.
// Deprecated: Use VrfMetaData.ABI instead.
var VrfABI = VrfMetaData.ABI

// Vrf is an auto generated Go binding around an Ethereum contract.
type Vrf struct {
	VrfCaller     // Read-only binding to the contract
	VrfTransactor // Write-only binding to the contract
	VrfFilterer   // Log filterer for contract events
}

// VrfCaller is an auto generated read-only Go binding around an Ethereum contract.
type VrfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VrfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VrfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VrfSession struct {
	Contract     *Vrf              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VrfCallerSession struct {
	Contract *VrfCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VrfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VrfTransactorSession struct {
	Contract     *VrfTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfRaw is an auto generated low-level Go binding around an Ethereum contract.
type VrfRaw struct {
	Contract *Vrf // Generic contract binding to access the raw methods on
}

// VrfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VrfCallerRaw struct {
	Contract *VrfCaller // Generic read-only contract binding to access the raw methods on
}

// VrfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VrfTransactorRaw struct {
	Contract *VrfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVrf creates a new instance of Vrf, bound to a specific deployed contract.
func NewVrf(address common.Address, backend bind.ContractBackend) (*Vrf, error) {
	contract, err := bindVrf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vrf{VrfCaller: VrfCaller{contract: contract}, VrfTransactor: VrfTransactor{contract: contract}, VrfFilterer: VrfFilterer{contract: contract}}, nil
}

// NewVrfCaller creates a new read-only instance of Vrf, bound to a specific deployed contract.
func NewVrfCaller(address common.Address, caller bind.ContractCaller) (*VrfCaller, error) {
	contract, err := bindVrf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VrfCaller{contract: contract}, nil
}

// NewVrfTransactor creates a new write-only instance of Vrf, bound to a specific deployed contract.
func NewVrfTransactor(address common.Address, transactor bind.ContractTransactor) (*VrfTransactor, error) {
	contract, err := bindVrf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VrfTransactor{contract: contract}, nil
}

// NewVrfFilterer creates a new log filterer instance of Vrf, bound to a specific deployed contract.
func NewVrfFilterer(address common.Address, filterer bind.ContractFilterer) (*VrfFilterer, error) {
	contract, err := bindVrf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VrfFilterer{contract: contract}, nil
}

// bindVrf binds a generic wrapper to an already deployed contract.
func bindVrf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VrfABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrf *VrfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vrf.Contract.VrfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrf *VrfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.Contract.VrfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrf *VrfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrf.Contract.VrfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vrf *VrfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vrf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vrf *VrfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vrf *VrfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vrf.Contract.contract.Transact(opts, method, params...)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_Vrf *VrfCaller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_Vrf *VrfSession) MAXCONSUMERS() (uint16, error) {
	return _Vrf.Contract.MAXCONSUMERS(&_Vrf.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_Vrf *VrfCallerSession) MAXCONSUMERS() (uint16, error) {
	return _Vrf.Contract.MAXCONSUMERS(&_Vrf.CallOpts)
}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_Vrf *VrfCaller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_Vrf *VrfSession) MAXNUMWORDS() (uint32, error) {
	return _Vrf.Contract.MAXNUMWORDS(&_Vrf.CallOpts)
}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_Vrf *VrfCallerSession) MAXNUMWORDS() (uint32, error) {
	return _Vrf.Contract.MAXNUMWORDS(&_Vrf.CallOpts)
}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_Vrf *VrfCaller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_Vrf *VrfSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _Vrf.Contract.MAXREQUESTCONFIRMATIONS(&_Vrf.CallOpts)
}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_Vrf *VrfCallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _Vrf.Contract.MAXREQUESTCONFIRMATIONS(&_Vrf.CallOpts)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 height) view returns(bytes32)
func (_Vrf *VrfCaller) GetBlockHash(opts *bind.CallOpts, height *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "getBlockHash", height)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 height) view returns(bytes32)
func (_Vrf *VrfSession) GetBlockHash(height *big.Int) ([32]byte, error) {
	return _Vrf.Contract.GetBlockHash(&_Vrf.CallOpts, height)
}

// GetBlockHash is a free data retrieval call binding the contract method 0xee82ac5e.
//
// Solidity: function getBlockHash(uint256 height) view returns(bytes32)
func (_Vrf *VrfCallerSession) GetBlockHash(height *big.Int) ([32]byte, error) {
	return _Vrf.Contract.GetBlockHash(&_Vrf.CallOpts, height)
}

// GetCommitment is a free data retrieval call binding the contract method 0x69bcdb7d.
//
// Solidity: function getCommitment(uint256 requestId) view returns(bytes32)
func (_Vrf *VrfCaller) GetCommitment(opts *bind.CallOpts, requestId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "getCommitment", requestId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCommitment is a free data retrieval call binding the contract method 0x69bcdb7d.
//
// Solidity: function getCommitment(uint256 requestId) view returns(bytes32)
func (_Vrf *VrfSession) GetCommitment(requestId *big.Int) ([32]byte, error) {
	return _Vrf.Contract.GetCommitment(&_Vrf.CallOpts, requestId)
}

// GetCommitment is a free data retrieval call binding the contract method 0x69bcdb7d.
//
// Solidity: function getCommitment(uint256 requestId) view returns(bytes32)
func (_Vrf *VrfCallerSession) GetCommitment(requestId *big.Int) ([32]byte, error) {
	return _Vrf.Contract.GetCommitment(&_Vrf.CallOpts, requestId)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit)
func (_Vrf *VrfCaller) GetConfig(opts *bind.CallOpts) (struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
}, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "getConfig")

	outstruct := new(struct {
		MinimumRequestConfirmations uint16
		MaxGasLimit                 uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit)
func (_Vrf *VrfSession) GetConfig() (struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
}, error) {
	return _Vrf.Contract.GetConfig(&_Vrf.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit)
func (_Vrf *VrfCallerSession) GetConfig() (struct {
	MinimumRequestConfirmations uint16
	MaxGasLimit                 uint32
}, error) {
	return _Vrf.Contract.GetConfig(&_Vrf.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_Vrf *VrfCaller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_Vrf *VrfSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _Vrf.Contract.GetRequestConfig(&_Vrf.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_Vrf *VrfCallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _Vrf.Contract.GetRequestConfig(&_Vrf.CallOpts)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_Vrf *VrfCaller) GetSubscription(opts *bind.CallOpts, subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance   *big.Int
		ReqCount  uint64
		Owner     common.Address
		Consumers []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_Vrf *VrfSession) GetSubscription(subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _Vrf.Contract.GetSubscription(&_Vrf.CallOpts, subId)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_Vrf *VrfCallerSession) GetSubscription(subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _Vrf.Contract.GetSubscription(&_Vrf.CallOpts, subId)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_Vrf *VrfCaller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_Vrf *VrfSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _Vrf.Contract.HashOfKey(&_Vrf.CallOpts, publicKey)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_Vrf *VrfCallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _Vrf.Contract.HashOfKey(&_Vrf.CallOpts, publicKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vrf *VrfCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vrf *VrfSession) Owner() (common.Address, error) {
	return _Vrf.Contract.Owner(&_Vrf.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vrf *VrfCallerSession) Owner() (common.Address, error) {
	return _Vrf.Contract.Owner(&_Vrf.CallOpts)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subId) view returns(bool)
func (_Vrf *VrfCaller) PendingRequestExists(opts *bind.CallOpts, subId uint64) (bool, error) {
	var out []interface{}
	err := _Vrf.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subId) view returns(bool)
func (_Vrf *VrfSession) PendingRequestExists(subId uint64) (bool, error) {
	return _Vrf.Contract.PendingRequestExists(&_Vrf.CallOpts, subId)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0xe82ad7d4.
//
// Solidity: function pendingRequestExists(uint64 subId) view returns(bool)
func (_Vrf *VrfCallerSession) PendingRequestExists(subId uint64) (bool, error) {
	return _Vrf.Contract.PendingRequestExists(&_Vrf.CallOpts, subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_Vrf *VrfTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_Vrf *VrfSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _Vrf.Contract.AcceptSubscriptionOwnerTransfer(&_Vrf.TransactOpts, subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_Vrf *VrfTransactorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _Vrf.Contract.AcceptSubscriptionOwnerTransfer(&_Vrf.TransactOpts, subId)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_Vrf *VrfTransactor) AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "addConsumer", subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_Vrf *VrfSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.AddConsumer(&_Vrf.TransactOpts, subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_Vrf *VrfTransactorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.AddConsumer(&_Vrf.TransactOpts, subId, consumer)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_Vrf *VrfTransactor) CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "cancelSubscription", subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_Vrf *VrfSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.CancelSubscription(&_Vrf.TransactOpts, subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_Vrf *VrfTransactorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.CancelSubscription(&_Vrf.TransactOpts, subId, to)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64 subId)
func (_Vrf *VrfTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "createSubscription")
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64 subId)
func (_Vrf *VrfSession) CreateSubscription() (*types.Transaction, error) {
	return _Vrf.Contract.CreateSubscription(&_Vrf.TransactOpts)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64 subId)
func (_Vrf *VrfTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _Vrf.Contract.CreateSubscription(&_Vrf.TransactOpts)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_Vrf *VrfTransactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_Vrf *VrfSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Vrf.Contract.DeregisterProvingKey(&_Vrf.TransactOpts, publicProvingKey)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_Vrf *VrfTransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Vrf.Contract.DeregisterProvingKey(&_Vrf.TransactOpts, publicProvingKey)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xaf198b97.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint64,uint32,uint32,address) rc) returns(uint96)
func (_Vrf *VrfTransactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFCoordinatorRequestCommitment) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "fulfillRandomWords", proof, rc)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xaf198b97.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint64,uint32,uint32,address) rc) returns(uint96)
func (_Vrf *VrfSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorRequestCommitment) (*types.Transaction, error) {
	return _Vrf.Contract.FulfillRandomWords(&_Vrf.TransactOpts, proof, rc)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0xaf198b97.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint64,uint32,uint32,address) rc) returns(uint96)
func (_Vrf *VrfTransactorSession) FulfillRandomWords(proof VRFProof, rc VRFCoordinatorRequestCommitment) (*types.Transaction, error) {
	return _Vrf.Contract.FulfillRandomWords(&_Vrf.TransactOpts, proof, rc)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vrf *VrfTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vrf *VrfSession) Initialize() (*types.Transaction, error) {
	return _Vrf.Contract.Initialize(&_Vrf.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Vrf *VrfTransactorSession) Initialize() (*types.Transaction, error) {
	return _Vrf.Contract.Initialize(&_Vrf.TransactOpts)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x7bce14d1.
//
// Solidity: function registerProvingKey(uint256[2] publicProvingKey) returns()
func (_Vrf *VrfTransactor) RegisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "registerProvingKey", publicProvingKey)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x7bce14d1.
//
// Solidity: function registerProvingKey(uint256[2] publicProvingKey) returns()
func (_Vrf *VrfSession) RegisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Vrf.Contract.RegisterProvingKey(&_Vrf.TransactOpts, publicProvingKey)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x7bce14d1.
//
// Solidity: function registerProvingKey(uint256[2] publicProvingKey) returns()
func (_Vrf *VrfTransactorSession) RegisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _Vrf.Contract.RegisterProvingKey(&_Vrf.TransactOpts, publicProvingKey)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_Vrf *VrfTransactor) RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "removeConsumer", subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_Vrf *VrfSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.RemoveConsumer(&_Vrf.TransactOpts, subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_Vrf *VrfTransactorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.RemoveConsumer(&_Vrf.TransactOpts, subId, consumer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vrf *VrfTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vrf *VrfSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vrf.Contract.RenounceOwnership(&_Vrf.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vrf *VrfTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vrf.Contract.RenounceOwnership(&_Vrf.TransactOpts)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 requestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256)
func (_Vrf *VrfTransactor) RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "requestRandomWords", keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 requestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256)
func (_Vrf *VrfSession) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _Vrf.Contract.RequestRandomWords(&_Vrf.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 requestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256)
func (_Vrf *VrfTransactorSession) RequestRandomWords(keyHash [32]byte, subId uint64, requestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _Vrf.Contract.RequestRandomWords(&_Vrf.TransactOpts, keyHash, subId, requestConfirmations, callbackGasLimit, numWords)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_Vrf *VrfTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_Vrf *VrfSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.RequestSubscriptionOwnerTransfer(&_Vrf.TransactOpts, subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_Vrf *VrfTransactorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.RequestSubscriptionOwnerTransfer(&_Vrf.TransactOpts, subId, newOwner)
}

// SetConfig is a paid mutator transaction binding the contract method 0x0fe9842a.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit) returns()
func (_Vrf *VrfTransactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit)
}

// SetConfig is a paid mutator transaction binding the contract method 0x0fe9842a.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit) returns()
func (_Vrf *VrfSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32) (*types.Transaction, error) {
	return _Vrf.Contract.SetConfig(&_Vrf.TransactOpts, minimumRequestConfirmations, maxGasLimit)
}

// SetConfig is a paid mutator transaction binding the contract method 0x0fe9842a.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit) returns()
func (_Vrf *VrfTransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32) (*types.Transaction, error) {
	return _Vrf.Contract.SetConfig(&_Vrf.TransactOpts, minimumRequestConfirmations, maxGasLimit)
}

// StoreBlockHash is a paid mutator transaction binding the contract method 0x15bf8212.
//
// Solidity: function storeBlockHash(uint256 height, bytes32 hash) returns()
func (_Vrf *VrfTransactor) StoreBlockHash(opts *bind.TransactOpts, height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "storeBlockHash", height, hash)
}

// StoreBlockHash is a paid mutator transaction binding the contract method 0x15bf8212.
//
// Solidity: function storeBlockHash(uint256 height, bytes32 hash) returns()
func (_Vrf *VrfSession) StoreBlockHash(height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _Vrf.Contract.StoreBlockHash(&_Vrf.TransactOpts, height, hash)
}

// StoreBlockHash is a paid mutator transaction binding the contract method 0x15bf8212.
//
// Solidity: function storeBlockHash(uint256 height, bytes32 hash) returns()
func (_Vrf *VrfTransactorSession) StoreBlockHash(height *big.Int, hash [32]byte) (*types.Transaction, error) {
	return _Vrf.Contract.StoreBlockHash(&_Vrf.TransactOpts, height, hash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vrf *VrfTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vrf.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vrf *VrfSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.TransferOwnership(&_Vrf.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vrf *VrfTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vrf.Contract.TransferOwnership(&_Vrf.TransactOpts, newOwner)
}

// VrfInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Vrf contract.
type VrfInitializedIterator struct {
	Event *VrfInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfInitialized represents a Initialized event raised by the Vrf contract.
type VrfInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vrf *VrfFilterer) FilterInitialized(opts *bind.FilterOpts) (*VrfInitializedIterator, error) {

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &VrfInitializedIterator{contract: _Vrf.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vrf *VrfFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *VrfInitialized) (event.Subscription, error) {

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfInitialized)
				if err := _Vrf.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Vrf *VrfFilterer) ParseInitialized(log types.Log) (*VrfInitialized, error) {
	event := new(VrfInitialized)
	if err := _Vrf.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vrf contract.
type VrfOwnershipTransferredIterator struct {
	Event *VrfOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfOwnershipTransferred represents a OwnershipTransferred event raised by the Vrf contract.
type VrfOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vrf *VrfFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VrfOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VrfOwnershipTransferredIterator{contract: _Vrf.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vrf *VrfFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VrfOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfOwnershipTransferred)
				if err := _Vrf.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vrf *VrfFilterer) ParseOwnershipTransferred(log types.Log) (*VrfOwnershipTransferred, error) {
	event := new(VrfOwnershipTransferred)
	if err := _Vrf.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfProvingKeyDeregisteredIterator is returned from FilterProvingKeyDeregistered and is used to iterate over the raw logs and unpacked data for ProvingKeyDeregistered events raised by the Vrf contract.
type VrfProvingKeyDeregisteredIterator struct {
	Event *VrfProvingKeyDeregistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfProvingKeyDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfProvingKeyDeregistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfProvingKeyDeregistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfProvingKeyDeregistered represents a ProvingKeyDeregistered event raised by the Vrf contract.
type VrfProvingKeyDeregistered struct {
	KeyHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProvingKeyDeregistered is a free log retrieval operation binding the contract event 0xbd242ec01625c15ecbc02cf700ac8b02c86f7346fa91a08e186810221ae509d0.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash)
func (_Vrf *VrfFilterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts) (*VrfProvingKeyDeregisteredIterator, error) {

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "ProvingKeyDeregistered")
	if err != nil {
		return nil, err
	}
	return &VrfProvingKeyDeregisteredIterator{contract: _Vrf.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

// WatchProvingKeyDeregistered is a free log subscription operation binding the contract event 0xbd242ec01625c15ecbc02cf700ac8b02c86f7346fa91a08e186810221ae509d0.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash)
func (_Vrf *VrfFilterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VrfProvingKeyDeregistered) (event.Subscription, error) {

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "ProvingKeyDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfProvingKeyDeregistered)
				if err := _Vrf.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvingKeyDeregistered is a log parse operation binding the contract event 0xbd242ec01625c15ecbc02cf700ac8b02c86f7346fa91a08e186810221ae509d0.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash)
func (_Vrf *VrfFilterer) ParseProvingKeyDeregistered(log types.Log) (*VrfProvingKeyDeregistered, error) {
	event := new(VrfProvingKeyDeregistered)
	if err := _Vrf.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfProvingKeyRegisteredIterator is returned from FilterProvingKeyRegistered and is used to iterate over the raw logs and unpacked data for ProvingKeyRegistered events raised by the Vrf contract.
type VrfProvingKeyRegisteredIterator struct {
	Event *VrfProvingKeyRegistered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfProvingKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfProvingKeyRegistered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfProvingKeyRegistered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfProvingKeyRegistered represents a ProvingKeyRegistered event raised by the Vrf contract.
type VrfProvingKeyRegistered struct {
	KeyHash [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProvingKeyRegistered is a free log retrieval operation binding the contract event 0xc9583fd3afa3d7f16eb0b88d0268e7d05c09bafa4b21e092cbd1320e1bc8089d.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash)
func (_Vrf *VrfFilterer) FilterProvingKeyRegistered(opts *bind.FilterOpts) (*VrfProvingKeyRegisteredIterator, error) {

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "ProvingKeyRegistered")
	if err != nil {
		return nil, err
	}
	return &VrfProvingKeyRegisteredIterator{contract: _Vrf.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchProvingKeyRegistered is a free log subscription operation binding the contract event 0xc9583fd3afa3d7f16eb0b88d0268e7d05c09bafa4b21e092cbd1320e1bc8089d.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash)
func (_Vrf *VrfFilterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VrfProvingKeyRegistered) (event.Subscription, error) {

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "ProvingKeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfProvingKeyRegistered)
				if err := _Vrf.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvingKeyRegistered is a log parse operation binding the contract event 0xc9583fd3afa3d7f16eb0b88d0268e7d05c09bafa4b21e092cbd1320e1bc8089d.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash)
func (_Vrf *VrfFilterer) ParseProvingKeyRegistered(log types.Log) (*VrfProvingKeyRegistered, error) {
	event := new(VrfProvingKeyRegistered)
	if err := _Vrf.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfRandomWordsFulfilledIterator is returned from FilterRandomWordsFulfilled and is used to iterate over the raw logs and unpacked data for RandomWordsFulfilled events raised by the Vrf contract.
type VrfRandomWordsFulfilledIterator struct {
	Event *VrfRandomWordsFulfilled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfRandomWordsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfRandomWordsFulfilled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfRandomWordsFulfilled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfRandomWordsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfRandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfRandomWordsFulfilled represents a RandomWordsFulfilled event raised by the Vrf contract.
type VrfRandomWordsFulfilled struct {
	RequestId  *big.Int
	OutputSeed *big.Int
	Payment    *big.Int
	Success    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsFulfilled is a free log retrieval operation binding the contract event 0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint96 payment, bool success)
func (_Vrf *VrfFilterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*VrfRandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfRandomWordsFulfilledIterator{contract: _Vrf.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomWordsFulfilled is a free log subscription operation binding the contract event 0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint96 payment, bool success)
func (_Vrf *VrfFilterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VrfRandomWordsFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfRandomWordsFulfilled)
				if err := _Vrf.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordsFulfilled is a log parse operation binding the contract event 0x7dffc5ae5ee4e2e4df1651cf6ad329a73cebdb728f37ea0187b9b17e036756e4.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint96 payment, bool success)
func (_Vrf *VrfFilterer) ParseRandomWordsFulfilled(log types.Log) (*VrfRandomWordsFulfilled, error) {
	event := new(VrfRandomWordsFulfilled)
	if err := _Vrf.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfRandomWordsRequestedIterator is returned from FilterRandomWordsRequested and is used to iterate over the raw logs and unpacked data for RandomWordsRequested events raised by the Vrf contract.
type VrfRandomWordsRequestedIterator struct {
	Event *VrfRandomWordsRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfRandomWordsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfRandomWordsRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfRandomWordsRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfRandomWordsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfRandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfRandomWordsRequested represents a RandomWordsRequested event raised by the Vrf contract.
type VrfRandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       uint64
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	Sender                      common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsRequested is a free log retrieval operation binding the contract event 0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint64 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, address indexed sender)
func (_Vrf *VrfFilterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []uint64, sender []common.Address) (*VrfRandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VrfRandomWordsRequestedIterator{contract: _Vrf.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordsRequested is a free log subscription operation binding the contract event 0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint64 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, address indexed sender)
func (_Vrf *VrfFilterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VrfRandomWordsRequested, keyHash [][32]byte, subId []uint64, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfRandomWordsRequested)
				if err := _Vrf.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRandomWordsRequested is a log parse operation binding the contract event 0x63373d1c4696214b898952999c9aaec57dac1ee2723cec59bea6888f489a9772.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint64 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, address indexed sender)
func (_Vrf *VrfFilterer) ParseRandomWordsRequested(log types.Log) (*VrfRandomWordsRequested, error) {
	event := new(VrfRandomWordsRequested)
	if err := _Vrf.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfSubscriptionCanceledIterator is returned from FilterSubscriptionCanceled and is used to iterate over the raw logs and unpacked data for SubscriptionCanceled events raised by the Vrf contract.
type VrfSubscriptionCanceledIterator struct {
	Event *VrfSubscriptionCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfSubscriptionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfSubscriptionCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfSubscriptionCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfSubscriptionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfSubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfSubscriptionCanceled represents a SubscriptionCanceled event raised by the Vrf contract.
type VrfSubscriptionCanceled struct {
	SubId  uint64
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCanceled is a free log retrieval operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subId, address to, uint256 amount)
func (_Vrf *VrfFilterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []uint64) (*VrfSubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfSubscriptionCanceledIterator{contract: _Vrf.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCanceled is a free log subscription operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subId, address to, uint256 amount)
func (_Vrf *VrfFilterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VrfSubscriptionCanceled, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfSubscriptionCanceled)
				if err := _Vrf.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionCanceled is a log parse operation binding the contract event 0xe8ed5b475a5b5987aa9165e8731bb78043f39eee32ec5a1169a89e27fcd49815.
//
// Solidity: event SubscriptionCanceled(uint64 indexed subId, address to, uint256 amount)
func (_Vrf *VrfFilterer) ParseSubscriptionCanceled(log types.Log) (*VrfSubscriptionCanceled, error) {
	event := new(VrfSubscriptionCanceled)
	if err := _Vrf.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfSubscriptionConsumerAddedIterator is returned from FilterSubscriptionConsumerAdded and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerAdded events raised by the Vrf contract.
type VrfSubscriptionConsumerAddedIterator struct {
	Event *VrfSubscriptionConsumerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfSubscriptionConsumerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfSubscriptionConsumerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfSubscriptionConsumerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfSubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfSubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfSubscriptionConsumerAdded represents a SubscriptionConsumerAdded event raised by the Vrf contract.
type VrfSubscriptionConsumerAdded struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerAdded is a free log retrieval operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subId, address consumer)
func (_Vrf *VrfFilterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []uint64) (*VrfSubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfSubscriptionConsumerAddedIterator{contract: _Vrf.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerAdded is a free log subscription operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subId, address consumer)
func (_Vrf *VrfFilterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VrfSubscriptionConsumerAdded, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfSubscriptionConsumerAdded)
				if err := _Vrf.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionConsumerAdded is a log parse operation binding the contract event 0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0.
//
// Solidity: event SubscriptionConsumerAdded(uint64 indexed subId, address consumer)
func (_Vrf *VrfFilterer) ParseSubscriptionConsumerAdded(log types.Log) (*VrfSubscriptionConsumerAdded, error) {
	event := new(VrfSubscriptionConsumerAdded)
	if err := _Vrf.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfSubscriptionConsumerRemovedIterator is returned from FilterSubscriptionConsumerRemoved and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerRemoved events raised by the Vrf contract.
type VrfSubscriptionConsumerRemovedIterator struct {
	Event *VrfSubscriptionConsumerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfSubscriptionConsumerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfSubscriptionConsumerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfSubscriptionConsumerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfSubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfSubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfSubscriptionConsumerRemoved represents a SubscriptionConsumerRemoved event raised by the Vrf contract.
type VrfSubscriptionConsumerRemoved struct {
	SubId    uint64
	Consumer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerRemoved is a free log retrieval operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer)
func (_Vrf *VrfFilterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []uint64) (*VrfSubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfSubscriptionConsumerRemovedIterator{contract: _Vrf.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerRemoved is a free log subscription operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer)
func (_Vrf *VrfFilterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VrfSubscriptionConsumerRemoved, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfSubscriptionConsumerRemoved)
				if err := _Vrf.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionConsumerRemoved is a log parse operation binding the contract event 0x182bff9831466789164ca77075fffd84916d35a8180ba73c27e45634549b445b.
//
// Solidity: event SubscriptionConsumerRemoved(uint64 indexed subId, address consumer)
func (_Vrf *VrfFilterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VrfSubscriptionConsumerRemoved, error) {
	event := new(VrfSubscriptionConsumerRemoved)
	if err := _Vrf.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfSubscriptionCreatedIterator is returned from FilterSubscriptionCreated and is used to iterate over the raw logs and unpacked data for SubscriptionCreated events raised by the Vrf contract.
type VrfSubscriptionCreatedIterator struct {
	Event *VrfSubscriptionCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VrfSubscriptionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfSubscriptionCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VrfSubscriptionCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VrfSubscriptionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfSubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfSubscriptionCreated represents a SubscriptionCreated event raised by the Vrf contract.
type VrfSubscriptionCreated struct {
	SubId uint64
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCreated is a free log retrieval operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subId, address owner)
func (_Vrf *VrfFilterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []uint64) (*VrfSubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfSubscriptionCreatedIterator{contract: _Vrf.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCreated is a free log subscription operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subId, address owner)
func (_Vrf *VrfFilterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VrfSubscriptionCreated, subId []uint64) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _Vrf.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfSubscriptionCreated)
				if err := _Vrf.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSubscriptionCreated is a log parse operation binding the contract event 0x464722b4166576d3dcbba877b999bc35cf911f4eaf434b7eba68fa113951d0bf.
//
// Solidity: event SubscriptionCreated(uint64 indexed subId, address owner)
func (_Vrf *VrfFilterer) ParseSubscriptionCreated(log types.Log) (*VrfSubscriptionCreated, error) {
	event := new(VrfSubscriptionCreated)
	if err := _Vrf.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
