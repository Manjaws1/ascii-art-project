package template

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetArt(banner string) []string {
	filePath := fmt.Sprintf("%s.txt", banner)

	var content []byte

	// Check if file exists locally
	if _, err := os.Stat(filePath); err == nil {

		content, err = os.ReadFile(filePath)
		if err != nil {
			log.Fatal("Error reading file", err)
		}

	} else {

		URL := fmt.Sprintf("https://acad.learn2earn.ng/api/content/root/01-edu_module/content/ascii-art/%s.txt", banner)
		response, err := http.Get(URL)

		if err != nil {
			log.Fatal("Error fetching file", err)
		}

		defer response.Body.Close()

		if response.StatusCode > 299 {
			log.Fatalf("Error with status code %d: and body: %s", response.StatusCode, response.Body)
		}

		content, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Error reading body", err)
		}

		// Write to file
		err = os.WriteFile(filePath, content, 0444)
		if err != nil {
			log.Fatal(err)
		}
	}

	clean := strings.ReplaceAll(string(content), "\r\n", "\n")
	clean = strings.ReplaceAll(clean, "\r", "\n")

	return strings.Split(clean, "\n")
}

func GetTemplateSlice(lines []string) [][8]string {

	templateSlice := make([][8]string, 95)
	char := 0

	for i := 1; i+7 < len(lines); i += 9 {
		for r := 0; r < 8 && char < len(templateSlice); r++ {
			templateSlice[char][r] = lines[i+r]
		}
		char++
	}
	return templateSlice
}
