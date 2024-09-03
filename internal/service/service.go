package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/flrn000/pokedexcli/internal/cache"
)

const baseURL = "https://pokeapi.co/api/v2"

var responseCache = cache.NewCache(15 * time.Second)

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

	if data, exists := responseCache.Get(apiURL); exists {
		var result LocationArea
		err := json.Unmarshal(data, &result)
		if err != nil {
			return result, fmt.Errorf("error decoding cached data: %v", err)
		}

		return result, nil
	}

	res, err := http.Get(apiURL)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error fetching maps: %v", err)
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("error: Status %v", res.Status)
	}

	bodyCopy := bytes.Buffer{}
	if _, err := io.Copy(&bodyCopy, res.Body); err != nil {
		return LocationArea{}, err
	}
	responseCache.Add(apiURL, bodyCopy.Bytes())

	var data LocationArea
	if err := json.Unmarshal(bodyCopy.Bytes(), &data); err != nil {
		return data, fmt.Errorf("error decoding response body: %v", err)
	}

	return data, nil
}
