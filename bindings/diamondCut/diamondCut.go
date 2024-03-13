// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamondCut

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

// IDiamondFacetCut is an auto generated low-level Go binding around an user-defined struct.
type IDiamondFacetCut struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondCutMetaData contains all meta data concerning the DiamondCut contract.
var DiamondCutMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"diamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"internalType\":\"structIDiamond.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamond.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DiamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamond.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DiamondCut\",\"inputs\":[{\"name\":\"_diamondCut\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structIDiamond.FacetCut[]\",\"components\":[{\"name\":\"facetAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"action\",\"type\":\"uint8\",\"internalType\":\"enumIDiamond.FacetCutAction\"},{\"name\":\"functionSelectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"name\":\"_init\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CannotAddFunctionToDiamondThatAlreadyExists\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotAddSelectorsToZeroAddress\",\"inputs\":[{\"name\":\"_selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"type\":\"error\",\"name\":\"CannotRemoveFunctionThatDoesNotExist\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotRemoveImmutableFunction\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceFunctionThatDoesNotExists\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceFunctionWithTheSameFunctionFromTheSameFacet\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceFunctionsFromFacetWithZeroAddress\",\"inputs\":[{\"name\":\"_selectors\",\"type\":\"bytes4[]\",\"internalType\":\"bytes4[]\"}]},{\"type\":\"error\",\"name\":\"CannotReplaceImmutableFunction\",\"inputs\":[{\"name\":\"_selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}]},{\"type\":\"error\",\"name\":\"IncorrectFacetCutAction\",\"inputs\":[{\"name\":\"_action\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]},{\"type\":\"error\",\"name\":\"InitializationFunctionReverted\",\"inputs\":[{\"name\":\"_initializationContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_calldata\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"NoBytecodeAtAddress\",\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NoSelectorsProvidedForFacetForCut\",\"inputs\":[{\"name\":\"_facetAddress\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"NotContractOwner\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_contractOwner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"RemoveFacetAddressMustBeZeroAddress\",\"inputs\":[{\"name\":\"_facetAddress\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061158c806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80631f931c1c14610030575b600080fd5b61004361003e366004610e65565b610045565b005b61004d61009e565b61009761005a8587610fe2565b8484848080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061015992505050565b5050505050565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c6003015473ffffffffffffffffffffffffffffffffffffffff163314610157577fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131f546040517fff4127cb00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff90911660248201526044015b60405180910390fd5b565b60005b83518110156102f25760008482815181106101795761017961113e565b6020026020010151604001519050600085838151811061019b5761019b61113e565b602002602001015160000151905081516000036101fc576040517fe767f91f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8216600482015260240161014e565b60008684815181106102105761021061113e565b6020026020010151602001519050600060028111156102315761023161116d565b8160028111156102435761024361116d565b0361025757610252828461033d565b6102e7565b600181600281111561026b5761026b61116d565b0361027a5761025282846105d0565b600281600281111561028e5761028e61116d565b0361029d5761025282846108ad565b8060028111156102af576102af61116d565b6040517f7fe9a41e00000000000000000000000000000000000000000000000000000000815260ff909116600482015260240161014e565b50505060010161015c565b507f8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb6738383836040516103269392919061120a565b60405180910390a16103388282610cb9565b505050565b73ffffffffffffffffffffffffffffffffffffffff821661038c57806040517f0ae3681c00000000000000000000000000000000000000000000000000000000815260040161014e9190611375565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d54604080516060810190915260248082527fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c92916103f5918691906114e36020830139610db2565b60005b83518110156100975760008482815181106104155761041561113e565b6020908102919091018101517fffffffff00000000000000000000000000000000000000000000000000000000811660009081529186905260409091205490915073ffffffffffffffffffffffffffffffffffffffff1680156104c8576040517febbf5d070000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161014e565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff808a16825261ffff80881660208085019182527fffffffff00000000000000000000000000000000000000000000000000000000881660009081528b825295862094518554925190931674010000000000000000000000000000000000000000027fffffffffffffffffffff00000000000000000000000000000000000000000000909216929093169190911717909155600180880180549182018155835291206008820401805460e085901c60046007909416939093026101000a92830263ffffffff9093021916919091179055836105bf8161140a565b945050600190920191506103f89050565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c73ffffffffffffffffffffffffffffffffffffffff831661064057816040517fcd98a96f00000000000000000000000000000000000000000000000000000000815260040161014e9190611375565b6106628360405180606001604052806028815260200161152f60289139610db2565b60005b82518110156108a75760008382815181106106825761068261113e565b6020908102919091018101517fffffffff00000000000000000000000000000000000000000000000000000000811660009081529185905260409091205490915073ffffffffffffffffffffffffffffffffffffffff16308103610736576040517f520300da0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161014e565b8573ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036107bf576040517f358d9d1a0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161014e565b73ffffffffffffffffffffffffffffffffffffffff8116610830576040517f7479f9390000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161014e565b507fffffffff0000000000000000000000000000000000000000000000000000000016600090815260208390526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8616179055600101610665565b50505050565b7fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131d547fc8fcad8db84d3cc18b4c41d551ea0ee66dd599cde068d998e57d5e09332c131c9073ffffffffffffffffffffffffffffffffffffffff841615610957576040517fd091bc8100000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8516600482015260240161014e565b60005b83518110156100975760008482815181106109775761097761113e565b6020908102919091018101517fffffffff000000000000000000000000000000000000000000000000000000008116600090815286835260409081902081518083019092525473ffffffffffffffffffffffffffffffffffffffff81168083527401000000000000000000000000000000000000000090910461ffff169382019390935290925090610a59576040517f7a08a22d0000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161014e565b80513073ffffffffffffffffffffffffffffffffffffffff90911603610acf576040517f6fafeb080000000000000000000000000000000000000000000000000000000081527fffffffff000000000000000000000000000000000000000000000000000000008316600482015260240161014e565b83610ad98161142b565b94505083816020015161ffff1614610bfa576000856001018581548110610b0257610b0261113e565b90600052602060002090600891828204019190066004029054906101000a900460e01b90508086600101836020015161ffff1681548110610b4557610b4561113e565b600091825260208083206008830401805463ffffffff60079094166004026101000a938402191660e09590951c92909202939093179055838201517fffffffff00000000000000000000000000000000000000000000000000000000939093168152908790526040902080547fffffffffffffffffffff0000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000061ffff909316929092029190911790555b84600101805480610c0d57610c0d611460565b6000828152602080822060087fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff90940193840401805463ffffffff600460078716026101000a0219169055919092557fffffffff00000000000000000000000000000000000000000000000000000000909316815291859052506040902080547fffffffffffffffffffff0000000000000000000000000000000000000000000016905560010161095a565b73ffffffffffffffffffffffffffffffffffffffff8216610cd8575050565b610cfa8260405180606001604052806028815260200161150760289139610db2565b6000808373ffffffffffffffffffffffffffffffffffffffff1683604051610d22919061148f565b600060405180830381855af49150503d8060008114610d5d576040519150601f19603f3d011682016040523d82523d6000602084013e610d62565b606091505b5091509150816108a757805115610d7c5780518082602001fd5b83836040517f192105d700000000000000000000000000000000000000000000000000000000815260040161014e9291906114ab565b813b60008190036103385782826040517f919834b900000000000000000000000000000000000000000000000000000000815260040161014e9291906114ab565b803573ffffffffffffffffffffffffffffffffffffffff81168114610e1757600080fd5b919050565b60008083601f840112610e2e57600080fd5b50813567ffffffffffffffff811115610e4657600080fd5b602083019150836020828501011115610e5e57600080fd5b9250929050565b600080600080600060608688031215610e7d57600080fd5b853567ffffffffffffffff80821115610e9557600080fd5b818801915088601f830112610ea957600080fd5b813581811115610eb857600080fd5b8960208260051b8501011115610ecd57600080fd5b60208301975080965050610ee360208901610df3565b94506040880135915080821115610ef957600080fd5b50610f0688828901610e1c565b969995985093965092949392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715610f6957610f69610f17565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610fb657610fb6610f17565b604052919050565b600067ffffffffffffffff821115610fd857610fd8610f17565b5060051b60200190565b6000610ff5610ff084610fbe565b610f6f565b83815260208082019190600586811b86013681111561101357600080fd5b865b8181101561113157803567ffffffffffffffff808211156110365760008081fd5b818a0191506060823603121561104c5760008081fd5b611054610f46565b61105d83610df3565b815286830135600381106110715760008081fd5b81880152604083810135838111156110895760008081fd5b939093019236601f8501126110a057600092508283fd5b833592506110b0610ff084610fbe565b83815292871b840188019288810190368511156110cd5760008081fd5b948901945b8486101561111a5785357fffffffff000000000000000000000000000000000000000000000000000000008116811461110b5760008081fd5b825294890194908901906110d2565b918301919091525088525050948301948301611015565b5092979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60005b838110156111b757818101518382015260200161119f565b50506000910152565b600081518084526111d881602086016020860161119c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006060808301606084528087518083526080925060808601915060808160051b8701016020808b0160005b84811015611338577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808a8503018652815188850173ffffffffffffffffffffffffffffffffffffffff825116865284820151600381106112bf577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b868601526040918201519186018a905281519081905290840190600090898701905b808310156113235783517fffffffff000000000000000000000000000000000000000000000000000000001682529286019260019290920191908601906112e1565b50978501979550505090820190600101611236565b505073ffffffffffffffffffffffffffffffffffffffff8a1690880152868103604088015261136781896111c0565b9a9950505050505050505050565b6020808252825182820181905260009190848201906040850190845b818110156113cf5783517fffffffff000000000000000000000000000000000000000000000000000000001683529284019291840191600101611391565b50909695505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061ffff808316818103611421576114216113db565b6001019392505050565b60008161143a5761143a6113db565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600082516114a181846020870161119c565b9190910192915050565b73ffffffffffffffffffffffffffffffffffffffff831681526040602082015260006114da60408301846111c0565b94935050505056fe4c69624469616d6f6e644375743a2041646420666163657420686173206e6f20636f64654c69624469616d6f6e644375743a205f696e6974206164647265737320686173206e6f20636f64654c69624469616d6f6e644375743a205265706c61636520666163657420686173206e6f20636f6465a2646970667358221220bd049ed4c5254da1c7ebc647d49824ac63b762a40fc52e857aaab67fa793db3f64736f6c63430008170033",
}

// DiamondCutABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondCutMetaData.ABI instead.
var DiamondCutABI = DiamondCutMetaData.ABI

// DiamondCutBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DiamondCutMetaData.Bin instead.
var DiamondCutBin = DiamondCutMetaData.Bin

// DeployDiamondCut deploys a new Ethereum contract, binding an instance of DiamondCut to it.
func DeployDiamondCut(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DiamondCut, error) {
	parsed, err := DiamondCutMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DiamondCutBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DiamondCut{DiamondCutCaller: DiamondCutCaller{contract: contract}, DiamondCutTransactor: DiamondCutTransactor{contract: contract}, DiamondCutFilterer: DiamondCutFilterer{contract: contract}}, nil
}

// DiamondCut is an auto generated Go binding around an Ethereum contract.
type DiamondCut struct {
	DiamondCutCaller     // Read-only binding to the contract
	DiamondCutTransactor // Write-only binding to the contract
	DiamondCutFilterer   // Log filterer for contract events
}

// DiamondCutCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCutCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondCutTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondCutFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondCutSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondCutSession struct {
	Contract     *DiamondCut       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCutCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCutCallerSession struct {
	Contract *DiamondCutCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DiamondCutTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondCutTransactorSession struct {
	Contract     *DiamondCutTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DiamondCutRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondCutRaw struct {
	Contract *DiamondCut // Generic contract binding to access the raw methods on
}

// DiamondCutCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCutCallerRaw struct {
	Contract *DiamondCutCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondCutTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondCutTransactorRaw struct {
	Contract *DiamondCutTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamondCut creates a new instance of DiamondCut, bound to a specific deployed contract.
func NewDiamondCut(address common.Address, backend bind.ContractBackend) (*DiamondCut, error) {
	contract, err := bindDiamondCut(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DiamondCut{DiamondCutCaller: DiamondCutCaller{contract: contract}, DiamondCutTransactor: DiamondCutTransactor{contract: contract}, DiamondCutFilterer: DiamondCutFilterer{contract: contract}}, nil
}

// NewDiamondCutCaller creates a new read-only instance of DiamondCut, bound to a specific deployed contract.
func NewDiamondCutCaller(address common.Address, caller bind.ContractCaller) (*DiamondCutCaller, error) {
	contract, err := bindDiamondCut(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutCaller{contract: contract}, nil
}

// NewDiamondCutTransactor creates a new write-only instance of DiamondCut, bound to a specific deployed contract.
func NewDiamondCutTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondCutTransactor, error) {
	contract, err := bindDiamondCut(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCutTransactor{contract: contract}, nil
}

// NewDiamondCutFilterer creates a new log filterer instance of DiamondCut, bound to a specific deployed contract.
func NewDiamondCutFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondCutFilterer, error) {
	contract, err := bindDiamondCut(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondCutFilterer{contract: contract}, nil
}

// bindDiamondCut binds a generic wrapper to an already deployed contract.
func bindDiamondCut(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondCutMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCut *DiamondCutRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCut.Contract.DiamondCutCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCut *DiamondCutRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCut.Contract.DiamondCutTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCut *DiamondCutRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCut.Contract.DiamondCutTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DiamondCut *DiamondCutCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DiamondCut.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DiamondCut *DiamondCutTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DiamondCut.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DiamondCut *DiamondCutTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DiamondCut.Contract.contract.Transact(opts, method, params...)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCut *DiamondCutTransactor) DiamondCut(opts *bind.TransactOpts, _diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCut.contract.Transact(opts, "diamondCut", _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCut *DiamondCutSession) DiamondCut(_diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCut.Contract.DiamondCut(&_DiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata) returns()
func (_DiamondCut *DiamondCutTransactorSession) DiamondCut(_diamondCut []IDiamondFacetCut, _init common.Address, _calldata []byte) (*types.Transaction, error) {
	return _DiamondCut.Contract.DiamondCut(&_DiamondCut.TransactOpts, _diamondCut, _init, _calldata)
}

// DiamondCutDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the DiamondCut contract.
type DiamondCutDiamondCutIterator struct {
	Event *DiamondCutDiamondCut // Event containing the contract specifics and raw log

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
func (it *DiamondCutDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutDiamondCut)
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
		it.Event = new(DiamondCutDiamondCut)
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
func (it *DiamondCutDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutDiamondCut represents a DiamondCut event raised by the DiamondCut contract.
type DiamondCutDiamondCut struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCut *DiamondCutFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*DiamondCutDiamondCutIterator, error) {

	logs, sub, err := _DiamondCut.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &DiamondCutDiamondCutIterator{contract: _DiamondCut.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCut *DiamondCutFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *DiamondCutDiamondCut) (event.Subscription, error) {

	logs, sub, err := _DiamondCut.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutDiamondCut)
				if err := _DiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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
func (_DiamondCut *DiamondCutFilterer) ParseDiamondCut(log types.Log) (*DiamondCutDiamondCut, error) {
	event := new(DiamondCutDiamondCut)
	if err := _DiamondCut.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DiamondCutDiamondCut0Iterator is returned from FilterDiamondCut0 and is used to iterate over the raw logs and unpacked data for DiamondCut0 events raised by the DiamondCut contract.
type DiamondCutDiamondCut0Iterator struct {
	Event *DiamondCutDiamondCut0 // Event containing the contract specifics and raw log

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
func (it *DiamondCutDiamondCut0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DiamondCutDiamondCut0)
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
		it.Event = new(DiamondCutDiamondCut0)
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
func (it *DiamondCutDiamondCut0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DiamondCutDiamondCut0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DiamondCutDiamondCut0 represents a DiamondCut0 event raised by the DiamondCut contract.
type DiamondCutDiamondCut0 struct {
	DiamondCut []IDiamondFacetCut
	Init       common.Address
	Calldata   []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut0 is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCut *DiamondCutFilterer) FilterDiamondCut0(opts *bind.FilterOpts) (*DiamondCutDiamondCut0Iterator, error) {

	logs, sub, err := _DiamondCut.contract.FilterLogs(opts, "DiamondCut0")
	if err != nil {
		return nil, err
	}
	return &DiamondCutDiamondCut0Iterator{contract: _DiamondCut.contract, event: "DiamondCut0", logs: logs, sub: sub}, nil
}

// WatchDiamondCut0 is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCut *DiamondCutFilterer) WatchDiamondCut0(opts *bind.WatchOpts, sink chan<- *DiamondCutDiamondCut0) (event.Subscription, error) {

	logs, sub, err := _DiamondCut.contract.WatchLogs(opts, "DiamondCut0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DiamondCutDiamondCut0)
				if err := _DiamondCut.contract.UnpackLog(event, "DiamondCut0", log); err != nil {
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

// ParseDiamondCut0 is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] _diamondCut, address _init, bytes _calldata)
func (_DiamondCut *DiamondCutFilterer) ParseDiamondCut0(log types.Log) (*DiamondCutDiamondCut0, error) {
	event := new(DiamondCutDiamondCut0)
	if err := _DiamondCut.contract.UnpackLog(event, "DiamondCut0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
