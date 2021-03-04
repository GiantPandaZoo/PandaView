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

// IOptionABI is the input ABI used to generate the binding from.
const IOptionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountUSDT\",\"type\":\"uint256\"}],\"name\":\"addPremium\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"expiryDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundAccOPASellerShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundAccPremiumShare\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getRoundBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundExpiryDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundSettlePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundStrikePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundTotalPremiums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"}],\"name\":\"getRoundTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getSettledRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUnclaimedProfitsRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"strikePrice_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSupply\",\"type\":\"uint256\"}],\"name\":\"resetOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sellerOPAShare\",\"type\":\"uint256\"}],\"name\":\"setRoundAccOPASellerShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"premiumShare\",\"type\":\"uint256\"}],\"name\":\"setRoundAccPremiumShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setSettledRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"r\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setUnclaimedProfitsRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"strikePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPremiums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOption is an auto generated Go binding around an Ethereum contract.
type IOption struct {
	IOptionCaller     // Read-only binding to the contract
	IOptionTransactor // Write-only binding to the contract
	IOptionFilterer   // Log filterer for contract events
}

// IOptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOptionSession struct {
	Contract     *IOption          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOptionCallerSession struct {
	Contract *IOptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IOptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOptionTransactorSession struct {
	Contract     *IOptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IOptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOptionRaw struct {
	Contract *IOption // Generic contract binding to access the raw methods on
}

// IOptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOptionCallerRaw struct {
	Contract *IOptionCaller // Generic read-only contract binding to access the raw methods on
}

// IOptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOptionTransactorRaw struct {
	Contract *IOptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOption creates a new instance of IOption, bound to a specific deployed contract.
func NewIOption(address common.Address, backend bind.ContractBackend) (*IOption, error) {
	contract, err := bindIOption(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOption{IOptionCaller: IOptionCaller{contract: contract}, IOptionTransactor: IOptionTransactor{contract: contract}, IOptionFilterer: IOptionFilterer{contract: contract}}, nil
}

// NewIOptionCaller creates a new read-only instance of IOption, bound to a specific deployed contract.
func NewIOptionCaller(address common.Address, caller bind.ContractCaller) (*IOptionCaller, error) {
	contract, err := bindIOption(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOptionCaller{contract: contract}, nil
}

// NewIOptionTransactor creates a new write-only instance of IOption, bound to a specific deployed contract.
func NewIOptionTransactor(address common.Address, transactor bind.ContractTransactor) (*IOptionTransactor, error) {
	contract, err := bindIOption(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOptionTransactor{contract: contract}, nil
}

// NewIOptionFilterer creates a new log filterer instance of IOption, bound to a specific deployed contract.
func NewIOptionFilterer(address common.Address, filterer bind.ContractFilterer) (*IOptionFilterer, error) {
	contract, err := bindIOption(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOptionFilterer{contract: contract}, nil
}

// bindIOption binds a generic wrapper to an already deployed contract.
func bindIOption(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOption *IOptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOption.Contract.IOptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOption *IOptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOption.Contract.IOptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOption *IOptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOption.Contract.IOptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOption *IOptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOption.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOption *IOptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOption.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOption *IOptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOption.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IOption *IOptionCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IOption *IOptionSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IOption.Contract.Allowance(&_IOption.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IOption *IOptionCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IOption.Contract.Allowance(&_IOption.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IOption *IOptionCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IOption *IOptionSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IOption.Contract.BalanceOf(&_IOption.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IOption *IOptionCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IOption.Contract.BalanceOf(&_IOption.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IOption *IOptionCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IOption *IOptionSession) Decimals() (uint8, error) {
	return _IOption.Contract.Decimals(&_IOption.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IOption *IOptionCallerSession) Decimals() (uint8, error) {
	return _IOption.Contract.Decimals(&_IOption.CallOpts)
}

// ExpiryDate is a free data retrieval call binding the contract method 0x516dde43.
//
// Solidity: function expiryDate() view returns(uint256)
func (_IOption *IOptionCaller) ExpiryDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "expiryDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExpiryDate is a free data retrieval call binding the contract method 0x516dde43.
//
// Solidity: function expiryDate() view returns(uint256)
func (_IOption *IOptionSession) ExpiryDate() (*big.Int, error) {
	return _IOption.Contract.ExpiryDate(&_IOption.CallOpts)
}

// ExpiryDate is a free data retrieval call binding the contract method 0x516dde43.
//
// Solidity: function expiryDate() view returns(uint256)
func (_IOption *IOptionCallerSession) ExpiryDate() (*big.Int, error) {
	return _IOption.Contract.ExpiryDate(&_IOption.CallOpts)
}

// GetDuration is a free data retrieval call binding the contract method 0xad2e8c9b.
//
// Solidity: function getDuration() view returns(uint256)
func (_IOption *IOptionCaller) GetDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDuration is a free data retrieval call binding the contract method 0xad2e8c9b.
//
// Solidity: function getDuration() view returns(uint256)
func (_IOption *IOptionSession) GetDuration() (*big.Int, error) {
	return _IOption.Contract.GetDuration(&_IOption.CallOpts)
}

// GetDuration is a free data retrieval call binding the contract method 0xad2e8c9b.
//
// Solidity: function getDuration() view returns(uint256)
func (_IOption *IOptionCallerSession) GetDuration() (*big.Int, error) {
	return _IOption.Contract.GetDuration(&_IOption.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_IOption *IOptionCaller) GetPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_IOption *IOptionSession) GetPool() (common.Address, error) {
	return _IOption.Contract.GetPool(&_IOption.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_IOption *IOptionCallerSession) GetPool() (common.Address, error) {
	return _IOption.Contract.GetPool(&_IOption.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_IOption *IOptionCaller) GetRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_IOption *IOptionSession) GetRound() (*big.Int, error) {
	return _IOption.Contract.GetRound(&_IOption.CallOpts)
}

// GetRound is a free data retrieval call binding the contract method 0x9f8743f7.
//
// Solidity: function getRound() view returns(uint256)
func (_IOption *IOptionCallerSession) GetRound() (*big.Int, error) {
	return _IOption.Contract.GetRound(&_IOption.CallOpts)
}

// GetRoundAccOPASellerShare is a free data retrieval call binding the contract method 0x0bcdcb35.
//
// Solidity: function getRoundAccOPASellerShare(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundAccOPASellerShare(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundAccOPASellerShare", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundAccOPASellerShare is a free data retrieval call binding the contract method 0x0bcdcb35.
//
// Solidity: function getRoundAccOPASellerShare(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundAccOPASellerShare(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundAccOPASellerShare(&_IOption.CallOpts, r)
}

// GetRoundAccOPASellerShare is a free data retrieval call binding the contract method 0x0bcdcb35.
//
// Solidity: function getRoundAccOPASellerShare(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundAccOPASellerShare(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundAccOPASellerShare(&_IOption.CallOpts, r)
}

// GetRoundAccPremiumShare is a free data retrieval call binding the contract method 0x5c6403f8.
//
// Solidity: function getRoundAccPremiumShare(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundAccPremiumShare(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundAccPremiumShare", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundAccPremiumShare is a free data retrieval call binding the contract method 0x5c6403f8.
//
// Solidity: function getRoundAccPremiumShare(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundAccPremiumShare(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundAccPremiumShare(&_IOption.CallOpts, r)
}

// GetRoundAccPremiumShare is a free data retrieval call binding the contract method 0x5c6403f8.
//
// Solidity: function getRoundAccPremiumShare(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundAccPremiumShare(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundAccPremiumShare(&_IOption.CallOpts, r)
}

// GetRoundBalanceOf is a free data retrieval call binding the contract method 0x518ff67e.
//
// Solidity: function getRoundBalanceOf(uint256 r, address account) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundBalanceOf(opts *bind.CallOpts, r *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundBalanceOf", r, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundBalanceOf is a free data retrieval call binding the contract method 0x518ff67e.
//
// Solidity: function getRoundBalanceOf(uint256 r, address account) view returns(uint256)
func (_IOption *IOptionSession) GetRoundBalanceOf(r *big.Int, account common.Address) (*big.Int, error) {
	return _IOption.Contract.GetRoundBalanceOf(&_IOption.CallOpts, r, account)
}

// GetRoundBalanceOf is a free data retrieval call binding the contract method 0x518ff67e.
//
// Solidity: function getRoundBalanceOf(uint256 r, address account) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundBalanceOf(r *big.Int, account common.Address) (*big.Int, error) {
	return _IOption.Contract.GetRoundBalanceOf(&_IOption.CallOpts, r, account)
}

// GetRoundExpiryDate is a free data retrieval call binding the contract method 0xaf51dbc3.
//
// Solidity: function getRoundExpiryDate(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundExpiryDate(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundExpiryDate", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundExpiryDate is a free data retrieval call binding the contract method 0xaf51dbc3.
//
// Solidity: function getRoundExpiryDate(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundExpiryDate(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundExpiryDate(&_IOption.CallOpts, r)
}

// GetRoundExpiryDate is a free data retrieval call binding the contract method 0xaf51dbc3.
//
// Solidity: function getRoundExpiryDate(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundExpiryDate(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundExpiryDate(&_IOption.CallOpts, r)
}

// GetRoundSettlePrice is a free data retrieval call binding the contract method 0x96dafd98.
//
// Solidity: function getRoundSettlePrice(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundSettlePrice(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundSettlePrice", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundSettlePrice is a free data retrieval call binding the contract method 0x96dafd98.
//
// Solidity: function getRoundSettlePrice(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundSettlePrice(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundSettlePrice(&_IOption.CallOpts, r)
}

// GetRoundSettlePrice is a free data retrieval call binding the contract method 0x96dafd98.
//
// Solidity: function getRoundSettlePrice(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundSettlePrice(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundSettlePrice(&_IOption.CallOpts, r)
}

// GetRoundStrikePrice is a free data retrieval call binding the contract method 0x04fb9726.
//
// Solidity: function getRoundStrikePrice(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundStrikePrice(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundStrikePrice", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundStrikePrice is a free data retrieval call binding the contract method 0x04fb9726.
//
// Solidity: function getRoundStrikePrice(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundStrikePrice(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundStrikePrice(&_IOption.CallOpts, r)
}

// GetRoundStrikePrice is a free data retrieval call binding the contract method 0x04fb9726.
//
// Solidity: function getRoundStrikePrice(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundStrikePrice(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundStrikePrice(&_IOption.CallOpts, r)
}

