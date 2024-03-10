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
	"github.com/spf13/cobra"
)

type Diamond struct {
	Chain   string `json:"chain"`
	Diamond struct {
		Name string `json:"name"`
		Abi  []struct {
			Inputs []struct {
				InternalType string `json:"internalType"`
				Name         string `json:"name"`
				Type         string `json:"type"`
			} `json:"inputs,omitempty"`
			StateMutability string `json:"stateMutability,omitempty"`
			Type            string `json:"type"`
			Name            string `json:"name,omitempty"`
		} `json:"abi"`
		Address string `json:"address"`
		Facets  []struct {
			Name string `json:"name"`
			Abi  []struct {
				Inputs          []any  `json:"inputs"`
				Name            string `json:"name"`
				Type            string `json:"type"`
				Anonymous       bool   `json:"anonymous,omitempty"`
				Outputs         []any  `json:"outputs,omitempty"`
				StateMutability string `json:"stateMutability,omitempty"`
			} `json:"abi"`
			Address string `json:"address"`
		} `json:"facets"`
	} `json:"diamond"`
	DiamondAbi []struct {
		Inputs          []any  `json:"inputs,omitempty"`
		Name            string `json:"name,omitempty"`
		Type            string `json:"type"`
		Anonymous       bool   `json:"anonymous,omitempty"`
		Outputs         []any  `json:"outputs,omitempty"`
		StateMutability string `json:"stateMutability,omitempty"`
	} `json:"diamondAbi"`
}

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		network, err := cmd.Flags().GetString("network")
		if err != nil {
			fmt.Println("Error getting network flag")
		}
		address, err := cmd.Flags().GetString("address")
		if err != nil {
			fmt.Println("Error getting address flag")
		}

		if address == "" {
			fmt.Println("Error: --address flag is required")
			return
		}

		var resp *http.Response
		var jsonData Diamond
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

		fmt.Println("Diamond Name: ", jsonData.Diamond.Name)
		fmt.Println("Diamond Address: ", jsonData.Diamond.Address)

		for _, facet := range jsonData.Diamond.Facets {
			fmt.Println("-----------------")
			fmt.Println("Facet Name: ", facet.Name)
			fmt.Println("Facet Address: ", facet.Address)
		}
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
	inspectCmd.Flags().StringP("network", "n", "mainnet", "The network the diamond contract is deployed to")
	inspectCmd.Flags().StringP("address", "a", "", "The address of the diamond contract to inspect")
}
