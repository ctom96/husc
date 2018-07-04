// Package huscutil -
// Provides functionality to convert a husc file to json
// or a json file to husc
package huscutil

import (
	"fmt"
)

// ConvertToJSON converts slice of .json files to .husc files
func ConvertToJSON(files []string) []string {
	fmt.Println("ConvertToJSON called")

	var retFiles = files

	return retFiles
}
