package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/mattn/go-tty"
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
	options := [4]string{"Select Pokemon", "Show all Pokemon", "Show Pokemons from an Element", "Quit"}
	count := 1
	c := exec.Command("clear")
	fmt.Println("Select what you want to do:")
	for _, option := range options {
		fmt.Printf("[%d] - %s\n", count, option)
		count++
	}

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	r, err := tty.ReadRune()

	for {

		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(r)
		if string(r) != "" {
			break
		}

	}
	if string(r) == "1" {
		c.Stdout = os.Stdout
		c.Run()
		println("You choose Option " + string(r))
		println("Which Pokemon do you want to select?")

		for _, pokemon := range pokemons {
			reader := bufio.NewReader(os.Stdin)
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
			var elements string
			for _, type0 := range pokemon.TYPE {
				elements += type0.TYPE + ", "
			}
			if string(line) == pokemon.Name {
				fmt.Printf("ID: %d\nName: %s\nHP: %s\nType: %s\n %s \n", pokemon.ID, pokemon.Name, pokemon.HP, elements, pokemon.ASCII)
			}
		}
	}
	if string(r) == "2" {
		c.Stdout = os.Stdout
		c.Run()
		println("You choose Option " + string(r))
		println("Here are all Pokemons")
		// Print all Pokemon
		for _, pokemon := range pokemons {
			var elements string
			for _, type0 := range pokemon.TYPE {
				elements += type0.TYPE + ", "

			}

			elements = elements[:len(elements)-1]
			fmt.Printf("ID: %d\nName: %s\nHP: %s\nType: %s\n %s \n", pokemon.ID, pokemon.Name, pokemon.HP, elements, pokemon.ASCII)
		}
	}

	if string(r) == "3" {
		c.Stdout = os.Stdout
		c.Run()
		println("You choose Option " + string(r))
		println("From which element do you want to see all Pokemon")

	}

	if string(r) == "4" {
		c.Stdout = os.Stdout
		c.Run()
		println("You choose Option " + string(r))
		println("Good Bye!")

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
