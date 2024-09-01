package main

import (
	"fmt"
	"io"
	"os"
)

type cliCommand struct {
	name        string
	description string
	action      func() error
}

func commandHelp() error {
	io.WriteString(os.Stdout, "\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, command := range getCLICommands() {
		io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", command.name, command.description))
	}

	return nil
}

func commandExit() error {
	os.Exit(0)
	return io.EOF
}

func getCLICommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			action:      commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			action:      commandExit,
		},
	}
}
