package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	if val, ok := c.cache.Get(url); ok {
		locationsResp := LocationAreaResponse{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationResponse := LocationAreaResponse{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	c.cache.Add(url, data)
	return locationResponse, nil
}

func (c *Client) LocationDetails(location *string) (LocationDetailResponse, error) {
	
	LocationDetailResponse := LocationDetailResponse{}
	err := errors.New("")
	url := baseURL + "/location-area/" + *location

	if val, ok := c.cache.Get(url); ok {
		
		err := json.Unmarshal(val, &LocationDetailResponse)
		if err != nil {
			return LocationDetailResponse, err
		}

		return LocationDetailResponse, nil
	}


	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationDetailResponse, nil
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationDetailResponse, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationDetailResponse, err
	}

	err = json.Unmarshal(data, &LocationDetailResponse)
	if err != nil {
		return LocationDetailResponse, err
	}
	c.cache.Add(url, data)

	return LocationDetailResponse, err
}