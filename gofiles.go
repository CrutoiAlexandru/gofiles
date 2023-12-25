package gofiles

import "github.com/CrutoiAlexandru/gofiles/internal/text"

func ReadTextFile(file_path string) ([]string, error) {
	var file_lines []string

	file_lines, err := text.Read(file_path)
	if err != nil {
		return file_lines, err
	}

	return file_lines, nil
}
