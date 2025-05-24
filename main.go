package main

import (
	"pokedex/internal/pokeapi"
	"time"
)



func main() {
	pokeClient := pokeapi.NewClient(5 *time.Second, time.Minute*5)
	// pokeCache := pokecache.NewCache(30 * time.Second)

	cfg := &Config{
		pokeapiClient: pokeClient,
		// Cache : pokeCache,
	}
	StartRepl(cfg)
}
