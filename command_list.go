package main

import "fmt"


func CommandList(config *Config) error {
	fmt.Println("Your Pokedex:")
	for _, v := range config.Pokedex {
		fmt.Printf(" - %s\n", v.Name)
	}

	return nil
}