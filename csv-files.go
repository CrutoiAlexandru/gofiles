package gofiles

import (
	"log"

	"github.com/CrutoiAlexandru/gofiles/internal/csv"
)

func ReadCsvFile(filePath string) []string {
	var fileLines []string

	fileLines, err := csv.Read(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return fileLines
}

func WriteCsvFile(filePath string, content string) {
	err := csv.Write(filePath, content)
	if err != nil {
		log.Fatal(err)
	}
}

func AppendCsvFile(filePath string, content string) {
	err := csv.Append(filePath, content)
	if err != nil {
		log.Fatal(err)
	}
}
