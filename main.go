package main

import (
	"fmt"
	"os"
	"os/exec"
)

var (
	version = "v0.1.0"

	cmd   Config
	flags []string
)

func main() {
	initPledge()
	initUnveil()

	if len(os.Args) == 1 {
		fmt.Println("Usage: pavo <command> <flags>")
		os.Exit(1)
	}

	if os.Args[1] == "version" {
		fmt.Println("Pavo", version)
		os.Exit(0)
	}

	// Parse the config file.
	parseConfig()

	// Block futher unveil calls.
	blockUnveil()

	// Get path of command.
	cmdPath, err := exec.LookPath(cmd.Command)
	if err != nil {
		fmt.Printf("%s %s :: %s\n",
			cmd.Command,
			"not found in $PATH",
			err.Error())
		os.Exit(1)
	}

	// TODO: Make the output realtime.
	out, err := exec.Command(cmdPath, flags...).Output()
	if err != nil {
		err = fmt.Errorf("%s :: %s",
			"Failed to execute command",
			err.Error())
		os.Exit(1)
	}

	// Print the output.
	fmt.Printf(string(out))
}
