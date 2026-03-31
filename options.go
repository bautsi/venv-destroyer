package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Load Venv Names Function
//
// 1. Load all option names and related names from options.json.
//
// 2. Load default venv names.
//
// 3. Use param passed keys as additional option keys to load additional names.
func loadVenvNames(addtKeys []string) ([]string, error) {
	data, err := os.ReadFile("options.json")
	if err != nil {
		fmt.Println("Options json file load error: ", err)
	}

	var options map[string][]string
	err = json.Unmarshal(data, &options)
	if err != nil {
		return nil, fmt.Errorf("Options data unmarshal error: %v", err)
	}

	var optionVals []string
	optionVals = append(optionVals, options["venv"]...)

	for _, key := range addtKeys {
		if names, ok := options[key]; !ok {
			return nil, fmt.Errorf("Key: '%v' not in options.json", key)
		} else {
			optionVals = append(optionVals, names...)
		}
	}

	return optionVals, nil
}
