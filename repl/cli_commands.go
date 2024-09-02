package repl

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/flrn000/pokedexcli/internal/service"
)

type cliCommand struct {
	name        string
	description string
	action      func(*Config) error
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
		"map": {
			name:        "map",
			description: "Displays the names of the next page of locations",
			action:      commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous page of locations",
			action:      commandMapB,
		},
	}
}

func commandHelp(cfg *Config) error {
	io.WriteString(os.Stdout, "\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, command := range getCLICommands() {
		io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", command.name, command.description))
	}

	return nil
}

func commandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	data, err := service.GetLocationAreaData(cfg.nextPageURl)
	if err != nil {
		return err
	}

	cfg.nextPageURl = data.Next
	cfg.previousPageURL = data.Previous

	for _, m := range data.Results {
		fmt.Println(m.Name)
	}

	return nil
}

func commandMapB(cfg *Config) error {
	if cfg.previousPageURL == nil {
		return errors.New("you are on the first page")
	}
	data, err := service.GetLocationAreaData(cfg.previousPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURl = data.Next
	cfg.previousPageURL = data.Previous

	for _, m := range data.Results {
		fmt.Println(m.Name)
	}

	return nil
}
