// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package simpledropbox

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

// SimpleDropboxFile is an auto generated low-level Go binding around an user-defined struct.
type SimpleDropboxFile struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}

// SimpledropboxMetaData contains all meta data concerning the Simpledropbox contract.
var SimpledropboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uploadTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"uploadTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"uploadTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"}],\"internalType\":\"structSimpleDropbox.File[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileDescription\",\"type\":\"string\"}],\"name\":\"uploadFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600d81526020017f53696d706c6544726f70626f78000000000000000000000000000000000000008152505f90816100479190610297565b505f600155348015610057575f80fd5b50610366565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d857607f821691505b6020821081036100eb576100ea610094565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261014d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610112565b6101578683610112565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019b6101966101918461016f565b610178565b61016f565b9050919050565b5f819050919050565b6101b483610181565b6101c86101c0826101a2565b84845461011e565b825550505050565b5f90565b6101dc6101d0565b6101e78184846101ab565b505050565b5b8181101561020a576101ff5f826101d4565b6001810190506101ed565b5050565b601f82111561024f57610220816100f1565b61022984610103565b81016020851015610238578190505b61024c61024485610103565b8301826101ec565b50505b505050565b5f82821c905092915050565b5f61026f5f1984600802610254565b1980831691505092915050565b5f6102878383610260565b9150826002028217905092915050565b6102a08261005d565b67ffffffffffffffff8111156102b9576102b8610067565b5b6102c382546100c1565b6102ce82828561020e565b5f60209050601f8311600181146102ff575f84156102ed578287015190505b6102f7858261027c565b86555061035e565b601f19841661030d866100f1565b5f5b828110156103345784890151825560018201915060208501945060208101905061030f565b86831015610351578489015161034d601f891682610260565b8355505b6001600288020188555050505b505050505050565b611401806103735f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806306fdde03146100645780634395351914610082578063b033d8fa146100a0578063bab50cc9146100be578063e82d9d33146100dc578063f4c714b4146100f8575b5f80fd5b61006c61012f565b6040516100799190610a8a565b60405180910390f35b61008a6101ba565b6040516100979190610ac2565b60405180910390f35b6100a86101c0565b6040516100b59190610cee565b60405180910390f35b6100c6610546565b6040516100d39190610ac2565b60405180910390f35b6100f660048036038101906100f19190610e75565b61054f565b005b610112600480360381019061010d9190610f5c565b610748565b604051610126989796959493929190610f96565b60405180910390f35b5f805461013b9061105b565b80601f01602080910402602001604051908101604052809291908181526020018280546101679061105b565b80156101b25780601f10610189576101008083540402835291602001916101b2565b820191905f5260205f20905b81548152906001019060200180831161019557829003601f168201915b505050505081565b60015481565b60605f6101cb610546565b90505f8167ffffffffffffffff8111156101e8576101e7610d27565b5b60405190808252806020026020018201604052801561022157816020015b61020e6109c3565b8152602001906001900390816102065790505b5090505f600190505b82811161053d5760025f8281526020019081526020015f20604051806101000160405290815f82015481526020016001820180546102679061105b565b80601f01602080910402602001604051908101604052809291908181526020018280546102939061105b565b80156102de5780601f106102b5576101008083540402835291602001916102de565b820191905f5260205f20905b8154815290600101906020018083116102c157829003601f168201915b50505050508152602001600282015481526020016003820180546103019061105b565b80601f016020809104026020016040519081016040528092919081815260200182805461032d9061105b565b80156103785780601f1061034f57610100808354040283529160200191610378565b820191905f5260205f20905b81548152906001019060200180831161035b57829003601f168201915b505050505081526020016004820180546103919061105b565b80601f01602080910402602001604051908101604052809291908181526020018280546103bd9061105b565b80156104085780601f106103df57610100808354040283529160200191610408565b820191905f5260205f20905b8154815290600101906020018083116103eb57829003601f168201915b505050505081526020016005820180546104219061105b565b80601f016020809104026020016040519081016040528092919081815260200182805461044d9061105b565b80156104985780601f1061046f57610100808354040283529160200191610498565b820191905f5260205f20905b81548152906001019060200180831161047b57829003601f168201915b5050505050815260200160068201548152602001600782015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250508260018361050e91906110b8565b8151811061051f5761051e6110eb565b5b6020026020010181905250808061053590611118565b91505061022a565b50809250505090565b5f600154905090565b5f85511161055b575f80fd5b5f835111610567575f80fd5b5f815111610573575f80fd5b5f82511161057f575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036105b6575f80fd5b5f84116105c1575f80fd5b60015f8154809291906105d390611118565b919050555060405180610100016040528060015481526020018681526020018581526020018481526020018381526020018281526020014281526020013373ffffffffffffffffffffffffffffffffffffffff1681525060025f60015481526020019081526020015f205f820151815f0155602082015181600101908161065a91906112fc565b5060408201518160020155606082015181600301908161067a91906112fc565b50608082015181600401908161069091906112fc565b5060a08201518160050190816106a691906112fc565b5060c0820151816006015560e0820151816007015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f6e95bbaa37a5fdec42a312e049dea2fe6160ebe659214d3409b4f91f164c2f5260015486868686864233604051610739989796959493929190610f96565b60405180910390a15050505050565b6002602052805f5260405f205f91509050805f01549080600101805461076d9061105b565b80601f01602080910402602001604051908101604052809291908181526020018280546107999061105b565b80156107e45780601f106107bb576101008083540402835291602001916107e4565b820191905f5260205f20905b8154815290600101906020018083116107c757829003601f168201915b5050505050908060020154908060030180546107ff9061105b565b80601f016020809104026020016040519081016040528092919081815260200182805461082b9061105b565b80156108765780601f1061084d57610100808354040283529160200191610876565b820191905f5260205f20905b81548152906001019060200180831161085957829003601f168201915b50505050509080600401805461088b9061105b565b80601f01602080910402602001604051908101604052809291908181526020018280546108b79061105b565b80156109025780601f106108d957610100808354040283529160200191610902565b820191905f5260205f20905b8154815290600101906020018083116108e557829003601f168201915b5050505050908060050180546109179061105b565b80601f01602080910402602001604051908101604052809291908181526020018280546109439061105b565b801561098e5780601f106109655761010080835404028352916020019161098e565b820191905f5260205f20905b81548152906001019060200180831161097157829003601f168201915b505050505090806006015490806007015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905088565b6040518061010001604052805f8152602001606081526020015f81526020016060815260200160608152602001606081526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610a5c82610a1a565b610a668185610a24565b9350610a76818560208601610a34565b610a7f81610a42565b840191505092915050565b5f6020820190508181035f830152610aa28184610a52565b905092915050565b5f819050919050565b610abc81610aaa565b82525050565b5f602082019050610ad55f830184610ab3565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b0d81610aaa565b82525050565b5f82825260208201905092915050565b5f610b2d82610a1a565b610b378185610b13565b9350610b47818560208601610a34565b610b5081610a42565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610b8482610b5b565b9050919050565b610b9481610b7a565b82525050565b5f61010083015f830151610bb05f860182610b04565b5060208301518482036020860152610bc88282610b23565b9150506040830151610bdd6040860182610b04565b5060608301518482036060860152610bf58282610b23565b91505060808301518482036080860152610c0f8282610b23565b91505060a083015184820360a0860152610c298282610b23565b91505060c0830151610c3e60c0860182610b04565b5060e0830151610c5160e0860182610b8b565b508091505092915050565b5f610c678383610b9a565b905092915050565b5f602082019050919050565b5f610c8582610adb565b610c8f8185610ae5565b935083602082028501610ca185610af5565b805f5b85811015610cdc5784840389528151610cbd8582610c5c565b9450610cc883610c6f565b925060208a01995050600181019050610ca4565b50829750879550505050505092915050565b5f6020820190508181035f830152610d068184610c7b565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610d5d82610a42565b810181811067ffffffffffffffff82111715610d7c57610d7b610d27565b5b80604052505050565b5f610d8e610d0e565b9050610d9a8282610d54565b919050565b5f67ffffffffffffffff821115610db957610db8610d27565b5b610dc282610a42565b9050602081019050919050565b828183375f83830152505050565b5f610def610dea84610d9f565b610d85565b905082815260208101848484011115610e0b57610e0a610d23565b5b610e16848285610dcf565b509392505050565b5f82601f830112610e3257610e31610d1f565b5b8135610e42848260208601610ddd565b91505092915050565b610e5481610aaa565b8114610e5e575f80fd5b50565b5f81359050610e6f81610e4b565b92915050565b5f805f805f60a08688031215610e8e57610e8d610d17565b5b5f86013567ffffffffffffffff811115610eab57610eaa610d1b565b5b610eb788828901610e1e565b9550506020610ec888828901610e61565b945050604086013567ffffffffffffffff811115610ee957610ee8610d1b565b5b610ef588828901610e1e565b935050606086013567ffffffffffffffff811115610f1657610f15610d1b565b5b610f2288828901610e1e565b925050608086013567ffffffffffffffff811115610f4357610f42610d1b565b5b610f4f88828901610e1e565b9150509295509295909350565b5f60208284031215610f7157610f70610d17565b5b5f610f7e84828501610e61565b91505092915050565b610f9081610b7a565b82525050565b5f61010082019050610faa5f83018b610ab3565b8181036020830152610fbc818a610a52565b9050610fcb6040830189610ab3565b8181036060830152610fdd8188610a52565b90508181036080830152610ff18187610a52565b905081810360a08301526110058186610a52565b905061101460c0830185610ab3565b61102160e0830184610f87565b9998505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061107257607f821691505b6020821081036110855761108461102e565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6110c282610aaa565b91506110cd83610aaa565b92508282039050818111156110e5576110e461108b565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f61112282610aaa565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036111545761115361108b565b5b600182019050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026111bb7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611180565b6111c58683611180565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6112006111fb6111f684610aaa565b6111dd565b610aaa565b9050919050565b5f819050919050565b611219836111e6565b61122d61122582611207565b84845461118c565b825550505050565b5f90565b611241611235565b61124c818484611210565b505050565b5b8181101561126f576112645f82611239565b600181019050611252565b5050565b601f8211156112b4576112858161115f565b61128e84611171565b8101602085101561129d578190505b6112b16112a985611171565b830182611251565b50505b505050565b5f82821c905092915050565b5f6112d45f19846008026112b9565b1980831691505092915050565b5f6112ec83836112c5565b9150826002028217905092915050565b61130582610a1a565b67ffffffffffffffff81111561131e5761131d610d27565b5b611328825461105b565b611333828285611273565b5f60209050601f831160018114611364575f8415611352578287015190505b61135c85826112e1565b8655506113c3565b601f1984166113728661115f565b5f5b8281101561139957848901518255600182019150602085019450602081019050611374565b868310156113b657848901516113b2601f8916826112c5565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220be6127a8a3dfcc22e75bc0702798af0b0b03558eb42e5e2719fa68e7e56ce2f164736f6c63430008190033",
}

// SimpledropboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SimpledropboxMetaData.ABI instead.
var SimpledropboxABI = SimpledropboxMetaData.ABI

// SimpledropboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SimpledropboxMetaData.Bin instead.
var SimpledropboxBin = SimpledropboxMetaData.Bin

// DeploySimpledropbox deploys a new Ethereum contract, binding an instance of Simpledropbox to it.
func DeploySimpledropbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Simpledropbox, error) {
	parsed, err := SimpledropboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SimpledropboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Simpledropbox{SimpledropboxCaller: SimpledropboxCaller{contract: contract}, SimpledropboxTransactor: SimpledropboxTransactor{contract: contract}, SimpledropboxFilterer: SimpledropboxFilterer{contract: contract}}, nil
}

// Simpledropbox is an auto generated Go binding around an Ethereum contract.
type Simpledropbox struct {
	SimpledropboxCaller     // Read-only binding to the contract
	SimpledropboxTransactor // Write-only binding to the contract
	SimpledropboxFilterer   // Log filterer for contract events
}

// SimpledropboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SimpledropboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpledropboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SimpledropboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpledropboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SimpledropboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SimpledropboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SimpledropboxSession struct {
	Contract     *Simpledropbox    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SimpledropboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SimpledropboxCallerSession struct {
	Contract *SimpledropboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SimpledropboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SimpledropboxTransactorSession struct {
	Contract     *SimpledropboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SimpledropboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SimpledropboxRaw struct {
	Contract *Simpledropbox // Generic contract binding to access the raw methods on
}

// SimpledropboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SimpledropboxCallerRaw struct {
	Contract *SimpledropboxCaller // Generic read-only contract binding to access the raw methods on
}

// SimpledropboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SimpledropboxTransactorRaw struct {
	Contract *SimpledropboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSimpledropbox creates a new instance of Simpledropbox, bound to a specific deployed contract.
func NewSimpledropbox(address common.Address, backend bind.ContractBackend) (*Simpledropbox, error) {
	contract, err := bindSimpledropbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Simpledropbox{SimpledropboxCaller: SimpledropboxCaller{contract: contract}, SimpledropboxTransactor: SimpledropboxTransactor{contract: contract}, SimpledropboxFilterer: SimpledropboxFilterer{contract: contract}}, nil
}

// NewSimpledropboxCaller creates a new read-only instance of Simpledropbox, bound to a specific deployed contract.
func NewSimpledropboxCaller(address common.Address, caller bind.ContractCaller) (*SimpledropboxCaller, error) {
	contract, err := bindSimpledropbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SimpledropboxCaller{contract: contract}, nil
}

// NewSimpledropboxTransactor creates a new write-only instance of Simpledropbox, bound to a specific deployed contract.
func NewSimpledropboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SimpledropboxTransactor, error) {
	contract, err := bindSimpledropbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SimpledropboxTransactor{contract: contract}, nil
}

