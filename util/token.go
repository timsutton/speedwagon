package util

// Return cookie header and value
func ADCCookieHeader(downloadPath string) (string, string) {
	return "ADCDownloadAuth", "mycookiehere"
}
