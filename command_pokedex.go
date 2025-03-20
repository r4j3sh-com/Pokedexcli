package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	// pokedex command will show the list of all caught Pokémon
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you haven't caught any Pokémon yet")
	}

	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", name)
	}

	return nil
}
