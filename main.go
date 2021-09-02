package main

import (
	"log"

	"github.com/freshly/environator/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}

	// cmd.Execute()
}
