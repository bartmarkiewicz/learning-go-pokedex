package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"pokedex/internal/pokecache"
)

type LocationArea struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

func GetLocationAreas(next string, cache *pokecache.PokeCache) (LocationAreaResponse, error) {
	result, cacheHit := cache.Get(next)

	if !cacheHit {
		resp, err := http.Get(next)
		locationAreaResponse := LocationAreaResponse{}

		body, err := io.ReadAll(resp.Body)

		cache.Add(next, body)

		if err != nil {
			return locationAreaResponse, err
		}

		if resp.StatusCode != 200 {
			return locationAreaResponse, fmt.Errorf("Error: %s", body)
		}

		if err = json.Unmarshal(body, &locationAreaResponse); err != nil {
			return locationAreaResponse, err
		}
		err = resp.Body.Close()

		return locationAreaResponse, nil
	}

	locationAreaResponse := LocationAreaResponse{}

	err := json.Unmarshal(result, &locationAreaResponse)

	return locationAreaResponse, err
}

type LocationDetailsResponse struct {
	Id                int                         `json:"id"`
	Name              string                      `json:"name"`
	PokemonEncounters []PokemonEncountersResponse `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
}

type PokemonEncountersResponse struct {
	Pokemon Pokemon `json:"pokemon"`
}

func GetLocationDetails(location string) (LocationDetailsResponse, error) {

	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", location))
	locationDetailsResponse := LocationDetailsResponse{}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return locationDetailsResponse, err
	}

	if resp.StatusCode != 200 {
		return locationDetailsResponse, fmt.Errorf("Error: %s", body)
	}

	if err = json.Unmarshal(body, &locationDetailsResponse); err != nil {
		return locationDetailsResponse, err
	}
	err = resp.Body.Close()

	return locationDetailsResponse, nil
}

// Pokemon -
type PokemonDetailsResponse struct {
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	HeldItems      []interface{} `json:"held_items"`
	ID             int           `json:"id"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int    `json:"weight"`
	Name   string `json:"name"`
}

func GetPokemonDetails(pokemon string) (PokemonDetailsResponse, error) {

	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", pokemon))
	pokemonDetailsResponse := PokemonDetailsResponse{}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return pokemonDetailsResponse, err
	}

	if resp.StatusCode != 200 {
		return pokemonDetailsResponse, fmt.Errorf("Error: %s", body)
	}

	if err = json.Unmarshal(body, &pokemonDetailsResponse); err != nil {
		return pokemonDetailsResponse, err
	}
	err = resp.Body.Close()

	return pokemonDetailsResponse, nil
}
