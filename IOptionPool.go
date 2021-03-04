// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// IOptionPoolABI is the input ABI used to generate the binding from.
const IOptionPoolABI = "[{\"inputs\":[],\"name\":\"NWA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newSigma\",\"type\":\"uint16\"}],\"name\":\"adjustSigma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIOption\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkOPA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"opa\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkPremium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkProfits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profits\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOPA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimProfits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentUtilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listOptions\",\"outputs\":[{\"internalType\":\"contractIOption[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOption\",\"name\":\"optionContract\",\"type\":\"address\"}],\"name\":\"optionsLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausePooler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIOption\",\"name\":\"optionContract\",\"type\":\"address\"}],\"name\":\"premiumCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxrate\",\"type\":\"uint8\"}],\"name\":\"setMaxUtilizationRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"OPAToken_\",\"type\":\"address\"}],\"name\":\"setOPAToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolManager\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rate\",\"type\":\"uint8\"}],\"name\":\"setUtilizationRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"settleBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"settlePooler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpausePooler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOptionPool is an auto generated Go binding around an Ethereum contract.
type IOptionPool struct {
	IOptionPoolCaller     // Read-only binding to the contract
	IOptionPoolTransactor // Write-only binding to the contract
	IOptionPoolFilterer   // Log filterer for contract events
}

// IOptionPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOptionPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOptionPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOptionPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOptionPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOptionPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOptionPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOptionPoolSession struct {
	Contract     *IOptionPool      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOptionPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOptionPoolCallerSession struct {
	Contract *IOptionPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IOptionPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOptionPoolTransactorSession struct {
	Contract     *IOptionPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IOptionPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOptionPoolRaw struct {
	Contract *IOptionPool // Generic contract binding to access the raw methods on
}

// IOptionPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOptionPoolCallerRaw struct {
	Contract *IOptionPoolCaller // Generic read-only contract binding to access the raw methods on
}

// IOptionPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOptionPoolTransactorRaw struct {
	Contract *IOptionPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOptionPool creates a new instance of IOptionPool, bound to a specific deployed contract.
func NewIOptionPool(address common.Address, backend bind.ContractBackend) (*IOptionPool, error) {
	contract, err := bindIOptionPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOptionPool{IOptionPoolCaller: IOptionPoolCaller{contract: contract}, IOptionPoolTransactor: IOptionPoolTransactor{contract: contract}, IOptionPoolFilterer: IOptionPoolFilterer{contract: contract}}, nil
}

// NewIOptionPoolCaller creates a new read-only instance of IOptionPool, bound to a specific deployed contract.
func NewIOptionPoolCaller(address common.Address, caller bind.ContractCaller) (*IOptionPoolCaller, error) {
	contract, err := bindIOptionPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolCaller{contract: contract}, nil
}

// NewIOptionPoolTransactor creates a new write-only instance of IOptionPool, bound to a specific deployed contract.
func NewIOptionPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*IOptionPoolTransactor, error) {
	contract, err := bindIOptionPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolTransactor{contract: contract}, nil
}

// NewIOptionPoolFilterer creates a new log filterer instance of IOptionPool, bound to a specific deployed contract.
func NewIOptionPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*IOptionPoolFilterer, error) {
	contract, err := bindIOptionPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolFilterer{contract: contract}, nil
}

// bindIOptionPool binds a generic wrapper to an already deployed contract.
func bindIOptionPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOptionPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOptionPool *IOptionPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOptionPool.Contract.IOptionPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOptionPool *IOptionPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.Contract.IOptionPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOptionPool *IOptionPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOptionPool.Contract.IOptionPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOptionPool *IOptionPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOptionPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOptionPool *IOptionPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOptionPool *IOptionPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOptionPool.Contract.contract.Transact(opts, method, params...)
}

// NWA is a free data retrieval call binding the contract method 0xf73c5b29.
//
// Solidity: function NWA() view returns(uint256)
func (_IOptionPool *IOptionPoolCaller) NWA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "NWA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NWA is a free data retrieval call binding the contract method 0xf73c5b29.
//
// Solidity: function NWA() view returns(uint256)
func (_IOptionPool *IOptionPoolSession) NWA() (*big.Int, error) {
	return _IOptionPool.Contract.NWA(&_IOptionPool.CallOpts)
}

