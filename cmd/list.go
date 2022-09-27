/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"howett.net/plist"
)

type DVTDownloadablePlist struct {
	SdkToSimulatorMappings []struct {
		SdkBuildUpdate       string `plist:"sdkBuildUpdate"`
		SimulatorBuildUpdate string `plist:"simulatorBuildUpdate"`
		SdkIdentifier        string `plist:"sdkIdentifier"`
	} `plist:"sdkToSimulatorMappings"`
	SdkToSeedMappings []struct {
		BuildUpdate string `plist:"buildUpdate"`
		Platform    string `plist:"platform"`
		SeedNumber  int    `plist:"seedNumber"`
	} `plist:"sdkToSeedMappings"`
	RefreshInterval int `plist:"refreshInterval"`
	Downloadables   []struct {
		Category         string `plist:"category"`
		SimulatorVersion struct {
			BuildUpdate string `plist:"buildUpdate"`
			Version     string `plist:"version"`
		} `plist:"simulatorVersion"`
		Source            string `plist:"source"`
		DictionaryVersion int    `plist:"dictionaryVersion"`
		ContentType       string `plist:"contentType"`
		Platform          string `plist:"platform"`
		Identifier        string `plist:"identifier"`
		Version           string `plist:"version"`
		FileSize          int64  `plist:"fileSize"`
		HostRequirements  struct {
			MaxHostVersion string `plist:"maxHostVersion"`
		} `plist:"hostRequirements,omitempty"`
		Name              string `plist:"name"`
		HostRequirements0 struct {
			ExcludedHostArchitectures []string `plist:"excludedHostArchitectures"`
			MaxHostVersion            string   `plist:"maxHostVersion"`
		} `plist:"hostRequirements,omitempty"`
		HostRequirements1 struct {
			ExcludedHostArchitectures []string `plist:"excludedHostArchitectures"`
			MaxHostVersion            string   `plist:"maxHostVersion"`
		} `plist:"hostRequirements,omitempty"`
		HostRequirements2 struct {
			ExcludedHostArchitectures []string `plist:"excludedHostArchitectures"`
		} `plist:"hostRequirements,omitempty"`
		HostRequirements3 struct {
			MinHostVersion  string `plist:"minHostVersion"`
			MinXcodeVersion string `plist:"minXcodeVersion"`
		} `plist:"hostRequirements,omitempty"`
		Authentication string `plist:"authentication,omitempty"`
	} `plist:"downloadables"`
	Version string `plist:"version"`
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available simulators from Apple's metadata",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Fetching the plist from Apple
		dvtResp, err := http.Get("https://devimages-cdn.apple.com/downloads/xcode/simulators/index2.dvtdownloadableindex")
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(dvtResp.Body)
		dvtResp.Body.Close()
		if dvtResp.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", dvtResp.StatusCode, body)
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%s", body)

		// DEBUG: reading from a local file instead
		// file, _ := os.ReadFile("dvt.plist")
		data := DVTDownloadablePlist{}

		_, err = plist.Unmarshal([]byte(body), &data)
		if err != nil {
			fmt.Println(err)
		}

		for i := 0; i < len(data.Downloadables); i++ {
			if data.Downloadables[i].ContentType == "diskImage" {
				fmt.Println("----")
				fmt.Println("Name: ", data.Downloadables[i].Name)
				fmt.Println("Version: ", data.Downloadables[i].SimulatorVersion.Version)
				fmt.Println("Build: ", data.Downloadables[i].SimulatorVersion.BuildUpdate)
				fmt.Println("----")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
