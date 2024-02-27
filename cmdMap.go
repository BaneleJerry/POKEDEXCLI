package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config) error {

	resp, err := c.pokeapiClient.LocationAreasResp(c.nextLocactinAreaURL)

	if err != nil {
		return err

	}

	fmt.Println("Location")

	for _, area := range resp.Results {
		fmt.Printf("-%s\n", area.Name)
	}
	c.nextLocactinAreaURL = resp.Next
	c.prevLocactinAreaURL = resp.Previous

	return nil
}

func commandBackMap(c *config) error {

	if c.prevLocactinAreaURL == nil {
		return errors.New("Your on the first page")
	}
	resp, err := c.pokeapiClient.LocationAreasResp(c.prevLocactinAreaURL)

	if err != nil {
		return err

	}

	fmt.Println("Location")

	for _, area := range resp.Results {
		fmt.Printf("-%s\n", area.Name)
	}

	c.nextLocactinAreaURL = resp.Next
	c.prevLocactinAreaURL = resp.Previous

	return nil
}
