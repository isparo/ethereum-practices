package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/josue/ethereum-practices/example3_dropbox_v2/dropboxv2"
)

const contratAddress string = "0xd19E0a116DC4181aBc633919c00EAde8D180E3FB"
const privateAccountAddress string = "56b0c19ed9894b37b620a402963e12af5b34f62d8859510c85e1044be30b353f"

const senderAccount string = "fd5df95c7826401a983cb33e0904a96b200b3bbb6aa5481f4f252f3ffbf8f271"

func main() {
	/* firmar el mensaje
	esto se debe de hacer desde el frontend pero para fines de este ejercicio usaremos golang
	regularmente las firmas van a ser hechas usando metamask o alguna wallet a la que nos conectamos
	desde el front*/

	// establecemos la conexion a la block chain
	address := common.HexToAddress(contratAddress)

	client, err := ethclient.Dial("http://127.0.0.1:7545") // URL de Ganache
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Conectado a la red Ethereum")

	instance, err := dropboxv2.NewDropboxv2(address, client)
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

	fmt.Println("_________________________________________________")
	fmt.Println("_________________________________________________")
	fmt.Println("_________________________________________________")

	GetFilesForAddress(instance)

}

func signMessage() (*big.Int, *big.Int, uint8, error) {
	privateKey, err := crypto.HexToECDSA(senderAccount)
	if err != nil {
		log.Fatal(err)
	}
	fileHash, _, _, _, _ := buildData()

	// Prefix the message hash according to the Ethereum standard
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n32%s", fileHash)
	ethSignedMessageHash := crypto.Keccak256Hash([]byte(message))

	signature, err := crypto.Sign(ethSignedMessageHash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Signature:", hexutil.Encode(signature))

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])
	v := uint8(signature[64]) // + 27 // ajustar v a 27 o 28

	if v < 27 {
		v += 27
	}

	// fmt.Println("signature r: ", r.Text(16))
	// fmt.Println("signature s: ", s.Text(16))
	// fmt.Println("signature v: ", v)
	// //fmt.Println("signature: ", messageHash.Hex())

	return r, s, v, nil
}

func buildData() (string, *big.Int, string, string, string) {
	var fileHash string = "67ea261420Dea06E2f5F77C204128E20fc40f02aafe14"
	var fileSize *big.Int = big.NewInt(1024)
	var fileType string = "text/plain"
	var fileName string = "documenttest1.txt"
	var fileDescription string = "este es un ejemplo de prueba"

	return fileHash,
		fileSize,
		fileType,
		fileName,
		fileDescription
}

func GetAllfiles(contractIns *dropboxv2.Dropboxv2) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	result, err := contractIns.ListFiles(callOpts)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, v := range result {
		fmt.Println(v)
	}
}

func upladFile(contractIns *dropboxv2.Dropboxv2, client *ethclient.Client) {
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
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	fileHash, fileSize, fileType, fileName, fileDescription := buildData()

	rr, ss, v, err := signMessage()
	if err != nil {
		log.Fatal(err)
		return
	}

	rBytes := rr.Bytes()
	sBytes := ss.Bytes()

	var r32 [32]byte
	copy(r32[:], rBytes)

	var s32 [32]byte
	copy(s32[:], sBytes)

	fmt.Printf("Valor en bytes32: %x\n", r32)
	fmt.Printf("Valor en bytes32: %x\n", s32)

	tx, err := contractIns.UploadFile(auth, fileHash, fileSize, fileType, fileName, fileDescription, v, r32, s32)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("transaction: ", tx.Cost())
	fmt.Println("transaction: ", tx.Hash())
}

func GetFilesForAddress(contractIns *dropboxv2.Dropboxv2) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	result, err := contractIns.ListFiles(callOpts)
	if err != nil {
		log.Fatal(err)
		return
	}

	privateKey, err := crypto.HexToECDSA(senderAccount)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	for _, v := range result {
		if v.Uploader.String() == fromAddress.String() {
			fmt.Println(v)
		}
	}
}
