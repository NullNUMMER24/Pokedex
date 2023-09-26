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
	TYPE  []TYPE `json:"type"`
}

type TYPE struct {
	TYPE string
}

func main() {
	pokemons, err := fetchPokemonData()
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}

	// Ask the user what he wants to do
	// c := exec.Command("clear")
	// fmt.Printf("Select what you want to do:")
	// for {
	// 	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
	// 		if key.Code == keys.ArrowDown {
	// 			c.Stdout = os.Stdout
	// 			c.Run()

	// 		}
	// 		return false, nil
	// 	})
	// 	keyInfo, _ := keyboard.GetKey()
	// 	key := keyInfo.Code

	// 	if key == keys.ArrowDown {
	// 		c.Stdout = os.Stdout
	// 		c.Run()

	// 	}
	// 	fmt.Printf("test")
	// 	keyboard.StopListener()

	// 	if key == keys.ENTER {
	// 		break
	// 	}
	// }

	// Print the fetched data
	for _, pokemon := range pokemons {
		var elements string
		for _, type0 := range pokemon.TYPE {
			elements += type0.TYPE + ", "

		}

		elements = elements[:len(elements)-1]
		fmt.Printf("ID: %d\nName: %s\nHP: %s\nType: %s\n %s \n", pokemon.ID, pokemon.Name, pokemon.HP, elements, pokemon.ASCII)
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
