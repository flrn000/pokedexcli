package repl

import (
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"os"

	"github.com/flrn000/pokedexcli/internal/service"
)

type cliCommand struct {
	name        string
	description string
	action      func(*Config, ...string) error
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
			name:        "explore <location_name>",
			description: "Displays a list of all the Pokemon in a given area",
			action:      commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			action:      commandCatch,
		},
	}
}

func commandHelp(cfg *Config, args ...string) error {
	io.WriteString(os.Stdout, "\nWelcome to the Pokedex!\nUsage:\n\n")

	for _, command := range getCLICommands() {
		io.WriteString(os.Stdout, fmt.Sprintf("%s: %s\n", command.name, command.description))
	}

	return nil
}

func commandExit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, args ...string) error {
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

func commandMapB(cfg *Config, args ...string) error {
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

func commandExplore(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no location area name provided")
	}

	data, err := service.Explore(args[0])
	if err != nil {
		return err
	}

	io.WriteString(os.Stdout, fmt.Sprintf("Exploring %v...\nFound Pokemon:\n", args[0]))

	for _, v := range data.PokemonEncounters {
		fmt.Println("- ", v.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no pokemon name provided")
	}
	pokemonName := args[0]

	io.WriteString(os.Stdout, fmt.Sprintf("Throwing a Pokeball at %v...\n", pokemonName))
	pokemonInfo, err := service.Catch(pokemonName)
	if err != nil {
		return err
	}

	catchChance := rand.IntN(pokemonInfo.BaseExperience)
	catchThreshold := pokemonInfo.BaseExperience/2 + 10

	if catchChance >= catchThreshold {
		pokedex[pokemonName] = pokemonInfo
		io.WriteString(os.Stdout, fmt.Sprintf("%v was caught!\n", pokemonName))
		return nil
	} else {
		io.WriteString(os.Stdout, fmt.Sprintf("%v escaped!\n", pokemonName))
		return nil
	}
}
