package main

import (
	"fmt"
	"pokedex/internal/cli"
	"strings"
)


func StartRepl(config *Config) {
	// scanner := bufio.NewScanner(os.Stdin)
	cli, err := cli.Setup()
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer term.Restore(int(os.Stdin.Fd()), cli.OldState)
	defer cli.CleanUp()
	// fmt.Printf("CLI Started %d\n", cli.HistoryIndex)
	// config := &Config{}
	for {
		// b, _ := cli.Reader.ReadByte()
		// fmt.Println(b)
		cli.Input()
		// fmt.Print("Pokedex >")
		// scanner.Scan()
		// input := cleanInput(scanner.Text())
		// if len(input) > 0 {
			
		// 	command, ok := GetCommands()[input[0]]
		// 	if ok {
		// 		if len(input) > 1 {
		// 			config.Args = &input[1]
		// 		}
		// 		err := command.callback(config)
		// 		if err != nil {
		// 			fmt.Println(err)
		// 		}
		// 		continue
		// 	} else {
		// 		fmt.Println("Unknown command")
		// 		continue
		// 	}
			
			
		// }
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
	"explore": {
		name: "explore",
		descirption: "Given a location or location id, it will list the pokemon avilable.",
		callback: CommandExplore,
	},
	"catch": {
		name: "catch",
		descirption: "Attempt to catch a pokemon.",
		callback: CommandCatch,
	},
	"pokedex": {
		name: "pokedex",
		descirption: "List caught Pokemon.",
		callback: CommandList,
	},
	"inspect": {
		name: "inspect",
		descirption: "Inspect a caught pokemons pokedex entry.",
		callback: commandInspect,
	},
	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
