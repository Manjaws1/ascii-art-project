package alignment

import (
	"testing"
)

// Test func for Left, Right and Centre Alignment
func TestBasicAlignments(t *testing.T) {

	tests := []struct {
		name 		string
		fn 			func(string, int) string
		word		string
		width		int
		expected	string		
	} {
		// Left Alignment
		{"left one word", AlignLeft, "hello", 11, "hello      "},
		{"left two words", AlignLeft, "Hello There", 21, "Hello There          "},
		{"left three words", AlignLeft, "How are you", 21, "How are you          "},

		// Right Alignment
		{"right one word", AlignRight, "hello", 11, "      hello"},
		{"right two words", AlignRight, "Hello There", 21, "          Hello There"},
		{"right three words", AlignRight, "How are you", 21, "          How are you"},


		// Center Alignment
		{"center one word", AlignCenter, "hello", 11, "   hello   "},
		{"center two words", AlignCenter, "Hello There", 21, "     Hello There     "},
		{"center three words", AlignCenter, "How are you", 21, "     How are you     "},

		// Edge Cases
		{"empty left", AlignLeft, "", 5, "     "},
		{"empty right", AlignRight, "", 5, "     "},
		{"empty center", AlignCenter, "", 5, "     "},

		// Exact Fit
		{"left exact fit", AlignLeft, "hello", 5, "hello"},
		{"right exact fit", AlignRight, "hello", 5, "hello"},
		{"center exact fit", AlignCenter, "hello", 5, "hello"},

		// Width Smaller than Text
		{"left small width", AlignLeft, "hello world", 5, "hello world"},
		{"right small width", AlignRight, "hello world", 5, "hello world"},
		{"center small width", AlignCenter, "hello world", 5, "hello world"},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			result := tt.fn(tt.word, tt.width)

			if result != tt.expected {
				t.Errorf("Expected '%s', got '%s'", tt.expected, result)
			}

			if len(tt.expected) == tt.width && len(result) != tt.width {
				t.Errorf("Expected width %d, go %d", tt.width, len(result))
			}
		})
	}
} 



func TestAlignJustifyRow(t *testing.T) {
	tests := []struct {
		name 		string 
		words 		[]string
		width 		int
		expectLen 	int
		expect		string
	} {
		// Normal cases
		{"two words", []string{"hello", "world"}, 20, 20, ""},
		{"three words", []string{"how", "are", "you"}, 30, 30, ""},
		{"mixed case", []string{"Hello", "There"}, 25, 25, ""},

		// Edge cases
		{"single word", []string{"hello"}, 20, 5, "hello"},
		{"empty slice", []string{}, 20, 0, ""},
		{"exact fit", []string{"hello", "world"}, 11, 11, "hello world"},

		// Width smaller than content
		{"width too small", []string{"helloooo", "worlddddd"}, 10, 0, "helloooo worlddddd",},

		// Words with empty strings
		{"empty word inside", []string{"hello", "", "world"}, 20, 20, ""},

		// Multiple small words
		{"many words", []string{"a", "b", "c", "d"}, 20, 20, ""},
	}


	for _, tt := range tests {
	t.Run(tt.name, func(t *testing.T) {
		result := AlignJustifyRow(tt.words, tt.width)

		// Exact match if provided
		if tt.expect != "" {
			if result != tt.expect {
				t.Errorf("Expected '%s', got '%s'", tt.expect, result)
			}
			return
		}

		// Otherwise check length
		if len(result) != tt.expectLen {
			t.Errorf("Expected length %d, got %d | '%s'",
				tt.expectLen, len(result), result)
		}
	})
}

}