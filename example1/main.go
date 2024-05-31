package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/josue/ethereum-practices/example1/simplestorage"
)

const contractAddress string = "0x5AF3aC19fE271E5E5B104Fe3094a9244aEE85781"
const accountPrivateKey string = "779d00cb4433fc1861af90505a6914a4a251b86160ea78e40b4986c0750f2465"

func main() {

	address := common.HexToAddress(contractAddress)

	client, err := ethclient.Dial("http://127.0.0.1:7545") // URL de Ganache
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado a la red Ethereum")

	fmt.Printf("Dirección del contrato: %s\n", address)

	instance, err := simplestorage.NewSimplestorage(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	for i := 0; i < 100; i++ {
		readFromContract(instance)
		writeToContract(instance, client)
		readFromContract(instance)

		time.Sleep(100 * time.Second)
	}

}

func readFromContract(contractIns *simplestorage.Simplestorage) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	res, err := contractIns.Get(callOpts)

	fmt.Println(res)
	fmt.Println(err)
}

func writeToContract(contractIns *simplestorage.Simplestorage, client *ethclient.Client) {

	privateKey, err := crypto.HexToECDSA(accountPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	fmt.Println("public key based on private key: ", publicKey)
	fmt.Println("public key adrees: ", fromAddress)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("gas price: ", gasPrice)

	chainID := params.AllDevChainProtocolChanges.ChainID
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei 0 cost
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	value := big.NewInt(101)
	tx, err := contractIns.Set(auth, value)

	fmt.Println(err)
	fmt.Println(tx)

}
