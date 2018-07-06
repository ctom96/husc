// HUSC - Hierarchical Universal Struct Core
//
// HUSC defines a data format, and the husc tool is used to
// convert from HUSC to JSON or JSON to HUSC.
package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {

	// Process command-line arguements with "flag" import
	var toJSON bool
	flag.BoolVar(&toJSON, "json", false, "Convert specified .husc files to JSON files")
	flag.BoolVar(&toJSON, "j", false, "Convert specified .husc files to JSON files (shorthand)")

	var toHUSC bool
	flag.BoolVar(&toHUSC, "husc", false, "Convert specified .husc files to HUSC files")
	flag.BoolVar(&toHUSC, "h", false, "Convert specified .husc files to HUSC files (shorthand)")

	flag.Parse()

	files := flag.Args()

	if toJSON {
		for _, file := range files {
			if !strings.HasSuffix(file, ".husc") {
				fmt.Println("Not every file given is of type .husc.\nExiting...")
				return
			}
		}
		fmt.Println("This is the point where we would convert your files to .json...")

	} else if toHUSC {
		for _, file := range files {
			if !strings.HasSuffix(file, ".json") {
				fmt.Println("Not every file given is of type .json.\nExiting...")
				return
			}
		}
		fmt.Println("This is the point where we would convert your files to .husc...")
	} else {
		fmt.Println("No -j or -h flag specified.")
		flag.Usage()
	}

}

// TODO: Write the library to convert from husc to json
// TODO: Write the library to convert from json to husc
