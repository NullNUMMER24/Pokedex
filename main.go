package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Pokemon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	HP    string `json:"hp"`
	ASCII string `json:"ASCII"`
}

func main() {
	pokemons, err := fetchPokemonData()
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	// Print the fetched data
	for _, pokemon := range pokemons {
		fmt.Printf("ID: %d\nName: %s\nHP: %s\n %s \n", pokemon.ID, pokemon.Name, pokemon.HP, pokemon.ASCII)
	}
}

func fetchPokemonData() ([]Pokemon, error) {
	// Make an HTTP GET request to the API
	resp, err := http.Get("http://localhost:3000/Pokemon")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Create a slice of Pokemon to hold the parsed data
	var pokemons []Pokemon

	// Unmarshal the JSON data into the pokemons slice
	err = json.Unmarshal(body, &pokemons)
	if err != nil {
		return nil, err
	}

	return pokemons, nil
}
