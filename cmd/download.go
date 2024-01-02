package cmd

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/hashicorp/go-version"
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
			cookieRequestUrl := util.AppleDownloadServicesURL(url)
			adcAuthCookie := util.ADCCookieHeader(cookieRequestUrl)
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

	var matchingRuntimes []util.PlatformDownloadable
	var matchedRuntime util.PlatformDownloadable

	for _, v := range data.Downloadables {
		if strings.HasPrefix(v.Name, runtimeName) {

			foundRuntime := util.PlatformDownloadable{
				Identifier: v.Identifier,
				Source:     v.Source,
				Name:       v.Name,
				Platform:   v.Platform,
			}

			ver, _ := version.NewVersion(v.Version)
			foundRuntime.Version = *ver

			if v.Authentication != "" {
				foundRuntime.AuthRequired = true
			}
			foundRuntime.DownloadFileName = v.Name + ".dmg"
			matchingRuntimes = append(matchingRuntimes, foundRuntime)
		}
	}

	// TODO: If there's 'Simulator' somewhere in the provided name, just try and match an exact slice elem
	// and bypass the
	// if strings.Contains(runtimeName, "Simulator") {

	// }

	// If there are multiple matches, take the highest-versioned one
	if len(matchingRuntimes) > 1 {
		slices.SortFunc(matchingRuntimes, func(a, b util.PlatformDownloadable) int {
			return a.Version.Compare(&b.Version)
		})
		matchedRuntime = matchingRuntimes[len(matchingRuntimes)-1]
	}

	idx := slices.IndexFunc(matchingRuntimes,
		func(r util.PlatformDownloadable) bool {
			// Every entry contains the words 'Simulator' or 'Simulator Runtime', and we
			// only care about the words leading up to that
			normalizedName := strings.Split(r.Name, " Simulator")[0]
			return normalizedName == runtimeName
		})
	if idx >= 0 {
		matchedRuntime = matchingRuntimes[idx]
	}

	fmt.Println(matchedRuntime)

	return matchedRuntime.DownloadFileName,
		matchedRuntime.Source,
		matchedRuntime.AuthRequired
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
