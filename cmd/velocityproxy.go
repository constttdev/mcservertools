package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(velocityProxyCmd)
	velocityProxyCmd.Flags().StringVarP(&velocityProxyAddress, "address", "a", "", "The address to bind the proxy to")
	velocityProxyCmd.MarkFlagRequired("address")
}

var velocityProxyAddress string

var velocityProxyCmd = &cobra.Command{
	Use:   "velocity",
	Short: "Create a fake velocity proxy",
	Long:  `Create a fake velocity proxy to test exploits`,
	Run: func(cmd *cobra.Command, args []string) {
		italic := color.New(color.Italic).SprintFunc()
		dir := filepath.Join(".", "velocityProxy")
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println(italic("[DEBUG] Created folder"))

		type versionResp struct{ Versions []string }
		type buildResp struct{ Builds []int }
		type downloadResp struct {
			Downloads struct {
				App struct{ Name string } `json:"application"`
			} `json:"downloads"`
		}

		getJSON := func(url string, target any) error {
			r, err := http.Get(url)
			if err != nil {
				return err
			}
			defer r.Body.Close()
			return json.NewDecoder(r.Body).Decode(target)
		}

		var vr versionResp
		if getJSON("https://api.papermc.io/v2/projects/velocity", &vr) != nil || len(vr.Versions) == 0 {
			fmt.Println("Failed to fetch version")
			return
		}
		version := vr.Versions[len(vr.Versions)-1]

		var br buildResp
		if getJSON(fmt.Sprintf("https://api.papermc.io/v2/projects/velocity/versions/%s", version), &br) != nil || len(br.Builds) == 0 {
			fmt.Println("Failed to fetch build")
			return
		}
		build := br.Builds[len(br.Builds)-1]

		var dr downloadResp
		if getJSON(fmt.Sprintf("https://api.papermc.io/v2/projects/velocity/versions/%s/builds/%d", version, build), &dr) != nil {
			fmt.Println("Failed to fetch download info")
			return
		}
		jar := dr.Downloads.App.Name
		downloadURL := fmt.Sprintf("https://api.papermc.io/v2/projects/velocity/versions/%s/builds/%d/downloads/%s", version, build, jar)

		jarPath := filepath.Join(dir, jar)
		resp, err := http.Get(downloadURL)
		if err != nil {
			fmt.Println("Download error:", err)
			return
		}
		defer resp.Body.Close()

		out, err := os.Create(jarPath)
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return
		}

		if _, err := io.Copy(out, resp.Body); err != nil {
			fmt.Println("Error saving file:", err)
			out.Close()
			return
		}

		if err := out.Close(); err != nil {
			fmt.Println("Error closing file:", err)
			return
		}

		newJarPath := filepath.Join(dir, "server.jar")
		if err := os.Rename(jarPath, newJarPath); err != nil {
			fmt.Println("Error renaming file:", err)
			return
		}
		fmt.Println(italic("[DEBUG] Renamed to server.jar"))

		jarExecute := exec.Command("java", "-jar", newJarPath)

		if err := jarExecute.Start(); err != nil {
			fmt.Println("Error starting command:", err)
			return
		} else {
			fmt.Println(italic("[DEBUG] Started server jar"))
		}

		time.Sleep(5 * time.Second)

		jarExecute.Process.Kill()
		fmt.Println(italic("[DEBUG] Killed server jar"))
	},
}
