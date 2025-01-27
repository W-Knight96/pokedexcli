package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanned := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanned.Scan()
		line := scanned.Text()
		slice := strings.Fields(strings.ToLower(strings.TrimSpace(line)))
		if len(slice) == 0 {
			continue
		}
		fmt.Printf("Your command was: %s\n", slice[0])
	}
}

func cleanInput(text string) []string {
	cleanString := strings.ToLower(strings.TrimSpace(text))
	return strings.Fields(cleanString)
}