// NWA is a free data retrieval call binding the contract method 0xf73c5b29.
//
// Solidity: function NWA() view returns(uint256)
func (_IOptionPool *IOptionPoolCallerSession) NWA() (*big.Int, error) {
	return _IOptionPool.Contract.NWA(&_IOptionPool.CallOpts)
}

// CheckOPA is a free data retrieval call binding the contract method 0x752ce231.
//
// Solidity: function checkOPA(address account) view returns(uint256 opa)
func (_IOptionPool *IOptionPoolCaller) CheckOPA(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "checkOPA", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckOPA is a free data retrieval call binding the contract method 0x752ce231.
//
// Solidity: function checkOPA(address account) view returns(uint256 opa)
func (_IOptionPool *IOptionPoolSession) CheckOPA(account common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.CheckOPA(&_IOptionPool.CallOpts, account)
}

// CheckOPA is a free data retrieval call binding the contract method 0x752ce231.
//
// Solidity: function checkOPA(address account) view returns(uint256 opa)
func (_IOptionPool *IOptionPoolCallerSession) CheckOPA(account common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.CheckOPA(&_IOptionPool.CallOpts, account)
}

// CheckPremium is a free data retrieval call binding the contract method 0xc9dd915f.
//
// Solidity: function checkPremium(address account) view returns(uint256 premium)
func (_IOptionPool *IOptionPoolCaller) CheckPremium(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "checkPremium", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckPremium is a free data retrieval call binding the contract method 0xc9dd915f.
//
// Solidity: function checkPremium(address account) view returns(uint256 premium)
func (_IOptionPool *IOptionPoolSession) CheckPremium(account common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.CheckPremium(&_IOptionPool.CallOpts, account)
}

// CheckPremium is a free data retrieval call binding the contract method 0xc9dd915f.
//
// Solidity: function checkPremium(address account) view returns(uint256 premium)
func (_IOptionPool *IOptionPoolCallerSession) CheckPremium(account common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.CheckPremium(&_IOptionPool.CallOpts, account)
}

// CheckProfits is a free data retrieval call binding the contract method 0xa7ba8360.
//
// Solidity: function checkProfits(address account) view returns(uint256 profits)
func (_IOptionPool *IOptionPoolCaller) CheckProfits(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "checkProfits", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckProfits is a free data retrieval call binding the contract method 0xa7ba8360.
//
// Solidity: function checkProfits(address account) view returns(uint256 profits)
func (_IOptionPool *IOptionPoolSession) CheckProfits(account common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.CheckProfits(&_IOptionPool.CallOpts, account)
}

// CheckProfits is a free data retrieval call binding the contract method 0xa7ba8360.
//
// Solidity: function checkProfits(address account) view returns(uint256 profits)
func (_IOptionPool *IOptionPoolCallerSession) CheckProfits(account common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.CheckProfits(&_IOptionPool.CallOpts, account)
}

// CurrentUtilizationRate is a free data retrieval call binding the contract method 0x48621a70.
//
// Solidity: function currentUtilizationRate() view returns(uint256)
func (_IOptionPool *IOptionPoolCaller) CurrentUtilizationRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "currentUtilizationRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentUtilizationRate is a free data retrieval call binding the contract method 0x48621a70.
//
// Solidity: function currentUtilizationRate() view returns(uint256)
func (_IOptionPool *IOptionPoolSession) CurrentUtilizationRate() (*big.Int, error) {
	return _IOptionPool.Contract.CurrentUtilizationRate(&_IOptionPool.CallOpts)
}

// CurrentUtilizationRate is a free data retrieval call binding the contract method 0x48621a70.
//
// Solidity: function currentUtilizationRate() view returns(uint256)
func (_IOptionPool *IOptionPoolCallerSession) CurrentUtilizationRate() (*big.Int, error) {
	return _IOptionPool.Contract.CurrentUtilizationRate(&_IOptionPool.CallOpts)
}

// GetNextUpdateTime is a free data retrieval call binding the contract method 0x7696a015.
//
// Solidity: function getNextUpdateTime() view returns(uint256)
func (_IOptionPool *IOptionPoolCaller) GetNextUpdateTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "getNextUpdateTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextUpdateTime is a free data retrieval call binding the contract method 0x7696a015.
//
// Solidity: function getNextUpdateTime() view returns(uint256)
func (_IOptionPool *IOptionPoolSession) GetNextUpdateTime() (*big.Int, error) {
	return _IOptionPool.Contract.GetNextUpdateTime(&_IOptionPool.CallOpts)
}

// GetNextUpdateTime is a free data retrieval call binding the contract method 0x7696a015.
//
// Solidity: function getNextUpdateTime() view returns(uint256)
func (_IOptionPool *IOptionPoolCallerSession) GetNextUpdateTime() (*big.Int, error) {
	return _IOptionPool.Contract.GetNextUpdateTime(&_IOptionPool.CallOpts)
}

// ListOptions is a free data retrieval call binding the contract method 0xd2bfdee6.
//
// Solidity: function listOptions() view returns(address[])
func (_IOptionPool *IOptionPoolCaller) ListOptions(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "listOptions")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ListOptions is a free data retrieval call binding the contract method 0xd2bfdee6.
//
// Solidity: function listOptions() view returns(address[])
func (_IOptionPool *IOptionPoolSession) ListOptions() ([]common.Address, error) {
	return _IOptionPool.Contract.ListOptions(&_IOptionPool.CallOpts)
}

// ListOptions is a free data retrieval call binding the contract method 0xd2bfdee6.
//
// Solidity: function listOptions() view returns(address[])
func (_IOptionPool *IOptionPoolCallerSession) ListOptions() ([]common.Address, error) {
	return _IOptionPool.Contract.ListOptions(&_IOptionPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IOptionPool *IOptionPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IOptionPool *IOptionPoolSession) Name() (string, error) {
	return _IOptionPool.Contract.Name(&_IOptionPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IOptionPool *IOptionPoolCallerSession) Name() (string, error) {
	return _IOptionPool.Contract.Name(&_IOptionPool.CallOpts)
}

// OptionsLeft is a free data retrieval call binding the contract method 0xaee88a84.
//
// Solidity: function optionsLeft(address optionContract) view returns(uint256 left, uint256 round)
func (_IOptionPool *IOptionPoolCaller) OptionsLeft(opts *bind.CallOpts, optionContract common.Address) (struct {
	Left  *big.Int
	Round *big.Int
}, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "optionsLeft", optionContract)

	outstruct := new(struct {
		Left  *big.Int
		Round *big.Int
	})

	outstruct.Left = out[0].(*big.Int)
	outstruct.Round = out[1].(*big.Int)

	return *outstruct, err

}

// OptionsLeft is a free data retrieval call binding the contract method 0xaee88a84.
//
// Solidity: function optionsLeft(address optionContract) view returns(uint256 left, uint256 round)
func (_IOptionPool *IOptionPoolSession) OptionsLeft(optionContract common.Address) (struct {
	Left  *big.Int
	Round *big.Int
}, error) {
	return _IOptionPool.Contract.OptionsLeft(&_IOptionPool.CallOpts, optionContract)
}

// OptionsLeft is a free data retrieval call binding the contract method 0xaee88a84.
//
// Solidity: function optionsLeft(address optionContract) view returns(uint256 left, uint256 round)
func (_IOptionPool *IOptionPoolCallerSession) OptionsLeft(optionContract common.Address) (struct {
	Left  *big.Int
	Round *big.Int
}, error) {
	return _IOptionPool.Contract.OptionsLeft(&_IOptionPool.CallOpts, optionContract)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOptionPool *IOptionPoolCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOptionPool *IOptionPoolSession) Owner() (common.Address, error) {
	return _IOptionPool.Contract.Owner(&_IOptionPool.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IOptionPool *IOptionPoolCallerSession) Owner() (common.Address, error) {
	return _IOptionPool.Contract.Owner(&_IOptionPool.CallOpts)
}

// PremiumCost is a free data retrieval call binding the contract method 0x7a1ca686.
//
// Solidity: function premiumCost(uint256 amount, address optionContract) view returns(uint256)
func (_IOptionPool *IOptionPoolCaller) PremiumCost(opts *bind.CallOpts, amount *big.Int, optionContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOptionPool.contract.Call(opts, &out, "premiumCost", amount, optionContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PremiumCost is a free data retrieval call binding the contract method 0x7a1ca686.
//
// Solidity: function premiumCost(uint256 amount, address optionContract) view returns(uint256)
func (_IOptionPool *IOptionPoolSession) PremiumCost(amount *big.Int, optionContract common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.PremiumCost(&_IOptionPool.CallOpts, amount, optionContract)
}

// PremiumCost is a free data retrieval call binding the contract method 0x7a1ca686.
//
// Solidity: function premiumCost(uint256 amount, address optionContract) view returns(uint256)
func (_IOptionPool *IOptionPoolCallerSession) PremiumCost(amount *big.Int, optionContract common.Address) (*big.Int, error) {
	return _IOptionPool.Contract.PremiumCost(&_IOptionPool.CallOpts, amount, optionContract)
}

// AdjustSigma is a paid mutator transaction binding the contract method 0x9a8deab2.
//
// Solidity: function adjustSigma(uint16 newSigma) returns()
func (_IOptionPool *IOptionPoolTransactor) AdjustSigma(opts *bind.TransactOpts, newSigma uint16) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "adjustSigma", newSigma)
}

// AdjustSigma is a paid mutator transaction binding the contract method 0x9a8deab2.
//
// Solidity: function adjustSigma(uint16 newSigma) returns()
func (_IOptionPool *IOptionPoolSession) AdjustSigma(newSigma uint16) (*types.Transaction, error) {
	return _IOptionPool.Contract.AdjustSigma(&_IOptionPool.TransactOpts, newSigma)
}

// AdjustSigma is a paid mutator transaction binding the contract method 0x9a8deab2.
//
// Solidity: function adjustSigma(uint16 newSigma) returns()
func (_IOptionPool *IOptionPoolTransactorSession) AdjustSigma(newSigma uint16) (*types.Transaction, error) {
	return _IOptionPool.Contract.AdjustSigma(&_IOptionPool.TransactOpts, newSigma)
}

// Buy is a paid mutator transaction binding the contract method 0x2afaca20.
//
// Solidity: function buy(uint256 amount, address optionContract, uint256 round) returns()
func (_IOptionPool *IOptionPoolTransactor) Buy(opts *bind.TransactOpts, amount *big.Int, optionContract common.Address, round *big.Int) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "buy", amount, optionContract, round)
}

// Buy is a paid mutator transaction binding the contract method 0x2afaca20.
//
// Solidity: function buy(uint256 amount, address optionContract, uint256 round) returns()
func (_IOptionPool *IOptionPoolSession) Buy(amount *big.Int, optionContract common.Address, round *big.Int) (*types.Transaction, error) {
	return _IOptionPool.Contract.Buy(&_IOptionPool.TransactOpts, amount, optionContract, round)
}

// Buy is a paid mutator transaction binding the contract method 0x2afaca20.
//
// Solidity: function buy(uint256 amount, address optionContract, uint256 round) returns()
func (_IOptionPool *IOptionPoolTransactorSession) Buy(amount *big.Int, optionContract common.Address, round *big.Int) (*types.Transaction, error) {
	return _IOptionPool.Contract.Buy(&_IOptionPool.TransactOpts, amount, optionContract, round)
}

// ClaimOPA is a paid mutator transaction binding the contract method 0x63700916.
//
// Solidity: function claimOPA() returns()
func (_IOptionPool *IOptionPoolTransactor) ClaimOPA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "claimOPA")
}

// ClaimOPA is a paid mutator transaction binding the contract method 0x63700916.
//
// Solidity: function claimOPA() returns()
func (_IOptionPool *IOptionPoolSession) ClaimOPA() (*types.Transaction, error) {
	return _IOptionPool.Contract.ClaimOPA(&_IOptionPool.TransactOpts)
}

// ClaimOPA is a paid mutator transaction binding the contract method 0x63700916.
//
// Solidity: function claimOPA() returns()
func (_IOptionPool *IOptionPoolTransactorSession) ClaimOPA() (*types.Transaction, error) {
	return _IOptionPool.Contract.ClaimOPA(&_IOptionPool.TransactOpts)
}

// ClaimPremium is a paid mutator transaction binding the contract method 0x529663db.
//
// Solidity: function claimPremium() returns()
func (_IOptionPool *IOptionPoolTransactor) ClaimPremium(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "claimPremium")
}

