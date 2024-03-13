package utils

import "github.com/ethereum/go-ethereum/accounts/abi/bind"

func GenerateCallDataFromAbi(meta *bind.MetaData, method string, args ...interface{}) ([]byte, error) {
	abi, err := meta.GetAbi()
	if err != nil {
		return nil, err
	}
	return abi.Pack(method, args...)
}
