/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/timsutton/learn-go/util"
	"howett.net/plist"
)

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
		util.RefreshDVTMetadata()
		// TODO: reading the file contents and returning the struct should probably be its own set of
		// helper functions in util package
		body, _ := os.ReadFile(util.DVTCacheFilePath())

		data := util.DVTDownloadablePlist{}
		_, err := plist.Unmarshal([]byte(body), &data)
		if err != nil {
			fmt.Println(err)
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Name", "Version", "Build"})

		for i := 0; i < len(data.Downloadables); i++ {
			if data.Downloadables[i].ContentType == "diskImage" {
				t.AppendRow([]interface{}{data.Downloadables[i].Name, data.Downloadables[i].SimulatorVersion.Version, data.Downloadables[i].SimulatorVersion.BuildUpdate})
				t.AppendSeparator()
			}
		}
		t.SetStyle(table.StyleLight)
		t.Render()
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
