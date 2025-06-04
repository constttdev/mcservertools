package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

type GeoInfo struct {
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	City        string  `json:"city"`
	ISP         string  `json:"isp"`
	Org         string  `json:"org"`
	AS          string  `json:"as"`
	Query       string  `json:"query"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

func init() {
	RootCmd.AddCommand(geolocationCmd)
	geolocationCmd.Flags().StringVarP(&ip, "ip", "", "", "IP address to locate")
	geolocationCmd.MarkFlagRequired("ip")
}

var ip string

var geolocationCmd = &cobra.Command{
	Use:   "locateip",
	Short: "See the approximate location of an ip",
	Long:  `See the approximate location of an ip via gelocation`,
	Run: func(cmd *cobra.Command, args []string) {
		green := color.New(color.FgGreen).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()

		fmt.Printf(green("Getting the approximate location of: %s\n\n"), ip)

		resp, err := http.Get("http://ip-api.com/json/" + ip)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		var geo GeoInfo
		if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
			panic(err)
		}

		fmt.Printf(yellow("IP: %s\nCountry: %s\nRegion: %s\nCity: %s\nISP: %s\nOrg: %s\nLat/Lon: %.2f, %.2f\n"),
			geo.Query, geo.Country, geo.Region, geo.City, geo.ISP, geo.Org, geo.Lat, geo.Lon)
	},
}
