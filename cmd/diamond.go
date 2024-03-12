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
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/charmbracelet/huh/spinner"
	"github.com/mark3labs/louper-cli/components"
	"github.com/mark3labs/louper-cli/types"
	"github.com/spf13/cobra"
)

// diamondCmd represents the diamond command
var diamondCmd = &cobra.Command{
	Use:   "diamond",
	Short: "Get detailed information about a dimaond contract",
	Run: func(_ *cobra.Command, _ []string) {
		var resp *http.Response
		var err error
		var jsonData types.Diamond
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

		fmt.Println(components.Table("Diamond Name", "Diamond Address").Row(jsonData.Diamond.Name, jsonData.Diamond.Address))

		facetsTable := components.Table("Facet Name", "Facet Address")
		for _, facet := range jsonData.Diamond.Facets {
			facetsTable.Row(facet.Name, facet.Address)
		}

		fmt.Println(facetsTable)
	},
}

func init() {
	inspectCmd.AddCommand(diamondCmd)
}
