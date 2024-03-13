// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamond

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

// DiamondArgs is an auto generated low-level Go binding around an user-defined struct.
type DiamondArgs struct {
	Owner        common.Address
	Init         common.Address
	InitCalldata []byte
}

// IDiamondFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondMetaData contains all meta data concerning the Diamond contract.
var DiamondMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"internalType\":\"structIDiamond.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamond.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_args\",\"type\":\"tuple\",\"internalType\":\"structDiamondArgs\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"init\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initCalldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"DiamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamond.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotAddFunctionToDiamondThatAlreadyExists\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotAddSelectorsToZeroAddress\",\"inputs\":[{\"name\":\"_selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"type\":\"error\",\"name\":\"CannotRemoveFunctionThatDoesNotExist\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotRemoveImmutableFunction\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceFunctionThatDoesNotExists\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceFunctionsFromFacetWithZeroAddress\",\"inputs\":[{\"name\":\"_selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceImmutableFunction\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"FunctionNotFound\",\"inputs\":[{\"name\":\"_functionSelector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"IncorrectFacetCutAction\",\"inputs\":[{\"name\":\"_action\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InitializationFunctionReverted\",\"inputs\":[{\"name\":\"_initializationContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NoBytecodeAtAddress\",\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NoSelectorsProvidedForFacetForCut\",\"inputs\":[{\"name\":\"_facetAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RemoveFacetAddressMustBeZeroAddress\",\"inputs\":[{\"name\":\"_facetAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405260405161106c38038061106c83398101604081905261002291610aab565b805161002d9061004d565b61004682826020015183604001516100d060201b60201c565b5050610e7a565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f80546001600160a01b031981166001600160a01b03848116918217909355604051600080516020610fb8833981519152939092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60005b835181101561022f5760008482815181106100f0576100f0610c4c565b6020026020010151604001519050600085838151811061011257610112610c4c565b602002602001015160000151905081516000036101525760405163e767f91f60e01b81526001600160a01b03821660048201526024015b60405180910390fd5b600086848151811061016657610166610c4c565b60200260200101516020015190506000600281111561018757610187610c62565b81600281111561019957610199610c62565b036101ad576101a8828461027a565b610224565b60018160028111156101c1576101c1610c62565b036101d0576101a88284610426565b60028160028111156101e4576101e4610c62565b036101f3576101a882846105af565b80600281111561020557610205610c62565b604051633ff4d20f60e11b815260ff9091166004820152602401610149565b5050506001016100d3565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb67383838360405161026393929190610cea565b60405180910390a1610275828261082e565b505050565b6001600160a01b0382166102a357806040516302b8da0760e21b81526004016101499190610db4565b6000805160206110248339815191525460408051606081019091526024808252600080516020610fb883398151915292916102e891869190610fd860208301396108f4565b60005b835181101561041f57600084828151811061030857610308610c4c565b6020908102919091018101516001600160e01b031981166000908152918690526040909120549091506001600160a01b031680156103655760405163ebbf5d0760e01b81526001600160e01b031983166004820152602401610149565b6040805180820182526001600160a01b03808a16825261ffff80881660208085019182526001600160e01b0319881660009081528b8252958620945185549251909316600160a01b026001600160b01b0319909216929093169190911717909155600180880180549182018155835291206008820401805460e085901c60046007909416939093026101000a92830263ffffffff90930219169190911790558361040e81610de4565b945050600190920191506102eb9050565b5050505050565b600080516020610fb88339815191526001600160a01b03831661045e578160405163cd98a96f60e01b81526004016101499190610db4565b61048083604051806060016040528060288152602001611044602891396108f4565b60005b82518110156105a95760008382815181106104a0576104a0610c4c565b6020908102919091018101516001600160e01b031981166000908152918590526040909120549091506001600160a01b03163081036104fe57604051632901806d60e11b81526001600160e01b031983166004820152602401610149565b856001600160a01b0316816001600160a01b03160361053c57604051631ac6ce8d60e11b81526001600160e01b031983166004820152602401610149565b6001600160a01b03811661056f57604051637479f93960e01b81526001600160e01b031983166004820152602401610149565b506001600160e01b031916600090815260208390526040902080546001600160a01b0319166001600160a01b038616179055600101610483565b50505050565b60008051602061102483398151915254600080516020610fb8833981519152906001600160a01b038416156106025760405163d091bc8160e01b81526001600160a01b0385166004820152602401610149565b60005b835181101561041f57600084828151811061062257610622610c4c565b6020908102919091018101516001600160e01b0319811660009081528683526040908190208151808301909252546001600160a01b038116808352600160a01b90910461ffff16938201939093529092509061069d57604051637a08a22d60e01b81526001600160e01b031983166004820152602401610149565b8051306001600160a01b03909116036106d557604051630df5fd6160e31b81526001600160e01b031983166004820152602401610149565b836106df81610e05565b94505083816020015161ffff16146107bd57600085600101858154811061070857610708610c4c565b90600052602060002090600891828204019190066004029054906101000a900460e01b90508086600101836020015161ffff168154811061074b5761074b610c4c565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c92909202939093179055838201516001600160e01b03199390931681529087905260409020805461ffff60a01b1916600160a01b61ffff909316929092029190911790555b846001018054806107d0576107d0610e1c565b60008281526020808220600860001990940193840401805463ffffffff600460078716026101000a0219169055919092556001600160e01b0319909316815291859052506040902080546001600160b01b0319169055600101610605565b6001600160a01b038216610840575050565b61086282604051806060016040528060288152602001610ffc602891396108f4565b600080836001600160a01b03168360405161087d9190610e32565b600060405180830381855af49150503d80600081146108b8576040519150601f19603f3d011682016040523d82523d6000602084013e6108bd565b606091505b5091509150816105a9578051156108d75780518082602001fd5b838360405163192105d760e01b8152600401610149929190610e4e565b813b600081900361027557828260405163919834b960e01b8152600401610149929190610e4e565b634e487b7160e01b600052604160045260246000fd5b604051606081016001600160401b03811182821017156109545761095461091c565b60405290565b604051601f8201601f191681016001600160401b03811182821017156109825761098261091c565b604052919050565b60006001600160401b038211156109a3576109a361091c565b5060051b60200190565b80516001600160a01b03811681146109c457600080fd5b919050565b60005b838110156109e45781810151838201526020016109cc565b50506000910152565b6000606082840312156109ff57600080fd5b610a07610932565b9050610a12826109ad565b81526020610a218184016109ad565b8282015260408301516001600160401b0380821115610a3f57600080fd5b818501915085601f830112610a5357600080fd5b815181811115610a6557610a6561091c565b610a77601f8201601f1916850161095a565b91508082528684828501011115610a8d57600080fd5b610a9c818584018686016109c9565b50604084015250909392505050565b60008060408385031215610abe57600080fd5b82516001600160401b0380821115610ad557600080fd5b818501915085601f830112610ae957600080fd5b81516020610afe610af98361098a565b61095a565b82815260059290921b84018101918181019089841115610b1d57600080fd5b8286015b84811015610c1b57805186811115610b3857600080fd5b87016060818d03601f19011215610b4e57600080fd5b610b56610932565b610b618683016109ad565b8152604082015160038110610b7557600080fd5b81870152606082015188811115610b8b57600080fd5b8083019250508c603f830112610ba057600080fd5b85820151610bb0610af98261098a565b81815260059190911b830160400190878101908f831115610bd057600080fd5b6040850194505b82851015610c065784516001600160e01b031981168114610bf757600080fd5b82529388019390880190610bd7565b60408401525050845250918301918301610b21565b5091880151919650909350505080821115610c3557600080fd5b50610c42858286016109ed565b9150509250929050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052602160045260246000fd5b60008151808452602080850194506020840160005b83811015610cb35781516001600160e01b03191687529582019590820190600101610c8d565b509495945050505050565b60008151808452610cd68160208601602086016109c9565b601f01601f19169290920160200192915050565b600060608083016060845280875180835260808601915060808160051b87010192506020808a016000805b84811015610d8457898703607f19018652825180516001600160a01b031688528481015160038110610d5557634e487b7160e01b84526021600452602484fd5b88860152604090810151908801899052610d7189890182610c78565b9750509483019491830191600101610d15565b5050506001600160a01b0389169087015250508381036040850152610da98186610cbe565b979650505050505050565b602081526000610dc76020830184610c78565b9392505050565b634e487b7160e01b600052601160045260246000fd5b600061ffff808316818103610dfb57610dfb610dce565b6001019392505050565b600081610e1457610e14610dce565b506000190190565b634e487b7160e01b600052603160045260246000fd5b60008251610e448184602087016109c9565b9190910192915050565b6001600160a01b0383168152604060208201819052600090610e7290830184610cbe565b949350505050565b61012f80610e896000396000f3fe608060405236600a57005b600080357fffffffff000000000000000000000000000000000000000000000000000000001681527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c6020819052604090912054819073ffffffffffffffffffffffffffffffffffffffff168060d6576040517f5416eb980000000000000000000000000000000000000000000000000000000081527fffffffff0000000000000000000000000000000000000000000000000000000060003516600482015260240160405180910390fd5b3660008037600080366000845af43d6000803e80801560f4573d6000f35b3d6000fdfea2646970667358221220d66d3fc31dd2ec1e2f2f7933f25afefbcd6f386b375bff041cb3088ac0de9b6d64736f6c63430008170033c8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c4c69624469616d6f6e644375743a2041646420666163657420686173206e6f20636f64654c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f6465c8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d4c69624469616d6f6e644375743a205265706c61636520666163657420686173206e6f20636f6465",
}

// DiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondMetaData.ABI instead.
var DiamondABI = DiamondMetaData.ABI

// DiamondBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondMetaData.Bin instead.
var DiamondBin = DiamondMetaData.Bin

// DeployDiamond deploys a new Ethereum contract, binding an instance of Diamond to it.
func DeployDiamond(auth *bind.TransactOpts, backend bind.ContractBackend, _diamondCut []IDiamondFacetCut, _args DiamondArgs) (common.Address, *types.Transaction, *Diamond, error) {
	parsed, err := DiamondMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondBin), backend, _diamondCut, _args)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// Diamond is an auto generated Go binding around an Ethereum contract.
type Diamond struct {
	DiamondCaller     // Read-only binding to the contract
	DiamondTransactor // Write-only binding to the contract
	DiamondFilterer   // Log filterer for contract events
}

// DiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondSession struct {
	Contract     *Diamond          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCallerSession struct {
	Contract *DiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondTransactorSession struct {
	Contract     *DiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondRaw struct {
	Contract *Diamond // Generic contract binding to access the raw methods on
}

// DiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCallerRaw struct {
	Contract *DiamondCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondTransactorRaw struct {
	Contract *DiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamond creates a new instance of Diamond, bound to a specific deployed contract.
func NewDiamond(address common.Address, backend bind.ContractBackend) (*Diamond, error) {
	contract, err := bindDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// NewDiamondCaller creates a new read-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondCaller(address common.Address, caller bind.ContractCaller) (*DiamondCaller, error) {
	contract, err := bindDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCaller{contract: contract}, nil
}

// NewDiamondTransactor creates a new write-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondTransactor, error) {
	contract, err := bindDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondTransactor{contract: contract}, nil
}

// NewDiamondFilterer creates a new log filterer instance of Diamond, bound to a specific deployed contract.
func NewDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondFilterer, error) {
	contract, err := bindDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondFilterer{contract: contract}, nil
}

// bindDiamond binds a generic wrapper to an already deployed contract.
func bindDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.DiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondSession) Receive() (*types.Transaction, error) {
	return _Diamond.Contract.Receive(&_Diamond.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Diamond *DiamondTransactorSession) Receive() (*types.Transaction, error) {
	return _Diamond.Contract.Receive(&_Diamond.TransactOpts)
}

// DiamondDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the Diamond contract.
type DiamondDiamondCutIterator struct {
	Event *DiamondDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondDiamondCut)
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
		it.Event = new(DiamondDiamondCut)
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
func (it *DiamondDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondDiamondCut represents a DiamondCut event raised by the Diamond contract.
type DiamondDiamondCut struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Diamond *DiamondFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondDiamondCutIterator, error) {

	logs, sub, err := _Diamond.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondDiamondCutIterator{contract: _Diamond.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Diamond *DiamondFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondDiamondCut) (event.Subscription, error) {

	logs, sub, err := _Diamond.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondDiamondCut)
				if err := _Diamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_Diamond *DiamondFilterer) ParseDiamondCut(log types.Log) (*DiamondDiamondCut, error) {
	event := new(DiamondDiamondCut)
	if err := _Diamond.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Diamond contract.
type DiamondOwnershipTransferredIterator struct {
	Event *DiamondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DiamondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondOwnershipTransferred)
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
		it.Event = new(DiamondOwnershipTransferred)
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
func (it *DiamondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondOwnershipTransferred represents a OwnershipTransferred event raised by the Diamond contract.
type DiamondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Diamond *DiamondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DiamondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Diamond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DiamondOwnershipTransferredIterator{contract: _Diamond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Diamond *DiamondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DiamondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Diamond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondOwnershipTransferred)
				if err := _Diamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Diamond *DiamondFilterer) ParseOwnershipTransferred(log types.Log) (*DiamondOwnershipTransferred, error) {
	event := new(DiamondOwnershipTransferred)
	if err := _Diamond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
