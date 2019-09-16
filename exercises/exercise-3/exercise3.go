// CST8333 Exercise 3 - Lucas Estienne

package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"encoding/csv"
)

// simple data structure containing a string
type Record struct {
	content string
}

// main function, this is the entrypoint
func main() {

	// declare variables
	var records []Record
	numRecords := 5 // number of records to get
	

	// Load lines from CSV
	lines, err := getLinesFromCSV("data/canadianCheeseDirectory.csv")
	check(err)

	// Get subslice of lines slice, not including column names
	for i := 0; i < numRecords; i++ {
		// append a Record object to Record slice
		records = append(records, Record { strings.Join(lines[i][:], ",") })
	}

	// loop through records object
	for _, v := range(records) {
		fmt.Println(v.content) // print record content
	}
	fmt.Println("Lucas Estienne")
}

// helper function to do error handling
func check(e error) {
    if e != nil {
		log.Fatal("Error", e)
        panic(e)
    }
}

// helper function to read CSV
func getLinesFromCSV(filePath string) (lines [][]string, err error) {
	// open file
	file, err := os.Open(filePath)
	check(err)
	defer file.Close() // defer closing the file until function returns

	// create CSV Reader from file
	reader := csv.NewReader(file)
	return reader.ReadAll()
}