package alignment

import (
	"fmt"
	"strings"
)

func AlignCenter(text string, width int) string {
	padding := width - len(text)

	if padding <= 0 {
		return text
	}

	left := padding / 2
	right := padding - left

	return fmt.Sprintf("%*s%s%*s", left, "", text, right, "")
}

// Func to align a row left
func AlignLeft(text string, width int) string {
	return fmt.Sprintf("%-*s", width, text)
}

// Function to align a row right
func AlignRight(text string, width int) string {
	return fmt.Sprintf("%*s", width, text)
}

func AlignJustifyRow(words []string, width int) string {
	if len(words) == 0 {
		return ""
	}

	if len(words) == 1 {
		return words[0]
	}

	totalChars := 0
	for _, w := range words {
		totalChars += len(w)
	}

	totalSpaces := width - totalChars
	gaps := len(words) - 1

	if totalSpaces <= 0 {
		return strings.Join(words, " ")
	}

	spacePerGap := totalSpaces / gaps
	extra := totalSpaces % gaps

	var builder strings.Builder
	for i, w := range words {
		builder.WriteString(w)

		if i < gaps {
			spaces := spacePerGap
			if extra > 0 {
				spaces++
				extra--
			}
			
			builder.WriteString(strings.Repeat(" ", spaces))
		}
	}

	return builder.String()
}
