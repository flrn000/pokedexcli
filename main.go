package main

import (
	"github.com/flrn000/pokedexcli/repl"
)

func main() {
	cfg := repl.Config{}
	repl.Start(&cfg)
}
