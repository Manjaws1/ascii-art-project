package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func asciiProcessor(input []string) {

	filePath := "standard.txt"
	link := "https://acad.learn2earn.ng/api/content/root/01-edu_module/content/ascii-art/standard.txt"

	// Step 1. Get the arguments
	/* if len(os.Args) != 2 {
		log.Fatal("Usage: go run . <text>")
	}
	textArgs := os.Args[1]

	textArgs = strings.ReplaceAll(textArgs, "\\n", "\n")

	if textArgs == "" {
		return
	}
	*/

	// Step 2. Get Template
	var template string

	// Check if template file exists locally
	if _, err := os.Stat(filePath); err == nil {
		data, err := os.ReadFile(filePath)

		if err != nil {
			log.Fatal("Error reading file", err)
		}

		template = string(data)
	} else {
		response, err := http.Get(link)
		if err != nil {
			log.Fatal(err)
		}
		data, err := io.ReadAll(response.Body)
		response.Body.Close()

		if response.StatusCode > 299 {
			log.Fatalf("Response failed with status code: %d and\nbody: %s\n", response.StatusCode, data)
		}

		if err != nil {
			log.Fatal(err)
		}

		err = os.WriteFile(filePath, data, 0444)

		if err != nil {
			log.Fatal(err)
		}

		template = string(data)
	}

	// Step 3. Converting to template Slice for Easy Retrieval
	lines := strings.Split(template, "\n")

	var templateSlice = make([][8]string, 95)

	char := 0

	for i := 1; i+7 < len(lines); i += 9 {
		for r := range 8 {
			templateSlice[char][r] = lines[i+r]
		}
		char++
	}

	// Step 4. Looping through the text argument

	// Checking if all are empty words
	allEmpty := true
	for _, word := range input {
		if word != "" {
			allEmpty = false
			break
		}
	}

	// Print corresponding number of newline characters
	if allEmpty && len(input) > 0 {
		for i := 0; i < len(input)-1; i++ {
			fmt.Println()
		}
		return
	}

	// Printing for non empty words
	for _, word := range input {
		if word == "" {
			fmt.Println()
			continue
		}
		for i := range 8 {
			for _, r := range word {
				if r >= 32 && r <= 126 {
					line := templateSlice[r-32][i]
					fmt.Print(line)
				}
			}
			fmt.Println()

		}
	}
}
