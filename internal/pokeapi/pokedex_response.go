package pokeapi

type PokeAPILocationAreaResponse struct {
	Count    int                   `json:"count"`
	Next     string                `json:"next"`
	Previous string                `json:"previous"`
	Results  []PokeAPILocationArea `json:"results"`
}

type PokeAPILocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeAPILocationAreaDetail struct {
	ID		int                        `json:"id"`
	Name 	string                     `json:"name"`
	PokemonEncounters []PokeAPILocationAreaPokemon `json:"pokemon_encounters"`
}

type PokeAPILocationAreaPokemon struct {
	Pokemon PokeAPILocationAreaPokemonDetail `json:"pokemon"`
}

type PokeAPILocationAreaPokemonDetail struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonDetail struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int         `json:"base_stat"`
		Stat     PokemonStat `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type PokemonType `json:"type"`
	} `json:"types"`
	BaseExperience int `json:"base_experience"`
}

type PokemonStat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Name string `json:"name"`
}