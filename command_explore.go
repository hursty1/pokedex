package main

import (
	"fmt"
	"pokedex/internal/pokeapi"
)


func CommandExplore(config *Config) error {
	if config.Args == nil {
		
		return fmt.Errorf("\rPlease provide a region to explore\r")
	}
	fmt.Printf("\rArguments are: %s \n\r", *config.Args)

	locationDetailResponse, err := config.pokeapiClient.LocationDetails(config.Args)
	if err != nil {
		return err
	}
	fmt.Print("\r" + locationDetailResponse.Name + "\r")
	PrintPokemonNames(locationDetailResponse)
	return nil
}


func PrintPokemonNames(locationDetails pokeapi.LocationDetailResponse) {
	encounters := locationDetails.PokemonEncounters
	for _, encounter := range encounters {
		fmt.Printf("- %s\n\r", encounter.Pokemon.Name)
	}
}