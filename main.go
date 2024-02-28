package main

import (
	"time"

	"github.com/BaneleJerry/POKEDEXCLI/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocactinAreaURL *string
	prevLocactinAreaURL *string
}

func main() {

	cfg := config{
        pokeapiClient: pokeapi.NewClient(time.Minute),
    }
	startRepl(&cfg)
}
