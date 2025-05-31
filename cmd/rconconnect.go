package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/gorcon/rcon"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rconConnectCmd)
	rconConnectCmd.Flags().StringVarP(&rconConnectAddress, "address", "a", "", "RCON address (e.g. 127.0.0.1)")
	rconConnectCmd.Flags().IntVarP(&rconConnectPort, "port", "p", 25575, "RCON port (default 25575)")
	rconConnectCmd.Flags().StringVarP(&rconPassword, "password", "", "", "RCON password (required)")
	rconConnectCmd.MarkFlagRequired("password")
	rconConnectCmd.MarkFlagRequired("address")
}

var rconConnectAddress string
var rconConnectPort int
var rconPassword string

var rconConnectCmd = &cobra.Command{
	Use:   "rconconnect",
	Short: "Connect to a Minecraft server via RCON",
	Long:  `Connect to a Minecraft server's RCON and execute console commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()

		addr := fmt.Sprintf("%s:%d", rconConnectAddress, rconConnectPort)
		client, err := rcon.Dial(addr, rconPassword)
		if err != nil {
			log.Fatalf("Failed to connect to RCON: %v", err)
		}
		defer client.Close()

		fmt.Println(green("Connected to RCON. Type commands or 'exit' to quit."))

		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")
			if !scanner.Scan() {
				break
			}
			text := scanner.Text()
			if text == "exit" {
				fmt.Println("Exiting RCON session.")
				break
			}
			resp, err := client.Execute(text)
			if err != nil {
				log.Printf("Command error: %v", err)
				continue
			}
			fmt.Println(resp)
		}
	},
}
