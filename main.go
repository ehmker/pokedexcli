package main

import (
	"time"

	"github.com/ehmker/pokedexcli/internal/pokecache"
)

func main(){
	startREPL()
}

type cliCommand struct {
	Name string
	Description string
	Config *Config
	Callback func(*Config, string) error
}

type Config struct{
	Next string
	Previous string
	Cache *pokecache.PokeCache
}

func getCommands() map[string]cliCommand {
	cache := pokecache.NewCache(5 * time.Minute)
	config := Config{
		Cache: &cache,
	}
	
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map":{
			Name: "map",
			Description: "Displays the next 20 locations",
			Config: &config,
			Callback: commandMap,
		},
		"mapb": {
			Name: "mapb",
			Description: "Displays the previous 20 locations",
			Config: &config,
			Callback: commandMapBack,
		},
		"explore": {
			Name: "explore",
			Description: "Called as 'explore <location name/id>'.\n  Displays the pokemon that appear in area.",
			Config: &config,
			Callback: commandExplore,
		},
		"catch": {
			Name: "catch",
			Description: "Called as 'catch <pokemon name/id>'.\n  Attempts to catch the pokemon and add to the pokedex",
			Config: &config,
			Callback: commandCatch,
		},
	}
}