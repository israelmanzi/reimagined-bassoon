package main

import (
	"fmt"
	"os"
)

func help() error {
	fmt.Println("Commands:")

	for _, command := range gCommands() {
		fmt.Printf(" =====> %s - %s\n", command.name, command.description)
	}

	return nil
}

func exit() error {
	os.Exit(0)

	return nil
}
