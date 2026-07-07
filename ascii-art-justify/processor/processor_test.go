package processor

import (
	"strings"
	"testing"
)

// // helper: create a fake templateSlice
func mockTemplateSlice() [][8]string {
	ts := make([][8]string, 95)

	// 	// fill each character with predictable values
	for i := 0; i < 95; i++ {
		for row := 0; row < 8; row++ {
			ts[i][row] = string(rune(i + 32)) // ASCII char
		}
	}
	return ts
}

func TestProcessor_AllEmpty(t *testing.T) {
	text := []string{"", "", ""}
	template := mockTemplateSlice()

	result := Processor(text, template, "")

	expectedLines := len(text) - 1
	actualLines := strings.Count(result, "\n")

	if actualLines != expectedLines {
		t.Errorf("Expected %d newlines, got %d", expectedLines, actualLines)
	}
}

func TestProcessor_NormalText(t *testing.T) {
	text := []string{"A"}
	template := mockTemplateSlice()

	result := Processor(text, template, "")

	lines := strings.Split(result, "\n")

	if len(lines) < 8 {
		t.Errorf("Expected at least 8 lines, got %d", len(lines))
	}
}

func TestBuildOutput_Default(t *testing.T) {
	template := mockTemplateSlice()

	result := buildOutput(template, "AB", "")

	lines := strings.Split(result, "\n")

	if len(lines) < 8 {
		t.Errorf("Expected 8 lines, got %d", len(lines))
	}
}

func TestBuildOutput_WithSpaces(t *testing.T) {
	template := mockTemplateSlice()

	result := buildOutput(template, "A B", "")

	if !strings.Contains(result, " ") {
		t.Errorf("Expected spaces between words")
	}
}
