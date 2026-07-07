package template

import (
	"os"
	"reflect"
	"testing"
)

func TestGetArt(t *testing.T) {
	err := os.WriteFile("testbanner.txt", []byte("A\nB\nC"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	result := GetArt("testbanner")

	expected := []string{"A", "B", "C"}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
	defer os.Remove("testbanner.txt")
}

func TestGetTemplateSlice(t *testing.T) {
	lines := []string{
		"skip",
		"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8", " ",
		"b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8", " ",
		"extra",
	}

	expected := make([][8]string, 2)
	expected[0] = [8]string{"a1", "a2", "a3", "a4", "a5", "a6", "a7", "a8"}
	expected[1] = [8]string{"b1", "b2", "b3", "b4", "b5", "b6", "b7", "b8"}

	result := GetTemplateSlice(lines)

	if !reflect.DeepEqual(result[:2], expected) {
		t.Errorf("Expected %v, got %v", expected, result[:2])
	}
}
