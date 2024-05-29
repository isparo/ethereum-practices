// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simplestorage

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
	_ = abi.ConvertType
)

// SimplestorageMetaData contains all meta data concerning the Simplestorage contract.
var SimplestorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storedData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506101718061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80632a1afcd91461004357806360fe47b1146100615780636d4ce63c1461007d575b5f80fd5b61004b61009b565b60405161005891906100c9565b60405180910390f35b61007b60048036038101906100769190610110565b6100a0565b005b6100856100a9565b60405161009291906100c9565b60405180910390f35b5f5481565b805f8190555050565b5f8054905090565b5f819050919050565b6100c3816100b1565b82525050565b5f6020820190506100dc5f8301846100ba565b92915050565b5f80fd5b6100ef816100b1565b81146100f9575f80fd5b50565b5f8135905061010a816100e6565b92915050565b5f60208284031215610125576101246100e2565b5b5f610132848285016100fc565b9150509291505056fea264697066735822122004968261bdf7ec1986bdb15829b8a1316bbacbac13a312706c937750ff86b67864736f6c63430008190033",
}

// SimplestorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SimplestorageMetaData.ABI instead.
var SimplestorageABI = SimplestorageMetaData.ABI

// SimplestorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimplestorageMetaData.Bin instead.
var SimplestorageBin = SimplestorageMetaData.Bin

// DeploySimplestorage deploys a new Ethereum contract, binding an instance of Simplestorage to it.
func DeploySimplestorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simplestorage, error) {
	parsed, err := SimplestorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimplestorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simplestorage{SimplestorageCaller: SimplestorageCaller{contract: contract}, SimplestorageTransactor: SimplestorageTransactor{contract: contract}, SimplestorageFilterer: SimplestorageFilterer{contract: contract}}, nil
}

// Simplestorage is an auto generated Go binding around an Ethereum contract.
type Simplestorage struct {
	SimplestorageCaller     // Read-only binding to the contract
	SimplestorageTransactor // Write-only binding to the contract
	SimplestorageFilterer   // Log filterer for contract events
}

// SimplestorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimplestorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimplestorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimplestorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimplestorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimplestorageSession struct {
	Contract     *Simplestorage    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimplestorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimplestorageCallerSession struct {
	Contract *SimplestorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimplestorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimplestorageTransactorSession struct {
	Contract     *SimplestorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimplestorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimplestorageRaw struct {
	Contract *Simplestorage // Generic contract binding to access the raw methods on
}

// SimplestorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimplestorageCallerRaw struct {
	Contract *SimplestorageCaller // Generic read-only contract binding to access the raw methods on
}

// SimplestorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimplestorageTransactorRaw struct {
	Contract *SimplestorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimplestorage creates a new instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorage(address common.Address, backend bind.ContractBackend) (*Simplestorage, error) {
	contract, err := bindSimplestorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simplestorage{SimplestorageCaller: SimplestorageCaller{contract: contract}, SimplestorageTransactor: SimplestorageTransactor{contract: contract}, SimplestorageFilterer: SimplestorageFilterer{contract: contract}}, nil
}

// NewSimplestorageCaller creates a new read-only instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorageCaller(address common.Address, caller bind.ContractCaller) (*SimplestorageCaller, error) {
	contract, err := bindSimplestorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimplestorageCaller{contract: contract}, nil
}

// NewSimplestorageTransactor creates a new write-only instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SimplestorageTransactor, error) {
	contract, err := bindSimplestorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimplestorageTransactor{contract: contract}, nil
}

// NewSimplestorageFilterer creates a new log filterer instance of Simplestorage, bound to a specific deployed contract.
func NewSimplestorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SimplestorageFilterer, error) {
	contract, err := bindSimplestorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimplestorageFilterer{contract: contract}, nil
}

// bindSimplestorage binds a generic wrapper to an already deployed contract.
func bindSimplestorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimplestorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplestorage *SimplestorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplestorage.Contract.SimplestorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplestorage *SimplestorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplestorage.Contract.SimplestorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplestorage *SimplestorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplestorage.Contract.SimplestorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simplestorage *SimplestorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simplestorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simplestorage *SimplestorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simplestorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simplestorage *SimplestorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simplestorage.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Simplestorage *SimplestorageCaller) Get(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simplestorage.contract.Call(opts, &out, "get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Simplestorage *SimplestorageSession) Get() (*big.Int, error) {
	return _Simplestorage.Contract.Get(&_Simplestorage.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() view returns(uint256)
func (_Simplestorage *SimplestorageCallerSession) Get() (*big.Int, error) {
	return _Simplestorage.Contract.Get(&_Simplestorage.CallOpts)
}

// StoredData is a free data retrieval call binding the contract method 0x2a1afcd9.
//
// Solidity: function storedData() view returns(uint256)
func (_Simplestorage *SimplestorageCaller) StoredData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simplestorage.contract.Call(opts, &out, "storedData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StoredData is a free data retrieval call binding the contract method 0x2a1afcd9.
//
// Solidity: function storedData() view returns(uint256)
func (_Simplestorage *SimplestorageSession) StoredData() (*big.Int, error) {
	return _Simplestorage.Contract.StoredData(&_Simplestorage.CallOpts)
}

// StoredData is a free data retrieval call binding the contract method 0x2a1afcd9.
//
// Solidity: function storedData() view returns(uint256)
func (_Simplestorage *SimplestorageCallerSession) StoredData() (*big.Int, error) {
	return _Simplestorage.Contract.StoredData(&_Simplestorage.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Simplestorage *SimplestorageTransactor) Set(opts *bind.TransactOpts, x *big.Int) (*types.Transaction, error) {
	return _Simplestorage.contract.Transact(opts, "set", x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Simplestorage *SimplestorageSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Simplestorage.Contract.Set(&_Simplestorage.TransactOpts, x)
}

// Set is a paid mutator transaction binding the contract method 0x60fe47b1.
//
// Solidity: function set(uint256 x) returns()
func (_Simplestorage *SimplestorageTransactorSession) Set(x *big.Int) (*types.Transaction, error) {
	return _Simplestorage.Contract.Set(&_Simplestorage.TransactOpts, x)
}
