// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package uniswapv2factory

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

// Uniswapv2factoryABI is the input ABI used to generate the binding from.
const Uniswapv2factoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"PairCreated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPairsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"createPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeToSetter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"getPair\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"setFeeToSetter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Uniswapv2factory is an auto generated Go binding around an Ethereum contract.
type Uniswapv2factory struct {
	Uniswapv2factoryCaller     // Read-only binding to the contract
	Uniswapv2factoryTransactor // Write-only binding to the contract
	Uniswapv2factoryFilterer   // Log filterer for contract events
}

// Uniswapv2factoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type Uniswapv2factoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Uniswapv2factoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Uniswapv2factoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Uniswapv2factorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Uniswapv2factorySession struct {
	Contract     *Uniswapv2factory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Uniswapv2factoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Uniswapv2factoryCallerSession struct {
	Contract *Uniswapv2factoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// Uniswapv2factoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Uniswapv2factoryTransactorSession struct {
	Contract     *Uniswapv2factoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// Uniswapv2factoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type Uniswapv2factoryRaw struct {
	Contract *Uniswapv2factory // Generic contract binding to access the raw methods on
}

// Uniswapv2factoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Uniswapv2factoryCallerRaw struct {
	Contract *Uniswapv2factoryCaller // Generic read-only contract binding to access the raw methods on
}

// Uniswapv2factoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Uniswapv2factoryTransactorRaw struct {
	Contract *Uniswapv2factoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapv2factory creates a new instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factory(address common.Address, backend bind.ContractBackend) (*Uniswapv2factory, error) {
	contract, err := bindUniswapv2factory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factory{Uniswapv2factoryCaller: Uniswapv2factoryCaller{contract: contract}, Uniswapv2factoryTransactor: Uniswapv2factoryTransactor{contract: contract}, Uniswapv2factoryFilterer: Uniswapv2factoryFilterer{contract: contract}}, nil
}

// NewUniswapv2factoryCaller creates a new read-only instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factoryCaller(address common.Address, caller bind.ContractCaller) (*Uniswapv2factoryCaller, error) {
	contract, err := bindUniswapv2factory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryCaller{contract: contract}, nil
}

// NewUniswapv2factoryTransactor creates a new write-only instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factoryTransactor(address common.Address, transactor bind.ContractTransactor) (*Uniswapv2factoryTransactor, error) {
	contract, err := bindUniswapv2factory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryTransactor{contract: contract}, nil
}

// NewUniswapv2factoryFilterer creates a new log filterer instance of Uniswapv2factory, bound to a specific deployed contract.
func NewUniswapv2factoryFilterer(address common.Address, filterer bind.ContractFilterer) (*Uniswapv2factoryFilterer, error) {
	contract, err := bindUniswapv2factory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryFilterer{contract: contract}, nil
}

// bindUniswapv2factory binds a generic wrapper to an already deployed contract.
func bindUniswapv2factory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Uniswapv2factoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2factory *Uniswapv2factoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2factory.Contract.Uniswapv2factoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2factory *Uniswapv2factoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.Uniswapv2factoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2factory *Uniswapv2factoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.Uniswapv2factoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Uniswapv2factory *Uniswapv2factoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Uniswapv2factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Uniswapv2factory *Uniswapv2factoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Uniswapv2factory *Uniswapv2factoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.contract.Transact(opts, method, params...)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryCaller) AllPairs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factory.contract.Call(opts, &out, "allPairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factorySession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Uniswapv2factory.Contract.AllPairs(&_Uniswapv2factory.CallOpts, arg0)
}

// AllPairs is a free data retrieval call binding the contract method 0x1e3dd18b.
//
// Solidity: function allPairs(uint256 ) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryCallerSession) AllPairs(arg0 *big.Int) (common.Address, error) {
	return _Uniswapv2factory.Contract.AllPairs(&_Uniswapv2factory.CallOpts, arg0)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapv2factory *Uniswapv2factoryCaller) AllPairsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Uniswapv2factory.contract.Call(opts, &out, "allPairsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapv2factory *Uniswapv2factorySession) AllPairsLength() (*big.Int, error) {
	return _Uniswapv2factory.Contract.AllPairsLength(&_Uniswapv2factory.CallOpts)
}