// ClaimPremium is a paid mutator transaction binding the contract method 0x529663db.
//
// Solidity: function claimPremium() returns()
func (_IOptionPool *IOptionPoolSession) ClaimPremium() (*types.Transaction, error) {
	return _IOptionPool.Contract.ClaimPremium(&_IOptionPool.TransactOpts)
}

// ClaimPremium is a paid mutator transaction binding the contract method 0x529663db.
//
// Solidity: function claimPremium() returns()
func (_IOptionPool *IOptionPoolTransactorSession) ClaimPremium() (*types.Transaction, error) {
	return _IOptionPool.Contract.ClaimPremium(&_IOptionPool.TransactOpts)
}

// ClaimProfits is a paid mutator transaction binding the contract method 0x9df51b89.
//
// Solidity: function claimProfits() returns()
func (_IOptionPool *IOptionPoolTransactor) ClaimProfits(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "claimProfits")
}

// ClaimProfits is a paid mutator transaction binding the contract method 0x9df51b89.
//
// Solidity: function claimProfits() returns()
func (_IOptionPool *IOptionPoolSession) ClaimProfits() (*types.Transaction, error) {
	return _IOptionPool.Contract.ClaimProfits(&_IOptionPool.TransactOpts)
}

// ClaimProfits is a paid mutator transaction binding the contract method 0x9df51b89.
//
// Solidity: function claimProfits() returns()
func (_IOptionPool *IOptionPoolTransactorSession) ClaimProfits() (*types.Transaction, error) {
	return _IOptionPool.Contract.ClaimProfits(&_IOptionPool.TransactOpts)
}

