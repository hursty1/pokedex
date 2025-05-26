package main

import "fmt"


func CommandList(config *Config) error {
	fmt.Print("Your Pokedex:\n\r")
	for _, v := range config.Pokedex {
		fmt.Printf(" - %s\n\r", v.Name)
	}

	return nil
}