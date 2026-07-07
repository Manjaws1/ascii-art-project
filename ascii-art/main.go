package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	// Check arguments
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run . <text>")
	}

	// Get input
	textArgs := os.Args[1]

	// Replace escaped newline with actual newline
	textArgs = strings.ReplaceAll(textArgs, "\\n", "\n")

	if textArgs == "" {
		return
	}

	// Split into lines
	text := strings.Split(textArgs, "\n")

	// Call processor (no return value)
	asciiProcessor(text)
}
