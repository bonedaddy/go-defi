// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package presalefactory

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

// PresalefactoryABI is the input ABI used to generate the binding from.
const PresalefactoryABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"presaleContract\",\"type\":\"address\"}],\"name\":\"presaleRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_allow\",\"type\":\"bool\"}],\"name\":\"adminAllowPresaleGenerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"presaleAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"presaleGeneratorAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleGeneratorsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_presaleAddress\",\"type\":\"address\"}],\"name\":\"presaleIsRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presalesLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_presaleAddress\",\"type\":\"address\"}],\"name\":\"registerPresale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Presalefactory is an auto generated Go binding around an Ethereum contract.
type Presalefactory struct {
	PresalefactoryCaller     // Read-only binding to the contract
	PresalefactoryTransactor // Write-only binding to the contract
	PresalefactoryFilterer   // Log filterer for contract events
}

// PresalefactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PresalefactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresalefactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PresalefactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresalefactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PresalefactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresalefactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PresalefactorySession struct {
	Contract     *Presalefactory   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PresalefactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PresalefactoryCallerSession struct {
	Contract *PresalefactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PresalefactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PresalefactoryTransactorSession struct {
	Contract     *PresalefactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PresalefactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PresalefactoryRaw struct {
	Contract *Presalefactory // Generic contract binding to access the raw methods on
}

// PresalefactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PresalefactoryCallerRaw struct {
	Contract *PresalefactoryCaller // Generic read-only contract binding to access the raw methods on
}

// PresalefactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PresalefactoryTransactorRaw struct {
	Contract *PresalefactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPresalefactory creates a new instance of Presalefactory, bound to a specific deployed contract.
func NewPresalefactory(address common.Address, backend bind.ContractBackend) (*Presalefactory, error) {
	contract, err := bindPresalefactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Presalefactory{PresalefactoryCaller: PresalefactoryCaller{contract: contract}, PresalefactoryTransactor: PresalefactoryTransactor{contract: contract}, PresalefactoryFilterer: PresalefactoryFilterer{contract: contract}}, nil
}

// NewPresalefactoryCaller creates a new read-only instance of Presalefactory, bound to a specific deployed contract.
func NewPresalefactoryCaller(address common.Address, caller bind.ContractCaller) (*PresalefactoryCaller, error) {
	contract, err := bindPresalefactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PresalefactoryCaller{contract: contract}, nil
}

// NewPresalefactoryTransactor creates a new write-only instance of Presalefactory, bound to a specific deployed contract.
func NewPresalefactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*PresalefactoryTransactor, error) {
	contract, err := bindPresalefactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PresalefactoryTransactor{contract: contract}, nil
}

// NewPresalefactoryFilterer creates a new log filterer instance of Presalefactory, bound to a specific deployed contract.
func NewPresalefactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*PresalefactoryFilterer, error) {
	contract, err := bindPresalefactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PresalefactoryFilterer{contract: contract}, nil
}

