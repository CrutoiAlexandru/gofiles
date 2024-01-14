package csv

import (
	"bufio"
	"os"
)

func Write(filePath string, content string) error {
	var file *os.File
	var writer *bufio.Writer
	var err error

	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer = bufio.NewWriter(file)
	_, err = writer.Write([]byte(content))
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
