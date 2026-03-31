package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"text/tabwriter"
)

// Debugger for [checkVenv] function
//
// Calling [printAllPaths] as [fs.WalkDirFunc]
func checkVenvDebug(names []string) error {
	_, _ = fmt.Println("Checking for existing virtual environments...")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	err := filepath.WalkDir(".", printAllPathsWrapper(w, names))
	if err != nil {
		return err
	}

	_, _ = fmt.Println("Check finished.")
	return nil
}

// Print All Paths Function
//
// Print all folders' and files' paths from current directory.
//
// If path name in names, point out.
func printAllPaths(path string, d fs.DirEntry, err error, w *tabwriter.Writer, names []string) error {
	if err != nil {
		return fmt.Errorf("WalkDirFunc error: %v", err)
	}

	if slices.Contains(names, d.Name()) {
		_, _ = fmt.Fprintln(w, logSplitter)
		_, _ = fmt.Fprintln(w, "Destroy target found.")
		_, _ = fmt.Fprintf(w, "\tName: %v\n", d.Name())

		if d.IsDir() {
			_, _ = fmt.Fprintln(w, "\tType: Folder")
		} else {
			_, _ = fmt.Fprintln(w, "\tType: File")
		}

		_, _ = fmt.Fprintf(w, "\tPath: %v\n", path)
		_, _ = fmt.Fprintln(w, logSplitter)
	} else if d.IsDir() {
		_, _ = fmt.Fprintf(w, "%v\n", d.Name())
		_, _ = fmt.Fprintf(w, "\tFolder Path: %v\n", path)
	} else {
		_, _ = fmt.Fprintf(w, "%v\n", d.Name())
		_, _ = fmt.Fprintf(w, "\tFile Path: %v\n", path)
	}

	return nil
}

// [printAllPaths] Factory wrapper
//
// Pass names to mock [findEnvFolder].
func printAllPathsWrapper(w *tabwriter.Writer, names []string) fs.WalkDirFunc {
	return func(path string, d fs.DirEntry, err error) error {
		return printAllPaths(path, d, err, w, names)
	}
}
