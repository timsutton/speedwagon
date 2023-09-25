package util

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"howett.net/plist"
)

const DVTDownloadableIndexUrl = "https://devimages-cdn.apple.com/downloads/xcode/simulators/index2.dvtdownloadableindex"

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

func RefreshDVTMetadata() {
	file, err := os.Create(DVTCacheFilePath())
	if err != nil {
		// TODO: bad form to exit the whole program from somewhere this deep, should instead return
		// an error out of these functions and bubble it up to the high-level command at least, or main
		log.Fatal(err)
	}

	dvtResp, err := http.Get(DVTDownloadableIndexUrl)
	if err != nil {
		// TODO: same as above
		log.Fatal(err)
	}
	_, err = io.Copy(file, dvtResp.Body)
	if err != nil {
		log.Fatal(err)
	}
	dvtResp.Body.Close()
	if dvtResp.StatusCode > 299 {
		// TODO: same as above
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", dvtResp.StatusCode, dvtResp.Body)
	}
	if err != nil {
		// TODO: same as above
		log.Fatal(err)
	}

	// DEBUG: reading from a local file instead, if there's no network available or to speed things up
	// file, _ := os.ReadFile("dvt.plist")
	// io.Copy(file, dvtResp.Body)
}

// Returns a complete, up-to-date DVTDownloadablePlist struct
func DVTMetadata() DVTDownloadablePlist {
	// For now, we just always call a refresh until we have a better way to cache its response
	RefreshDVTMetadata()

	body, _ := os.ReadFile(DVTCacheFilePath())

	data := DVTDownloadablePlist{}
	_, err := plist.Unmarshal([]byte(body), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func AppCacheDir() string {
	userCacheDir, _ := os.UserCacheDir()
	return filepath.Join(userCacheDir, "ca.macops.speedwagon")
}

func DVTCacheFilePath() string {
	return filepath.Join(AppCacheDir(), "dvt_index.plist")
}
