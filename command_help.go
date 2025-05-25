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
	fmt.Print("\r\n") // reset and newline
	fmt.Print("\rWelcome to the Pokedex!\r\n")
	fmt.Print("\rUsage:\r\n\r\n")

	for _, commands := range GetCommands() {
		fmt.Printf("\r%-10s : %s\r\n", commands.name, commands.descirption)
	}

	fmt.Print("\r\n")
	return nil
}