package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func checkVenvDebug() error {
	_, _ = fmt.Println("Checking for existing virtual environments...")
	// Add logic to check for existing virtual environments here

	err := filepath.WalkDir(".", printAllPaths)
	if err != nil {
		return err
	}

	_, _ = fmt.Println("Check finished.")
	return nil
}

func printAllPaths(path string, d fs.DirEntry, err error) error {
	if d.IsDir() {
		_, _ = fmt.Printf("Folder: %s\n", path)
	} else {
		_, _ = fmt.Printf("File: %s\n", path)
	}

	return nil
}