// GetRoundTotalPremiums is a free data retrieval call binding the contract method 0xf27f77ff.
//
// Solidity: function getRoundTotalPremiums(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundTotalPremiums(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundTotalPremiums", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundTotalPremiums is a free data retrieval call binding the contract method 0xf27f77ff.
//
// Solidity: function getRoundTotalPremiums(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundTotalPremiums(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundTotalPremiums(&_IOption.CallOpts, r)
}

// GetRoundTotalPremiums is a free data retrieval call binding the contract method 0xf27f77ff.
//
// Solidity: function getRoundTotalPremiums(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundTotalPremiums(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundTotalPremiums(&_IOption.CallOpts, r)
}

// GetRoundTotalSupply is a free data retrieval call binding the contract method 0xb3e32158.
//
// Solidity: function getRoundTotalSupply(uint256 r) view returns(uint256)
func (_IOption *IOptionCaller) GetRoundTotalSupply(opts *bind.CallOpts, r *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getRoundTotalSupply", r)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundTotalSupply is a free data retrieval call binding the contract method 0xb3e32158.
//
// Solidity: function getRoundTotalSupply(uint256 r) view returns(uint256)
func (_IOption *IOptionSession) GetRoundTotalSupply(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundTotalSupply(&_IOption.CallOpts, r)
}

// GetRoundTotalSupply is a free data retrieval call binding the contract method 0xb3e32158.
//
// Solidity: function getRoundTotalSupply(uint256 r) view returns(uint256)
func (_IOption *IOptionCallerSession) GetRoundTotalSupply(r *big.Int) (*big.Int, error) {
	return _IOption.Contract.GetRoundTotalSupply(&_IOption.CallOpts, r)
}

// GetSettledRound is a free data retrieval call binding the contract method 0xbc37d4da.
//
// Solidity: function getSettledRound(address account) view returns(uint256)
func (_IOption *IOptionCaller) GetSettledRound(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getSettledRound", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSettledRound is a free data retrieval call binding the contract method 0xbc37d4da.
//
// Solidity: function getSettledRound(address account) view returns(uint256)
func (_IOption *IOptionSession) GetSettledRound(account common.Address) (*big.Int, error) {
	return _IOption.Contract.GetSettledRound(&_IOption.CallOpts, account)
}

// GetSettledRound is a free data retrieval call binding the contract method 0xbc37d4da.
//
// Solidity: function getSettledRound(address account) view returns(uint256)
func (_IOption *IOptionCallerSession) GetSettledRound(account common.Address) (*big.Int, error) {
	return _IOption.Contract.GetSettledRound(&_IOption.CallOpts, account)
}

// GetUnclaimedProfitsRound is a free data retrieval call binding the contract method 0x75882609.
//
// Solidity: function getUnclaimedProfitsRound(address account) view returns(uint256)
func (_IOption *IOptionCaller) GetUnclaimedProfitsRound(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "getUnclaimedProfitsRound", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnclaimedProfitsRound is a free data retrieval call binding the contract method 0x75882609.
//
// Solidity: function getUnclaimedProfitsRound(address account) view returns(uint256)
func (_IOption *IOptionSession) GetUnclaimedProfitsRound(account common.Address) (*big.Int, error) {
	return _IOption.Contract.GetUnclaimedProfitsRound(&_IOption.CallOpts, account)
}

// GetUnclaimedProfitsRound is a free data retrieval call binding the contract method 0x75882609.
//
// Solidity: function getUnclaimedProfitsRound(address account) view returns(uint256)
func (_IOption *IOptionCallerSession) GetUnclaimedProfitsRound(account common.Address) (*big.Int, error) {
	return _IOption.Contract.GetUnclaimedProfitsRound(&_IOption.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IOption *IOptionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IOption *IOptionSession) Name() (string, error) {
	return _IOption.Contract.Name(&_IOption.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IOption *IOptionCallerSession) Name() (string, error) {
	return _IOption.Contract.Name(&_IOption.CallOpts)
}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256)
func (_IOption *IOptionCaller) StrikePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "strikePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256)
func (_IOption *IOptionSession) StrikePrice() (*big.Int, error) {
	return _IOption.Contract.StrikePrice(&_IOption.CallOpts)
}

// StrikePrice is a free data retrieval call binding the contract method 0xc52987cf.
//
// Solidity: function strikePrice() view returns(uint256)
func (_IOption *IOptionCallerSession) StrikePrice() (*big.Int, error) {
	return _IOption.Contract.StrikePrice(&_IOption.CallOpts)
}

// TotalPremiums is a free data retrieval call binding the contract method 0x5bda1584.
//
// Solidity: function totalPremiums() view returns(uint256)
func (_IOption *IOptionCaller) TotalPremiums(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "totalPremiums")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPremiums is a free data retrieval call binding the contract method 0x5bda1584.
//
// Solidity: function totalPremiums() view returns(uint256)
func (_IOption *IOptionSession) TotalPremiums() (*big.Int, error) {
	return _IOption.Contract.TotalPremiums(&_IOption.CallOpts)
}

// TotalPremiums is a free data retrieval call binding the contract method 0x5bda1584.
//
// Solidity: function totalPremiums() view returns(uint256)
func (_IOption *IOptionCallerSession) TotalPremiums() (*big.Int, error) {
	return _IOption.Contract.TotalPremiums(&_IOption.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IOption *IOptionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IOption.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IOption *IOptionSession) TotalSupply() (*big.Int, error) {
	return _IOption.Contract.TotalSupply(&_IOption.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IOption *IOptionCallerSession) TotalSupply() (*big.Int, error) {
	return _IOption.Contract.TotalSupply(&_IOption.CallOpts)
}

// AddPremium is a paid mutator transaction binding the contract method 0xea4cc359.
//
// Solidity: function addPremium(uint256 amountUSDT) returns()
func (_IOption *IOptionTransactor) AddPremium(opts *bind.TransactOpts, amountUSDT *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "addPremium", amountUSDT)
}

// AddPremium is a paid mutator transaction binding the contract method 0xea4cc359.
//
// Solidity: function addPremium(uint256 amountUSDT) returns()
func (_IOption *IOptionSession) AddPremium(amountUSDT *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.AddPremium(&_IOption.TransactOpts, amountUSDT)
}

// AddPremium is a paid mutator transaction binding the contract method 0xea4cc359.
//
// Solidity: function addPremium(uint256 amountUSDT) returns()
func (_IOption *IOptionTransactorSession) AddPremium(amountUSDT *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.AddPremium(&_IOption.TransactOpts, amountUSDT)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IOption *IOptionTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IOption *IOptionSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.Approve(&_IOption.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IOption *IOptionTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.Approve(&_IOption.TransactOpts, spender, amount)
}

// ResetOption is a paid mutator transaction binding the contract method 0xa230f178.
//
// Solidity: function resetOption(uint256 strikePrice_, uint256 newSupply) returns()
func (_IOption *IOptionTransactor) ResetOption(opts *bind.TransactOpts, strikePrice_ *big.Int, newSupply *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "resetOption", strikePrice_, newSupply)
}

// ResetOption is a paid mutator transaction binding the contract method 0xa230f178.
//
// Solidity: function resetOption(uint256 strikePrice_, uint256 newSupply) returns()
func (_IOption *IOptionSession) ResetOption(strikePrice_ *big.Int, newSupply *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.ResetOption(&_IOption.TransactOpts, strikePrice_, newSupply)
}

// ResetOption is a paid mutator transaction binding the contract method 0xa230f178.
//
// Solidity: function resetOption(uint256 strikePrice_, uint256 newSupply) returns()
func (_IOption *IOptionTransactorSession) ResetOption(strikePrice_ *big.Int, newSupply *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.ResetOption(&_IOption.TransactOpts, strikePrice_, newSupply)
}

// SetRoundAccOPASellerShare is a paid mutator transaction binding the contract method 0x49d37d7a.
//
// Solidity: function setRoundAccOPASellerShare(uint256 r, uint256 sellerOPAShare) returns()
func (_IOption *IOptionTransactor) SetRoundAccOPASellerShare(opts *bind.TransactOpts, r *big.Int, sellerOPAShare *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "setRoundAccOPASellerShare", r, sellerOPAShare)
}

// SetRoundAccOPASellerShare is a paid mutator transaction binding the contract method 0x49d37d7a.
//
// Solidity: function setRoundAccOPASellerShare(uint256 r, uint256 sellerOPAShare) returns()
func (_IOption *IOptionSession) SetRoundAccOPASellerShare(r *big.Int, sellerOPAShare *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.SetRoundAccOPASellerShare(&_IOption.TransactOpts, r, sellerOPAShare)
}

// SetRoundAccOPASellerShare is a paid mutator transaction binding the contract method 0x49d37d7a.
//
// Solidity: function setRoundAccOPASellerShare(uint256 r, uint256 sellerOPAShare) returns()
func (_IOption *IOptionTransactorSession) SetRoundAccOPASellerShare(r *big.Int, sellerOPAShare *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.SetRoundAccOPASellerShare(&_IOption.TransactOpts, r, sellerOPAShare)
}

// SetRoundAccPremiumShare is a paid mutator transaction binding the contract method 0x00eec150.
//
// Solidity: function setRoundAccPremiumShare(uint256 r, uint256 premiumShare) returns()
func (_IOption *IOptionTransactor) SetRoundAccPremiumShare(opts *bind.TransactOpts, r *big.Int, premiumShare *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "setRoundAccPremiumShare", r, premiumShare)
}

// SetRoundAccPremiumShare is a paid mutator transaction binding the contract method 0x00eec150.
//
// Solidity: function setRoundAccPremiumShare(uint256 r, uint256 premiumShare) returns()
func (_IOption *IOptionSession) SetRoundAccPremiumShare(r *big.Int, premiumShare *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.SetRoundAccPremiumShare(&_IOption.TransactOpts, r, premiumShare)
}

// SetRoundAccPremiumShare is a paid mutator transaction binding the contract method 0x00eec150.
//
// Solidity: function setRoundAccPremiumShare(uint256 r, uint256 premiumShare) returns()
func (_IOption *IOptionTransactorSession) SetRoundAccPremiumShare(r *big.Int, premiumShare *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.SetRoundAccPremiumShare(&_IOption.TransactOpts, r, premiumShare)
}

// SetSettledRound is a paid mutator transaction binding the contract method 0x243da684.
//
// Solidity: function setSettledRound(uint256 r, address account) returns()
func (_IOption *IOptionTransactor) SetSettledRound(opts *bind.TransactOpts, r *big.Int, account common.Address) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "setSettledRound", r, account)
}

// SetSettledRound is a paid mutator transaction binding the contract method 0x243da684.
//
// Solidity: function setSettledRound(uint256 r, address account) returns()
func (_IOption *IOptionSession) SetSettledRound(r *big.Int, account common.Address) (*types.Transaction, error) {
	return _IOption.Contract.SetSettledRound(&_IOption.TransactOpts, r, account)
}

// SetSettledRound is a paid mutator transaction binding the contract method 0x243da684.
//
// Solidity: function setSettledRound(uint256 r, address account) returns()
func (_IOption *IOptionTransactorSession) SetSettledRound(r *big.Int, account common.Address) (*types.Transaction, error) {
	return _IOption.Contract.SetSettledRound(&_IOption.TransactOpts, r, account)
}

// SetUnclaimedProfitsRound is a paid mutator transaction binding the contract method 0xf3681b8a.
//
// Solidity: function setUnclaimedProfitsRound(uint256 r, address account) returns()
func (_IOption *IOptionTransactor) SetUnclaimedProfitsRound(opts *bind.TransactOpts, r *big.Int, account common.Address) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "setUnclaimedProfitsRound", r, account)
}

