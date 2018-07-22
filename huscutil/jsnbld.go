// Package huscutil -
// Provides functionality to convert a husc file to json
// or a json file to husc
package huscutil

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// Parse the file
	var rootHusc huscObject
	parentHusc := &rootHusc
	indentLevel := 0

	fmt.Println("Root's name:", rootHusc.name, "Type:", rootHusc.t, "Value:", rootHusc.value)

}

// Recursively parse each huscObject
func parseHuscObject(scanner *bufio.Scanner, level int) huscObject {

	var currObject huscObject // current husc object
	levelSpaces := level * 4

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":") {
			// we are at a new husc object (or array)

		}
	}
}

func countSpaces(line string) int {
	i := 0
	for _, runeval := range line {
		if runeval == ' ' {
			i++
		} else {
			break
		}
	}
	return i
}
