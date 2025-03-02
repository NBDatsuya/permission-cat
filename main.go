package main

import (
	"log"
	"permission-cat/cmd"
)

func main() {
	err := cmd.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err.Error())
	}
}
