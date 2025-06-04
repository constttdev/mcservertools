package cmd

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(uuidCmd)
	uuidCmd.Flags().StringVarP(&player, "player", "a", "", "Player name to uuid lookup (e.g. Notch)")
	uuidCmd.MarkFlagRequired("player")
}

var player string

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "See an uuid of an player",
	Long:  `See the premium and minimized uuid of an specified player`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()

		fmt.Println(green("Looking up UUIDs for:"), player)
		fmt.Println(" ")

		premiumUUID := getPremiumUUID(player)
		if premiumUUID == "" {
			fmt.Println(yellow("Premium UUID:"), red("Not found (likely cracked or offline mode)"))
		} else {
			fmt.Println(yellow("Premium UUID (formatted):"), formatUUID(premiumUUID))
			fmt.Println(yellow("Premium UUID (minimized):"), premiumUUID)
		}

		offlineUUID := generateOfflineUUID(player)
		fmt.Println(yellow("Offline UUID (formatted):"), offlineUUID.String())
		fmt.Println(yellow("Offline UUID (minimized):"), strings.ReplaceAll(offlineUUID.String(), "-", ""))
	},
}

func getPremiumUUID(name string) string {
	resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
	if err != nil || resp.StatusCode != 200 {
		return ""
	}
	defer resp.Body.Close()

	var data struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return ""
	}
	return data.ID
}

func generateOfflineUUID(name string) uuid.UUID {
	hash := md5.Sum([]byte("OfflinePlayer:" + name))
	return uuid.Must(uuid.FromBytes(hash[:]))
}

func formatUUID(id string) string {
	return fmt.Sprintf("%s-%s-%s-%s-%s", id[0:8], id[8:12], id[12:16], id[16:20], id[20:])
}
