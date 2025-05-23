package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type LocationAreaResponse struct {
	Count 			int 	`json:"count"`
	Next			string	`json:"next"`
	Previous		string	`json:"previous"`
	Results			[]LocationArea `json:"results"`
}
type LocationArea struct {
	Id	 			int 	`json:"id"`
	Name	 		string 	`json:"name"`
	Game_Index	 	int 	`json:"game_index"`
	// encounter_method_rates	 []EncounterMethodRate
	// location	 NamedAPIResource
	// names	 []Name
	// pokemon_encounters	[]PokemonEncounter
}

func CommandMap(config *Config) error {
	base_url := "https://pokeapi.co/api/v2/location-area/"
	offset:=0
	limit:=20
	offset_str := fmt.Sprintf("offset=%d", offset)
	limit_str := fmt.Sprintf("limit=%d", limit)
	
	base_url = base_url+"?"+offset_str+"&"+limit_str

	if config.Next != "" {
		base_url = config.Next
	}
	GetLocationArea(base_url, config)
	return nil
}

func CommandMapB(config *Config) error {
	base_url := "https://pokeapi.co/api/v2/location-area/"
	offset:=0
	limit:=20
	offset_str := fmt.Sprintf("offset=%d", offset)
	limit_str := fmt.Sprintf("limit=%d", limit)
	
	base_url = base_url+"?"+offset_str+"&"+limit_str

	if config.Previous != "" {
		base_url = config.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	GetLocationArea(base_url, config)

	return nil
}

func GetLocationArea(base_url string, config *Config) error {
	

	res, err := http.Get(base_url)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// fmt.Println(string(body))

	lar := LocationAreaResponse{}
	if err := json.Unmarshal(body, &lar); err != nil {
		return err
	}
	// fmt.Println("Name is: ", lar.Results[0].Name)

	config.Next = lar.Next
	config.Previous = lar.Previous
	
	PrintNames(lar.Results)

	return nil
}


func PrintNames(locationAreas []LocationArea) {
	for _, location := range locationAreas {
		fmt.Println(location.Name)
	}
}