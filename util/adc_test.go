package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppleDownloadServicesURL(t *testing.T) {
	assert.Equal(t,
		appleDownloadServicesURL("https://fake.apple.com/more/paths/sim.dmg"),
		"https://developerservices2.apple.com/services/download?path=/more/paths/sim.dmg")
}

// func TestADCCookieHeader(t *testing.T) {
// 	fmt.Println("TODO")
// 	ADCCookieHeader("https://foo.com/")
// }
