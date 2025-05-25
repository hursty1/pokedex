package main

import "fmt"


func commandInspect(config *Config) error {
	if config.Args == nil {
		
		return fmt.Errorf("Please provide a pokemon to inspect.")
	}
	pokemonDetails, ok := config.Pokedex[*config.Args]
	if !ok {
		return fmt.Errorf("Plese inspect a already caught pokemon")
	}
	fmt.Printf("Name: %s\n",pokemonDetails.Name)
	fmt.Printf("Height: %d\n", pokemonDetails.Height)
	fmt.Printf("Weight: %d\n", pokemonDetails.Weight)
	fmt.Printf("Stats:\n")
	for _, v := range pokemonDetails.Stats {
		fmt.Printf("\t-%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemonDetails.Types {
		fmt.Printf("\t-%s\n", v.Type.Name)
	}
	return nil
}