// PauseBuyer is a paid mutator transaction binding the contract method 0xc876ce68.
//
// Solidity: function pauseBuyer() returns()
func (_IOptionPool *IOptionPoolTransactor) PauseBuyer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "pauseBuyer")
}

// PauseBuyer is a paid mutator transaction binding the contract method 0xc876ce68.
//
// Solidity: function pauseBuyer() returns()
func (_IOptionPool *IOptionPoolSession) PauseBuyer() (*types.Transaction, error) {
	return _IOptionPool.Contract.PauseBuyer(&_IOptionPool.TransactOpts)
}

// PauseBuyer is a paid mutator transaction binding the contract method 0xc876ce68.
//
// Solidity: function pauseBuyer() returns()
func (_IOptionPool *IOptionPoolTransactorSession) PauseBuyer() (*types.Transaction, error) {
	return _IOptionPool.Contract.PauseBuyer(&_IOptionPool.TransactOpts)
}

// PausePooler is a paid mutator transaction binding the contract method 0xbbff2f76.
//
// Solidity: function pausePooler() returns()
func (_IOptionPool *IOptionPoolTransactor) PausePooler(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "pausePooler")
}

// PausePooler is a paid mutator transaction binding the contract method 0xbbff2f76.
//
// Solidity: function pausePooler() returns()
func (_IOptionPool *IOptionPoolSession) PausePooler() (*types.Transaction, error) {
	return _IOptionPool.Contract.PausePooler(&_IOptionPool.TransactOpts)
}

