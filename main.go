package main

import (
	"flag"
	"fmt"
)

func main() {
	debugFlag := flag.Bool("debug", false, "Debug mode, .")

	destroyFlag := flag.Bool("destroy", false, "Destroy all virtual environments.")
	checkFlag := flag.Bool("check", false, "Check if any virtual environment exists.")

	flag.Parse()

	// Debug Mode
	if *debugFlag {
		if *destroyFlag {
			destroyVenvDebug()
			return
		}
		if *checkFlag {
			err := checkVenvDebug()
			if err != nil {
				fmt.Println("Check Virtual Environment Debug Error: ", err)
			}
			return
		}
		introduceDebug()
		return
	}

	// Default Mode
	if *destroyFlag {
		destroyVenv()
		return
	}
	if *checkFlag {
		checkVenv()
		return
	}
	introduce()
}
