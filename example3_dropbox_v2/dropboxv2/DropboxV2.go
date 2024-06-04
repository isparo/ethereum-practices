// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dropboxv2

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

// DropboxV2File is an auto generated low-level Go binding around an user-defined struct.
type DropboxV2File struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}

// Dropboxv2MetaData contains all meta data concerning the Dropboxv2 contract.
var Dropboxv2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"uploadTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"}],\"name\":\"FileUploaded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"uploadTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFileCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listFiles\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fileId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"fileDescription\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"uploadTime\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"uploader\",\"type\":\"address\"}],\"internalType\":\"structDropboxV2.File[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_fileHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_fileSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_fileType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fileDescription\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"uploadFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526040518060400160405280600981526020017f44726f70626f78563200000000000000000000000000000000000000000000008152505f90816100479190610297565b505f600155348015610057575f80fd5b50610366565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806100d857607f821691505b6020821081036100eb576100ea610094565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261014d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610112565b6101578683610112565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61019b6101966101918461016f565b610178565b61016f565b9050919050565b5f819050919050565b6101b483610181565b6101c86101c0826101a2565b84845461011e565b825550505050565b5f90565b6101dc6101d0565b6101e78184846101ab565b505050565b5b8181101561020a576101ff5f826101d4565b6001810190506101ed565b5050565b601f82111561024f57610220816100f1565b61022984610103565b81016020851015610238578190505b61024c61024485610103565b8301826101ec565b50505b505050565b5f82821c905092915050565b5f61026f5f1984600802610254565b1980831691505092915050565b5f6102878383610260565b9150826002028217905092915050565b6102a08261005d565b67ffffffffffffffff8111156102b9576102b8610067565b5b6102c382546100c1565b6102ce82828561020e565b5f60209050601f8311600181146102ff575f84156102ed578287015190505b6102f7858261027c565b86555061035e565b601f19841661030d866100f1565b5f5b828110156103345784890151825560018201915060208501945060208101905061030f565b86831015610351578489015161034d601f891682610260565b8355505b6001600288020188555050505b505050505050565b61177f806103735f395ff3fe608060405234801561000f575f80fd5b5060043610610060575f3560e01c806306fdde03146100645780634395351914610082578063b033d8fa146100a0578063bab50cc9146100be578063be46386f146100dc578063f4c714b4146100f8575b5f80fd5b61006c61012f565b6040516100799190610bbd565b60405180910390f35b61008a6101ba565b6040516100979190610bf5565b60405180910390f35b6100a86101c0565b6040516100b59190610e21565b60405180910390f35b6100c6610546565b6040516100d39190610bf5565b60405180910390f35b6100f660048036038101906100f19190611011565b61054f565b005b610112600480360381019061010d9190611132565b610786565b60405161012698979695949392919061116c565b60405180910390f35b5f805461013b90611231565b80601f016020809104026020016040519081016040528092919081815260200182805461016790611231565b80156101b25780601f10610189576101008083540402835291602001916101b2565b820191905f5260205f20905b81548152906001019060200180831161019557829003601f168201915b505050505081565b60015481565b60605f6101cb610546565b90505f8167ffffffffffffffff8111156101e8576101e7610e5a565b5b60405190808252806020026020018201604052801561022157816020015b61020e610af6565b8152602001906001900390816102065790505b5090505f600190505b82811161053d5760025f8281526020019081526020015f20604051806101000160405290815f820154815260200160018201805461026790611231565b80601f016020809104026020016040519081016040528092919081815260200182805461029390611231565b80156102de5780601f106102b5576101008083540402835291602001916102de565b820191905f5260205f20905b8154815290600101906020018083116102c157829003601f168201915b505050505081526020016002820154815260200160038201805461030190611231565b80601f016020809104026020016040519081016040528092919081815260200182805461032d90611231565b80156103785780601f1061034f57610100808354040283529160200191610378565b820191905f5260205f20905b81548152906001019060200180831161035b57829003601f168201915b5050505050815260200160048201805461039190611231565b80601f01602080910402602001604051908101604052809291908181526020018280546103bd90611231565b80156104085780601f106103df57610100808354040283529160200191610408565b820191905f5260205f20905b8154815290600101906020018083116103eb57829003601f168201915b5050505050815260200160058201805461042190611231565b80601f016020809104026020016040519081016040528092919081815260200182805461044d90611231565b80156104985780601f1061046f57610100808354040283529160200191610498565b820191905f5260205f20905b81548152906001019060200180831161047b57829003601f168201915b5050505050815260200160068201548152602001600782015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250508260018361050e919061128e565b8151811061051f5761051e6112c1565b5b60200260200101819052508080610535906112ee565b91505061022a565b50809250505090565b5f600154905090565b5f88604051602001610561919061136f565b6040516020818303038152906040528051906020012090505f61058682868686610a01565b90505f8a5111610594575f80fd5b5f8851116105a0575f80fd5b5f8651116105ac575f80fd5b5f8751116105b8575f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16036105ef575f80fd5b5f89116105fa575f80fd5b60015f81548092919061060c906112ee565b919050555060405180610100016040528060015481526020018b81526020018a81526020018981526020018881526020018781526020014281526020018273ffffffffffffffffffffffffffffffffffffffff1681525060025f60015481526020019081526020015f205f820151815f015560208201518160010190816106939190611522565b506040820151816002015560608201518160030190816106b39190611522565b5060808201518160040190816106c99190611522565b5060a08201518160050190816106df9190611522565b5060c0820151816006015560e0820151816007015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050507f6e95bbaa37a5fdec42a312e049dea2fe6160ebe659214d3409b4f91f164c2f526001548b8b8b8b8b428860405161077298979695949392919061116c565b60405180910390a150505050505050505050565b6002602052805f5260405f205f91509050805f0154908060010180546107ab90611231565b80601f01602080910402602001604051908101604052809291908181526020018280546107d790611231565b80156108225780601f106107f957610100808354040283529160200191610822565b820191905f5260205f20905b81548152906001019060200180831161080557829003601f168201915b50505050509080600201549080600301805461083d90611231565b80601f016020809104026020016040519081016040528092919081815260200182805461086990611231565b80156108b45780601f1061088b576101008083540402835291602001916108b4565b820191905f5260205f20905b81548152906001019060200180831161089757829003601f168201915b5050505050908060040180546108c990611231565b80601f01602080910402602001604051908101604052809291908181526020018280546108f590611231565b80156109405780601f1061091757610100808354040283529160200191610940565b820191905f5260205f20905b81548152906001019060200180831161092357829003601f168201915b50505050509080600501805461095590611231565b80601f016020809104026020016040519081016040528092919081815260200182805461098190611231565b80156109cc5780601f106109a3576101008083540402835291602001916109cc565b820191905f5260205f20905b8154815290600101906020018083116109af57829003601f168201915b505050505090806006015490806007015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905088565b5f8085604051602001610a14919061165b565b6040516020818303038152906040528051906020012090505f6001828787876040515f8152602001604052604051610a4f949392919061169e565b6020604051602081039080840390855afa158015610a6f573d5f803e3d5ffd5b5050506020604051035190505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ae9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ae09061172b565b60405180910390fd5b8092505050949350505050565b6040518061010001604052805f8152602001606081526020015f81526020016060815260200160608152602001606081526020015f81526020015f73ffffffffffffffffffffffffffffffffffffffff1681525090565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610b8f82610b4d565b610b998185610b57565b9350610ba9818560208601610b67565b610bb281610b75565b840191505092915050565b5f6020820190508181035f830152610bd58184610b85565b905092915050565b5f819050919050565b610bef81610bdd565b82525050565b5f602082019050610c085f830184610be6565b92915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610c4081610bdd565b82525050565b5f82825260208201905092915050565b5f610c6082610b4d565b610c6a8185610c46565b9350610c7a818560208601610b67565b610c8381610b75565b840191505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610cb782610c8e565b9050919050565b610cc781610cad565b82525050565b5f61010083015f830151610ce35f860182610c37565b5060208301518482036020860152610cfb8282610c56565b9150506040830151610d106040860182610c37565b5060608301518482036060860152610d288282610c56565b91505060808301518482036080860152610d428282610c56565b91505060a083015184820360a0860152610d5c8282610c56565b91505060c0830151610d7160c0860182610c37565b5060e0830151610d8460e0860182610cbe565b508091505092915050565b5f610d9a8383610ccd565b905092915050565b5f602082019050919050565b5f610db882610c0e565b610dc28185610c18565b935083602082028501610dd485610c28565b805f5b85811015610e0f5784840389528151610df08582610d8f565b9450610dfb83610da2565b925060208a01995050600181019050610dd7565b50829750879550505050505092915050565b5f6020820190508181035f830152610e398184610dae565b905092915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610e9082610b75565b810181811067ffffffffffffffff82111715610eaf57610eae610e5a565b5b80604052505050565b5f610ec1610e41565b9050610ecd8282610e87565b919050565b5f67ffffffffffffffff821115610eec57610eeb610e5a565b5b610ef582610b75565b9050602081019050919050565b828183375f83830152505050565b5f610f22610f1d84610ed2565b610eb8565b905082815260208101848484011115610f3e57610f3d610e56565b5b610f49848285610f02565b509392505050565b5f82601f830112610f6557610f64610e52565b5b8135610f75848260208601610f10565b91505092915050565b610f8781610bdd565b8114610f91575f80fd5b50565b5f81359050610fa281610f7e565b92915050565b5f60ff82169050919050565b610fbd81610fa8565b8114610fc7575f80fd5b50565b5f81359050610fd881610fb4565b92915050565b5f819050919050565b610ff081610fde565b8114610ffa575f80fd5b50565b5f8135905061100b81610fe7565b92915050565b5f805f805f805f80610100898b03121561102e5761102d610e4a565b5b5f89013567ffffffffffffffff81111561104b5761104a610e4e565b5b6110578b828c01610f51565b98505060206110688b828c01610f94565b975050604089013567ffffffffffffffff81111561108957611088610e4e565b5b6110958b828c01610f51565b965050606089013567ffffffffffffffff8111156110b6576110b5610e4e565b5b6110c28b828c01610f51565b955050608089013567ffffffffffffffff8111156110e3576110e2610e4e565b5b6110ef8b828c01610f51565b94505060a06111008b828c01610fca565b93505060c06111118b828c01610ffd565b92505060e06111228b828c01610ffd565b9150509295985092959890939650565b5f6020828403121561114757611146610e4a565b5b5f61115484828501610f94565b91505092915050565b61116681610cad565b82525050565b5f610100820190506111805f83018b610be6565b8181036020830152611192818a610b85565b90506111a16040830189610be6565b81810360608301526111b38188610b85565b905081810360808301526111c78187610b85565b905081810360a08301526111db8186610b85565b90506111ea60c0830185610be6565b6111f760e083018461115d565b9998505050505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061124857607f821691505b60208210810361125b5761125a611204565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61129882610bdd565b91506112a383610bdd565b92508282039050818111156112bb576112ba611261565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f6112f882610bdd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361132a57611329611261565b5b600182019050919050565b5f81905092915050565b5f61134982610b4d565b6113538185611335565b9350611363818560208601610b67565b80840191505092915050565b5f61137a828461133f565b915081905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026113e17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826113a6565b6113eb86836113a6565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61142661142161141c84610bdd565b611403565b610bdd565b9050919050565b5f819050919050565b61143f8361140c565b61145361144b8261142d565b8484546113b2565b825550505050565b5f90565b61146761145b565b611472818484611436565b505050565b5b818110156114955761148a5f8261145f565b600181019050611478565b5050565b601f8211156114da576114ab81611385565b6114b484611397565b810160208510156114c3578190505b6114d76114cf85611397565b830182611477565b50505b505050565b5f82821c905092915050565b5f6114fa5f19846008026114df565b1980831691505092915050565b5f61151283836114eb565b9150826002028217905092915050565b61152b82610b4d565b67ffffffffffffffff81111561154457611543610e5a565b5b61154e8254611231565b611559828285611499565b5f60209050601f83116001811461158a575f8415611578578287015190505b6115828582611507565b8655506115e9565b601f19841661159886611385565b5f5b828110156115bf5784890151825560018201915060208501945060208101905061159a565b868310156115dc57848901516115d8601f8916826114eb565b8355505b6001600288020188555050505b505050505050565b7f19457468657265756d205369676e6564204d6573736167653a0a3332000000005f82015250565b5f611625601c83611335565b9150611630826115f1565b601c82019050919050565b5f819050919050565b61165561165082610fde565b61163b565b82525050565b5f61166582611619565b91506116718284611644565b60208201915081905092915050565b61168981610fde565b82525050565b61169881610fa8565b82525050565b5f6080820190506116b15f830187611680565b6116be602083018661168f565b6116cb6040830185611680565b6116d86060830184611680565b95945050505050565b7f496e76616c6964207369676e61747572650000000000000000000000000000005f82015250565b5f611715601183610b57565b9150611720826116e1565b602082019050919050565b5f6020820190508181035f83015261174281611709565b905091905056fea2646970667358221220d6ec62f1fb207a3454ea2368c7d07b523f5cd10b9ab077e86fc887143d40119764736f6c63430008190033",
}