// PausePooler is a paid mutator transaction binding the contract method 0xbbff2f76.
//
// Solidity: function pausePooler() returns()
func (_IOptionPool *IOptionPoolTransactorSession) PausePooler() (*types.Transaction, error) {
	return _IOptionPool.Contract.PausePooler(&_IOptionPool.TransactOpts)
}

// SetMaxUtilizationRate is a paid mutator transaction binding the contract method 0x9d3627c6.
//
// Solidity: function setMaxUtilizationRate(uint8 maxrate) returns()
func (_IOptionPool *IOptionPoolTransactor) SetMaxUtilizationRate(opts *bind.TransactOpts, maxrate uint8) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "setMaxUtilizationRate", maxrate)
}

// SetMaxUtilizationRate is a paid mutator transaction binding the contract method 0x9d3627c6.
//
// Solidity: function setMaxUtilizationRate(uint8 maxrate) returns()
func (_IOptionPool *IOptionPoolSession) SetMaxUtilizationRate(maxrate uint8) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetMaxUtilizationRate(&_IOptionPool.TransactOpts, maxrate)
}

// SetMaxUtilizationRate is a paid mutator transaction binding the contract method 0x9d3627c6.
//
// Solidity: function setMaxUtilizationRate(uint8 maxrate) returns()
func (_IOptionPool *IOptionPoolTransactorSession) SetMaxUtilizationRate(maxrate uint8) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetMaxUtilizationRate(&_IOptionPool.TransactOpts, maxrate)
}

