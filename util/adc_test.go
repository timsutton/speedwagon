package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppleDownloadServicesURL(t *testing.T) {
	assert.Equal(t,
		AppleDownloadServicesURL("https://fake.apple.com/more/paths/sim.dmg"),
		"https://developerservices2.apple.com/services/download?path=/more/paths/sim.dmg")
}

func TestADCCookieHeader(t *testing.T) {
	// mock developerservices2.apple.com endpoint offering 'virtual' auth cookie
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(
			"Set-Cookie",
			"ADCDownloadAuth=xa%2FSZcGbduzUu%2BB3Q9G%0D%0A;Version=1;Comment=;Domain=apple.com;Path=/;Max-Age=108000;HttpOnly;Expires=Mon, 25 Sep 2023 02:40:01 GMT")
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	value := ADCCookieHeader(server.URL)
	assert.Equal(t, value, "xa%2FSZcGbduzUu%2BB3Q9G%0D%0A")
}
