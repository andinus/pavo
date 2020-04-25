package main

import (
	"fmt"
	"os"
)

// configFile returns the config file for pavo. It is
// configDir/pavo.json by default but can be changed by setting
// PAVO_CONFIG_FILE environment variable. If you set this then no need
// to set PAVO_CONFIG_DIR.
func configFile() (configFile string) {
	configFile = os.Getenv("PAVO_CONFIG_FILE")

	if len(configFile) == 0 {
		configFile = fmt.Sprintf("%s/%s",
			configDir(),
			"pavo.json")
	}

	return
}

// configDir returns the config directory for pavo. First
// PAVO_CONFIG_DIR environment variable is checked, if it doesn't
// exist then XDG_CONFIG_HOME & if that is not set either then we
// assume it to be the default value which is $HOME/.config according
// to XDG Base Directory Specification.
func configDir() (configDir string) {
	configDir = os.Getenv("PAVO_CONFIG_DIR")

	if len(configDir) == 0 {
		configDir = os.Getenv("XDG_CONFIG_HOME")
	}

	if len(configDir) == 0 {
		configDir = fmt.Sprintf("%s/%s",
			os.Getenv("HOME"),
			".config")
	}

	return
}
