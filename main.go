package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	userInput := bufio.NewScanner(os.Stdin)

	io.WriteString(os.Stdout, "Pokedex > ")

	for userInput.Scan() {
		if err := userInput.Err(); err != nil {
			log.Fatalf("error reading user input: %v", err)
		}

		text := normalizeInput(userInput.Text())
		cliCommands := getCLICommands()

		if command, exists := cliCommands[text]; exists {
			command.action()
		} else {
			io.WriteString(os.Stdout, fmt.Sprintf("unknown command: %s\n", text))
		}

		io.WriteString(os.Stdout, "\nPokedex > ")
	}
}
