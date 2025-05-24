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
// func ListPokemonLocations(url *string, config *Config) (pokeapi.LocationAreaResponse, error) {
// 	var (
// 			locationResponse pokeapi.LocationAreaResponse
// 			err error
// 			data []byte
// 		)
// 	if url == nil { // first run
// 		locationResponse, err = config.pokeapiClient.ListLocations(url)
// 		if err != nil {
// 			LocationAreaResponse := pokeapi.LocationAreaResponse{}
// 			return LocationAreaResponse, err
// 		}
// 		//marshal data
// 		data, err = json.Marshal(locationResponse)
// 		if err != nil {
// 			return locationResponse, err
// 		}
// 		config.Cache.Add("https://pokeapi.co/api/v2/location-area", data)
// 		return locationResponse, err
	

// 	} else { //sec
// 		//check cahce
// 		if data, found := config.Cache.Get(*url); found {
// 			err = json.Unmarshal(data, &locationResponse)
// 			if err != nil {
// 				return locationResponse, err
// 			}
// 			return locationResponse, err
		
// 		} else {
// 			locationResponse, err = config.pokeapiClient.ListLocations(url)
// 			if err != nil {
// 				LocationAreaResponse := pokeapi.LocationAreaResponse{}
// 				return LocationAreaResponse, err
// 			}
// 			//marshal data
// 			data, err = json.Marshal(locationResponse)
// 			if err != nil {
// 				return locationResponse, err
// 			}
// 			config.Cache.Add(*url, data)
// 			return locationResponse, err
// 		}
// 	}
// }

func CommandMapB(config *Config) error {
	if config.Previous == nil {
		return errors.New("you're on the first page")
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
		fmt.Println(location.Name)
	}
}