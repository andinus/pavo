package main

import (
	"fmt"
	"os"

	"tildegit.org/andinus/lynx"
)

// blockUnveil func blocks further unveil calls.
func blockUnveil() {
	err := lynx.UnveilBlock()
	if err != nil {
		fmt.Printf("%s :: %s",
			"UnveilBlock() failed",
			err.Error())
		os.Exit(1)
	}

	// We drop unveil from promises after blocking it. We drop
	// rpath too because the config file has been read.
	err = lynx.PledgePromises("stdio exec")
	if err != nil {
		fmt.Printf("%s :: %s",
			"blockUnveil failed",
			err.Error())
		os.Exit(1)
	}
}

// initUnveil initializes unveil for inital use.
func initUnveil() {
	err := lynx.Unveil(configFile, "rc")
	if err != nil {
		fmt.Printf("%s :: %s",
			"Unveil configFile failed",
			err.Error())
		os.Exit(1)
	}
}
