package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/flrn000/pokedexcli/internal/service"
	"github.com/flrn000/pokedexcli/internal/utils"
)

type Config struct {
	nextPageURl, previousPageURL *string
}

var pokedex = make(map[string]service.PokemonInfo)

func Start(cfg *Config) {
	userInput := bufio.NewScanner(os.Stdin)

	io.WriteString(os.Stdout, "Pokedex > ")

	for userInput.Scan() {
		if err := userInput.Err(); err != nil {
			log.Fatalf("error reading user input: %v", err)
		}

		words := utils.NormalizeInput(userInput.Text())
		if len(words) == 0 {
			io.WriteString(os.Stdout, "\nPokedex > ")

			continue
		}
		cliCommands := getCLICommands()
		commandName := words[0]

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		if command, exists := cliCommands[commandName]; exists {
			err := command.action(cfg, args...)
			if err != nil {
				fmt.Println(err)
				io.WriteString(os.Stdout, "\nPokedex > ")

				continue
			}
		} else {
			io.WriteString(os.Stdout, fmt.Sprintf("unknown command: %s\n", words))
		}

		io.WriteString(os.Stdout, "\nPokedex > ")
	}
}