// Dropboxv2ABI is the input ABI used to generate the binding from.
// Deprecated: Use Dropboxv2MetaData.ABI instead.
var Dropboxv2ABI = Dropboxv2MetaData.ABI

// Dropboxv2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use Dropboxv2MetaData.Bin instead.
var Dropboxv2Bin = Dropboxv2MetaData.Bin

// DeployDropboxv2 deploys a new Ethereum contract, binding an instance of Dropboxv2 to it.
func DeployDropboxv2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Dropboxv2, error) {
	parsed, err := Dropboxv2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(Dropboxv2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dropboxv2{Dropboxv2Caller: Dropboxv2Caller{contract: contract}, Dropboxv2Transactor: Dropboxv2Transactor{contract: contract}, Dropboxv2Filterer: Dropboxv2Filterer{contract: contract}}, nil
}

// Dropboxv2 is an auto generated Go binding around an Ethereum contract.
type Dropboxv2 struct {
	Dropboxv2Caller     // Read-only binding to the contract
	Dropboxv2Transactor // Write-only binding to the contract
	Dropboxv2Filterer   // Log filterer for contract events
}

// Dropboxv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Dropboxv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dropboxv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Dropboxv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dropboxv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Dropboxv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Dropboxv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Dropboxv2Session struct {
	Contract     *Dropboxv2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Dropboxv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Dropboxv2CallerSession struct {
	Contract *Dropboxv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Dropboxv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Dropboxv2TransactorSession struct {
	Contract     *Dropboxv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Dropboxv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Dropboxv2Raw struct {
	Contract *Dropboxv2 // Generic contract binding to access the raw methods on
}

// Dropboxv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Dropboxv2CallerRaw struct {
	Contract *Dropboxv2Caller // Generic read-only contract binding to access the raw methods on
}

// Dropboxv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Dropboxv2TransactorRaw struct {
	Contract *Dropboxv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDropboxv2 creates a new instance of Dropboxv2, bound to a specific deployed contract.
func NewDropboxv2(address common.Address, backend bind.ContractBackend) (*Dropboxv2, error) {
	contract, err := bindDropboxv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dropboxv2{Dropboxv2Caller: Dropboxv2Caller{contract: contract}, Dropboxv2Transactor: Dropboxv2Transactor{contract: contract}, Dropboxv2Filterer: Dropboxv2Filterer{contract: contract}}, nil
}

// NewDropboxv2Caller creates a new read-only instance of Dropboxv2, bound to a specific deployed contract.
func NewDropboxv2Caller(address common.Address, caller bind.ContractCaller) (*Dropboxv2Caller, error) {
	contract, err := bindDropboxv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Dropboxv2Caller{contract: contract}, nil
}

// NewDropboxv2Transactor creates a new write-only instance of Dropboxv2, bound to a specific deployed contract.
func NewDropboxv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Dropboxv2Transactor, error) {
	contract, err := bindDropboxv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Dropboxv2Transactor{contract: contract}, nil
}

// NewDropboxv2Filterer creates a new log filterer instance of Dropboxv2, bound to a specific deployed contract.
func NewDropboxv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Dropboxv2Filterer, error) {
	contract, err := bindDropboxv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Dropboxv2Filterer{contract: contract}, nil
}

