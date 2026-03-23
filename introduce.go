package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func introduce() {
	splitter := "======================================================"

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	fmt.Fprintln(w, "Virtual Environment Destroyer")
	fmt.Fprintln(w, "\tCLI tool for quick destroying virtual environments.")
	fmt.Fprintln(w, splitter)
	fmt.Fprintln(w, "Usage:")
	fmt.Fprintln(w, "\t-check")
	fmt.Fprintln(w, "\t\t\t\tCheck if any virtual environment exists.")
	fmt.Fprintln(w, "\t-destroy")
	fmt.Fprintln(w, "\t\t\t\tDestroy all virtual environments.")

	w.Flush()
}