// SetOPAToken is a paid mutator transaction binding the contract method 0x3f4e9515.
//
// Solidity: function setOPAToken(address OPAToken_) returns()
func (_IOptionPool *IOptionPoolTransactor) SetOPAToken(opts *bind.TransactOpts, OPAToken_ common.Address) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "setOPAToken", OPAToken_)
}

// SetOPAToken is a paid mutator transaction binding the contract method 0x3f4e9515.
//
// Solidity: function setOPAToken(address OPAToken_) returns()
func (_IOptionPool *IOptionPoolSession) SetOPAToken(OPAToken_ common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetOPAToken(&_IOptionPool.TransactOpts, OPAToken_)
}

// SetOPAToken is a paid mutator transaction binding the contract method 0x3f4e9515.
//
// Solidity: function setOPAToken(address OPAToken_) returns()
func (_IOptionPool *IOptionPoolTransactorSession) SetOPAToken(OPAToken_ common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetOPAToken(&_IOptionPool.TransactOpts, OPAToken_)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address poolManager) returns()
func (_IOptionPool *IOptionPoolTransactor) SetPoolManager(opts *bind.TransactOpts, poolManager common.Address) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "setPoolManager", poolManager)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address poolManager) returns()
func (_IOptionPool *IOptionPoolSession) SetPoolManager(poolManager common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetPoolManager(&_IOptionPool.TransactOpts, poolManager)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address poolManager) returns()
func (_IOptionPool *IOptionPoolTransactorSession) SetPoolManager(poolManager common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetPoolManager(&_IOptionPool.TransactOpts, poolManager)
}

// SetUtilizationRate is a paid mutator transaction binding the contract method 0xc6b8ffb1.
//
// Solidity: function setUtilizationRate(uint8 rate) returns()
func (_IOptionPool *IOptionPoolTransactor) SetUtilizationRate(opts *bind.TransactOpts, rate uint8) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "setUtilizationRate", rate)
}

// SetUtilizationRate is a paid mutator transaction binding the contract method 0xc6b8ffb1.
//
// Solidity: function setUtilizationRate(uint8 rate) returns()
func (_IOptionPool *IOptionPoolSession) SetUtilizationRate(rate uint8) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetUtilizationRate(&_IOptionPool.TransactOpts, rate)
}

// SetUtilizationRate is a paid mutator transaction binding the contract method 0xc6b8ffb1.
//
// Solidity: function setUtilizationRate(uint8 rate) returns()
func (_IOptionPool *IOptionPoolTransactorSession) SetUtilizationRate(rate uint8) (*types.Transaction, error) {
	return _IOptionPool.Contract.SetUtilizationRate(&_IOptionPool.TransactOpts, rate)
}

// SettleBuyer is a paid mutator transaction binding the contract method 0xc9c44038.
//
// Solidity: function settleBuyer(address account) returns()
func (_IOptionPool *IOptionPoolTransactor) SettleBuyer(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "settleBuyer", account)
}

