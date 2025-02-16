package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func DeleteFile(fileName string) error {
	dir := "file_manager"
	filePath := filepath.Join(dir, fileName)
	if _, err := os.Stat(filePath); err != nil {
		return errors.New("No file found")
	}
	err := os.Remove(filePath)
	if err != nil {
		errors.New("Unable to delete file")
	}
	fmt.Println("File deleted successfully ✅")
	return nil
}
