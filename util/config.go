package util

import (
	"log"
	"os"
)

// Basic setup to ensure our working dir exists, etc.
func SetupEnvironment() {
	err := os.MkdirAll(AppCacheDir(), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
