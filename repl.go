package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func StartRepl(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	// config := &Config{}
	for {

		fmt.Print("Pokedex >")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) > 0 {
			
			command, ok := GetCommands()[input[0]]
			if ok {
				err := command.callback(config)
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

func GetCommands() map[string]cliCommands {
	return map[string]cliCommands{
	"exit": {
		name: "exit",
		descirption: "Exit the pokedex",
		callback: CommandExit,
	},
	"help": {
		name: "help",
		descirption:"Displays a help message",
		callback: CommandHelp,
	},
	"map": {
		name: "map",
		descirption: "Displays names of locations in the pokemon world",
		callback: CommandMap,
	},
	"mapb": {
		name: "map back",
		descirption: "Displays the last 20 locations in the pokemon world",
		callback: CommandMapB,
	},
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
