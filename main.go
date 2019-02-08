package main

import (
	"github.com/mvcatsifma/golang-wire/events"
)

func main() {
	e := events.InitializeEvent()

	e.Start()
}