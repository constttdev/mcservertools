package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	mcpinger "github.com/Raqbit/mc-pinger"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type BlacklistResponse struct {
	Blacklisted bool   `json:"blacklisted"`
	LastUpdate  string `json:"lastUpdate"`
}

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
		red := color.New(color.FgRed).SprintFunc()

		blacklisted, err := isServerBlacklisted(serverInfoAddress)

		if err != nil {
			fmt.Printf("Error checking blacklist status: %v\n", err)
			return
		}

		host, port := resolveMinecraftSRV(serverInfoAddress, serverInfoPort)

		fmt.Println(green("Pinging:"), fmt.Sprintf("%s:%d \n", serverInfoAddress, serverInfoPort))

		pinger := mcpinger.New(host, uint16(port), mcpinger.WithTimeout(10*time.Second))

		start := time.Now()
		info, pErr := pinger.Ping()
		latency := time.Since(start)

		ips, ipErr := net.LookupHost(host)

		if pErr != nil {
			log.Println(pErr)
			return
		}

		fmt.Printf(cyan("Description: \"%s\"\n"), cleanMOTD(info.Description.Text))
		fmt.Printf(green("Players: %d/%d\n"), info.Players.Online, info.Players.Max)
		fmt.Printf("%s %s\n", yellow("Version(s):"), info.Version.Name)
		fmt.Printf(green("Ping: %d ms\n"), latency.Milliseconds())
		if blacklisted {
			fmt.Printf(red("Blocked by mojang: true ('%s')\n"), serverInfoAddress)
		} else {
			fmt.Printf(green("Blocked by mojang: false ('%s')\n"), serverInfoAddress)
		}

		if ipErr == nil && len(ips) > 0 {
			fmt.Println(cyan("IP: "), ips[0])
		} else {
			fmt.Println("IP: (Could not resolve)")
		}
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

func isServerBlacklisted(serverAddress string) (bool, error) {
	url := fmt.Sprintf("https://eu.mc-api.net/v3/server/blacklisted/%s", serverAddress)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return false, fmt.Errorf("failed to fetch blacklist status: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	var result BlacklistResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to parse response: %v", err)
	}

	return result.Blacklisted, nil
}