// bindDropboxv2 binds a generic wrapper to an already deployed contract.
func bindDropboxv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Dropboxv2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dropboxv2 *Dropboxv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dropboxv2.Contract.Dropboxv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dropboxv2 *Dropboxv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dropboxv2.Contract.Dropboxv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dropboxv2 *Dropboxv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dropboxv2.Contract.Dropboxv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dropboxv2 *Dropboxv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dropboxv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dropboxv2 *Dropboxv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dropboxv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dropboxv2 *Dropboxv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dropboxv2.Contract.contract.Transact(opts, method, params...)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_Dropboxv2 *Dropboxv2Caller) FileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dropboxv2.contract.Call(opts, &out, "fileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_Dropboxv2 *Dropboxv2Session) FileCount() (*big.Int, error) {
	return _Dropboxv2.Contract.FileCount(&_Dropboxv2.CallOpts)
}

// FileCount is a free data retrieval call binding the contract method 0x43953519.
//
// Solidity: function fileCount() view returns(uint256)
func (_Dropboxv2 *Dropboxv2CallerSession) FileCount() (*big.Int, error) {
	return _Dropboxv2.Contract.FileCount(&_Dropboxv2.CallOpts)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Dropboxv2 *Dropboxv2Caller) Files(opts *bind.CallOpts, arg0 *big.Int) (struct {
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
	err := _Dropboxv2.contract.Call(opts, &out, "files", arg0)

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
func (_Dropboxv2 *Dropboxv2Session) Files(arg0 *big.Int) (struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}, error) {
	return _Dropboxv2.Contract.Files(&_Dropboxv2.CallOpts, arg0)
}

// Files is a free data retrieval call binding the contract method 0xf4c714b4.
//
// Solidity: function files(uint256 ) view returns(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Dropboxv2 *Dropboxv2CallerSession) Files(arg0 *big.Int) (struct {
	FileId          *big.Int
	FileHash        string
	FileSize        *big.Int
	FileType        string
	FileName        string
	FileDescription string
	UploadTime      *big.Int
	Uploader        common.Address
}, error) {
	return _Dropboxv2.Contract.Files(&_Dropboxv2.CallOpts, arg0)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Dropboxv2 *Dropboxv2Caller) GetFileCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dropboxv2.contract.Call(opts, &out, "getFileCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Dropboxv2 *Dropboxv2Session) GetFileCount() (*big.Int, error) {
	return _Dropboxv2.Contract.GetFileCount(&_Dropboxv2.CallOpts)
}