// bindPresalefactory binds a generic wrapper to an already deployed contract.
func bindPresalefactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PresalefactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Presalefactory *PresalefactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Presalefactory.Contract.PresalefactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Presalefactory *PresalefactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presalefactory.Contract.PresalefactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Presalefactory *PresalefactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Presalefactory.Contract.PresalefactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Presalefactory *PresalefactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Presalefactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Presalefactory *PresalefactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presalefactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Presalefactory *PresalefactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Presalefactory.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presalefactory *PresalefactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Presalefactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presalefactory *PresalefactorySession) Owner() (common.Address, error) {
	return _Presalefactory.Contract.Owner(&_Presalefactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Presalefactory *PresalefactoryCallerSession) Owner() (common.Address, error) {
	return _Presalefactory.Contract.Owner(&_Presalefactory.CallOpts)
}

// PresaleAtIndex is a free data retrieval call binding the contract method 0x65384f36.
//
// Solidity: function presaleAtIndex(uint256 _index) view returns(address)
func (_Presalefactory *PresalefactoryCaller) PresaleAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Presalefactory.contract.Call(opts, &out, "presaleAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PresaleAtIndex is a free data retrieval call binding the contract method 0x65384f36.
//
// Solidity: function presaleAtIndex(uint256 _index) view returns(address)
func (_Presalefactory *PresalefactorySession) PresaleAtIndex(_index *big.Int) (common.Address, error) {
	return _Presalefactory.Contract.PresaleAtIndex(&_Presalefactory.CallOpts, _index)
}

// PresaleAtIndex is a free data retrieval call binding the contract method 0x65384f36.
//
// Solidity: function presaleAtIndex(uint256 _index) view returns(address)
func (_Presalefactory *PresalefactoryCallerSession) PresaleAtIndex(_index *big.Int) (common.Address, error) {
	return _Presalefactory.Contract.PresaleAtIndex(&_Presalefactory.CallOpts, _index)
}

// PresaleGeneratorAtIndex is a free data retrieval call binding the contract method 0x0a014fbc.
//
// Solidity: function presaleGeneratorAtIndex(uint256 _index) view returns(address)
func (_Presalefactory *PresalefactoryCaller) PresaleGeneratorAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Presalefactory.contract.Call(opts, &out, "presaleGeneratorAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PresaleGeneratorAtIndex is a free data retrieval call binding the contract method 0x0a014fbc.
//
// Solidity: function presaleGeneratorAtIndex(uint256 _index) view returns(address)
func (_Presalefactory *PresalefactorySession) PresaleGeneratorAtIndex(_index *big.Int) (common.Address, error) {
	return _Presalefactory.Contract.PresaleGeneratorAtIndex(&_Presalefactory.CallOpts, _index)
}

// PresaleGeneratorAtIndex is a free data retrieval call binding the contract method 0x0a014fbc.
//
// Solidity: function presaleGeneratorAtIndex(uint256 _index) view returns(address)
func (_Presalefactory *PresalefactoryCallerSession) PresaleGeneratorAtIndex(_index *big.Int) (common.Address, error) {
	return _Presalefactory.Contract.PresaleGeneratorAtIndex(&_Presalefactory.CallOpts, _index)
}

// PresaleGeneratorsLength is a free data retrieval call binding the contract method 0x9ff20793.
//
// Solidity: function presaleGeneratorsLength() view returns(uint256)
func (_Presalefactory *PresalefactoryCaller) PresaleGeneratorsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presalefactory.contract.Call(opts, &out, "presaleGeneratorsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleGeneratorsLength is a free data retrieval call binding the contract method 0x9ff20793.
//
// Solidity: function presaleGeneratorsLength() view returns(uint256)
func (_Presalefactory *PresalefactorySession) PresaleGeneratorsLength() (*big.Int, error) {
	return _Presalefactory.Contract.PresaleGeneratorsLength(&_Presalefactory.CallOpts)
}

// PresaleGeneratorsLength is a free data retrieval call binding the contract method 0x9ff20793.
//
// Solidity: function presaleGeneratorsLength() view returns(uint256)
func (_Presalefactory *PresalefactoryCallerSession) PresaleGeneratorsLength() (*big.Int, error) {
	return _Presalefactory.Contract.PresaleGeneratorsLength(&_Presalefactory.CallOpts)
}

// PresaleIsRegistered is a free data retrieval call binding the contract method 0xe2fc90ca.
//
// Solidity: function presaleIsRegistered(address _presaleAddress) view returns(bool)
func (_Presalefactory *PresalefactoryCaller) PresaleIsRegistered(opts *bind.CallOpts, _presaleAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Presalefactory.contract.Call(opts, &out, "presaleIsRegistered", _presaleAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleIsRegistered is a free data retrieval call binding the contract method 0xe2fc90ca.
//
// Solidity: function presaleIsRegistered(address _presaleAddress) view returns(bool)
func (_Presalefactory *PresalefactorySession) PresaleIsRegistered(_presaleAddress common.Address) (bool, error) {
	return _Presalefactory.Contract.PresaleIsRegistered(&_Presalefactory.CallOpts, _presaleAddress)
}

// PresaleIsRegistered is a free data retrieval call binding the contract method 0xe2fc90ca.
//
// Solidity: function presaleIsRegistered(address _presaleAddress) view returns(bool)
func (_Presalefactory *PresalefactoryCallerSession) PresaleIsRegistered(_presaleAddress common.Address) (bool, error) {
	return _Presalefactory.Contract.PresaleIsRegistered(&_Presalefactory.CallOpts, _presaleAddress)
}

// PresalesLength is a free data retrieval call binding the contract method 0x4e76edbb.
//
// Solidity: function presalesLength() view returns(uint256)
func (_Presalefactory *PresalefactoryCaller) PresalesLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Presalefactory.contract.Call(opts, &out, "presalesLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresalesLength is a free data retrieval call binding the contract method 0x4e76edbb.
//
// Solidity: function presalesLength() view returns(uint256)
func (_Presalefactory *PresalefactorySession) PresalesLength() (*big.Int, error) {
	return _Presalefactory.Contract.PresalesLength(&_Presalefactory.CallOpts)
}

// PresalesLength is a free data retrieval call binding the contract method 0x4e76edbb.
//
// Solidity: function presalesLength() view returns(uint256)
func (_Presalefactory *PresalefactoryCallerSession) PresalesLength() (*big.Int, error) {
	return _Presalefactory.Contract.PresalesLength(&_Presalefactory.CallOpts)
}

// AdminAllowPresaleGenerator is a paid mutator transaction binding the contract method 0xd348c964.
//
// Solidity: function adminAllowPresaleGenerator(address _address, bool _allow) returns()
func (_Presalefactory *PresalefactoryTransactor) AdminAllowPresaleGenerator(opts *bind.TransactOpts, _address common.Address, _allow bool) (*types.Transaction, error) {
	return _Presalefactory.contract.Transact(opts, "adminAllowPresaleGenerator", _address, _allow)
}

// AdminAllowPresaleGenerator is a paid mutator transaction binding the contract method 0xd348c964.
//
// Solidity: function adminAllowPresaleGenerator(address _address, bool _allow) returns()
func (_Presalefactory *PresalefactorySession) AdminAllowPresaleGenerator(_address common.Address, _allow bool) (*types.Transaction, error) {
	return _Presalefactory.Contract.AdminAllowPresaleGenerator(&_Presalefactory.TransactOpts, _address, _allow)
}

// AdminAllowPresaleGenerator is a paid mutator transaction binding the contract method 0xd348c964.
//
// Solidity: function adminAllowPresaleGenerator(address _address, bool _allow) returns()
func (_Presalefactory *PresalefactoryTransactorSession) AdminAllowPresaleGenerator(_address common.Address, _allow bool) (*types.Transaction, error) {
	return _Presalefactory.Contract.AdminAllowPresaleGenerator(&_Presalefactory.TransactOpts, _address, _allow)
}

// RegisterPresale is a paid mutator transaction binding the contract method 0x11c065b7.
//
// Solidity: function registerPresale(address _presaleAddress) returns()
func (_Presalefactory *PresalefactoryTransactor) RegisterPresale(opts *bind.TransactOpts, _presaleAddress common.Address) (*types.Transaction, error) {
	return _Presalefactory.contract.Transact(opts, "registerPresale", _presaleAddress)
}

// RegisterPresale is a paid mutator transaction binding the contract method 0x11c065b7.
//
// Solidity: function registerPresale(address _presaleAddress) returns()
func (_Presalefactory *PresalefactorySession) RegisterPresale(_presaleAddress common.Address) (*types.Transaction, error) {
	return _Presalefactory.Contract.RegisterPresale(&_Presalefactory.TransactOpts, _presaleAddress)
}

// RegisterPresale is a paid mutator transaction binding the contract method 0x11c065b7.
//
// Solidity: function registerPresale(address _presaleAddress) returns()
func (_Presalefactory *PresalefactoryTransactorSession) RegisterPresale(_presaleAddress common.Address) (*types.Transaction, error) {
	return _Presalefactory.Contract.RegisterPresale(&_Presalefactory.TransactOpts, _presaleAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presalefactory *PresalefactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Presalefactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presalefactory *PresalefactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _Presalefactory.Contract.RenounceOwnership(&_Presalefactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Presalefactory *PresalefactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Presalefactory.Contract.RenounceOwnership(&_Presalefactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presalefactory *PresalefactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Presalefactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presalefactory *PresalefactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Presalefactory.Contract.TransferOwnership(&_Presalefactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Presalefactory *PresalefactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Presalefactory.Contract.TransferOwnership(&_Presalefactory.TransactOpts, newOwner)
}

// PresalefactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Presalefactory contract.
type PresalefactoryOwnershipTransferredIterator struct {
	Event *PresalefactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PresalefactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresalefactoryOwnershipTransferred)
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
		it.Event = new(PresalefactoryOwnershipTransferred)
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
func (it *PresalefactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresalefactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresalefactoryOwnershipTransferred represents a OwnershipTransferred event raised by the Presalefactory contract.
type PresalefactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presalefactory *PresalefactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PresalefactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Presalefactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresalefactoryOwnershipTransferredIterator{contract: _Presalefactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Presalefactory *PresalefactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PresalefactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Presalefactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresalefactoryOwnershipTransferred)
				if err := _Presalefactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Presalefactory *PresalefactoryFilterer) ParseOwnershipTransferred(log types.Log) (*PresalefactoryOwnershipTransferred, error) {
	event := new(PresalefactoryOwnershipTransferred)
	if err := _Presalefactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PresalefactoryPresaleRegisteredIterator is returned from FilterPresaleRegistered and is used to iterate over the raw logs and unpacked data for PresaleRegistered events raised by the Presalefactory contract.
type PresalefactoryPresaleRegisteredIterator struct {
	Event *PresalefactoryPresaleRegistered // Event containing the contract specifics and raw log

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
func (it *PresalefactoryPresaleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresalefactoryPresaleRegistered)
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
		it.Event = new(PresalefactoryPresaleRegistered)
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
func (it *PresalefactoryPresaleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresalefactoryPresaleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresalefactoryPresaleRegistered represents a PresaleRegistered event raised by the Presalefactory contract.
type PresalefactoryPresaleRegistered struct {
	PresaleContract common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPresaleRegistered is a free log retrieval operation binding the contract event 0xa62fce43cb61612c50d7b3485c2fb44000803dacc3472b8ce4c638f235e97a1f.
//
// Solidity: event presaleRegistered(address presaleContract)
func (_Presalefactory *PresalefactoryFilterer) FilterPresaleRegistered(opts *bind.FilterOpts) (*PresalefactoryPresaleRegisteredIterator, error) {

	logs, sub, err := _Presalefactory.contract.FilterLogs(opts, "presaleRegistered")
	if err != nil {
		return nil, err
	}
	return &PresalefactoryPresaleRegisteredIterator{contract: _Presalefactory.contract, event: "presaleRegistered", logs: logs, sub: sub}, nil
}

// WatchPresaleRegistered is a free log subscription operation binding the contract event 0xa62fce43cb61612c50d7b3485c2fb44000803dacc3472b8ce4c638f235e97a1f.
//
// Solidity: event presaleRegistered(address presaleContract)
func (_Presalefactory *PresalefactoryFilterer) WatchPresaleRegistered(opts *bind.WatchOpts, sink chan<- *PresalefactoryPresaleRegistered) (event.Subscription, error) {

	logs, sub, err := _Presalefactory.contract.WatchLogs(opts, "presaleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresalefactoryPresaleRegistered)
				if err := _Presalefactory.contract.UnpackLog(event, "presaleRegistered", log); err != nil {
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

// ParsePresaleRegistered is a log parse operation binding the contract event 0xa62fce43cb61612c50d7b3485c2fb44000803dacc3472b8ce4c638f235e97a1f.
//
// Solidity: event presaleRegistered(address presaleContract)
func (_Presalefactory *PresalefactoryFilterer) ParsePresaleRegistered(log types.Log) (*PresalefactoryPresaleRegistered, error) {
	event := new(PresalefactoryPresaleRegistered)
	if err := _Presalefactory.contract.UnpackLog(event, "presaleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
