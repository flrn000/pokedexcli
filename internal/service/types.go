package service

type Result struct {
	Name string `json:"name"`
}

type LocationArea struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type PokemonEncounter struct {
	Pokemon struct {
		Name string `json:"name"`
	} `json:"pokemon"`
}

type Encounters struct {
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}
