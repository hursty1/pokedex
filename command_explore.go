package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)


func CommandExplore(config *Config) error {
	if config.Args == nil {
		
		return fmt.Errorf("Please provide a region to explore")
	}
	fmt.Printf("Arguments are: %s \n", *config.Args)

	locationDetailResponse, err := config.pokeapiClient.LocationDetails(config.Args)
	if err != nil {
		return err
	}
	fmt.Println(locationDetailResponse.Name)
	PrintPokemonNames(locationDetailResponse)
	return nil
}


func PrintPokemonNames(locationDetails pokeapi.LocationDetailResponse) {
	encounters := locationDetails.PokemonEncounters
	for _, encounter := range encounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
}