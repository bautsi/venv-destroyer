package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Global Variables
var logSplitter = "======================================================"

func main() {
	// Mode Switch
	debugFlag := flag.Bool("debug", false, "Debug mode, use with other flags.")

	// Options
	allOptionsFlag := flag.Bool("all", false, "Activate all types of files, use with other flags.")

	withEnvFlag := flag.Bool("env", false, "Also handle env files, use with other flags.")
	withCacheFlag := flag.Bool("cache", false, "Also handle cache files, use with other flags.")

	// Actions
	destroyFlag := flag.Bool("destroy", false, "Destroy all virtual environments.")
	checkFlag := flag.Bool("check", false, "Check if any virtual environment exists.")

	flag.Parse()

	// Option Flags Handle
	var addtKeys []string
	if *allOptionsFlag {
		*withEnvFlag = true
		*withCacheFlag = true
	}

	if *withEnvFlag {
		addtKeys = append(addtKeys, "env")
	}
	if *withCacheFlag {
		addtKeys = append(addtKeys, "cache")
	}

	// Load Options
	names, err := loadVenvNames(addtKeys)
	if err != nil {
		os.Exit(2)
	}
	fmt.Println(strings.Join(names, ", "))

	// Debug Mode
	if *debugFlag {
		if *destroyFlag {
			destroyVenvDebug()
			return
		}
		if *checkFlag {
			err := checkVenvDebug(names)
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
