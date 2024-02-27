package main

import (
	"fmt"
	"log"

	"github.com/BaneleJerry/POKEDEXCLI/internal/pokeapi"
)

func commandMap() error {

	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.LocationAreasResp()

	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Location")

	for _, area := range resp.Results {
		fmt.Printf("-%s\n", area.Name)
	}

	return nil
}
