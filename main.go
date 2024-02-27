package main

import "github.com/BaneleJerry/POKEDEXCLI/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocactinAreaURL *string
	prevLocactinAreaURL *string
}

func main() {

	cfg := config{
        pokeapiClient: pokeapi.NewClient(),
    }
	startRepl(&cfg)
}
