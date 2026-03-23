package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func introduceDebug() {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	_, _ = fmt.Fprintln(w, "Please use debug flag with other flags.")
	_, _ = fmt.Fprintln(w, "Usage:")
	_, _ = fmt.Fprintln(w, "\t-check -debug")

	_ = w.Flush()
}
