package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/r4j3sh-com/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a Pokémon name")
	}
	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemonByName(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Add code to simulate catching the Pokémon
	// Generate a random number based on the Pokémon's base experience
	// Calculate catch rate
	catchRate := calculateCatchRate(pokemon)

	// Generate a random number between 0 and 100
	randomNum := rand.Intn(101)

	//fmt.Printf("Catch rate: %d, Random number: %d\n", catchRate, randomNum)

	if randomNum <= catchRate {
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func calculateCatchRate(pokemon pokeapi.Pokemon) int {
	// Base Catch rate (we can adjust as needed)
	baseCatchRate := 50

	// Adjust catch rate based on Pokémon's base experience
	// The higher the base experience, the lower the catch rate
	catchRate := baseCatchRate - int(float64(pokemon.BaseExperience)*0.2)

	// Add some randomness to the catch rate
	randomFactor := rand.Intn(21) - 10 // Random number between -10 and 10
	catchRate += randomFactor

	// Ensure catch rate is between 10 and 90
	if catchRate < 5 {
		catchRate = 5
	} else if catchRate > 80 {
		catchRate = 80
	}
	return catchRate
}
