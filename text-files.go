package main

import (
	"log"

	"github.com/CrutoiAlexandru/gofiles/internal/text"
)

func ReadTextFile(filePath string) []string {
	var fileLines []string

	fileLines, err := text.Read(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return fileLines
}

func WriteTextFile(filePath string, content string) {
	err := text.Write(filePath, content)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendTextFile(filePath string, content string) {
	err := text.Append(filePath, content)
	if err != nil {
		log.Fatal(err)
	}
}

func RemoveTextFile(filePath string) {
	err := text.Remove(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filePath := "tst"
	RemoveTextFile(filePath)
}