// SetUnclaimedProfitsRound is a paid mutator transaction binding the contract method 0xf3681b8a.
//
// Solidity: function setUnclaimedProfitsRound(uint256 r, address account) returns()
func (_IOption *IOptionSession) SetUnclaimedProfitsRound(r *big.Int, account common.Address) (*types.Transaction, error) {
	return _IOption.Contract.SetUnclaimedProfitsRound(&_IOption.TransactOpts, r, account)
}

// SetUnclaimedProfitsRound is a paid mutator transaction binding the contract method 0xf3681b8a.
//
// Solidity: function setUnclaimedProfitsRound(uint256 r, address account) returns()
func (_IOption *IOptionTransactorSession) SetUnclaimedProfitsRound(r *big.Int, account common.Address) (*types.Transaction, error) {
	return _IOption.Contract.SetUnclaimedProfitsRound(&_IOption.TransactOpts, r, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IOption *IOptionTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IOption *IOptionSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.Transfer(&_IOption.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IOption *IOptionTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.Transfer(&_IOption.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IOption *IOptionTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IOption *IOptionSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.TransferFrom(&_IOption.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IOption *IOptionTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IOption.Contract.TransferFrom(&_IOption.TransactOpts, sender, recipient, amount)
}

// IOptionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IOption contract.
type IOptionApprovalIterator struct {
	Event *IOptionApproval // Event containing the contract specifics and raw log

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
func (it *IOptionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionApproval)
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
		it.Event = new(IOptionApproval)
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
func (it *IOptionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionApproval represents a Approval event raised by the IOption contract.
type IOptionApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IOption *IOptionFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IOptionApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IOption.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IOptionApprovalIterator{contract: _IOption.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IOption *IOptionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IOptionApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IOption.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionApproval)
				if err := _IOption.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IOption *IOptionFilterer) ParseApproval(log types.Log) (*IOptionApproval, error) {
	event := new(IOptionApproval)
	if err := _IOption.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOptionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IOption contract.
type IOptionTransferIterator struct {
	Event *IOptionTransfer // Event containing the contract specifics and raw log

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
func (it *IOptionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOptionTransfer)
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
		it.Event = new(IOptionTransfer)
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
func (it *IOptionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOptionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOptionTransfer represents a Transfer event raised by the IOption contract.
type IOptionTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IOption *IOptionFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IOptionTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IOption.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IOptionTransferIterator{contract: _IOption.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IOption *IOptionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IOptionTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IOption.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOptionTransfer)
				if err := _IOption.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IOption *IOptionFilterer) ParseTransfer(log types.Log) (*IOptionTransfer, error) {
	event := new(IOptionTransfer)
	if err := _IOption.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
