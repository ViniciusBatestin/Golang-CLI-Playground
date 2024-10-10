package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type Pokemon struct {
	Name   string `json:"name"`
	Image  string `json:"image"`
	Weight int    `json:weight`
}

var SamplePokemon = &Pokemon{
	Name:   "Charizard",
	Image:  "https://www.pokemon.com/static-assets/content-assets/cms2/img/pokedex/full/006.png",
	Weight: 100,
}

func main() {

	// Download list of pokemons
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=100000&offset=0")
	if err != nil {
		panic(err)
	}

	var response map[string]any

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		panic(err)
	}
	res.Body.Close()

	pokemonUrls := []string{}
	results, ok := response["results"].([]any)
	if !ok {
		panic(fmt.Sprint(response))
	}

	for _, v := range results {
		result := v.(map[string]any)
		pokemonUrls = append(pokemonUrls, result["url"].(string))
	}

	if len(pokemonUrls) < 100 {
		panic("We are missing pokemons")
	}

	pokemons := []Pokemon{}

	// Downlaod 10 pokemons
	for _, url := range pokemonUrls[:10] {
		res, err = http.Get(url)
		if err != nil {
			panic(err)
		}
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			panic(err)
		}
		res.Body.Close()
		name := response["name"].(string)
		weight := int(response["weight"].(float64))
		sprints := response["sprites"].(map[string]any)
		image := sprints["front_default"].(string)

		pokemons = append(pokemons, Pokemon{Name: name, Weight: weight, Image: image})

	}

	http.HandleFunc("/pokemons", func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(pokemons)
		if err != nil {
			slog.Error(err.Error())
		}

	})
	fmt.Println("Starting the server")

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		panic(err)
	}

}

// curl http://localhost:9090
