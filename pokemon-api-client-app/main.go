package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client *http.Client

type PokemonTypes struct {
	Data []string `json:"data`
}

func GetPokemonTypes() {
	url := "https://api.pokemontcg.io/v2/types"

	var pokemonTypes PokemonTypes

	err := GetJson(url, &pokemonTypes)
	if err != nil {
		fmt.Println("Error getting Pokemon Types - %s\n", err.Error)
	} else {
		fmt.Println("Pokemon Types - \n")
		for _, pType := range pokemonTypes.Data{
			fmt.Println(pType)
		}
	}
}

func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	GetPokemonTypes()
}