// NewSimpledropboxFilterer creates a new log filterer instance of Simpledropbox, bound to a specific deployed contract.
func NewSimpledropboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SimpledropboxFilterer, error) {
	contract, err := bindSimpledropbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SimpledropboxFilterer{contract: contract}, nil
}

// bindSimpledropbox binds a generic wrapper to an already deployed contract.
func bindSimpledropbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SimpledropboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpledropbox *SimpledropboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpledropbox.Contract.SimpledropboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpledropbox *SimpledropboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpledropbox.Contract.SimpledropboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpledropbox *SimpledropboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpledropbox.Contract.SimpledropboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Simpledropbox *SimpledropboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Simpledropbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Simpledropbox *SimpledropboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Simpledropbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Simpledropbox *SimpledropboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Simpledropbox.Contract.contract.Transact(opts, method, params...)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_Simpledropbox *SimpledropboxCaller) FileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simpledropbox.contract.Call(opts, &out, "fileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_Simpledropbox *SimpledropboxSession) FileCount() (*big.Int, error) {
	return _Simpledropbox.Contract.FileCount(&_Simpledropbox.CallOpts)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_Simpledropbox *SimpledropboxCallerSession) FileCount() (*big.Int, error) {
	return _Simpledropbox.Contract.FileCount(&_Simpledropbox.CallOpts)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Simpledropbox *SimpledropboxCaller) Files(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}, error) {
	var out []interface{}
	err := _Simpledropbox.contract.Call(opts, &out, "files", arg0)

	outstruct := new(struct {
		FileId          *big.Int
		FileHash        string
		FileSize        *big.Int
		FileType        string
		FileName        string
		FileDescription string
		UploadTime      *big.Int
		Uploader        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FileId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FileHash = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.FileSize = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FileType = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.FileName = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.FileDescription = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.UploadTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Uploader = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Simpledropbox *SimpledropboxSession) Files(arg0 *big.Int) (struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}, error) {
	return _Simpledropbox.Contract.Files(&_Simpledropbox.CallOpts, arg0)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Simpledropbox *SimpledropboxCallerSession) Files(arg0 *big.Int) (struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}, error) {
	return _Simpledropbox.Contract.Files(&_Simpledropbox.CallOpts, arg0)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Simpledropbox *SimpledropboxCaller) GetFileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Simpledropbox.contract.Call(opts, &out, "getFileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Simpledropbox *SimpledropboxSession) GetFileCount() (*big.Int, error) {
	return _Simpledropbox.Contract.GetFileCount(&_Simpledropbox.CallOpts)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Simpledropbox *SimpledropboxCallerSession) GetFileCount() (*big.Int, error) {
	return _Simpledropbox.Contract.GetFileCount(&_Simpledropbox.CallOpts)
}

