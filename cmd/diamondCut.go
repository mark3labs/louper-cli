/*
Copyright © 2024 Ed Zynda <ezynda3@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/mark3labs/louper-cli/bindings/diamondCut"
	"github.com/mark3labs/louper-cli/types"
	"github.com/mark3labs/louper-cli/utils"
	"github.com/spf13/cobra"
)

var (
	dumpCallData bool
	privateKey   string
	rawCuts      []string
)

// diamondCutCmd represents the diamondCut command
var diamondCutCmd = &cobra.Command{
	Use:   "diamond-cut",
	Short: "Add, Remove or Replace facets in a diamond contract",
	Run: func(_ *cobra.Command, _ []string) {
		cuts := []types.DiamondCut{}
		for _, rawCut := range rawCuts {
			cut, err := parseDiamondCut(rawCut)
			if err != nil {
				fmt.Println(err)
				return
			}
			cuts = append(cuts, cut)
		}

		if dumpCallData {
			displayCalldata(cuts)
			return
		} else {
			fmt.Println("Not implemented yet!")
		}
	},
}

func init() {
	rootCmd.AddCommand(diamondCutCmd)

	diamondCutCmd.Flags().StringVarP(&network, "network", "n", "mainnet", "The network the diamond contract is deployed to")
	diamondCutCmd.Flags().StringVarP(&address, "address", "a", "", "The address of the diamond contract to cut")
	diamondCutCmd.Flags().StringVarP(&privateKey, "private-key", "p", "", "The private key to use to sign the transaction")
	diamondCutCmd.Flags().BoolVar(&dumpCallData, "calldata", false, "Dump the call data for the diamond cut")
	diamondCutCmd.Flags().StringArrayVar(&rawCuts, "cut", []string{}, "A list of raw diamond cut data to apply to the contract")
	diamondCutCmd.MarkFlagsOneRequired("private-key", "calldata")
	diamondCutCmd.MarkFlagsMutuallyExclusive("private-key", "calldata")
	diamondCutCmd.MarkFlagRequired("address")
}

func parseDiamondCut(rawCut string) (types.DiamondCut, error) {
	parts := strings.Split(rawCut, "|")
	if len(parts) != 3 {
		return types.DiamondCut{}, fmt.Errorf("invalid raw cut data: %s", rawCut)
	}

	action := strings.ToLower(parts[0])
	var facetCutAction uint8
	switch action {
	case "add":
		facetCutAction = types.Add
	case "replace":
		facetCutAction = types.Replace
	case "remove":
		facetCutAction = types.Remove
	}

	address := common.HexToAddress(parts[1])

	if parts[2] == "" {
		return types.DiamondCut{}, fmt.Errorf("no function selectors provided: %s", rawCut)
	}

	selectors := strings.Split(parts[2], ",")
	var functionSelectors [][4]byte
	for _, selector := range selectors {
		selector = strings.TrimPrefix(selector, "0x")
		selectorBytes, err := hex.DecodeString(selector)
		if err != nil {
			return types.DiamondCut{}, fmt.Errorf("invalid function selector: 0x%s", selector)
		}
		var functionSelector [4]byte
		copy(functionSelector[:], selectorBytes[:4])
		functionSelectors = append(functionSelectors, functionSelector)
	}

	return types.DiamondCut{
		FunctionSelectors: functionSelectors,
		FacetAddress:      address,
		Action:            facetCutAction,
	}, nil
}

func displayCalldata(cut []types.DiamondCut) {
	calldata, err := utils.GenerateCallDataFromAbi(diamondCut.DiamondCutMetaData, "diamondCut", cut, common.Address{}, []byte{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("0x%s", hex.EncodeToString(calldata))
}
