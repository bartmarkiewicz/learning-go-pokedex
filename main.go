package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/commands"
	"pokedex/internal/pokecache"
	"pokedex/pokeapi"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	bufioScanner := bufio.NewScanner(reader)

	possibleCommands := map[string]commands.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 locations in the Pokemon world",
			Callback:    commands.CommandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 locations in the Pokemon world",
			Callback:    commands.CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the Pokemon world",
			Callback:    commands.CommandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a Pokemon",
			Callback:    commands.CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a Pokemon",
			Callback:    commands.CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Displays your Pokedex",
			Callback:    commands.CommandPokedex,
		},
	}

	pokedexConfig := commands.PokedexConfig{
		Next:             "https://pokeapi.co/api/v2/location-area/",
		Previous:         "https://pokeapi.co/api/v2/location-area/",
		PokeDexCache:     pokecache.NewCache(10 * time.Second),
		UserPokedex:      make(map[string]pokeapi.PokemonDetailsResponse),
		PossibleCommands: possibleCommands,
	}

	for true {
		fmt.Print("Pokedex > ")
		bufioScanner.Scan()
		userInput := bufioScanner.Text()
		cleanedInput := cleanInput(userInput)
		if cmd, success := possibleCommands[cleanedInput[0]]; success {
			extraArgs := cleanedInput[1:]

			err := cmd.Callback(&pokedexConfig, extraArgs)
			if err != nil {
				continue
			}
			continue
		}

		fmt.Println("Unknown command")
	}
}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{}
	}
	splitString := strings.Split(text, " ")
	returnedSlice := []string{}

	for _, word := range splitString {
		if strings.TrimSpace(word) != "" {
			returnedSlice = append(returnedSlice, strings.ToLower(strings.TrimSpace(word)))
		}
	}

	return returnedSlice
}
