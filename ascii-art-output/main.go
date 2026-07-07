// main.go
package main

import (
	"fmt"
	"os"
	"strings"
)

// BannerType represents available banner styles
type BannerType string

const (
	Standard   BannerType = "standard"
	Shadow     BannerType = "shadow"
	Thinkertoy BannerType = "thinkertoy"
)

func main() {
	config, err := parseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Load banner file
	bannerFile := fmt.Sprintf("banners/%s.txt", config.Banner)
	banner, err := loadBanner(bannerFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Build ASCII art
	asciiLines := buildASCII(banner, config.Text)

	// Output
	if config.OutputFile != "" {
		if err := writeToFile(config.OutputFile, asciiLines); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("ASCII art saved to %s\n", config.OutputFile)
	} else {
		// Print to stdout
		for _, line := range asciiLines {
			fmt.Println(line)
		}
	}
}

// Config holds parsed command line arguments
type Config struct {
	OutputFile string
	Text       string
	Banner     BannerType
}

// parseArgs parses command line arguments according to the project specification
func parseArgs(args []string) (Config, error) {
	config := Config{
		Banner: Standard, // default banner
	}

	if len(args) < 2 {
		return config, fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}

	// Case 1: Only string → go run . "hello"
	if len(args) == 2 {
		config.Text = args[1]
		return config, nil
	}

	// Case 2: With --output flag
	if strings.HasPrefix(args[1], "--output=") {
		outputPart := strings.TrimPrefix(args[1], "--output=")
		if outputPart == "" {
			return config, fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
		}
		config.OutputFile = outputPart

		// Remaining arguments: text and optional banner
		remaining := args[2:]
		if len(remaining) == 0 {
			return config, fmt.Errorf("missing string argument")
		}

		config.Text = remaining[0]

		if len(remaining) >= 2 {
			bannerName := remaining[1]
			if !isValidBanner(bannerName) {
				return config, fmt.Errorf("invalid banner: %s. Allowed: standard, shadow, thinkertoy", bannerName)
			}
			config.Banner = BannerType(bannerName)
		}
		return config, nil
	}

	// Case 3: No output flag, but with banner → go run . "hello" standard
	if len(args) == 3 {
		config.Text = args[1]
		bannerName := args[2]
		if !isValidBanner(bannerName) {
			return config, fmt.Errorf("invalid banner: %s. Allowed: standard, shadow, thinkertoy", bannerName)
		}
		config.Banner = BannerType(bannerName)
		return config, nil
	}

	return config, fmt.Errorf("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
}

// isValidBanner checks if banner name is supported
func isValidBanner(name string) bool {
	switch name {
	case "standard", "shadow", "thinkertoy":
		return true
	}
	return false
}

// loadBanner loads the banner file and returns map of characters to their ASCII lines
func loadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read banner file %s: %w", filename, err)
	}

	chars := make(map[rune][]string)
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")

	charCode := ' '
	lineIndex := 0

	for _, line := range lines {
		if lineIndex == 8 {
			lineIndex = 0
			charCode++
			continue
		}
		if lineIndex < 8 {
			chars[charCode] = append(chars[charCode], line)
		}
		lineIndex++
	}

	if len(chars) == 0 {
		return nil, fmt.Errorf("no characters loaded from banner")
	}

	return chars, nil
}

// buildASCII builds the ASCII art lines from the input text
func buildASCII(banner map[rune][]string, text string) []string {
	if text == "" {
		return []string{""}
	}

	inputLines := strings.Split(text, "\n")
	var result []string

	for i, line := range inputLines {
		if line == "" {
			result = append(result, "")
			continue
		}

		var outputLines [8]string

		for _, char := range line {
			charLines := getCharLines(banner, char)
			for j := 0; j < 8; j++ {
				outputLines[j] += charLines[j]
			}
		}

		for _, out := range outputLines {
			result = append(result, out)
		}

		// Add separator line between input lines
		if i < len(inputLines)-1 || strings.HasSuffix(text, "\n") {
			result = append(result, "")
		}
	}

	return result
}

// getCharLines returns 8 lines for a character (blank if not found)
func getCharLines(banner map[rune][]string, r rune) []string {
	if lines, ok := banner[r]; ok && len(lines) == 8 {
		return lines
	}
	return make([]string, 8)
}

// writeToFile writes ASCII art to the specified file
func writeToFile(filename string, lines []string) error {
	content := strings.Join(lines, "\n")
	if !strings.HasSuffix(content, "\n") {
		content += "\n"
	}
	return os.WriteFile(filename, []byte(content), 0644)
}