package types

import "github.com/ethereum/go-ethereum/common"

const (
	Add = iota
	Replace
	Remove
)

type DiamondCut struct {
	FunctionSelectors [][4]byte
	FacetAddress      common.Address
	Action            uint8
}
