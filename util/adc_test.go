package util

import "testing"

func TestAppleDownloadServicesURL(t *testing.T) {
	got := appleDownloadServicesURL("https://fake.apple.com/more/paths/sim.dmg")
	if got != "https://developerservices2.apple.com/services/download?path=/more/paths/sim.dmg" {
		t.Errorf("appleDownloadServicesURL(\"https://fake.apple.com/more/paths/sim.dmg\") = %v; want 1", got)
	}
}
