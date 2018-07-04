// HUSC - Hierarchical Universal Struct Core
//
// HUSC defines a data format, and the husc tool is used to
// convert from HUSC to JSON or JSON to HUSC.

package main

import (
	"fmt"
	"husc/huscutil"
	"os"
	"strings"
)

func printUsage() {
	fmt.Println("ERROR: Incorrect Usage!")
	fmt.Println("Usage: husc [options...] [files...]")
}

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		printUsage()
		return
	}

	// TODO Add verbose mode
	//verbose := false
	var files []string

	// Process args
	switch strings.ToLower(args[0]) {
	case "-j", "-json":
		// Convert list of files into json formatted files
		files := append(files, args[1:]...)

		fmt.Println(files)
		fmt.Println(huscutil.ConvertToJSON(files))

	case "-h", "-husc":
		// Convert json to husc formatted files of the same name
		files := append(files, args[1:]...)

		fmt.Println(files)
	}
}
