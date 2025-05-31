package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(kickCmd)
}

var kickCmd = &cobra.Command{
	Use:   "kick",
	Short: "Kicks a user from a server",
	Long:  `Kicks a specified user from a server by joining with a cracked account.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Command is still WIP")
	},
}
