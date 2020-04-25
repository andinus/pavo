package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"tildegit.org/andinus/lynx"
)

type Pavo struct {
	Config []Config `json:"config"`
}

type Unveil struct {
	Permissions string `json:"permissions"`
	Path        string `json:"path"`
}

type Config struct {
	Command          string   `json:"command"`
	Unveil           []Unveil `json:"unveil"`
	CommandsToUnveil []string `json:"commandstounveil"`
	Execpromises     string   `json:"execpromises"`
}

func parseConfig() {
	// Read the config file
	data, err := ioutil.ReadFile(configFile())
	if err != nil {
		fmt.Printf("%s :: %s",
			"Failed to read the config file",
			err.Error())
		os.Exit(1)
	}

	var Pavo Pavo

	// Unmarshal data to struct.
	err = json.Unmarshal(data, &Pavo)
	if err != nil {
		fmt.Printf("%s :: %s\n",
			"Failed to unmarshal the config file",
			err.Error())
		os.Exit(1)
	}

	// Get command & flags from flag.Args().
	flag.Parse()
	flags = flag.Args()[1:]

	for k, v := range Pavo.Config {
		// If we find the command's config then break.
		if v.Command == flag.Args()[0] {
			cmd = v
			break
		}

		// If we don't find the config in file then return an
		// error.
		if k == len(Pavo.Config)-1 {
			fmt.Printf("%s %s",
				"Failed to find config for",
				flag.Args()[0])
			os.Exit(1)
		}
	}

	// Commands to unveil.
	commands := []string{cmd.Command}

	// Get commands to unveil from config.
	for _, v := range cmd.CommandsToUnveil {
		commands = append(commands, v)
	}

	// Unveil all values of $PATH/command.
	err = lynx.UnveilCommands(commands)
	if err != nil {
		fmt.Printf("%s :: %s",
			"Failed to unveil command",
			err.Error())
		os.Exit(1)
	}

	// Set execpromises from config.
	err = lynx.PledgeExecpromises(cmd.Execpromises)
	if err != nil {
		fmt.Printf("%s :: %s",
			"PledgeExecpromises failed",
			err.Error())
		os.Exit(1)
	}

	// Unveil all paths from config.
	paths := make(map[string]string)
	for _, v := range cmd.Unveil {
		paths[v.Path] = v.Permissions
	}

	err = lynx.UnveilPaths(paths)
	if err != nil {
		fmt.Printf("%s :: %s",
			"Failed to unveil paths",
			err.Error())
		os.Exit(1)
	}
}
