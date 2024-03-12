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
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/mark3labs/louper-cli/components"
	"github.com/mark3labs/louper-cli/constants"
	"github.com/mark3labs/louper-cli/types"
	"github.com/spf13/cobra"
)

// selectorsCmd represents the selectors command
var facetCmd = &cobra.Command{
	Use:   "facet",
	Short: "Get detailed information about a facet contract",
	Run: func(_ *cobra.Command, _ []string) {
		facetTable := components.Table("Facet Name", "Facet Address", "Network")
		selectorsTable := components.Table("Function Signature", "Selector")
		var status int

		ctx, cancel := context.WithCancel(context.Background())
		go getSelectors(cancel, facetTable, selectorsTable, &status)

		spinner.New().Title("Fetching Facet Details...").Context(ctx).Run()

		if status != 200 {
			fmt.Println(components.ErrorBox("Facet not found"))
			return
		}

		fmt.Println(facetTable)
		fmt.Println(selectorsTable)
	},
}

func init() {
	inspectCmd.AddCommand(facetCmd)
}

func getSelectors(cancel context.CancelFunc, facetTable, selectorsTable *table.Table, status *int) {
	defer cancel()

	// Fetch the ABI from https://anyabi.xyz/api/get-abi/<chain_id>/<contract_address>
	chainID := constants.ChainNamesToID[network]
	resp, err := http.Get(fmt.Sprintf("https://anyabi.xyz/api/get-abi/%d/%s", chainID, address))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	*status = resp.StatusCode

	var abiResp types.AbiResponse
	err = json.NewDecoder(resp.Body).Decode(&abiResp)
	if err != nil {
		log.Fatal(err)
	}

	abiString, err := json.Marshal(abiResp.Abi)
	if err != nil {
		log.Fatal(err)
	}
	ABI, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatal(err)
	}

	facetTable.Row(abiResp.Name, address, network)
	for _, me := range ABI.Methods {
		selectorsTable.Row(me.Sig, hex.EncodeToString(me.ID))
	}
}
