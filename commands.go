package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func commandExit(conf *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *Config) error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range supportedCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(conf *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if conf.Next != "" {
		url = conf.Next
	} else if conf.Next == "" && conf.Previous != "" {
		fmt.Println("No more pages, returning to the first page.")
		conf.Next = ""
		conf.Previous = ""
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data := PokeAPILocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	return nil
}

func commandPreviousMap(conf *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if conf.Previous != "" {
		url = conf.Previous
	} else if conf.Previous == "" && conf.Next == "" {
		fmt.Println("No more pages, returning to the first page.")
		conf.Next = ""
		conf.Previous = ""
		return nil
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	data := PokeAPILocationAreaResponse{}
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return err
	}

	for _, location := range data.Results {
		fmt.Println(location.Name)
	}

	conf.Next = data.Next
	conf.Previous = data.Previous

	return nil
}