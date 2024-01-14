package csv

import (
	"bufio"
	"os"
)

func Read(filePath string) ([]string, error) {
	var fileLines []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return fileLines, nil
}
