package gofiles

import (
	"log"

	"github.com/CrutoiAlexandru/gofiles/internal/text"
)

func ReadTextFile(file_path string) []string {
	var file_lines []string

	file_lines, err := text.Read(file_path)
	if err != nil {
		log.Fatal(err)
	}

	return file_lines
}