// SettleBuyer is a paid mutator transaction binding the contract method 0xc9c44038.
//
// Solidity: function settleBuyer(address account) returns()
func (_IOptionPool *IOptionPoolSession) SettleBuyer(account common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SettleBuyer(&_IOptionPool.TransactOpts, account)
}

// SettleBuyer is a paid mutator transaction binding the contract method 0xc9c44038.
//
// Solidity: function settleBuyer(address account) returns()
func (_IOptionPool *IOptionPoolTransactorSession) SettleBuyer(account common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SettleBuyer(&_IOptionPool.TransactOpts, account)
}

// SettlePooler is a paid mutator transaction binding the contract method 0x7aa37c1c.
//
// Solidity: function settlePooler(address account) returns()
func (_IOptionPool *IOptionPoolTransactor) SettlePooler(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "settlePooler", account)
}

// SettlePooler is a paid mutator transaction binding the contract method 0x7aa37c1c.
//
// Solidity: function settlePooler(address account) returns()
func (_IOptionPool *IOptionPoolSession) SettlePooler(account common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SettlePooler(&_IOptionPool.TransactOpts, account)
}

// SettlePooler is a paid mutator transaction binding the contract method 0x7aa37c1c.
//
// Solidity: function settlePooler(address account) returns()
func (_IOptionPool *IOptionPoolTransactorSession) SettlePooler(account common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.SettlePooler(&_IOptionPool.TransactOpts, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IOptionPool *IOptionPoolTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IOptionPool *IOptionPoolSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.TransferOwnership(&_IOptionPool.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IOptionPool *IOptionPoolTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IOptionPool.Contract.TransferOwnership(&_IOptionPool.TransactOpts, newOwner)
}

// UnpauseBuyer is a paid mutator transaction binding the contract method 0x3e05d047.
//
// Solidity: function unpauseBuyer() returns()
func (_IOptionPool *IOptionPoolTransactor) UnpauseBuyer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "unpauseBuyer")
}

// UnpauseBuyer is a paid mutator transaction binding the contract method 0x3e05d047.
//
// Solidity: function unpauseBuyer() returns()
func (_IOptionPool *IOptionPoolSession) UnpauseBuyer() (*types.Transaction, error) {
	return _IOptionPool.Contract.UnpauseBuyer(&_IOptionPool.TransactOpts)
}

// UnpauseBuyer is a paid mutator transaction binding the contract method 0x3e05d047.
//
// Solidity: function unpauseBuyer() returns()
func (_IOptionPool *IOptionPoolTransactorSession) UnpauseBuyer() (*types.Transaction, error) {
	return _IOptionPool.Contract.UnpauseBuyer(&_IOptionPool.TransactOpts)
}

// UnpausePooler is a paid mutator transaction binding the contract method 0x6de133a4.
//
// Solidity: function unpausePooler() returns()
func (_IOptionPool *IOptionPoolTransactor) UnpausePooler(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "unpausePooler")
}

// UnpausePooler is a paid mutator transaction binding the contract method 0x6de133a4.
//
// Solidity: function unpausePooler() returns()
func (_IOptionPool *IOptionPoolSession) UnpausePooler() (*types.Transaction, error) {
	return _IOptionPool.Contract.UnpausePooler(&_IOptionPool.TransactOpts)
}

// UnpausePooler is a paid mutator transaction binding the contract method 0x6de133a4.
//
// Solidity: function unpausePooler() returns()
func (_IOptionPool *IOptionPoolTransactorSession) UnpausePooler() (*types.Transaction, error) {
	return _IOptionPool.Contract.UnpausePooler(&_IOptionPool.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_IOptionPool *IOptionPoolTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOptionPool.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_IOptionPool *IOptionPoolSession) Update() (*types.Transaction, error) {
	return _IOptionPool.Contract.Update(&_IOptionPool.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_IOptionPool *IOptionPoolTransactorSession) Update() (*types.Transaction, error) {
	return _IOptionPool.Contract.Update(&_IOptionPool.TransactOpts)
}
