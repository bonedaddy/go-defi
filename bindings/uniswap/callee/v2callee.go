// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2callee

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Uniswapv2calleeABI is the input ABI used to generate the binding from.
const Uniswapv2calleeABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"uniswapV2Call\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Uniswapv2callee is an auto generated Go binding around an Ethereum contract.
type Uniswapv2callee struct {
	Uniswapv2calleeCaller     // Read-only binding to the contract
	Uniswapv2calleeTransactor // Write-only binding to the contract
	Uniswapv2calleeFilterer   // Log filterer for contract events
}

// Uniswapv2calleeCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2calleeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2calleeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2calleeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2calleeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2calleeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2calleeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2calleeSession struct {
	Contract     *Uniswapv2callee  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv2calleeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2calleeCallerSession struct {
	Contract *Uniswapv2calleeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// Uniswapv2calleeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2calleeTransactorSession struct {
	Contract     *Uniswapv2calleeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// Uniswapv2calleeRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2calleeRaw struct {
	Contract *Uniswapv2callee // Generic contract binding to access the raw methods on
}

// Uniswapv2calleeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2calleeCallerRaw struct {
	Contract *Uniswapv2calleeCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2calleeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2calleeTransactorRaw struct {
	Contract *Uniswapv2calleeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2callee creates a new instance of Uniswapv2callee, bound to a specific deployed contract.
func NewUniswapv2callee(address common.Address, backend bind.ContractBackend) (*Uniswapv2callee, error) {
	contract, err := bindUniswapv2callee(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2callee{Uniswapv2calleeCaller: Uniswapv2calleeCaller{contract: contract}, Uniswapv2calleeTransactor: Uniswapv2calleeTransactor{contract: contract}, Uniswapv2calleeFilterer: Uniswapv2calleeFilterer{contract: contract}}, nil
}

// NewUniswapv2calleeCaller creates a new read-only instance of Uniswapv2callee, bound to a specific deployed contract.
func NewUniswapv2calleeCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2calleeCaller, error) {
	contract, err := bindUniswapv2callee(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2calleeCaller{contract: contract}, nil
}

// NewUniswapv2calleeTransactor creates a new write-only instance of Uniswapv2callee, bound to a specific deployed contract.
func NewUniswapv2calleeTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2calleeTransactor, error) {
	contract, err := bindUniswapv2callee(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2calleeTransactor{contract: contract}, nil
}

// NewUniswapv2calleeFilterer creates a new log filterer instance of Uniswapv2callee, bound to a specific deployed contract.
func NewUniswapv2calleeFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2calleeFilterer, error) {
	contract, err := bindUniswapv2callee(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2calleeFilterer{contract: contract}, nil
}

// bindUniswapv2callee binds a generic wrapper to an already deployed contract.
func bindUniswapv2callee(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswapv2calleeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2callee *Uniswapv2calleeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2callee.Contract.Uniswapv2calleeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2callee *Uniswapv2calleeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2callee.Contract.Uniswapv2calleeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2callee *Uniswapv2calleeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2callee.Contract.Uniswapv2calleeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2callee *Uniswapv2calleeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2callee.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2callee *Uniswapv2calleeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2callee.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2callee *Uniswapv2calleeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2callee.Contract.contract.Transact(opts, method, params...)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Uniswapv2callee *Uniswapv2calleeTransactor) UniswapV2Call(opts *bind.TransactOpts, sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Uniswapv2callee.contract.Transact(opts, "uniswapV2Call", sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Uniswapv2callee *Uniswapv2calleeSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Uniswapv2callee.Contract.UniswapV2Call(&_Uniswapv2callee.TransactOpts, sender, amount0, amount1, data)
}

// UniswapV2Call is a paid mutator transaction binding the contract method 0x10d1e85c.
//
// Solidity: function uniswapV2Call(address sender, uint256 amount0, uint256 amount1, bytes data) returns()
func (_Uniswapv2callee *Uniswapv2calleeTransactorSession) UniswapV2Call(sender common.Address, amount0 *big.Int, amount1 *big.Int, data []byte) (*types.Transaction, error) {
	return _Uniswapv2callee.Contract.UniswapV2Call(&_Uniswapv2callee.TransactOpts, sender, amount0, amount1, data)
}
