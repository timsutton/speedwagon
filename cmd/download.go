package cmd

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/schollz/progressbar/v3"
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
		client := grab.NewClient()
		req, _ := grab.NewRequest(".", url)
		if authRequired {
			adcAuthCookie := util.ADCCookieHeader(url)
			req.HTTPRequest.Header.Set("Cookie", "ADCDownloadAuth="+adcAuthCookie)
		}

		fmt.Printf("Downloading %v...\n", req.URL())
		resp := client.Do(req)

		// start UI loop
		t := time.NewTicker(500 * time.Millisecond)
		defer t.Stop()

		bar := progressbar.DefaultBytes(
			resp.Size(),
			"",
		)

	Loop:
		for {
			select {
			case <-t.C:
				err := bar.Set(int(resp.BytesComplete()))
				if err != nil {
					log.Fatal(err)
				}

			case <-resp.Done:
				// download is complete
				break Loop
			}
		}

		if err := resp.Err(); err != nil {
			fmt.Printf("\n\n")
			panic(err)
		}
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
