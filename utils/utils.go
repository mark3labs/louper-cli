package utils

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

func GenerateCallDataFromAbi(meta *bind.MetaData, method string, args ...interface{}) ([]byte, error) {
	abi, err := meta.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Pack(method, args...)
}

func GetPrivateKey(private string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(private)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func EstimateGas(client ethclient.Client, from, to common.Address, data []byte) (uint64, error) {
	callMsg := ethereum.CallMsg{
		From:  from,
		To:    &to,
		Data:  data,
		Value: big.NewInt(0),
	}

	gas, err := client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return 0, err
	}
	return gas, nil
}

func IsContract(client *ethclient.Client, address common.Address) bool {
	code, err := client.CodeAt(context.Background(), address, nil)
	if err != nil {
		return false
	}
	return len(code) > 0
}
