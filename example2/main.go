package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/josue/ethereum-practices/example2/simpledropbox"
)

const contratAddress string = "0x40167ea261420Dea06E2f5F77C204128EA50B42c"
const privateAccountAddress string = "daf006ccc420fc40f02aafe44e832ce83788d98f0c481d58fea92ecf079e078f"

func main() {
	address := common.HexToAddress(contratAddress)

	client, err := ethclient.Dial("http://127.0.0.1:7545") // URL de Ganache
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a la red Ethereum")
	fmt.Printf("Dirección del contrato: %s\n", address)
	fmt.Printf("Client: %s\n", client)

	instance, err := simpledropbox.NewSimpledropbox(address, client)
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	contractName, _ := instance.Name(callOpts)
	fmt.Println("contract is loaded: ", contractName)

	GetAllfiles(instance)
	upladFile(instance, client)

	GetAllfiles(instance)
}

func GetAllfiles(contractIns *simpledropbox.Simpledropbox) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	result, err := contractIns.ListFiles(callOpts)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range result {
		fmt.Println(v.FileName)
	}
}

func upladFile(contractIns *simpledropbox.Simpledropbox, client *ethclient.Client) {

	privateKey, err := crypto.HexToECDSA(privateAccountAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("public key based on private key: ", publicKey)
	fmt.Println("public key adrees: ", fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("nonce: ", nonce)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("gas price: ", gasPrice)

	chainID := params.AllDevChainProtocolChanges.ChainID

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
		return
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei 0 cost
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fileHash, fileSize, fileType, fileName, fileDescription := buildData()

	tx, err := contractIns.UploadFile(auth, fileHash, fileSize, fileType, fileName, fileDescription)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("transaction: ", tx.Cost())
	fmt.Println("transaction: ", tx.Hash())
}

func buildData() (string, *big.Int, string, string, string) {
	var fileHash string = "67ea261420Dea06E2f5F77C204128E20fc40f02aafe44"
	var fileSize *big.Int = big.NewInt(101)
	var fileType string = "mp3"
	var fileName string = "temasos"
	var fileDescription string = "cancion"

	return fileHash,
		fileSize,
		fileType,
		fileName,
		fileDescription
}
