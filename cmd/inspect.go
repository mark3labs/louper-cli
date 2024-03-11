/*
Copyright © 2024 Ed Zynda <ezynda3@gmail.com>
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

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, _ []string) {
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
	rootCmd.AddCommand(inspectCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// inspectCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// inspectCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	inspectCmd.Flags().StringVarP(&network, "network", "n", "mainnet", "The network the diamond contract is deployed to")
	inspectCmd.Flags().StringVarP(&address, "address", "a", "", "The address of the diamond contract to inspect")
	inspectCmd.MarkFlagRequired("address")
}
