package main

import (
	"fmt"
	"github.com/mvcatsifma/golang-wire/events"
	"os"
)

func main() {
	e, err := events.InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}

	e.Start()
}