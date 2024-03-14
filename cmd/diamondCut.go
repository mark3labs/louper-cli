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
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/mark3labs/louper-cli/bindings/diamondCut"
	"github.com/mark3labs/louper-cli/components"
	"github.com/mark3labs/louper-cli/constants"
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
		// Parse all the cuts
		cuts := []diamondCut.IDiamondFacetCut{}
		for _, rawCut := range rawCuts {
			cut, err := parseDiamondCut(rawCut)
			if err != nil {
				fmt.Println(err)
				return
			}
			cuts = append(cuts, cut)
		}

		if dumpCallData {
			// Display Calldata if set
			displayCalldata(cuts)
			return
		} else {
			// Otherwise execute the TX
			executeCut(cuts)
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

// parseDiamondCut
// @param rawCut - The raw cut data to parse
func parseDiamondCut(rawCut string) (diamondCut.IDiamondFacetCut, error) {
	// Split the raw cut data into parts
	parts := strings.Split(rawCut, "|")
	if len(parts) != 3 {
		return diamondCut.IDiamondFacetCut{}, fmt.Errorf("invalid raw cut data: %s", rawCut)
	}

	// Parse the action
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

	// Parse the address
	address := common.HexToAddress(parts[1])

	// Parse the function selectors
	if parts[2] == "" {
		return diamondCut.IDiamondFacetCut{}, fmt.Errorf("no function selectors provided: %s", rawCut)
	}

	selectors := strings.Split(parts[2], ",")
	var functionSelectors [][4]byte
	for _, selector := range selectors {
		selector = strings.TrimPrefix(selector, "0x")
		selectorBytes, err := hex.DecodeString(selector)
		if err != nil {
			return diamondCut.IDiamondFacetCut{}, fmt.Errorf("invalid function selector: 0x%s", selector)
		}
		var functionSelector [4]byte
		copy(functionSelector[:], selectorBytes[:4])
		functionSelectors = append(functionSelectors, functionSelector)
	}

	// Return the parsed cut struct
	return diamondCut.IDiamondFacetCut{
		FunctionSelectors: functionSelectors,
		FacetAddress:      address,
		Action:            facetCutAction,
	}, nil
}

// displayCalldata
// @param cuts - The cuts to display the calldata for
func displayCalldata(cuts []diamondCut.IDiamondFacetCut) {
	calldata, err := utils.GenerateCallDataFromAbi(diamondCut.DiamondCutMetaData, "diamondCut", cuts, common.Address{}, []byte{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("0x%s", hex.EncodeToString(calldata))
}

// executeCut
// @param cuts - The cuts to execute
func executeCut(cuts []diamondCut.IDiamondFacetCut) {
	// Connect to the network
	client, err := ethclient.Dial(constants.RPCUrls[network])
	if err != nil {
		log.Fatal(err)
	}

	// Get the private key
	pKey, err := utils.GetPrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Create the transactor
	auth, err := bind.NewKeyedTransactorWithChainID(pKey, big.NewInt(int64(constants.ChainNamesToID[network])))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the nonce
	fromAddress := crypto.PubkeyToAddress(pKey.PublicKey)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
		return
	}

	calldata, err := utils.GenerateCallDataFromAbi(diamondCut.DiamondCutMetaData, "diamondCut", cuts, common.Address{}, []byte{})
	if err != nil {
		fmt.Println(components.ErrorBox(err.Error()))
		return
	}

	gasLimit, err := utils.EstimateGas(*client, fromAddress, common.HexToAddress(address), calldata)
	if err != nil {
		fmt.Println(components.ErrorBox(fmt.Sprintf("Failed to estimate gas: %s", err)))
		return
	}

	// Set the transactor values
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	// Create the diamond cut contract
	diamond, err := diamondCut.NewDiamondCut(common.HexToAddress(address), client)
	if err != nil {
		log.Fatal(err)
	}

	detailsTable := components.Table("Action", "Facet", "Selectors")
	for _, cut := range cuts {
		var actionStr string
		switch cut.Action {
		case types.Add:
			actionStr = "Add"
		case types.Remove:
			actionStr = "Remove"
		case types.Replace:
			actionStr = "Replace"
		}

		selectors := []string{}
		for _, selector := range cut.FunctionSelectors {
			selectors = append(selectors, fmt.Sprintf("0x%x", selector))
		}
		detailsTable.Row(actionStr, cut.FacetAddress.Hex(), strings.Join(selectors, ","))
	}

	txTable := components.Table("", "Transaction Details")
	txTable.Row("From", fromAddress.Hex())
	txTable.Row("To", address)
	txTable.Row("Gas Price", auth.GasPrice.String())
	txTable.Row("Gas Limit", fmt.Sprint(auth.GasLimit))
	txTable.Row("Nonce", auth.Nonce.String())
	txTable.Row("Value", auth.Value.String())

	fmt.Println(components.Box("Diamond Cut Details"))
	fmt.Println(txTable)
	fmt.Println(detailsTable)

	var execute bool
	huh.NewForm(
		huh.NewGroup(huh.NewConfirm().
			Title("Are you sure you want to execute this transaction?").
			Affirmative("Yes").
			Negative("No").
			Value(&execute),
		),
	).Run()
	if !execute {
		return
	}

	// Execute the cut
	tx, err := diamond.DiamondCut(auth, cuts, common.Address{}, []byte{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
}
