package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

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
