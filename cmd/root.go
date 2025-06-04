package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mcst",
	Short: "A CLI tool for testing Minecraft servers",
	Long:  `This tool allows you to perform various tests on Minecraft servers.`,
}
