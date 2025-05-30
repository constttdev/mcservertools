package cmd

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"strconv"
)

var address string

var rootCmd = &cobra.Command{
	Use:   "portscann",
	Short: "See open ports of an server",
	Long:  `Using this command you can see how many ports and what ports are open on an server`,
	Run: func(cmd *cobra.Command, args []string) {
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()

		if address == "" {
			fmt.Fprintln(os.Stderr, red("[ERROR] --address flag is required.\n"))
			fmt.Fprintln(os.Stderr, "Use '--help' to see usage instructions.")
			os.Exit(1)
		}

		fmt.Printf("Scanning: %s\n\n", address)

		var ports = map[int]string{
			25565: "Minecraft Java (default)",
			19132: "Minecraft Bedrock (UDP)",
			19133: "Minecraft Bedrock IPv6 (UDP)",
			8123:  "Dynmap (Web Plugin)",
			25575: "RCON (Remote Console)",
			25566: "BungeeCord/Waterfall Secondary",
			24454: "GeyserMC (Bedrock Bridge)",
			64738: "Mumble Voice Integration",
			43594: "Custom Plugin API",
			8080:  "Admin Panel / Dynmap HTTP",
			8443:  "Admin Panel / Dynmap HTTPS",
			3000:  "Custom Dashboard / Monitoring",
			3306:  "MySQL (Plugin Database)",
			5432:  "PostgreSQL (Plugin Database)",
		}

		for port, desc := range ports {
			address := address + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout("tcp", address, 2*time.Second)
			if err != nil {
				fmt.Printf(red("%-7s (%s) | Is closed or filterd\n"), strconv.Itoa(port), desc)
				continue
			}
			conn.Close()
			fmt.Printf(green("%-7s (%s) | Is open\n"), strconv.Itoa(port), desc)
		}
	},
}

func Execute() {
	rootCmd.Flags().StringVarP(&address, "address", "a", "", "Server address to scan")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
