/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download <runtime name>",
	Short: "Download the desired simulator to a local path",
	Long: `Given a name/version specifier of a simulator runtime, download
it to a local destination. It will be saved to the current
working directory, named '<name of runtime>.(dmg|pkg)'.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		// args[0] is safe because we're requiring exactly one arg
		desiredRuntimeName := args[0]
		fmt.Println("here is where we would refresh")
		// TODO: call the func in refresh.go
		fmt.Printf("..for desired runtime %v\n", desiredRuntimeName)
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
