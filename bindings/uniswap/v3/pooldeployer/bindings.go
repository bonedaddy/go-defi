// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package v3pooldeployer

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

// V3pooldeployerABI is the input ABI used to generate the binding from.
const V3pooldeployerABI = "[{\"inputs\":[],\"name\":\"parameters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// V3pooldeployerBin is the compiled bytecode used for deploying new contracts.
var V3pooldeployerBin = "0x608060405234801561001057600080fd5b5060f78061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80638903573014602d575b600080fd5b60336082565b6040805173ffffffffffffffffffffffffffffffffffffffff96871681529486166020860152929094168383015262ffffff16606083015260029290920b608082015290519081900360a00190f35b6000546001546002805473ffffffffffffffffffffffffffffffffffffffff938416939283169281169162ffffff7401000000000000000000000000000000000000000083041691770100000000000000000000000000000000000000000000009004900b8556fea164736f6c6343000706000a"

// DeployV3pooldeployer deploys a new Ethereum contract, binding an instance of V3pooldeployer to it.
func DeployV3pooldeployer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *V3pooldeployer, error) {
	parsed, err := abi.JSON(strings.NewReader(V3pooldeployerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(V3pooldeployerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &V3pooldeployer{V3pooldeployerCaller: V3pooldeployerCaller{contract: contract}, V3pooldeployerTransactor: V3pooldeployerTransactor{contract: contract}, V3pooldeployerFilterer: V3pooldeployerFilterer{contract: contract}}, nil
}

// V3pooldeployer is an auto generated Go binding around an Ethereum contract.
type V3pooldeployer struct {
	V3pooldeployerCaller     // Read-only binding to the contract
	V3pooldeployerTransactor // Write-only binding to the contract
	V3pooldeployerFilterer   // Log filterer for contract events
}

// V3pooldeployerCaller is an auto generated read-only Go binding around an Ethereum contract.
type V3pooldeployerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3pooldeployerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type V3pooldeployerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3pooldeployerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type V3pooldeployerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// V3pooldeployerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type V3pooldeployerSession struct {
	Contract     *V3pooldeployer   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// V3pooldeployerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type V3pooldeployerCallerSession struct {
	Contract *V3pooldeployerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// V3pooldeployerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type V3pooldeployerTransactorSession struct {
	Contract     *V3pooldeployerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// V3pooldeployerRaw is an auto generated low-level Go binding around an Ethereum contract.
type V3pooldeployerRaw struct {
	Contract *V3pooldeployer // Generic contract binding to access the raw methods on
}

// V3pooldeployerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type V3pooldeployerCallerRaw struct {
	Contract *V3pooldeployerCaller // Generic read-only contract binding to access the raw methods on
}

// V3pooldeployerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type V3pooldeployerTransactorRaw struct {
	Contract *V3pooldeployerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewV3pooldeployer creates a new instance of V3pooldeployer, bound to a specific deployed contract.
func NewV3pooldeployer(address common.Address, backend bind.ContractBackend) (*V3pooldeployer, error) {
	contract, err := bindV3pooldeployer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &V3pooldeployer{V3pooldeployerCaller: V3pooldeployerCaller{contract: contract}, V3pooldeployerTransactor: V3pooldeployerTransactor{contract: contract}, V3pooldeployerFilterer: V3pooldeployerFilterer{contract: contract}}, nil
}

// NewV3pooldeployerCaller creates a new read-only instance of V3pooldeployer, bound to a specific deployed contract.
func NewV3pooldeployerCaller(address common.Address, caller bind.ContractCaller) (*V3pooldeployerCaller, error) {
	contract, err := bindV3pooldeployer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &V3pooldeployerCaller{contract: contract}, nil
}

// NewV3pooldeployerTransactor creates a new write-only instance of V3pooldeployer, bound to a specific deployed contract.
func NewV3pooldeployerTransactor(address common.Address, transactor bind.ContractTransactor) (*V3pooldeployerTransactor, error) {
	contract, err := bindV3pooldeployer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &V3pooldeployerTransactor{contract: contract}, nil
}

// NewV3pooldeployerFilterer creates a new log filterer instance of V3pooldeployer, bound to a specific deployed contract.
func NewV3pooldeployerFilterer(address common.Address, filterer bind.ContractFilterer) (*V3pooldeployerFilterer, error) {
	contract, err := bindV3pooldeployer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &V3pooldeployerFilterer{contract: contract}, nil
}

// bindV3pooldeployer binds a generic wrapper to an already deployed contract.
func bindV3pooldeployer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(V3pooldeployerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3pooldeployer *V3pooldeployerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3pooldeployer.Contract.V3pooldeployerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3pooldeployer *V3pooldeployerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3pooldeployer.Contract.V3pooldeployerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3pooldeployer *V3pooldeployerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3pooldeployer.Contract.V3pooldeployerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_V3pooldeployer *V3pooldeployerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _V3pooldeployer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_V3pooldeployer *V3pooldeployerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _V3pooldeployer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_V3pooldeployer *V3pooldeployerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _V3pooldeployer.Contract.contract.Transact(opts, method, params...)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_V3pooldeployer *V3pooldeployerCaller) Parameters(opts *bind.CallOpts) (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	var out []interface{}
	err := _V3pooldeployer.contract.Call(opts, &out, "parameters")

	outstruct := new(struct {
		Factory     common.Address
		Token0      common.Address
		Token1      common.Address
		Fee         *big.Int
		TickSpacing *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Factory = out[0].(common.Address)
	outstruct.Token0 = out[1].(common.Address)
	outstruct.Token1 = out[2].(common.Address)
	outstruct.Fee = out[3].(*big.Int)
	outstruct.TickSpacing = out[4].(*big.Int)

	return *outstruct, err

}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_V3pooldeployer *V3pooldeployerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _V3pooldeployer.Contract.Parameters(&_V3pooldeployer.CallOpts)
}

// Parameters is a free data retrieval call binding the contract method 0x89035730.
//
// Solidity: function parameters() view returns(address factory, address token0, address token1, uint24 fee, int24 tickSpacing)
func (_V3pooldeployer *V3pooldeployerCallerSession) Parameters() (struct {
	Factory     common.Address
	Token0      common.Address
	Token1      common.Address
	Fee         *big.Int
	TickSpacing *big.Int
}, error) {
	return _V3pooldeployer.Contract.Parameters(&_V3pooldeployer.CallOpts)
}
