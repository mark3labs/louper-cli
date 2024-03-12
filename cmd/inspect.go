/*
Copyright © 2024 Ed Zynda <ezynda3@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Get detailed information about a diamond or facet contract",
}

func init() {
	rootCmd.AddCommand(inspectCmd)

	inspectCmd.PersistentFlags().StringVarP(&network, "network", "n", "mainnet", "The network the contract is deployed to")
	inspectCmd.PersistentFlags().StringVarP(&address, "address", "a", "", "The address of the diamond contract to inspect")
	inspectCmd.MarkPersistentFlagRequired("address")
}
