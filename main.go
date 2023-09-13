package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var filename string
	var contents any
	var file *os.File
	var err error

	// Take in cli arg for file containing JSON data
	flag.StringVar(&filename, "f", "", "File containing JSON data")
	flag.Parse()

	// Read file and unmarshal contents into var type any
	file, err = os.Open(filename) // For read access.
	if err != nil {
		fmt.Println("error with opening file")
	}

	if err = file.Close(); err != nil {
		fmt.Println("error with closing file")
	}

}
