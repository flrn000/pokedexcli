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
	action      func(*Config, string) error
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
		"explore": {
			name:        "explore",
			description: "Displays a list of all the Pokemon in a given area",
			action:      commandExplore,
		},
	}
}

func commandHelp(cfg *Config, commandArg string) error {
	io.WriteString(os.Stdout, "\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, command := range getCLICommands() {
		io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", command.name, command.description))
	}

	return nil
}

func commandExit(cfg *Config, commandArg string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, commandArg string) error {
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

func commandMapB(cfg *Config, commandArg string) error {
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

func commandExplore(cfg *Config, commandArg string) error {
	if len(commandArg) == 0 {
		return fmt.Errorf("no location area name provided")
	}

	data, err := service.Explore(commandArg)
	if err != nil {
		return err
	}

	io.WriteString(os.Stdout, fmt.Sprintf("Exploring %v...\nFound Pokemon:\n", commandArg))

	for _, v := range data.PokemonEncounters {
		fmt.Println("- ", v.Pokemon.Name)
	}

	return nil
}
