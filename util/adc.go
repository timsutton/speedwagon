package util

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Return cookie header and value
func ADCCookieHeader(downloadUrl string) string {
	cookieRequestUrl := appleDownloadServicesURL(downloadUrl)
	resp, err := http.Get(cookieRequestUrl)
	if err != nil {
		// TODO: pass this error back
		log.Fatal(err)
	}
	// naively assuming we have a single Set-Cookie and no problem scanning for the cookie name,
	// no error-checking here yet!
	header := resp.Header["Set-Cookie"][0]
	firstSegment, _, _ := strings.Cut(header, ";")
	_, valueOnly, _ := strings.Cut(firstSegment, "ADCDownloadAuth=")

	return valueOnly
}

// Given an ADC download URL that requires authentication, return the /services/download
// service URL we can use to retrieve the needed auth header for that ADC download URL
// i.e. https://developerservices2.apple.com/services/download?path=/path/to/the.dmg
func appleDownloadServicesURL(adcUrl string) string {
	newUrl := "https://developerservices2.apple.com/services/download?path="
	// blindly assuming Apple's URLs are well-formed
	adcUrlPath, _ := url.Parse(adcUrl)
	newUrl += adcUrlPath.Path
	return newUrl
}
