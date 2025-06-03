package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mcst",
	Short: "A CLI tool for testing Minecraft servers",
	Long:  `This tool allows you to perform various tests on Minecraft servers.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to mcst! Use --help for usage.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
