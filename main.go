package main

import (
	"fmt"
	"os"
	"time"

	shell "github.com/brianstrauch/cobra-shell"
	"github.com/constt/mcservertools/cmd"
	"github.com/hugolgst/rich-go/client"
)

func main() {
	startDRP()
	shellCmd := shell.New(cmd.RootCmd, nil)

	if err := shellCmd.Execute(); err != nil {
		fmt.Println("Error starting shell:", err)
		os.Exit(1)
	}
}

func startDRP() {
	err := client.Login("1379820182662414478")
	if err != nil {
		fmt.Println("Error initializing Discord Rich Presence:", err)
		return
	}

	now := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "Using mcservertools",
		Details:    "Testing Minecraft servers",
		LargeImage: "largeimageid",
		LargeText:  "MC Server Tools",
		SmallImage: "smallimageid",
		SmallText:  "CLI Tool",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	if err != nil {
		fmt.Println("Error setting Discord activity:", err)
	}
}
