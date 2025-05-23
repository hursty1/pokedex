package main

import (
	"fmt"
)

// func commandHelp() error {
// 	fmt.Println(`Welcome to the Pokedex!
// Usage:

// help: Displays a help message
// exit: Exit the Pokedex`)
// 	return errors.New("")
// }

func CommandHelp(config *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, commands := range GetCommands() {
		fmt.Printf("%s : %s \n", commands.name, commands.descirption)
	}
	fmt.Println()
	return nil
}