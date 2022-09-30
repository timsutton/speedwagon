package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/timsutton/speedwagon/util"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Output the program version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(util.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
