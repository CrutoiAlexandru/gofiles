package text

import "os"

func Remove(filePath string) error {
	var file *os.File
	var err error

	file, err = os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
