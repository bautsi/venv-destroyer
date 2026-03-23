package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func introduce() {
	splitter := "======================================================"

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintln(w, "Virtual Environment Destroyer")
	_, _ = fmt.Fprintln(w, "\tCLI tool for quick destroying virtual environments.")
	_, _ = fmt.Fprintln(w, splitter)
	_, _ = fmt.Fprintln(w, "Usage:")
	_, _ = fmt.Fprintln(w, "\t-check")
	_, _ = fmt.Fprintln(w, "\t\t\t\tCheck if any virtual environment exists.")
	_, _ = fmt.Fprintln(w, "\t-destroy")
	_, _ = fmt.Fprintln(w, "\t\t\t\tDestroy all virtual environments.")
	_, _ = fmt.Fprintln(w, "Debug:")
	_, _ = fmt.Fprintln(w, "\t-[any-flag] -debug")

	_ = w.Flush()
}