// GetFileCount is a free data retrieval call binding the contract method 0xbab50cc9.
//
// Solidity: function getFileCount() view returns(uint256)
func (_Dropboxv2 *Dropboxv2CallerSession) GetFileCount() (*big.Int, error) {
	return _Dropboxv2.Contract.GetFileCount(&_Dropboxv2.CallOpts)
}

// ListFiles is a free data retrieval call binding the contract method 0xb033d8fa.
//
// Solidity: function listFiles() view returns((uint256,string,uint256,string,string,string,uint256,address)[])
func (_Dropboxv2 *Dropboxv2Caller) ListFiles(opts *bind.CallOpts) ([]DropboxV2File, error) {
	var out []interface{}
	err := _Dropboxv2.contract.Call(opts, &out, "listFiles")

	if err != nil {
		return *new([]DropboxV2File), err
	}

	out0 := *abi.ConvertType(out[0], new([]DropboxV2File)).(*[]DropboxV2File)

	return out0, err

}

// ListFiles is a free data retrieval call binding the contract method 0xb033d8fa.
//
// Solidity: function listFiles() view returns((uint256,string,uint256,string,string,string,uint256,address)[])
func (_Dropboxv2 *Dropboxv2Session) ListFiles() ([]DropboxV2File, error) {
	return _Dropboxv2.Contract.ListFiles(&_Dropboxv2.CallOpts)
}

