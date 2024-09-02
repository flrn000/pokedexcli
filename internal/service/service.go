package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const baseURL = "https://pokeapi.co/api/v2"

type Result struct {
	Name string
}

type LocationArea struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

func GetLocationAreaData(pageURL *string) (LocationArea, error) {
	apiURL := baseURL + "/location-area"
	if pageURL != nil {
		apiURL = *pageURL
	}

	res, err := http.Get(apiURL)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error fetching maps: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("error: Status %v", res.Status)
	}

	var data LocationArea
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("error decoding response body: %v", err)
	}

	return data, nil
}
