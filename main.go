package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	cfg := &Config{} // Add this at the start
	scanned := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned.Scan()
		line := scanned.Text()
		slice := strings.Fields(strings.ToLower(strings.TrimSpace(line)))
		if len(slice) == 0 {
			continue
		}

		commandName := slice[0]
		args := []string{}
		if len(slice) > 1 {
			args = slice[1:]
		}

		command, exists := commands[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args ...string) error
}

type Config struct {
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Location `json:"results"`
}

type Location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func cleanInput(text string) []string {
	cleanString := strings.ToLower(strings.TrimSpace(text))
	return strings.Fields(cleanString)
}

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func callMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if cfg.Next != "" {
		url = cfg.Next
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return err
	}

	for _, location := range cfg.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func callMapb(cfg *Config) error {
	// Check if we're on first page
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	url := cfg.Previous

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &cfg)
	if err != nil {
		return err
	}

	for _, location := range cfg.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMap(cfg *Config, args ...string) error {
	return callMap(cfg)
}

func commandMapb(cfg *Config, args ...string) error {
	return callMapb(cfg)
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous page of locations",
			callback:    commandMapb,
		},
	}
}
