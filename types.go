package main

import "pokedex/internal/pokeapi"


type cliCommands struct {
	name string
	descirption string
	callback func(*Config) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next *string
	Previous *string
	// Cache *pokecache.Cache
}