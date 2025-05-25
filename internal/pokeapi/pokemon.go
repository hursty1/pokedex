package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)


func (c *Client) Pokemon(pokemon *string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + *pokemon

	if val, ok := c.cache.Get(url); ok {
		PokemonResponse := PokemonResponse{}
		err := json.Unmarshal(val, &PokemonResponse)
		if err != nil {
			return PokemonResponse, err
		}

		return PokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	PokemonResponse := PokemonResponse{}
	err = json.Unmarshal(data, &PokemonResponse)
	if err != nil {
		return PokemonResponse, err
	}
	c.cache.Add(url, data)
	return PokemonResponse, nil
} 