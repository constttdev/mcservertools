package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cmds",
	Short: "List all available commands",
	Long:  `Using this command you can see all available commands`,
	Run: func(cmd *cobra.Command, args []string) {
		hiBlue := color.New(color.FgHiBlue).SprintFunc()

		fmt.Println(hiBlue("Available commands: \n\nrconconnect <address> <password>\nserverinfo <address>\nportscan <address>\nuuid <player>\nlocateip <ip>\nkick <player>\n\n"))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
