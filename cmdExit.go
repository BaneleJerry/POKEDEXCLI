package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Bye Bye!!!")
	os.Exit(0)

	return nil
}