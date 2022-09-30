/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/timsutton/speedwagon/util"

	"github.com/dustin/go-humanize"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var Data util.DVTDownloadablePlist

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available simulators from Apple's metadata",
	Run: func(cmd *cobra.Command, args []string) {
		Data = util.DVTMetadata()
		TableOutput(Data)
	},
}

func TableOutput(data util.DVTDownloadablePlist) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Version", "Build", "Type", "Size"})

	for i := 0; i < len(Data.Downloadables); i++ {
		t.AppendRow([]interface{}{Data.Downloadables[i].Name,
			Data.Downloadables[i].SimulatorVersion.Version,
			Data.Downloadables[i].SimulatorVersion.BuildUpdate,
			Data.Downloadables[i].ContentType,
			humanize.Bytes(uint64(Data.Downloadables[i].FileSize))})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleLight)
	t.Render()

}

func init() {
	rootCmd.AddCommand(listCmd)

	// Ideas for supporting structured output formats, for later
	// listCmd.Flags().BoolVar(&u, "json", false, "Output in JSON")
	// listCmd.Flags().BoolVar(&pw, "yaml", false, "Output in YAML")
	// listCmd.Flags().StringVarP(&Format, "format", "f", "ascii", "Output format")
}
