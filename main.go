package main

import (
	"flag"
)

func main() {
	destroy := flag.Bool("destroy", false, "Destroy all virtual environments.")

	check := flag.Bool("check", false, "Check if any virtual environment exists.")

	flag.Parse()

	if *destroy {
		destroyVirtualEnvironment()
	}
	if *check {
		checkVirtualEnvironment()
	}
	if !*destroy && !*check {
		introduce()
	}
}
