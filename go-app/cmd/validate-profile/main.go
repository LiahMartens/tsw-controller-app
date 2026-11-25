package main

import (
	"fmt"
	"os"
	"path/filepath"

	"tsw_controller_app/config"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <profile.tswprofile> [...]\n", os.Args[0])
		os.Exit(1)
	}

	exitCode := 1
	for _, arg := range os.Args[1:] {
		path, err := filepath.Abs(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			continue
		}

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error reading %s: %v\n", path, err)
			continue
		}

		metadata := config.Config_Controller_Profile_Metadata{
			Path:       filepath.Base(path),
			IsEmbedded: false,
		}

		fmt.Printf("validating: %s\n", path)

		profile, err := config.ControllerProfileFromJSON(string(data), metadata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "  FAIL: %v\n", err)
			continue
		}

		// Also validate the JSON schema constraints manually (Go validator
		// covers struct tags but not things like min lengths on nested objects).
		if len(profile.Controls) == 0 {
			fmt.Fprintf(os.Stderr, "  FAIL: controls array is empty\n")
			continue
		}

		controlNames := make(map[string]bool)
		for _, c := range profile.Controls {
			if controlNames[c.Name] {
				fmt.Fprintf(os.Stderr, "  FAIL: duplicate control name %q\n", c.Name)
				exitCode = 1
				break
			}
			controlNames[c.Name] = true
		}

		fmt.Printf("  OK  (%d controls)\n", len(profile.Controls))
		exitCode = 0
	}

	if exitCode != 0 {
		os.Exit(1)
	}
}
