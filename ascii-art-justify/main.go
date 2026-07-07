package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"acad.learn2earn.ng/git/oogbebor/ascii-art-justify/processor"
	"acad.learn2earn.ng/git/oogbebor/ascii-art-justify/template"
)

func getArgs() (option, textArgs, banner string) {

	if len(os.Args) != 4 {
		log.Fatal("Usage: go run . [OPTION] [STRING] [BANNER]")
	}

	options := []string{"--align=center", "--align=right", "--align=left", "--align=justify"}
	banners := []string{"standard", "thinkertoy", "shadow"}

	// Checking arguments meet criteria

	// Correct option check
	correctOption := false

	for _, opt := range options {
		if os.Args[1] == opt {
			correctOption = true
			break
		}
	}

	if !correctOption {
		log.Fatal("Options: --align=center, --align=right, --align=left, --align=justify")
	}

	// Correct banner check
	correctBanner := false

	for _, banner := range banners {
		if os.Args[3] == banner {
			correctBanner = true
			break
		}
	}

	if !correctBanner {
		log.Fatal("Banners: standard, thinkertoy, shadow")
	}

	option = os.Args[1][8:]
	textArgs = os.Args[2]
	banner = os.Args[3]

	return
}

func main() {

	// Get arguments
	option, textArgs, banner := getArgs()

	// Clean and split textargs
	textArgs = strings.ReplaceAll(textArgs, "\\n", "\n")
	text := strings.Split(textArgs, "\n")

	templateString := template.GetArt(banner)
	templateSlice := template.GetTemplateSlice(templateString)

	fmt.Println(processor.Processor(text, templateSlice, option))

}
