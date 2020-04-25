package main

import (
	"fmt"
	"os"

	"tildegit.org/andinus/lynx"
)

// initPledge initializes pledge for initial use.
func initPledge() {
	// Pledge promises can only be dropped & we cannot add
	// anything so this call adds everything that maybe used in
	// program later. We don't define execpromises here because
	// that comes from the user.
	//
	// Note: Don't forget to change blockUnveil() if you add
	// anything new here.
	err := lynx.PledgePromises("unveil stdio rpath exec")
	if err != nil {
		fmt.Printf("%s :: %s",
			"initPledge failed",
			err.Error())
		os.Exit(1)
	}
}
