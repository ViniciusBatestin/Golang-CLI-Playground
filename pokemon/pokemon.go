package main

import(
	"fmt"
	"net/http"
)

type response struct {
	Name    string    `jason:name`
	Pokemon []Pokemon `json:"pokemon_entries`
}

type Pokemon struct {
	EntryNo int            `json:"entry_number"`
	Species PokemonSpecies `json:"pokemon_species"`
}

func main() {
	response, err := http.Get(https://pokeapi.co/api/v2/ability/3/)
}
