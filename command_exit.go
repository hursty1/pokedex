package main

import (
	"fmt"
	"os"
)


func CommandExit(config *Config) error {
	fmt.Println("\rClosing the Pokedex... Goodbye!\r")
	os.Exit(0)
	return nil
}