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
const IOptionPoolABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premiumCost\",\"type\":\"uint256\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"PremiumClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountCollateral\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premiumSettled\",\"type\":\"uint256\"}],\"name\":\"PremiumSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProfitsClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"profitsSettled\",\"type\":\"uint256\"}],\"name\":\"ProfitsSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"SettleLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sigma\",\"type\":\"uint256\"}],\"name\":\"SigmaSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sigma\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"SigmaUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NWA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newSigma\",\"type\":\"uint16\"}],\"name\":\"adjustSigma\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIOption\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkOPA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"opa\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkPremium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"checkProfits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profits\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOPA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimProfits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentUtilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextUpdateTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listOptions\",\"outputs\":[{\"internalType\":\"contractIOption[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOption\",\"name\":\"optionContract\",\"type\":\"address\"}],\"name\":\"optionsLeft\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pausePooler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIOption\",\"name\":\"optionContract\",\"type\":\"address\"}],\"name\":\"premiumCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"maxrate\",\"type\":\"uint8\"}],\"name\":\"setMaxUtilizationRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"OPAToken_\",\"type\":\"address\"}],\"name\":\"setOPAToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolManager\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"rate\",\"type\":\"uint8\"}],\"name\":\"setUtilizationRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"settleBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"settlePooler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseBuyer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpausePooler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// IOptionPoolBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the IOptionPool contract.
type IOptionPoolBuyIterator struct {
	Event *IOptionPoolBuy // Event containing the contract specifics and raw log

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
func (it *IOptionPoolBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolBuy)
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
		it.Event = new(IOptionPoolBuy)
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
func (it *IOptionPoolBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolBuy represents a Buy event raised by the IOptionPool contract.
type IOptionPoolBuy struct {
	Account        common.Address
	OptionContract common.Address
	Round          *big.Int
	Amount         *big.Int
	PremiumCost    *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0x00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4.
//
// Solidity: event Buy(address indexed account, address indexed optionContract, uint256 round, uint256 amount, uint256 premiumCost)
func (_IOptionPool *IOptionPoolFilterer) FilterBuy(opts *bind.FilterOpts, account []common.Address, optionContract []common.Address) (*IOptionPoolBuyIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "Buy", accountRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolBuyIterator{contract: _IOptionPool.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0x00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4.
//
// Solidity: event Buy(address indexed account, address indexed optionContract, uint256 round, uint256 amount, uint256 premiumCost)
func (_IOptionPool *IOptionPoolFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *IOptionPoolBuy, account []common.Address, optionContract []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "Buy", accountRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolBuy)
				if err := _IOptionPool.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0x00f93dbdb72854b6b6fb35433086556f2635fc83c37080c667496fecfa650fb4.
//
// Solidity: event Buy(address indexed account, address indexed optionContract, uint256 round, uint256 amount, uint256 premiumCost)
func (_IOptionPool *IOptionPoolFilterer) ParseBuy(log types.Log) (*IOptionPoolBuy, error) {
	event := new(IOptionPoolBuy)
	if err := _IOptionPool.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the IOptionPool contract.
type IOptionPoolDepositIterator struct {
	Event *IOptionPoolDeposit // Event containing the contract specifics and raw log

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
func (it *IOptionPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolDeposit)
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
		it.Event = new(IOptionPoolDeposit)
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
func (it *IOptionPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolDeposit represents a Deposit event raised by the IOptionPool contract.
type IOptionPoolDeposit struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) FilterDeposit(opts *bind.FilterOpts, account []common.Address) (*IOptionPoolDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolDepositIterator{contract: _IOptionPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *IOptionPoolDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "Deposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolDeposit)
				if err := _IOptionPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) ParseDeposit(log types.Log) (*IOptionPoolDeposit, error) {
	event := new(IOptionPoolDeposit)
	if err := _IOptionPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IOptionPool contract.
type IOptionPoolOwnershipTransferredIterator struct {
	Event *IOptionPoolOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IOptionPoolOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolOwnershipTransferred)
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
		it.Event = new(IOptionPoolOwnershipTransferred)
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
func (it *IOptionPoolOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolOwnershipTransferred represents a OwnershipTransferred event raised by the IOptionPool contract.
type IOptionPoolOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IOptionPool *IOptionPoolFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IOptionPoolOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolOwnershipTransferredIterator{contract: _IOptionPool.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IOptionPool *IOptionPoolFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IOptionPoolOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolOwnershipTransferred)
				if err := _IOptionPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IOptionPool *IOptionPoolFilterer) ParseOwnershipTransferred(log types.Log) (*IOptionPoolOwnershipTransferred, error) {
	event := new(IOptionPoolOwnershipTransferred)
	if err := _IOptionPool.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolPremiumClaimIterator is returned from FilterPremiumClaim and is used to iterate over the raw logs and unpacked data for PremiumClaim events raised by the IOptionPool contract.
type IOptionPoolPremiumClaimIterator struct {
	Event *IOptionPoolPremiumClaim // Event containing the contract specifics and raw log

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
func (it *IOptionPoolPremiumClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolPremiumClaim)
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
		it.Event = new(IOptionPoolPremiumClaim)
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
func (it *IOptionPoolPremiumClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolPremiumClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolPremiumClaim represents a PremiumClaim event raised by the IOptionPool contract.
type IOptionPoolPremiumClaim struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPremiumClaim is a free log retrieval operation binding the contract event 0xa3cf67a6fbca8be8363258a2b1374651d27418f73c29131fafa1d1be4c31754c.
//
// Solidity: event PremiumClaim(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) FilterPremiumClaim(opts *bind.FilterOpts, account []common.Address) (*IOptionPoolPremiumClaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "PremiumClaim", accountRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolPremiumClaimIterator{contract: _IOptionPool.contract, event: "PremiumClaim", logs: logs, sub: sub}, nil
}

// WatchPremiumClaim is a free log subscription operation binding the contract event 0xa3cf67a6fbca8be8363258a2b1374651d27418f73c29131fafa1d1be4c31754c.
//
// Solidity: event PremiumClaim(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) WatchPremiumClaim(opts *bind.WatchOpts, sink chan<- *IOptionPoolPremiumClaim, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "PremiumClaim", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolPremiumClaim)
				if err := _IOptionPool.contract.UnpackLog(event, "PremiumClaim", log); err != nil {
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

// ParsePremiumClaim is a log parse operation binding the contract event 0xa3cf67a6fbca8be8363258a2b1374651d27418f73c29131fafa1d1be4c31754c.
//
// Solidity: event PremiumClaim(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) ParsePremiumClaim(log types.Log) (*IOptionPoolPremiumClaim, error) {
	event := new(IOptionPoolPremiumClaim)
	if err := _IOptionPool.contract.UnpackLog(event, "PremiumClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolPremiumSettledIterator is returned from FilterPremiumSettled and is used to iterate over the raw logs and unpacked data for PremiumSettled events raised by the IOptionPool contract.
type IOptionPoolPremiumSettledIterator struct {
	Event *IOptionPoolPremiumSettled // Event containing the contract specifics and raw log

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
func (it *IOptionPoolPremiumSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolPremiumSettled)
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
		it.Event = new(IOptionPoolPremiumSettled)
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
func (it *IOptionPoolPremiumSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolPremiumSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolPremiumSettled represents a PremiumSettled event raised by the IOptionPool contract.
type IOptionPoolPremiumSettled struct {
	Account           common.Address
	AccountCollateral *big.Int
	PremiumSettled    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterPremiumSettled is a free log retrieval operation binding the contract event 0xdffb4674a87003aa761a6d45ac70251ed9f719d76b6b430731508c0bf309a7ee.
//
// Solidity: event PremiumSettled(address indexed account, uint256 accountCollateral, uint256 premiumSettled)
func (_IOptionPool *IOptionPoolFilterer) FilterPremiumSettled(opts *bind.FilterOpts, account []common.Address) (*IOptionPoolPremiumSettledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "PremiumSettled", accountRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolPremiumSettledIterator{contract: _IOptionPool.contract, event: "PremiumSettled", logs: logs, sub: sub}, nil
}

// WatchPremiumSettled is a free log subscription operation binding the contract event 0xdffb4674a87003aa761a6d45ac70251ed9f719d76b6b430731508c0bf309a7ee.
//
// Solidity: event PremiumSettled(address indexed account, uint256 accountCollateral, uint256 premiumSettled)
func (_IOptionPool *IOptionPoolFilterer) WatchPremiumSettled(opts *bind.WatchOpts, sink chan<- *IOptionPoolPremiumSettled, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "PremiumSettled", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolPremiumSettled)
				if err := _IOptionPool.contract.UnpackLog(event, "PremiumSettled", log); err != nil {
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

// ParsePremiumSettled is a log parse operation binding the contract event 0xdffb4674a87003aa761a6d45ac70251ed9f719d76b6b430731508c0bf309a7ee.
//
// Solidity: event PremiumSettled(address indexed account, uint256 accountCollateral, uint256 premiumSettled)
func (_IOptionPool *IOptionPoolFilterer) ParsePremiumSettled(log types.Log) (*IOptionPoolPremiumSettled, error) {
	event := new(IOptionPoolPremiumSettled)
	if err := _IOptionPool.contract.UnpackLog(event, "PremiumSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolProfitsClaimIterator is returned from FilterProfitsClaim and is used to iterate over the raw logs and unpacked data for ProfitsClaim events raised by the IOptionPool contract.
type IOptionPoolProfitsClaimIterator struct {
	Event *IOptionPoolProfitsClaim // Event containing the contract specifics and raw log

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
func (it *IOptionPoolProfitsClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolProfitsClaim)
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
		it.Event = new(IOptionPoolProfitsClaim)
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
func (it *IOptionPoolProfitsClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolProfitsClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolProfitsClaim represents a ProfitsClaim event raised by the IOptionPool contract.
type IOptionPoolProfitsClaim struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProfitsClaim is a free log retrieval operation binding the contract event 0xcd633c8447bd8df29b34e850deee0899eb6969c36309ea71fc3dd14795021fe0.
//
// Solidity: event ProfitsClaim(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) FilterProfitsClaim(opts *bind.FilterOpts, account []common.Address) (*IOptionPoolProfitsClaimIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "ProfitsClaim", accountRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolProfitsClaimIterator{contract: _IOptionPool.contract, event: "ProfitsClaim", logs: logs, sub: sub}, nil
}

// WatchProfitsClaim is a free log subscription operation binding the contract event 0xcd633c8447bd8df29b34e850deee0899eb6969c36309ea71fc3dd14795021fe0.
//
// Solidity: event ProfitsClaim(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) WatchProfitsClaim(opts *bind.WatchOpts, sink chan<- *IOptionPoolProfitsClaim, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "ProfitsClaim", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolProfitsClaim)
				if err := _IOptionPool.contract.UnpackLog(event, "ProfitsClaim", log); err != nil {
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

// ParseProfitsClaim is a log parse operation binding the contract event 0xcd633c8447bd8df29b34e850deee0899eb6969c36309ea71fc3dd14795021fe0.
//
// Solidity: event ProfitsClaim(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) ParseProfitsClaim(log types.Log) (*IOptionPoolProfitsClaim, error) {
	event := new(IOptionPoolProfitsClaim)
	if err := _IOptionPool.contract.UnpackLog(event, "ProfitsClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolProfitsSettledIterator is returned from FilterProfitsSettled and is used to iterate over the raw logs and unpacked data for ProfitsSettled events raised by the IOptionPool contract.
type IOptionPoolProfitsSettledIterator struct {
	Event *IOptionPoolProfitsSettled // Event containing the contract specifics and raw log

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
func (it *IOptionPoolProfitsSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolProfitsSettled)
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
		it.Event = new(IOptionPoolProfitsSettled)
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
func (it *IOptionPoolProfitsSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolProfitsSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolProfitsSettled represents a ProfitsSettled event raised by the IOptionPool contract.
type IOptionPoolProfitsSettled struct {
	Account        common.Address
	OptionContract common.Address
	Round          *big.Int
	ProfitsSettled *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProfitsSettled is a free log retrieval operation binding the contract event 0x0910ab8c8186073fe965240d819ee831f91f9f5850a906cb8e6a3e0f001bc2f9.
//
// Solidity: event ProfitsSettled(address indexed account, address indexed optionContract, uint256 round, uint256 profitsSettled)
func (_IOptionPool *IOptionPoolFilterer) FilterProfitsSettled(opts *bind.FilterOpts, account []common.Address, optionContract []common.Address) (*IOptionPoolProfitsSettledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "ProfitsSettled", accountRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolProfitsSettledIterator{contract: _IOptionPool.contract, event: "ProfitsSettled", logs: logs, sub: sub}, nil
}

// WatchProfitsSettled is a free log subscription operation binding the contract event 0x0910ab8c8186073fe965240d819ee831f91f9f5850a906cb8e6a3e0f001bc2f9.
//
// Solidity: event ProfitsSettled(address indexed account, address indexed optionContract, uint256 round, uint256 profitsSettled)
func (_IOptionPool *IOptionPoolFilterer) WatchProfitsSettled(opts *bind.WatchOpts, sink chan<- *IOptionPoolProfitsSettled, account []common.Address, optionContract []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "ProfitsSettled", accountRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolProfitsSettled)
				if err := _IOptionPool.contract.UnpackLog(event, "ProfitsSettled", log); err != nil {
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

// ParseProfitsSettled is a log parse operation binding the contract event 0x0910ab8c8186073fe965240d819ee831f91f9f5850a906cb8e6a3e0f001bc2f9.
//
// Solidity: event ProfitsSettled(address indexed account, address indexed optionContract, uint256 round, uint256 profitsSettled)
func (_IOptionPool *IOptionPoolFilterer) ParseProfitsSettled(log types.Log) (*IOptionPoolProfitsSettled, error) {
	event := new(IOptionPoolProfitsSettled)
	if err := _IOptionPool.contract.UnpackLog(event, "ProfitsSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolSettleLogIterator is returned from FilterSettleLog and is used to iterate over the raw logs and unpacked data for SettleLog events raised by the IOptionPool contract.
type IOptionPoolSettleLogIterator struct {
	Event *IOptionPoolSettleLog // Event containing the contract specifics and raw log

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
func (it *IOptionPoolSettleLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolSettleLog)
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
		it.Event = new(IOptionPoolSettleLog)
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
func (it *IOptionPoolSettleLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolSettleLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolSettleLog represents a SettleLog event raised by the IOptionPool contract.
type IOptionPoolSettleLog struct {
	OptionContract common.Address
	Round          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSettleLog is a free log retrieval operation binding the contract event 0xe912251a7f354a1f885407986a1b27d0e00bf0ca7ea6cce88699e894a24b663b.
//
// Solidity: event SettleLog(address indexed optionContract, uint256 round)
func (_IOptionPool *IOptionPoolFilterer) FilterSettleLog(opts *bind.FilterOpts, optionContract []common.Address) (*IOptionPoolSettleLogIterator, error) {

	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "SettleLog", optionContractRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolSettleLogIterator{contract: _IOptionPool.contract, event: "SettleLog", logs: logs, sub: sub}, nil
}

// WatchSettleLog is a free log subscription operation binding the contract event 0xe912251a7f354a1f885407986a1b27d0e00bf0ca7ea6cce88699e894a24b663b.
//
// Solidity: event SettleLog(address indexed optionContract, uint256 round)
func (_IOptionPool *IOptionPoolFilterer) WatchSettleLog(opts *bind.WatchOpts, sink chan<- *IOptionPoolSettleLog, optionContract []common.Address) (event.Subscription, error) {

	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "SettleLog", optionContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolSettleLog)
				if err := _IOptionPool.contract.UnpackLog(event, "SettleLog", log); err != nil {
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

// ParseSettleLog is a log parse operation binding the contract event 0xe912251a7f354a1f885407986a1b27d0e00bf0ca7ea6cce88699e894a24b663b.
//
// Solidity: event SettleLog(address indexed optionContract, uint256 round)
func (_IOptionPool *IOptionPoolFilterer) ParseSettleLog(log types.Log) (*IOptionPoolSettleLog, error) {
	event := new(IOptionPoolSettleLog)
	if err := _IOptionPool.contract.UnpackLog(event, "SettleLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolSigmaSetIterator is returned from FilterSigmaSet and is used to iterate over the raw logs and unpacked data for SigmaSet events raised by the IOptionPool contract.
type IOptionPoolSigmaSetIterator struct {
	Event *IOptionPoolSigmaSet // Event containing the contract specifics and raw log

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
func (it *IOptionPoolSigmaSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolSigmaSet)
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
		it.Event = new(IOptionPoolSigmaSet)
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
func (it *IOptionPoolSigmaSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolSigmaSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolSigmaSet represents a SigmaSet event raised by the IOptionPool contract.
type IOptionPoolSigmaSet struct {
	Sigma *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSigmaSet is a free log retrieval operation binding the contract event 0xc93b1126a4bf5c02c80817951d87454debac4f5ce8377394b0132eaee78ffaff.
//
// Solidity: event SigmaSet(uint256 sigma)
func (_IOptionPool *IOptionPoolFilterer) FilterSigmaSet(opts *bind.FilterOpts) (*IOptionPoolSigmaSetIterator, error) {

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "SigmaSet")
	if err != nil {
		return nil, err
	}
	return &IOptionPoolSigmaSetIterator{contract: _IOptionPool.contract, event: "SigmaSet", logs: logs, sub: sub}, nil
}

// WatchSigmaSet is a free log subscription operation binding the contract event 0xc93b1126a4bf5c02c80817951d87454debac4f5ce8377394b0132eaee78ffaff.
//
// Solidity: event SigmaSet(uint256 sigma)
func (_IOptionPool *IOptionPoolFilterer) WatchSigmaSet(opts *bind.WatchOpts, sink chan<- *IOptionPoolSigmaSet) (event.Subscription, error) {

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "SigmaSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolSigmaSet)
				if err := _IOptionPool.contract.UnpackLog(event, "SigmaSet", log); err != nil {
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

// ParseSigmaSet is a log parse operation binding the contract event 0xc93b1126a4bf5c02c80817951d87454debac4f5ce8377394b0132eaee78ffaff.
//
// Solidity: event SigmaSet(uint256 sigma)
func (_IOptionPool *IOptionPoolFilterer) ParseSigmaSet(log types.Log) (*IOptionPoolSigmaSet, error) {
	event := new(IOptionPoolSigmaSet)
	if err := _IOptionPool.contract.UnpackLog(event, "SigmaSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolSigmaUpdateIterator is returned from FilterSigmaUpdate and is used to iterate over the raw logs and unpacked data for SigmaUpdate events raised by the IOptionPool contract.
type IOptionPoolSigmaUpdateIterator struct {
	Event *IOptionPoolSigmaUpdate // Event containing the contract specifics and raw log

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
func (it *IOptionPoolSigmaUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolSigmaUpdate)
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
		it.Event = new(IOptionPoolSigmaUpdate)
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
func (it *IOptionPoolSigmaUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolSigmaUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolSigmaUpdate represents a SigmaUpdate event raised by the IOptionPool contract.
type IOptionPoolSigmaUpdate struct {
	Sigma *big.Int
	Rate  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSigmaUpdate is a free log retrieval operation binding the contract event 0x826623d38d576410a772393eb082789d6b5f1aa92739907fa948c4f7edd45477.
//
// Solidity: event SigmaUpdate(uint256 sigma, uint256 rate)
func (_IOptionPool *IOptionPoolFilterer) FilterSigmaUpdate(opts *bind.FilterOpts) (*IOptionPoolSigmaUpdateIterator, error) {

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "SigmaUpdate")
	if err != nil {
		return nil, err
	}
	return &IOptionPoolSigmaUpdateIterator{contract: _IOptionPool.contract, event: "SigmaUpdate", logs: logs, sub: sub}, nil
}

// WatchSigmaUpdate is a free log subscription operation binding the contract event 0x826623d38d576410a772393eb082789d6b5f1aa92739907fa948c4f7edd45477.
//
// Solidity: event SigmaUpdate(uint256 sigma, uint256 rate)
func (_IOptionPool *IOptionPoolFilterer) WatchSigmaUpdate(opts *bind.WatchOpts, sink chan<- *IOptionPoolSigmaUpdate) (event.Subscription, error) {

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "SigmaUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolSigmaUpdate)
				if err := _IOptionPool.contract.UnpackLog(event, "SigmaUpdate", log); err != nil {
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

// ParseSigmaUpdate is a log parse operation binding the contract event 0x826623d38d576410a772393eb082789d6b5f1aa92739907fa948c4f7edd45477.
//
// Solidity: event SigmaUpdate(uint256 sigma, uint256 rate)
func (_IOptionPool *IOptionPoolFilterer) ParseSigmaUpdate(log types.Log) (*IOptionPoolSigmaUpdate, error) {
	event := new(IOptionPoolSigmaUpdate)
	if err := _IOptionPool.contract.UnpackLog(event, "SigmaUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the IOptionPool contract.
type IOptionPoolWithdrawIterator struct {
	Event *IOptionPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *IOptionPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionPoolWithdraw)
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
		it.Event = new(IOptionPoolWithdraw)
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
func (it *IOptionPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionPoolWithdraw represents a Withdraw event raised by the IOptionPool contract.
type IOptionPoolWithdraw struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, account []common.Address) (*IOptionPoolWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.FilterLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &IOptionPoolWithdrawIterator{contract: _IOptionPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *IOptionPoolWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _IOptionPool.contract.WatchLogs(opts, "Withdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionPoolWithdraw)
				if err := _IOptionPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed account, uint256 amount)
func (_IOptionPool *IOptionPoolFilterer) ParseWithdraw(log types.Log) (*IOptionPoolWithdraw, error) {
	event := new(IOptionPoolWithdraw)
	if err := _IOptionPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
