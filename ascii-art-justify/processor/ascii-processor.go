package processor

import (
	"fmt"
	"strings"
	"os/exec"
	"os"
	"strconv"

	align "acad.learn2earn.ng/git/oogbebor/ascii-art-justify/alignment"
)

func getTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	s, _ := cmd.Output()

	width, err := strconv.Atoi(strings.Fields(string(s))[1])

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return width
}


func Processor(text []string, templateSlice [][8]string, option string) string {
	var builder strings.Builder

	// check for when all words are empty
	allEmpty := true
	for _, word := range text {
		if word != "" {
			allEmpty = false
			break
		}
	}

	// Print newline for every empty word
	if allEmpty {
		for i := 0; i < len(text)-1; i++ {
			builder.WriteString(fmt.Sprintln())
			continue
		}
		return builder.String()
	}

	// Build art for when text isn't all empty
	for _, word := range text {
		if word == "" {
			builder.WriteString(fmt.Sprintln())
			continue
		}
		builder.WriteString(buildOutput(templateSlice, word, option))
	}

	return builder.String()
}

func buildOutput(templateSlice [][8]string, text, option string) string {

	width := getTerminalWidth()
	words := strings.Fields(text)

	var builder strings.Builder

	for row := 0; row < 8; row++ {
		rowPieces := make([]string, len(words))
		for w, word := range words {
			for _, r := range word {
				if r >= 32 && r <= 126 {
					rowPieces[w] += templateSlice[r-32][row]
				}
			}
		}

		line := ""
		switch option {
		case "right":
			// join with single spaces
			line = strings.Join(rowPieces, " ")
			line = align.AlignRight(line, width)
		case "center":
			line = strings.Join(rowPieces, " ")
			line = align.AlignCenter(line, width)
		case "justify":
			// justify by spreading spaces between word blocks
			line = align.AlignJustifyRow(rowPieces, width)
		default:
			line = strings.Join(rowPieces, " ")
		}

		builder.WriteString(line + "\n")
	}

	return builder.String()
}
