package main

import (
	"fmt"
	"math/rand"
)


func CommandCatch(config *Config) error {
	if config.Args == nil {
		
		return fmt.Errorf("Please provide a pokemon to catch")
	}

	fmt.Printf("Arguments are: %s \n", *config.Args)

	fmt.Printf("Throwing a Pokeball at %s...\n", *config.Args)

	pokemonResponse, err := config.pokeapiClient.Pokemon(config.Args)
	if err != nil {
		return err
	}

	baseExperience := pokemonResponse.BaseExperience
	chance := float32(rand.Intn(baseExperience)) / float32(baseExperience)
	if chance > 0.4 {
		fmt.Printf("%s was caught!\n", *config.Args)
		config.Pokedex[*config.Args] = pokemonResponse
	} else {
		fmt.Printf("%s excaped!\n", *config.Args)
	}
	return nil
}