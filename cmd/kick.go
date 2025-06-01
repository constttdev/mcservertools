package cmd

import (
	"log"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/spf13/cobra"
)

var (
	kickAddress string
	kickPort    int
	kickPlayer  string
)

func init() {
	rootCmd.AddCommand(kickCmd)
	kickCmd.Flags().StringVarP(&kickAddress, "address", "a", "", "Server address to join (e.g., 127.0.0.1)")
	kickCmd.Flags().IntVarP(&kickPort, "port", "P", 25565, "Server port (default 25565)")
	kickCmd.Flags().StringVarP(&kickPlayer, "player", "p", "", "Player name to use for login (e.g., Notch)")
	kickCmd.MarkFlagRequired("address")
	kickCmd.MarkFlagRequired("player")
}

var kickCmd = &cobra.Command{
	Use:   "kick",
	Short: "Kicks a user from a server",
	Long:  `Kicks a specified user from a server by joining with a cracked account.`,
	Run: func(cmd *cobra.Command, args []string) {
		c := bot.NewClient()
		c.Auth.Name = kickPlayer

		if err := c.JoinServer("play.blockfun.gg:25565"); err != nil {
			log.Fatalf("Failed to join server: %v", err)
		}

		time.Sleep(2 * time.Second)

		c.Close()
	},
}
