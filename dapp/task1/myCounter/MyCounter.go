// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package myCounter

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

// MyCounterMetaData contains all meta data concerning the MyCounter contract.
var MyCounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b505f5f81905550610177806100225f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c806306661abd14610043578063a87d942c14610061578063d09de08a1461007f575b5f5ffd5b61004b610089565b60405161005891906100c8565b60405180910390f35b61006961008e565b60405161007691906100c8565b60405180910390f35b610087610096565b005b5f5481565b5f5f54905090565b60015f5f8282546100a7919061010e565b92505081905550565b5f819050919050565b6100c2816100b0565b82525050565b5f6020820190506100db5f8301846100b9565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610118826100b0565b9150610123836100b0565b925082820190508082111561013b5761013a6100e1565b5b9291505056fea26469706673582212203d4bf70bb4f2cb3de74fdb068bdadf69ff2a96d6f4fb65e00b14b3657d4cc2e364736f6c634300081e0033",
}

// MyCounterABI is the input ABI used to generate the binding from.
// Deprecated: Use MyCounterMetaData.ABI instead.
var MyCounterABI = MyCounterMetaData.ABI

// MyCounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MyCounterMetaData.Bin instead.
var MyCounterBin = MyCounterMetaData.Bin

// DeployMyCounter deploys a new Ethereum contract, binding an instance of MyCounter to it.
func DeployMyCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MyCounter, error) {
	parsed, err := MyCounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MyCounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyCounter{MyCounterCaller: MyCounterCaller{contract: contract}, MyCounterTransactor: MyCounterTransactor{contract: contract}, MyCounterFilterer: MyCounterFilterer{contract: contract}}, nil
}

// MyCounter is an auto generated Go binding around an Ethereum contract.
type MyCounter struct {
	MyCounterCaller     // Read-only binding to the contract
	MyCounterTransactor // Write-only binding to the contract
	MyCounterFilterer   // Log filterer for contract events
}

// MyCounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MyCounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyCounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MyCounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyCounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyCounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyCounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyCounterSession struct {
	Contract     *MyCounter        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyCounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyCounterCallerSession struct {
	Contract *MyCounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MyCounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyCounterTransactorSession struct {
	Contract     *MyCounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MyCounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MyCounterRaw struct {
	Contract *MyCounter // Generic contract binding to access the raw methods on
}

// MyCounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyCounterCallerRaw struct {
	Contract *MyCounterCaller // Generic read-only contract binding to access the raw methods on
}

// MyCounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyCounterTransactorRaw struct {
	Contract *MyCounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMyCounter creates a new instance of MyCounter, bound to a specific deployed contract.
func NewMyCounter(address common.Address, backend bind.ContractBackend) (*MyCounter, error) {
	contract, err := bindMyCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyCounter{MyCounterCaller: MyCounterCaller{contract: contract}, MyCounterTransactor: MyCounterTransactor{contract: contract}, MyCounterFilterer: MyCounterFilterer{contract: contract}}, nil
}

// NewMyCounterCaller creates a new read-only instance of MyCounter, bound to a specific deployed contract.
func NewMyCounterCaller(address common.Address, caller bind.ContractCaller) (*MyCounterCaller, error) {
	contract, err := bindMyCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyCounterCaller{contract: contract}, nil
}

// NewMyCounterTransactor creates a new write-only instance of MyCounter, bound to a specific deployed contract.
func NewMyCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*MyCounterTransactor, error) {
	contract, err := bindMyCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyCounterTransactor{contract: contract}, nil
}

// NewMyCounterFilterer creates a new log filterer instance of MyCounter, bound to a specific deployed contract.
func NewMyCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*MyCounterFilterer, error) {
	contract, err := bindMyCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyCounterFilterer{contract: contract}, nil
}

// bindMyCounter binds a generic wrapper to an already deployed contract.
func bindMyCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MyCounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyCounter *MyCounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyCounter.Contract.MyCounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyCounter *MyCounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyCounter.Contract.MyCounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyCounter *MyCounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyCounter.Contract.MyCounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyCounter *MyCounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MyCounter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyCounter *MyCounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyCounter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyCounter *MyCounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyCounter.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MyCounter *MyCounterCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyCounter.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MyCounter *MyCounterSession) Count() (*big.Int, error) {
	return _MyCounter.Contract.Count(&_MyCounter.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_MyCounter *MyCounterCallerSession) Count() (*big.Int, error) {
	return _MyCounter.Contract.Count(&_MyCounter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_MyCounter *MyCounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MyCounter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_MyCounter *MyCounterSession) GetCount() (*big.Int, error) {
	return _MyCounter.Contract.GetCount(&_MyCounter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_MyCounter *MyCounterCallerSession) GetCount() (*big.Int, error) {
	return _MyCounter.Contract.GetCount(&_MyCounter.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_MyCounter *MyCounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyCounter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_MyCounter *MyCounterSession) Increment() (*types.Transaction, error) {
	return _MyCounter.Contract.Increment(&_MyCounter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_MyCounter *MyCounterTransactorSession) Increment() (*types.Transaction, error) {
	return _MyCounter.Contract.Increment(&_MyCounter.TransactOpts)
}