// ListFiles is a free data retrieval call binding the contract method 0xb033d8fa.
//
// Solidity: function listFiles() view returns((uint256,string,uint256,string,string,string,uint256,address)[])
func (_Dropboxv2 *Dropboxv2CallerSession) ListFiles() ([]DropboxV2File, error) {
	return _Dropboxv2.Contract.ListFiles(&_Dropboxv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dropboxv2 *Dropboxv2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dropboxv2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dropboxv2 *Dropboxv2Session) Name() (string, error) {
	return _Dropboxv2.Contract.Name(&_Dropboxv2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dropboxv2 *Dropboxv2CallerSession) Name() (string, error) {
	return _Dropboxv2.Contract.Name(&_Dropboxv2.CallOpts)
}

// UploadFile is a paid mutator transaction binding the contract method 0xbe46386f.
//
// Solidity: function uploadFile(string _fileHash, uint256 _fileSize, string _fileType, string _fileName, string _fileDescription, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dropboxv2 *Dropboxv2Transactor) UploadFile(opts *bind.TransactOpts, _fileHash string, _fileSize *big.Int, _fileType string, _fileName string, _fileDescription string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dropboxv2.contract.Transact(opts, "uploadFile", _fileHash, _fileSize, _fileType, _fileName, _fileDescription, v, r, s)
}

// UploadFile is a paid mutator transaction binding the contract method 0xbe46386f.
//
// Solidity: function uploadFile(string _fileHash, uint256 _fileSize, string _fileType, string _fileName, string _fileDescription, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dropboxv2 *Dropboxv2Session) UploadFile(_fileHash string, _fileSize *big.Int, _fileType string, _fileName string, _fileDescription string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dropboxv2.Contract.UploadFile(&_Dropboxv2.TransactOpts, _fileHash, _fileSize, _fileType, _fileName, _fileDescription, v, r, s)
}

