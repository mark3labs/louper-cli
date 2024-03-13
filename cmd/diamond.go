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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/huh/spinner"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/mark3labs/louper-cli/components"
	"github.com/mark3labs/louper-cli/types"
	"github.com/spf13/cobra"
)

// diamondCmd represents the diamond command
var diamondCmd = &cobra.Command{
	Use:   "diamond",
	Short: "Get detailed information about a dimaond contract",
	Run: func(_ *cobra.Command, _ []string) {
		var err error
		var diamond types.Diamond
		var diamondTable, facetsTable *table.Table
		var diamondJson types.DiamondJson
		ctx, cancel := context.WithCancel(context.Background())

		go func() {
			defer cancel()
			diamond, err = getDiamond()
			if jsonFormat {
				diamondJson = buildDiamondJson(diamond)
			} else {
				diamondTable, facetsTable = buildDiamondTable(diamond)
			}
		}()

		spinner.New().Title("Fetching Diamond Details...").Context(ctx).Run()

		if err != nil {
			fmt.Println(components.ErrorBox(err.Error()))
			return
		}

		if jsonFormat {
			fmt.Println(diamondJson)
		} else {
			fmt.Println(diamondTable)
			fmt.Println(facetsTable)
		}
	},
}

func init() {
	inspectCmd.AddCommand(diamondCmd)
}

func getDiamond() (types.Diamond, error) {
	var resp *http.Response
	var err error
	var jsonData types.DiamondResponse
	spinner.New().Title("Fetching Diamond Details...").Action(func() {
		// Fetch JSON from https://louper.dev/diamond/{address}/json?network={network}
		resp, err = http.Get("https://louper.dev/diamond/" + address + "/json?network=" + network)
		if err != nil {
			fmt.Println("Error fetching JSON from Louper API")
		}
		defer resp.Body.Close()
		rawBody, _ := io.ReadAll(resp.Body)
		// Marshal JSON into struct
		json.Unmarshal(rawBody, &jsonData)
	}).Run()
	return jsonData.Diamond, nil
}

func buildDiamondTable(diamond types.Diamond) (*table.Table, *table.Table) {
	diamondTable := components.Table("Diamond Name", "Diamond Address", "Network").Row(diamond.Name, diamond.Address, network)
	facetsTable := components.Table("Facet Name", "Facet Address")
	for _, facet := range diamond.Facets {
		facetsTable.Row(facet.Name, facet.Address)
	}
	return diamondTable, facetsTable
}

func buildDiamondJson(diamond types.Diamond) types.DiamondJson {
	diamondJson := types.DiamondJson{
		Name:    diamond.Name,
		Address: diamond.Address,
	}
	for _, facet := range diamond.Facets {
		f := struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		}{
			Name:    facet.Name,
			Address: facet.Address,
		}
		diamondJson.Facets = append(diamondJson.Facets, f)
	}

	return diamondJson
}
