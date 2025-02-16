package services

import (
	"fmt"
	"os"
	"path/filepath"
)

func tree(dir string, prefix string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", dir)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println(prefix + "📁 " + file.Name())
			tree(filepath.Join(dir, file.Name()), prefix+"  ")
		} else {
			fmt.Println(prefix + "📑 " + file.Name())
		}
	}
}

func ListFiles() error {
	dir := "file_manager"

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("directory does not exist: %s", dir)
	}
	tree(dir, "")

	return nil
}
