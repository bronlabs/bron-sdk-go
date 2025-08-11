package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// Read version from VERSION file
	versionBytes, err := ioutil.ReadFile("VERSION")
	if err != nil {
		fmt.Printf("Error reading VERSION file: %v\n", err)
		os.Exit(1)
	}

	version := strings.TrimSpace(string(versionBytes))

	// Update version in version.go
	versionContent := fmt.Sprintf("package version\n\nconst SDK_VERSION = \"%s\"\n", version)

	err = ioutil.WriteFile("sdk/version/version.go", []byte(versionContent), 0644)
	if err != nil {
		fmt.Printf("Error updating version.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated SDK_VERSION to %s\n", version)
}
