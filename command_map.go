package main

import (
	"errors"
	"fmt"
	"pokedex/internal/pokeapi"
)


func CommandMap(config *Config) error {
	locationResponse, err := config.pokeapiClient.ListLocations(config.Next)
	if err != nil {
		return nil
	}
	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous

	PrintNames(locationResponse.Results)

	return nil
}


func CommandMapB(config *Config) error {
	if config.Previous == nil {
		return errors.New("\ryou're on the first page\r")
	}
	var (
		locationResponse pokeapi.LocationAreaResponse
		err error
	)
	locationResponse, err = config.pokeapiClient.ListLocations(config.Previous)
	if err != nil {
		return nil
	}
	config.Next = locationResponse.Next
	config.Previous = locationResponse.Previous

	PrintNames(locationResponse.Results)
	
	return nil
}


func PrintNames(locationAreas []pokeapi.LocationArea) {
	for _, location := range locationAreas {
		fmt.Print(location.Name + "\r\n")
	}
}