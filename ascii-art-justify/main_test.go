package main

import (
	"os"
	"testing"
)

func TestGetArgs_Success(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	os.Args = []string{"cmd", "--align=center", "hello", "standard"}
	option, text, banner := getArgs()

	if option != "center" || text != "hello" || banner != "standard" {
		t.Errorf("Expected center, hello, standard; got %s, %s, %s", option, text, banner)
	}
}
