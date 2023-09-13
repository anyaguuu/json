package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var jsonDataFilename string
	var contents []byte
	var reader *Reader
	var file *os.File
	var err error

	// Take in cli arg for file containing JSON data
	flag.StringVar(&jsonDataFilename, "f", "", "File containing JSON data")
	flag.Parse()

	// Read file and unmarshal contents into var type any
	file, err = os.Open(jsonDataFilename) // For read access.
	if err != nil {
		fmt.Println("error with opening file")
	}
	// contents =

	if err = file.Close(); err != nil {
		fmt.Println("error with closing file")
	}

}
