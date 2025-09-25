package main

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