package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"golang.org/x/crypto/sha3"
)

const (
	EthereumTransactionDefaultGasLimit uint64 = 21000
	EthereumDecimals                          = 18
)

func main() {

}

type Ethereum struct {
	rpcclient *rpc.Client
	client    *ethclient.Client
	from      struct {
		address    common.Address
		privateKey ecdsa.PrivateKey
	}
}

func (e *Ethereum) SetFrom(privateKeyString string) error {
	if privateKeyString[:2] == "0x" {
		privateKeyString = privateKeyString[2:]
	}

	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return err
	}

	e.from.privateKey = *privateKey

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	e.from.address = crypto.PubkeyToAddress(*publicKeyECDSA)

	return nil
}

func (e *Ethereum) SendEtherTransaction(toAddress string, amount float64, gasPriceMultiplier int64) (string, error) {
	valueBigInt, ok := math.ParseBig256(strconv.FormatFloat(EthereumDecimals*10*amount, 'f', 0, 64))
	if !ok {
		log.Fatal("parse decimal scale failed")
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", err
	}

	gasPrice.Mul(gasPrice, big.NewInt(gasPriceMultiplier))

	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	address := common.HexToAddress(toAddress)
	var data []byte
	fee, _ := math.ParseBig256(strconv.FormatUint(EthereumTransactionDefaultGasLimit, 10))

	nonce, err := e.client.PendingNonceAt(context.Background(), e.from.address)
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, address, valueBigInt.Sub(valueBigInt, fee.Mul(fee, gasPrice)), EthereumTransactionDefaultGasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), &e.from.privateKey)
	if err != nil {
		return "", err
	}

	err = e.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}

func (e *Ethereum) SendTokenTransaction(contractAddress, toAddress string, amount float64, gasPriceMultiplier int64) (string, error) {
	nonce, err := e.client.PendingNonceAt(context.Background(), e.from.address)
	if err != nil {
		log.Fatalf("%v", err)
	}

	gasPrice, err := e.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("%v", err)
	}
	gasPrice = big.Int{}.Mul(gasPrice, big.NewInt(gasPriceMultiplier))

	address := common.HexToAddress(toAddress)

	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	paddedAddress := common.LeftPadBytes(address.Bytes(), 32)

	decimals := float64(6) //TODO get contract decimals
	value, ok := math.ParseBig256(strconv.FormatFloat(decimals*10*amount, 'f', 0, 64))
	if !ok {
		return "", errors.New("parse decimal scale failed")
	}

	paddedAmount := common.LeftPadBytes(value.Bytes(), 32)

	var data []byte
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	contractAddressHex := common.HexToAddress(contractAddress)
	gasLimit, err := e.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:     e.from.address,
		To:       &contractAddressHex,
		GasPrice: gasPrice,
		Value:    big.NewInt(0),
		Data:     data,
	})
	if err != nil {
		return "", err
	}

	tx := types.NewTransaction(nonce, contractAddressHex, value, gasLimit, gasPrice, data)

	chainID, err := e.client.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), &e.from.privateKey)
	if err != nil {
		return "", err
	}

	if err = e.client.SendTransaction(context.Background(), signedTx); err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
