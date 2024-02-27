package main

import (
	"fmt"
	"os"
)

func commandExit(c *config) error {
	fmt.Println("Bye Bye!!!")
	os.Exit(0)

	return nil
}
