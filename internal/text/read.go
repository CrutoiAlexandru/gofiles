package text

import (
	"bufio"
	"os"
)

func Read(file_path string) ([]string, error) {
	var file_lines []string

	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_lines = append(file_lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return file_lines, nil
}
