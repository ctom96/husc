// Package huscutil -
// Provides functionality to convert a husc file to json
// or a json file to husc
package huscutil

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) bool {
	if e != nil {
		return false
	}
	return true
}

// ConvertToJSON converts slice of .husc files to .json files
func ConvertToJSON(files []string) {
	// Loop through the list of files, and generate a .husc for each
	for _, file := range files {
		parseFile(file)
	}

}

func parseFile(filePath string) {
	// attempt to open file
	file, err := os.Open(filePath)
	if !check(err) {
		// Will exit function when opening file fails
		fmt.Printf("Coould not open file %#v for parsing\n", filePath)
		return
	}
	defer file.Close()

	// create scanner to scan file
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
