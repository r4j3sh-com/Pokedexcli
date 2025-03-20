package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokémon name")
	}
	name := args[0]

	if _, caught := cfg.caughtPokemon[name]; !caught {
		return errors.New("you haven't caught that Pokémon yet")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemonByName(name)
	if err != nil {
		return err
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return nil

	/*
			Name: pidgey
		Height: 3
		Weight: 18
		Stats:
		  -hp: 40
		  -attack: 45
		  -defense: 40
		  -special-attack: 35
		  -special-defense: 35
		  -speed: 56
		Types:
		  - normal
		  - flying
	*/
}
