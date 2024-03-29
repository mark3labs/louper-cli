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

var justSelectors bool

// selectorsCmd represents the selectors command
var facetCmd = &cobra.Command{
	Use:   "facet",
	Short: "Get detailed information about a facet contract",
	Run: func(_ *cobra.Command, _ []string) {
		var facetTable, selectorsTable *table.Table
		var err error
		var facet types.AbiResponse
		var facetJson types.FacetJson

		ctx, cancel := context.WithCancel(context.Background())
		go func() {
			defer cancel()
			facet, err = getFacet()
			if err != nil {
				return
			}
			if jsonFormat || justSelectors {
				facetJson = buildFacetJson(facet)
			} else {
				facetTable, selectorsTable = buildFacetTable(facet)
			}
		}()

		spinner.New().Title("Fetching Facet Details...").Context(ctx).Run()

		if err != nil {
			fmt.Println(components.ErrorBox(err.Error()))
			return
		}

		if justSelectors {
			var selectors []string
			for _, me := range facetJson.Functions {
				selectors = append(selectors, "0x"+me.Selector)
			}
			fmt.Println(strings.Join(selectors, ","))
			return
		}

		if jsonFormat {
			fmt.Println(facetJson)
		} else {
			fmt.Println(facetTable)
			fmt.Println(selectorsTable)
		}
	},
}

func init() {
	inspectCmd.AddCommand(facetCmd)
	facetCmd.Flags().BoolVar(&justSelectors, "selectors", false, "Only display a comma separated list of selectors")
}

func getFacet() (types.AbiResponse, error) {
	chainID := constants.ChainNamesToID[network]
	resp, err := http.Get(fmt.Sprintf("https://anyabi.xyz/api/get-abi/%d/%s", chainID, address))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var abiResp types.AbiResponse
	err = json.NewDecoder(resp.Body).Decode(&abiResp)
	if err != nil {
		return types.AbiResponse{}, fmt.Errorf("failed to fetch facet")
	}

	if resp.StatusCode != 200 {
		return types.AbiResponse{}, fmt.Errorf("facet not found")
	}

	return abiResp, nil
}

func buildFacetTable(facet types.AbiResponse) (*table.Table, *table.Table) {
	abiString, err := json.Marshal(facet.Abi)
	if err != nil {
		log.Fatal(err)
	}
	ABI, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatal(err)
	}

	facetTable := components.Table("Facet Name", "Facet Address", "Network").Row(facet.Name, address, network)
	selectorsTable := components.Table("Method Signature", "Selector")
	for _, me := range ABI.Methods {
		selectorsTable.Row(me.Sig, "0x"+hex.EncodeToString(me.ID))
	}
	if len(ABI.Methods) == 0 {
		selectorsTable.Row("No methods found", "No selectors found")
	}

	return facetTable, selectorsTable
}

func buildFacetJson(facet types.AbiResponse) types.FacetJson {
	abiString, err := json.Marshal(facet.Abi)
	if err != nil {
		log.Fatal(err)
	}
	ABI, err := abi.JSON(strings.NewReader(string(abiString)))
	if err != nil {
		log.Fatal(err)
	}

	var facetJson types.FacetJson
	facetJson.Name = facet.Name
	facetJson.Address = address
	for _, me := range ABI.Methods {
		var function types.FunctionJson
		function.Name = me.Name
		function.Signature = me.Sig
		function.Selector = hex.EncodeToString(me.ID)
		facetJson.Functions = append(facetJson.Functions, function)
	}
	return facetJson
}
