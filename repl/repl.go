package repl

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/flrn000/pokedexcli/internal/utils"
)

type Config struct {
	nextPageURl, previousPageURL *string
}

func Start(cfg *Config) {
	userInput := bufio.NewScanner(os.Stdin)

	io.WriteString(os.Stdout, "Pokedex > ")

	for userInput.Scan() {
		if err := userInput.Err(); err != nil {
			log.Fatalf("error reading user input: %v", err)
		}

		text := utils.NormalizeInput(userInput.Text())
		cliCommands := getCLICommands()

		if command, exists := cliCommands[text]; exists {
			err := command.action(cfg)
			if err != nil {
				fmt.Println(err)
				io.WriteString(os.Stdout, "\nPokedex > ")

				continue
			}
		} else {
			io.WriteString(os.Stdout, fmt.Sprintf("unknown command: %s\n", text))
		}

		io.WriteString(os.Stdout, "\nPokedex > ")
	}
}