// UploadFile is a paid mutator transaction binding the contract method 0xbe46386f.
//
// Solidity: function uploadFile(string _fileHash, uint256 _fileSize, string _fileType, string _fileName, string _fileDescription, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dropboxv2 *Dropboxv2TransactorSession) UploadFile(_fileHash string, _fileSize *big.Int, _fileType string, _fileName string, _fileDescription string, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dropboxv2.Contract.UploadFile(&_Dropboxv2.TransactOpts, _fileHash, _fileSize, _fileType, _fileName, _fileDescription, v, r, s)
}

// Dropboxv2FileUploadedIterator is returned from FilterFileUploaded and is used to iterate over the raw logs and unpacked data for FileUploaded events raised by the Dropboxv2 contract.
type Dropboxv2FileUploadedIterator struct {
	Event *Dropboxv2FileUploaded // Event containing the contract specifics and raw log

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
func (it *Dropboxv2FileUploadedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Dropboxv2FileUploaded)
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
		it.Event = new(Dropboxv2FileUploaded)
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
func (it *Dropboxv2FileUploadedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Dropboxv2FileUploadedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Dropboxv2FileUploaded represents a FileUploaded event raised by the Dropboxv2 contract.
type Dropboxv2FileUploaded struct {
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
func (_Dropboxv2 *Dropboxv2Filterer) FilterFileUploaded(opts *bind.FilterOpts) (*Dropboxv2FileUploadedIterator, error) {

	logs, sub, err := _Dropboxv2.contract.FilterLogs(opts, "FileUploaded")
	if err != nil {
		return nil, err
	}
	return &Dropboxv2FileUploadedIterator{contract: _Dropboxv2.contract, event: "FileUploaded", logs: logs, sub: sub}, nil
}

// WatchFileUploaded is a free log subscription operation binding the contract event 0x6e95bbaa37a5fdec42a312e049dea2fe6160ebe659214d3409b4f91f164c2f52.
//
// Solidity: event FileUploaded(uint256 fileId, string fileHash, uint256 fileSize, string fileType, string fileName, string fileDescription, uint256 uploadTime, address uploader)
func (_Dropboxv2 *Dropboxv2Filterer) WatchFileUploaded(opts *bind.WatchOpts, sink chan<- *Dropboxv2FileUploaded) (event.Subscription, error) {

	logs, sub, err := _Dropboxv2.contract.WatchLogs(opts, "FileUploaded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Dropboxv2FileUploaded)
				if err := _Dropboxv2.contract.UnpackLog(event, "FileUploaded", log); err != nil {
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
func (_Dropboxv2 *Dropboxv2Filterer) ParseFileUploaded(log types.Log) (*Dropboxv2FileUploaded, error) {
	event := new(Dropboxv2FileUploaded)
	if err := _Dropboxv2.contract.UnpackLog(event, "FileUploaded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
