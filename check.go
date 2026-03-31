package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func checkVenv() error {
	_, _ = fmt.Println("Checking for existing virtual environments...")

	err := filepath.WalkDir(".", findEnvFolder)
	if err != nil {
		return err
	}

	_, _ = fmt.Println("Check finished.")
	return nil
}

func findEnvFolder(path string, d fs.DirEntry, err error) error {
	if d.IsDir() {
		_, _ = fmt.Printf("Folder: %s\n", path)
	} else {
		_, _ = fmt.Printf("File: %s\n", path)
	}

	return nil
}
