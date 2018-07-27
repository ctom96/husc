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
	rootHusc := parseHuscObject(scanner, 0)

	fmt.Println(rootHusc.toString())

}

func parseHuscArray(scanner *bufio.Scanner, level int) (huscArray, int) {
	var retArray huscArray
	tmpScanner := &(*scanner)

	numScans := 0

	for tmpScanner.Scan() {
		line := tmpScanner.Text()
		elems := strings.Fields(line)

		if countSpaces(line) == level*4 {
			// we are at the first line of this object, get the name
			retArray.name = elems[1][:len(elems[1])-1]
		}

		// add everything else as if it was a huscObject
		if countSpaces(line) == (level+1)*4 {
			retArray.values = append(retArray.values, parseHuscObject(tmpScanner, level+1))
		}

		numScans++
	}

	return retArray, numScans - 1
}

func parseHuscObject(scanner *bufio.Scanner, level int) huscObject {
	var retObject huscObject

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		// get name, when there
		if countSpaces(line) == level*4 {
			// we are at the first line of this object, get the name
			elems := strings.Fields(line)
			if len(elems) == 2 {
				retObject.name = elems[1][:len(elems[1])-1]
			} else {
				// len(elems) == 1, so no o
				retObject.name = elems[0][:len(elems[0])-1]
			}
			fmt.Println("retObject name =", retObject.name)
		}

		// Loot at the type of the next item to identify what to do, only when
		// indentation is 4 + level*4
		if countSpaces(line) == (level+1)*4 {
			elems := strings.Fields(line)
			if len(elems) == 2 {
				// check for array, object, or string
				switch elems[0] {
				case "a":
					// parse array
					ary, num := parseHuscArray(scanner, level+1)
					retObject.values = append(retObject.values, ary)
					fmt.Println(num)
					// advance this scanner the appriate amount
					// for i := 0; i < num; i++ {
					// 	scanner.Scan()
					// }
				case "o":
					// parse object
					retObject.values = append(retObject.values, parseHuscObject(scanner, level+1))
				default:
					// default to string type
					var stringHusc huscSingle
					stringHusc.dType = s
					stringHusc.name = elems[0]
					startValue := countSpaces(line) + len(elems[0]) + 1
					stringHusc.value = line[startValue:]
					retObject.values = append(retObject.values, stringHusc)
				}
			} else {
				// determine type
				switch elems[0] {
				case "s": // string type
					var stringHusc huscSingle
					stringHusc.dType = s
					stringHusc.name = elems[1]
					startValue := countSpaces(line) + len(elems[0]) + len(elems[1]) + 2
					stringHusc.value = line[startValue:]
					retObject.values = append(retObject.values, stringHusc)
				case "n": // number type
					var numberHusc huscSingle
					numberHusc.dType = n
					numberHusc.name = elems[1]
					numberHusc.value = elems[2]
					retObject.values = append(retObject.values, numberHusc)
				case "b": // bool type
					var boolHusc huscSingle
					boolHusc.dType = b
					boolHusc.name = elems[1]
					boolHusc.value = elems[2]
					retObject.values = append(retObject.values, boolHusc)
				case "N": // null type
					var nullHusc huscSingle
					nullHusc.dType = N
					nullHusc.name = elems[1]
					nullHusc.value = "null"
					retObject.values = append(retObject.values, nullHusc)
				}
			}
		}

	}

	return retObject
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
