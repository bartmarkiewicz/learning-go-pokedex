package commands

import (
	"fmt"
	"math/rand"
	"os"
	"pokedex/internal/pokecache"
	"pokedex/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *PokedexConfig, params []string) error
}

type PokedexConfig struct {
	Next             string
	Previous         string
	PokeDexCache     *pokecache.PokeCache
	PossibleCommands map[string]CliCommand
	UserPokedex      map[string]pokeapi.PokemonDetailsResponse
}

func CommandHelp(pokedexConfig *PokedexConfig, params []string) error {
	fmt.Println(
		`Welcome to the Pokedex!
Usage:	
`)
	for _, cmd := range pokedexConfig.PossibleCommands {
		fmt.Printf(" - %v: %v\n", cmd.Name, cmd.Description)
	}

	return nil
}

func CommandExit(pokedexConfig *PokedexConfig, params []string) error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return err
}

func CommandMap(pokedexConfig *PokedexConfig, params []string) error {
	locAreas, err := pokeapi.GetLocationAreas(pokedexConfig.Next, pokedexConfig.PokeDexCache)
	if err != nil {
		return err
	}
	pokedexConfig.Next = locAreas.Next
	pokedexConfig.Previous = locAreas.Previous
	for _, locArea := range locAreas.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}

func CommandMapb(pokedexConfig *PokedexConfig, params []string) error {
	locAreas, err := pokeapi.GetLocationAreas(pokedexConfig.Previous, pokedexConfig.PokeDexCache)
	if err != nil {
		return err
	}
	pokedexConfig.Next = locAreas.Next
	pokedexConfig.Previous = locAreas.Previous

	for _, locArea := range locAreas.Results {
		fmt.Println(locArea.Name)
	}
	return nil
}

func CommandExplore(pokedexConfig *PokedexConfig, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a location you want to explore")
		return nil
	}

	fmt.Printf("Exploring %v...\n", params[0])
	locationDetails, err := pokeapi.GetLocationDetails(params[0])

	if err != nil {
		return err
	}
	fmt.Printf("Found Pokemon:\n")

	for _, v := range locationDetails.PokemonEncounters {
		fmt.Printf(" - %v\n", v.Pokemon.Name)
	}
	return err
}

func CommandCatch(pokedexConfig *PokedexConfig, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a pokemon you want to catch")
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", params[0])
	pokemonDetails, err := pokeapi.GetPokemonDetails(params[0])

	if err != nil {
		return err
	}
	randomNumb := rand.Float64() - (float64(pokemonDetails.BaseExperience) / 1000)

	if randomNumb > 0.5 {
		fmt.Printf("%v was caught!\n", pokemonDetails.Name)
		pokedexConfig.UserPokedex[pokemonDetails.Name] = pokemonDetails
	} else {
		fmt.Printf("%v escaped!\n", pokemonDetails.Name)
	}

	return err
}

func CommandInspect(pokedexConfig *PokedexConfig, params []string) error {
	if len(params) == 0 {
		fmt.Println("Please provide a pokemon you want to inspect")
		return nil
	}

	pokemonDetails, found := pokedexConfig.UserPokedex[params[0]]

	if !found {
		fmt.Println("You don't have that Pokemon")
	}

	fmt.Printf("Name: %v\n", pokemonDetails.Name)
	fmt.Printf("Height: %v\n", pokemonDetails.Height)
	fmt.Printf("Weight: %v\n", pokemonDetails.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemonDetails.Stats {
		fmt.Printf("  -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, v := range pokemonDetails.Types {
		fmt.Printf("  -%v\n", v.Type.Name)
	}
	return nil
}

func CommandPokedex(pokedexConfig *PokedexConfig, params []string) error {
	fmt.Println("Your Pokedex:")

	for k, _ := range pokedexConfig.UserPokedex {
		fmt.Printf("  - %v\n", k)
	}

	return nil
}
