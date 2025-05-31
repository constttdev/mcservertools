package cmd

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
	"time"

	mcpinger "github.com/Raqbit/mc-pinger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverInfoCmd)
	serverInfoCmd.Flags().StringVarP(&serverInfoAddress, "address", "a", "", "Server address to query (e.g. hypixel.net)")
	serverInfoCmd.Flags().IntVarP(&serverInfoPort, "port", "p", 25565, "Server port (default 25565)")

	serverInfoCmd.MarkFlagRequired("address")
}

var serverInfoAddress string
var serverInfoPort int

var serverInfoCmd = &cobra.Command{
	Use:   "serverinfo",
	Short: "See info about a Minecraft server",
	Long:  `Fetches MOTD, version, and online players of a Minecraft Java server.`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		yellow := color.New(color.FgYellow).SprintfFunc()

		host, port := resolveMinecraftSRV(serverInfoAddress, serverInfoPort)

		fmt.Println(green("Pinging:"), fmt.Sprintf("%s:%d \n", serverInfoAddress, serverInfoPort))

		pinger := mcpinger.New(host, uint16(port), mcpinger.WithTimeout(10*time.Second))

		info, err := pinger.Ping()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf(cyan("Description: \"%s\"\n"), cleanMOTD(info.Description.Text))
		fmt.Printf(green("Players: %d/%d\n"), info.Players.Online, info.Players.Max)
		fmt.Printf("%s %s\n", yellow("Version(s):"), info.Version.Name)
	},
}

func resolveMinecraftSRV(address string, fallbackPort int) (string, int) {
	_, addrs, err := net.LookupSRV("minecraft", "tcp", address)
	if err != nil || len(addrs) == 0 {
		return address, fallbackPort
	}
	return addrs[0].Target, int(addrs[0].Port)
}

func cleanMOTD(motd string) string {
	re := regexp.MustCompile(`§[0-9a-fk-or]`)
	clean := re.ReplaceAllString(motd, "")

	clean = strings.ReplaceAll(clean, "\n", " | ")
	clean = strings.TrimSpace(clean)
	return clean
}