// ListFiles is a free data retrieval call binding the contract method 0xb033d8fa.
//
// Solidity: function listFiles() view returns((uint256,string,uint256,string,string,string,uint256,address)[])
func (_Simpledropbox *SimpledropboxCaller) ListFiles(opts *bind.CallOpts) ([]SimpleDropboxFile, error) {
	var out []interface{}
	err := _Simpledropbox.contract.Call(opts, &out, "listFiles")

	if err != nil {
		return *new([]SimpleDropboxFile), err
	}

	out0 := *abi.ConvertType(out[0], new([]SimpleDropboxFile)).(*[]SimpleDropboxFile)

	return out0, err

}

// ListFiles is a free data retrieval call binding the contract method 0xb033d8fa.
//
// Solidity: function listFiles() view returns((uint256,string,uint256,string,string,string,uint256,address)[])
func (_Simpledropbox *SimpledropboxSession) ListFiles() ([]SimpleDropboxFile, error) {
	return _Simpledropbox.Contract.ListFiles(&_Simpledropbox.CallOpts)
}

// ListFiles is a free data retrieval call binding the contract method 0xb033d8fa.
//
// Solidity: function listFiles() view returns((uint256,string,uint256,string,string,string,uint256,address)[])
func (_Simpledropbox *SimpledropboxCallerSession) ListFiles() ([]SimpleDropboxFile, error) {
	return _Simpledropbox.Contract.ListFiles(&_Simpledropbox.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Simpledropbox *SimpledropboxCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Simpledropbox.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Simpledropbox *SimpledropboxSession) Name() (string, error) {
	return _Simpledropbox.Contract.Name(&_Simpledropbox.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Simpledropbox *SimpledropboxCallerSession) Name() (string, error) {
	return _Simpledropbox.Contract.Name(&_Simpledropbox.CallOpts)
}

// UploadFile is a paid mutator transaction binding the contract method 0xe82d9d33.
//
// Solidity: function uploadFile(string _fileHash, uint256 _fileSize, string _fileType, string _fileName, string _fileDescription) returns()
func (_Simpledropbox *SimpledropboxTransactor) UploadFile(opts *bind.TransactOpts, _fileHash string, _fileSize *big.Int, _fileType string, _fileName string, _fileDescription string) (*types.Transaction, error) {
	return _Simpledropbox.contract.Transact(opts, "uploadFile", _fileHash, _fileSize, _fileType, _fileName, _fileDescription)
}

// UploadFile is a paid mutator transaction binding the contract method 0xe82d9d33.
//
// Solidity: function uploadFile(string _fileHash, uint256 _fileSize, string _fileType, string _fileName, string _fileDescription) returns()
func (_Simpledropbox *SimpledropboxSession) UploadFile(_fileHash string, _fileSize *big.Int, _fileType string, _fileName string, _fileDescription string) (*types.Transaction, error) {
	return _Simpledropbox.Contract.UploadFile(&_Simpledropbox.TransactOpts, _fileHash, _fileSize, _fileType, _fileName, _fileDescription)
}

// UploadFile is a paid mutator transaction binding the contract method 0xe82d9d33.
//
// Solidity: function uploadFile(string _fileHash, uint256 _fileSize, string _fileType, string _fileName, string _fileDescription) returns()
func (_Simpledropbox *SimpledropboxTransactorSession) UploadFile(_fileHash string, _fileSize *big.Int, _fileType string, _fileName string, _fileDescription string) (*types.Transaction, error) {
	return _Simpledropbox.Contract.UploadFile(&_Simpledropbox.TransactOpts, _fileHash, _fileSize, _fileType, _fileName, _fileDescription)
}

// SimpledropboxFileUploadedIterator is returned from FilterFileUploaded and is used to iterate over the raw logs and unpacked data for FileUploaded events raised by the Simpledropbox contract.
type SimpledropboxFileUploadedIterator struct {
	Event *SimpledropboxFileUploaded // Event containing the contract specifics and raw log

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
func (it *SimpledropboxFileUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SimpledropboxFileUploaded)
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
		it.Event = new(SimpledropboxFileUploaded)
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
func (it *SimpledropboxFileUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SimpledropboxFileUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SimpledropboxFileUploaded represents a FileUploaded event raised by the Simpledropbox contract.
type SimpledropboxFileUploaded struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFileUploaded is a free log retrieval operation binding the contract event 0x6e95bbaa37a5fdec42a312e049dea2fe6160ebe659214d3409b4f91f164c2f52.
//
// Solidity: event FileUploaded(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Simpledropbox *SimpledropboxFilterer) FilterFileUploaded(opts *bind.FilterOpts) (*SimpledropboxFileUploadedIterator, error) {

	logs, sub, err := _Simpledropbox.contract.FilterLogs(opts, "FileUploaded")
	if err != nil {
		return nil, err
	}
	return &SimpledropboxFileUploadedIterator{contract: _Simpledropbox.contract, event: "FileUploaded", logs: logs, sub: sub}, nil
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0x6e95bbaa37a5fdec42a312e049dea2fe6160ebe659214d3409b4f91f164c2f52.
//
// Solidity: event FileUploaded(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Simpledropbox *SimpledropboxFilterer) WatchFileUploaded(opts *bind.WatchOpts, sink chan<- *SimpledropboxFileUploaded) (event.Subscription, error) {

	logs, sub, err := _Simpledropbox.contract.WatchLogs(opts, "FileUploaded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SimpledropboxFileUploaded)
				if err := _Simpledropbox.contract.UnpackLog(event, "FileUploaded", log); err != nil {
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

// ParseFileUploaded is a log parse operation binding the contract event 0x6e95bbaa37a5fdec42a312e049dea2fe6160ebe659214d3409b4f91f164c2f52.
//
// Solidity: event FileUploaded(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Simpledropbox *SimpledropboxFilterer) ParseFileUploaded(log types.Log) (*SimpledropboxFileUploaded, error) {
	event := new(SimpledropboxFileUploaded)
	if err := _Simpledropbox.contract.UnpackLog(event, "FileUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
