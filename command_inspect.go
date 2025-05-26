package main

import "fmt"


func commandInspect(config *Config) error {
	if config.Args == nil {
		
		return fmt.Errorf("Please provide a pokemon to inspect.\n\r")
	}
	pokemonDetails, ok := config.Pokedex[*config.Args]
	if !ok {
		return fmt.Errorf("Plese inspect a already caught pokemon\n\r")
	}
	fmt.Printf("Name: %s\n\r",pokemonDetails.Name)
	fmt.Printf("Height: %d\n\r", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n\r", pokemonDetails.Weight)
	fmt.Printf("Stats:\n\r")
	for _, v := range pokemonDetails.Stats {
		fmt.Printf("\t-%s: %d\n\r", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n\r")
	for _, v := range pokemonDetails.Types {
		fmt.Printf("\t-%s\n\r", v.Type.Name)
	}
	return nil
}