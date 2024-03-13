// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamondLoupe

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

// IDiamondLoupeFacet is an auto generated low-level Go binding around an user-defined struct.
type IDiamondLoupeFacet struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// DiamondLoupeMetaData contains all meta data concerning the DiamondLoupe contract.
var DiamondLoupeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"facetAddress\",\"inputs\":[{\"name\":\"_functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"facetAddress_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"facetAddresses\",\"inputs\":[],\"outputs\":[{\"name\":\"facetAddresses_\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"facetFunctionSelectors\",\"inputs\":[{\"name\":\"_facet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"_facetFunctionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"facets\",\"inputs\":[],\"outputs\":[{\"name\":\"facets_\",\"type\":\"tuple[]\",\"internalType\":\"structIDiamondLoupe.Facet[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610cd3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100675760003560e01c80637a0ed627116100505780637a0ed627146100fa578063adfca15e1461010f578063cdffacc61461012f57600080fd5b806301ffc9a71461006c57806352ef6b2c146100e5575b600080fd5b6100d061007a3660046109c3565b7fffffffff000000000000000000000000000000000000000000000000000000001660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131e602052604090205460ff1690565b60405190151581526020015b60405180910390f35b6100ed6101cb565b6040516100dc9190610a0c565b6101026103b6565b6040516100dc9190610ac4565b61012261011d366004610b6e565b610836565b6040516100dc9190610ba4565b6101a661013d3660046109c3565b7fffffffff000000000000000000000000000000000000000000000000000000001660009081527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100dc565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d546060907fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c908067ffffffffffffffff81111561022b5761022b610bb7565b604051908082528060200260200182016040528015610254578160200160208202803683370190505b5092506000805b828110156103ac57600084600101828154811061027a5761027a610be6565b6000918252602080832060088304015460079092166004026101000a90910460e01b7fffffffff000000000000000000000000000000000000000000000000000000008116835290879052604082205490925073ffffffffffffffffffffffffffffffffffffffff1690805b8581101561034b5788818151811061030057610300610be6565b602002602001015173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610343576001915061034b565b6001016102e6565b50801561035b57506103a4915050565b8188868151811061036e5761036e610be6565b73ffffffffffffffffffffffffffffffffffffffff909216602092830291909101909101528461039d81610c44565b9550505050505b60010161025b565b5080845250505090565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d546060907fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c908067ffffffffffffffff81111561041657610416610bb7565b60405190808252806020026020018201604052801561045c57816020015b6040805180820190915260008152606060208201528152602001906001900390816104345790505b50925060008167ffffffffffffffff81111561047a5761047a610bb7565b6040519080825280602002602001820160405280156104a3578160200160208202803683370190505b5090506000805b838110156107ce5760008560010182815481106104c9576104c9610be6565b6000918252602080832060088304015460079092166004026101000a90910460e01b7fffffffff000000000000000000000000000000000000000000000000000000008116835290889052604082205490925073ffffffffffffffffffffffffffffffffffffffff1690805b85811015610665578273ffffffffffffffffffffffffffffffffffffffff168a828151811061056657610566610be6565b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff160361065d57838a82815181106105a0576105a0610be6565b6020026020010151602001518883815181106105be576105be610be6565b602002602001015161ffff16815181106105da576105da610be6565b60200260200101907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505086818151811061063857610638610be6565b60200260200101805180919061064d90610c7c565b61ffff1690525060019150610665565b600101610535565b50801561067557506107c6915050565b8189868151811061068857610688610be6565b602090810291909101015173ffffffffffffffffffffffffffffffffffffffff90911690528667ffffffffffffffff8111156106c6576106c6610bb7565b6040519080825280602002602001820160405280156106ef578160200160208202803683370190505b5089868151811061070257610702610be6565b6020026020010151602001819052508289868151811061072457610724610be6565b60200260200101516020015160008151811061074257610742610be6565b60200260200101907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505060018686815181106107a2576107a2610be6565b61ffff90921660209283029190910190910152846107bf81610c44565b9550505050505b6001016104aa565b5060005b8181101561082b5760008382815181106107ee576107ee610be6565b602002602001015161ffff169050600087838151811061081057610810610be6565b602090810291909101810151015191909152506001016107d2565b508085525050505090565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d546060907fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c9060008167ffffffffffffffff81111561089857610898610bb7565b6040519080825280602002602001820160405280156108c1578160200160208202803683370190505b50935060005b828110156109b85760008460010182815481106108e6576108e6610be6565b6000918252602080832060088304015460079092166004026101000a90910460e01b7fffffffff00000000000000000000000000000000000000000000000000000000811683529087905260409091205490915073ffffffffffffffffffffffffffffffffffffffff9081169088168190036109ae578187858151811061096f5761096f610be6565b7fffffffff0000000000000000000000000000000000000000000000000000000090921660209283029190910190910152836109aa81610c44565b9450505b50506001016108c7565b508352509092915050565b6000602082840312156109d557600080fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610a0557600080fd5b9392505050565b6020808252825182820181905260009190848201906040850190845b81811015610a5a57835173ffffffffffffffffffffffffffffffffffffffff1683529284019291840191600101610a28565b50909695505050505050565b60008151808452602080850194506020840160005b83811015610ab95781517fffffffff000000000000000000000000000000000000000000000000000000001687529582019590820190600101610a7b565b509495945050505050565b600060208083018184528085518083526040925060408601915060408160051b87010184880160005b83811015610b60578883037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc00185528151805173ffffffffffffffffffffffffffffffffffffffff168452870151878401879052610b4d87850182610a66565b9588019593505090860190600101610aed565b509098975050505050505050565b600060208284031215610b8057600080fd5b813573ffffffffffffffffffffffffffffffffffffffff81168114610a0557600080fd5b602081526000610a056020830184610a66565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c7557610c75610c15565b5060010190565b600061ffff808316818103610c9357610c93610c15565b600101939250505056fea2646970667358221220ca638061611a284108fe31d372045e9d68187bc0bbc42e050280f7f8a8ee2cdb64736f6c63430008170033",
}

// DiamondLoupeABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondLoupeMetaData.ABI instead.
var DiamondLoupeABI = DiamondLoupeMetaData.ABI

// DiamondLoupeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondLoupeMetaData.Bin instead.
var DiamondLoupeBin = DiamondLoupeMetaData.Bin

// DeployDiamondLoupe deploys a new Ethereum contract, binding an instance of DiamondLoupe to it.
func DeployDiamondLoupe(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondLoupe, error) {
	parsed, err := DiamondLoupeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondLoupeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondLoupe{DiamondLoupeCaller: DiamondLoupeCaller{contract: contract}, DiamondLoupeTransactor: DiamondLoupeTransactor{contract: contract}, DiamondLoupeFilterer: DiamondLoupeFilterer{contract: contract}}, nil
}

// DiamondLoupe is an auto generated Go binding around an Ethereum contract.
type DiamondLoupe struct {
	DiamondLoupeCaller     // Read-only binding to the contract
	DiamondLoupeTransactor // Write-only binding to the contract
	DiamondLoupeFilterer   // Log filterer for contract events
}

// DiamondLoupeCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondLoupeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondLoupeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondLoupeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondLoupeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondLoupeSession struct {
	Contract     *DiamondLoupe     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondLoupeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondLoupeCallerSession struct {
	Contract *DiamondLoupeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DiamondLoupeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondLoupeTransactorSession struct {
	Contract     *DiamondLoupeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DiamondLoupeRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondLoupeRaw struct {
	Contract *DiamondLoupe // Generic contract binding to access the raw methods on
}

// DiamondLoupeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondLoupeCallerRaw struct {
	Contract *DiamondLoupeCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondLoupeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondLoupeTransactorRaw struct {
	Contract *DiamondLoupeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondLoupe creates a new instance of DiamondLoupe, bound to a specific deployed contract.
func NewDiamondLoupe(address common.Address, backend bind.ContractBackend) (*DiamondLoupe, error) {
	contract, err := bindDiamondLoupe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupe{DiamondLoupeCaller: DiamondLoupeCaller{contract: contract}, DiamondLoupeTransactor: DiamondLoupeTransactor{contract: contract}, DiamondLoupeFilterer: DiamondLoupeFilterer{contract: contract}}, nil
}

// NewDiamondLoupeCaller creates a new read-only instance of DiamondLoupe, bound to a specific deployed contract.
func NewDiamondLoupeCaller(address common.Address, caller bind.ContractCaller) (*DiamondLoupeCaller, error) {
	contract, err := bindDiamondLoupe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeCaller{contract: contract}, nil
}

// NewDiamondLoupeTransactor creates a new write-only instance of DiamondLoupe, bound to a specific deployed contract.
func NewDiamondLoupeTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondLoupeTransactor, error) {
	contract, err := bindDiamondLoupe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeTransactor{contract: contract}, nil
}

// NewDiamondLoupeFilterer creates a new log filterer instance of DiamondLoupe, bound to a specific deployed contract.
func NewDiamondLoupeFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondLoupeFilterer, error) {
	contract, err := bindDiamondLoupe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondLoupeFilterer{contract: contract}, nil
}

// bindDiamondLoupe binds a generic wrapper to an already deployed contract.
func bindDiamondLoupe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondLoupeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupe *DiamondLoupeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupe.Contract.DiamondLoupeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupe *DiamondLoupeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupe.Contract.DiamondLoupeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupe *DiamondLoupeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupe.Contract.DiamondLoupeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondLoupe *DiamondLoupeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondLoupe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondLoupe *DiamondLoupeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondLoupe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondLoupe *DiamondLoupeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondLoupe.Contract.contract.Transact(opts, method, params...)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupe *DiamondLoupeCaller) FacetAddress(opts *bind.CallOpts, _functionSelector [4]byte) (common.Address, error) {
	var out []interface{}
	err := _DiamondLoupe.contract.Call(opts, &out, "facetAddress", _functionSelector)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupe *DiamondLoupeSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupe.Contract.FacetAddress(&_DiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 _functionSelector) view returns(address facetAddress_)
func (_DiamondLoupe *DiamondLoupeCallerSession) FacetAddress(_functionSelector [4]byte) (common.Address, error) {
	return _DiamondLoupe.Contract.FacetAddress(&_DiamondLoupe.CallOpts, _functionSelector)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupe *DiamondLoupeCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _DiamondLoupe.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupe *DiamondLoupeSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupe.Contract.FacetAddresses(&_DiamondLoupe.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facetAddresses_)
func (_DiamondLoupe *DiamondLoupeCallerSession) FacetAddresses() ([]common.Address, error) {
	return _DiamondLoupe.Contract.FacetAddresses(&_DiamondLoupe.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupe *DiamondLoupeCaller) FacetFunctionSelectors(opts *bind.CallOpts, _facet common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _DiamondLoupe.contract.Call(opts, &out, "facetFunctionSelectors", _facet)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupe *DiamondLoupeSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupe.Contract.FacetFunctionSelectors(&_DiamondLoupe.CallOpts, _facet)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address _facet) view returns(bytes4[] _facetFunctionSelectors)
func (_DiamondLoupe *DiamondLoupeCallerSession) FacetFunctionSelectors(_facet common.Address) ([][4]byte, error) {
	return _DiamondLoupe.Contract.FacetFunctionSelectors(&_DiamondLoupe.CallOpts, _facet)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupe *DiamondLoupeCaller) Facets(opts *bind.CallOpts) ([]IDiamondLoupeFacet, error) {
	var out []interface{}
	err := _DiamondLoupe.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]IDiamondLoupeFacet), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDiamondLoupeFacet)).(*[]IDiamondLoupeFacet)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupe *DiamondLoupeSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupe.Contract.Facets(&_DiamondLoupe.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_DiamondLoupe *DiamondLoupeCallerSession) Facets() ([]IDiamondLoupeFacet, error) {
	return _DiamondLoupe.Contract.Facets(&_DiamondLoupe.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupe *DiamondLoupeCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DiamondLoupe.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupe *DiamondLoupeSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupe.Contract.SupportsInterface(&_DiamondLoupe.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) view returns(bool)
func (_DiamondLoupe *DiamondLoupeCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _DiamondLoupe.Contract.SupportsInterface(&_DiamondLoupe.CallOpts, _interfaceId)
}
