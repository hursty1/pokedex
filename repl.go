package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommands struct {
	name string
	descirption string
	callback func() error
}


func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("Pokedex >")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			
			command, ok := getCommands()[input[0]]
			if ok {
				err := command.callback()
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
			
			
		}
	}
}

func getCommands() map[string]cliCommands {
	return map[string]cliCommands{
	"exit": {
		name: "exit",
		descirption: "Exit the pokedex",
		callback: commandExit,
	},
	"help": {
		name: "help",
		descirption:"Displays a help message",
		callback: commandHelp,
	},
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
