package main

import (
	"fmt"
	"os"
	"time"

	shell "github.com/brianstrauch/cobra-shell"
	"github.com/common-nighthawk/go-figure"
	"github.com/constt/mcservertools/cmd"
	"github.com/hugolgst/rich-go/client"
)

func main() {
	startDRP()

	figure := figure.NewColorFigure("MC Server Tools", "slant", "blue", true)
	figure.Print()
	fmt.Printf("\n\n\n\n\n\n\n")

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
		LargeImage: "logo_test_mod_scaled_16x_pngcrushed",
		LargeText:  "MC Server Tools",
		Timestamps: &client.Timestamps{
			Start: &now,
		},
	})

	if err != nil {
		fmt.Println("Error setting Discord activity:", err)
	}
}
