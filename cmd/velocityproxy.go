package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(velocityProxyCmd)
	velocityProxyCmd.Flags().StringVarP(&velocityProxyAddress, "address", "a", "", "The address to bind the proxy to")
	velocityProxyCmd.MarkFlagRequired("address")
}

var velocityProxyAddress string

var velocityProxyCmd = &cobra.Command{
	Use:   "velocity",
	Short: "Create a fake velocity proxy",
	Long:  `Create a fake velocity proxy to test exploits`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Command is currently WIP!")
	},
}