// AllPairsLength is a free data retrieval call binding the contract method 0x574f2ba3.
//
// Solidity: function allPairsLength() view returns(uint256)
func (_Uniswapv2factory *Uniswapv2factoryCallerSession) AllPairsLength() (*big.Int, error) {
	return _Uniswapv2factory.Contract.AllPairsLength(&_Uniswapv2factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapv2factory *Uniswapv2factoryCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factory.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapv2factory *Uniswapv2factorySession) FeeTo() (common.Address, error) {
	return _Uniswapv2factory.Contract.FeeTo(&_Uniswapv2factory.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Uniswapv2factory *Uniswapv2factoryCallerSession) FeeTo() (common.Address, error) {
	return _Uniswapv2factory.Contract.FeeTo(&_Uniswapv2factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapv2factory *Uniswapv2factoryCaller) FeeToSetter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factory.contract.Call(opts, &out, "feeToSetter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapv2factory *Uniswapv2factorySession) FeeToSetter() (common.Address, error) {
	return _Uniswapv2factory.Contract.FeeToSetter(&_Uniswapv2factory.CallOpts)
}

// FeeToSetter is a free data retrieval call binding the contract method 0x094b7415.
//
// Solidity: function feeToSetter() view returns(address)
func (_Uniswapv2factory *Uniswapv2factoryCallerSession) FeeToSetter() (common.Address, error) {
	return _Uniswapv2factory.Contract.FeeToSetter(&_Uniswapv2factory.CallOpts)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryCaller) GetPair(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (common.Address, error) {
	var out []interface{}
	err := _Uniswapv2factory.contract.Call(opts, &out, "getPair", tokenA, tokenB)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factorySession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Uniswapv2factory.Contract.GetPair(&_Uniswapv2factory.CallOpts, tokenA, tokenB)
}

// GetPair is a free data retrieval call binding the contract method 0xe6a43905.
//
// Solidity: function getPair(address tokenA, address tokenB) view returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryCallerSession) GetPair(tokenA common.Address, tokenB common.Address) (common.Address, error) {
	return _Uniswapv2factory.Contract.GetPair(&_Uniswapv2factory.CallOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryTransactor) CreatePair(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.contract.Transact(opts, "createPair", tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapv2factory *Uniswapv2factorySession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.CreatePair(&_Uniswapv2factory.TransactOpts, tokenA, tokenB)
}

// CreatePair is a paid mutator transaction binding the contract method 0xc9c65396.
//
// Solidity: function createPair(address tokenA, address tokenB) returns(address pair)
func (_Uniswapv2factory *Uniswapv2factoryTransactorSession) CreatePair(tokenA common.Address, tokenB common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.CreatePair(&_Uniswapv2factory.TransactOpts, tokenA, tokenB)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Uniswapv2factory *Uniswapv2factoryTransactor) SetFeeTo(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.contract.Transact(opts, "setFeeTo", arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Uniswapv2factory *Uniswapv2factorySession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.SetFeeTo(&_Uniswapv2factory.TransactOpts, arg0)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address ) returns()
func (_Uniswapv2factory *Uniswapv2factoryTransactorSession) SetFeeTo(arg0 common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.SetFeeTo(&_Uniswapv2factory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Uniswapv2factory *Uniswapv2factoryTransactor) SetFeeToSetter(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.contract.Transact(opts, "setFeeToSetter", arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Uniswapv2factory *Uniswapv2factorySession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.SetFeeToSetter(&_Uniswapv2factory.TransactOpts, arg0)
}

// SetFeeToSetter is a paid mutator transaction binding the contract method 0xa2e74af6.
//
// Solidity: function setFeeToSetter(address ) returns()
func (_Uniswapv2factory *Uniswapv2factoryTransactorSession) SetFeeToSetter(arg0 common.Address) (*types.Transaction, error) {
	return _Uniswapv2factory.Contract.SetFeeToSetter(&_Uniswapv2factory.TransactOpts, arg0)
}

// Uniswapv2factoryPairCreatedIterator is returned from FilterPairCreated and is used to iterate over the raw logs and unpacked data for PairCreated events raised by the Uniswapv2factory contract.
type Uniswapv2factoryPairCreatedIterator struct {
	Event *Uniswapv2factoryPairCreated // Event containing the contract specifics and raw log

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
func (it *Uniswapv2factoryPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Uniswapv2factoryPairCreated)
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
		it.Event = new(Uniswapv2factoryPairCreated)
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
func (it *Uniswapv2factoryPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Uniswapv2factoryPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Uniswapv2factoryPairCreated represents a PairCreated event raised by the Uniswapv2factory contract.
type Uniswapv2factoryPairCreated struct {
	Token0 common.Address
	Token1 common.Address
	Pair   common.Address
	Arg3   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPairCreated is a free log retrieval operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapv2factory *Uniswapv2factoryFilterer) FilterPairCreated(opts *bind.FilterOpts, token0 []common.Address, token1 []common.Address) (*Uniswapv2factoryPairCreatedIterator, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Uniswapv2factory.contract.FilterLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return &Uniswapv2factoryPairCreatedIterator{contract: _Uniswapv2factory.contract, event: "PairCreated", logs: logs, sub: sub}, nil
}

// WatchPairCreated is a free log subscription operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapv2factory *Uniswapv2factoryFilterer) WatchPairCreated(opts *bind.WatchOpts, sink chan<- *Uniswapv2factoryPairCreated, token0 []common.Address, token1 []common.Address) (event.Subscription, error) {

	var token0Rule []interface{}
	for _, token0Item := range token0 {
		token0Rule = append(token0Rule, token0Item)
	}
	var token1Rule []interface{}
	for _, token1Item := range token1 {
		token1Rule = append(token1Rule, token1Item)
	}

	logs, sub, err := _Uniswapv2factory.contract.WatchLogs(opts, "PairCreated", token0Rule, token1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Uniswapv2factoryPairCreated)
				if err := _Uniswapv2factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
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

// ParsePairCreated is a log parse operation binding the contract event 0x0d3648bd0f6ba80134a33ba9275ac585d9d315f0ad8355cddefde31afa28d0e9.
//
// Solidity: event PairCreated(address indexed token0, address indexed token1, address pair, uint256 arg3)
func (_Uniswapv2factory *Uniswapv2factoryFilterer) ParsePairCreated(log types.Log) (*Uniswapv2factoryPairCreated, error) {
	event := new(Uniswapv2factoryPairCreated)
	if err := _Uniswapv2factory.contract.UnpackLog(event, "PairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
