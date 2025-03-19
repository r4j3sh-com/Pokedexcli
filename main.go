package main

import (
	"time"

	"github.com/r4j3sh-com/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: pokeClient,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
