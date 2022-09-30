package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/timsutton/speedwagon/util"
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

		Data = util.DVTMetadata()
		filename, url, authRequired := findMatchingRuntime(desiredRuntimeName, Data)
		if filename == "" {
			fmt.Printf("Found no runtime matching name: '%s'. Available runtimes:\n", desiredRuntimeName)
			TableOutput(Data)
			// TODO: probably makes sense to exit nonzero here?
			return
		}

		// set up the download request
		client := &http.Client{}
		req, _ := http.NewRequest("GET", url, nil)
		if authRequired {
			adcAuthCookie := util.ADCCookieHeader(url)
			req.Header.Set("Cookie", "ADCDownloadAuth="+adcAuthCookie)
		}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Downloading '%v'...\n", filename)
		// dumb copy from HTTP response to output file
		// TODO: start these downloads in the app support dir, moving them here
		// only when they are complete.
		io.Copy(file, resp.Body)
	},
}

// return the following for the first matching runtime by the requested name:
// 1. string: a sensible filename
// 2. string: source URL for a matching runtime
// 3. bool: whether auth is required
func findMatchingRuntime(runtimeName string, data util.DVTDownloadablePlist) (string, string, bool) {

	var runtimeFilename string
	var runtimeUrl string
	authRequired := false

	foundMatchingRuntime := false
	for _, v := range data.Downloadables {
		if strings.HasPrefix(v.Name, runtimeName) {
			foundMatchingRuntime = true
			runtimeFilename = v.Name + ".dmg"
			runtimeUrl = v.Source
			if v.Authentication != "" {
				authRequired = true
			}
		}
		if foundMatchingRuntime {
			break
		}
	}

	return runtimeFilename, runtimeUrl, authRequired
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
