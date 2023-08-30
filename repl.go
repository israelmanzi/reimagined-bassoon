package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replStart() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("[>] ")

		scanner.Scan()

		text := parseInput(scanner.Text())

		if len(text) == 0 {
			continue
		}

		command := text[0]

		_comm, ok := gCommands()[command]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		_comm.callback()
	}
}

type command struct {
	name string
	description string
	callback func() error
}

func gCommands() map[string]command {
	commands := map[string]command{
		"help": {
			name: "help",
			description: "Show this help message",
			callback: help,
		},
		"exit": {
			name: "exit",
			description: "Exit the program",
			callback: exit,
		},
	}
	
	return commands
}

func parseInput(value string) []string {
	output := strings.Fields(strings.ToLower(value))

	return output
}
