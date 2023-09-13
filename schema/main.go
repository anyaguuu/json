package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var jsonDataFilename string
	var jsonData []byte
	var jsonFile *os.File
	var err error
	var contents any

	// Take in cli arg for file containing JSON data
	flag.StringVar(&jsonDataFilename, "f", "", "File containing JSON data")
	flag.Parse()

	// Read file contents
	jsonFile, err = os.Open(jsonDataFilename) // For read access.
	if err != nil {
		fmt.Println("error with opening file")
	}
	jsonData, err = os.ReadFile(jsonDataFilename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %q", jsonData)

	// Unmarshal contents, apply visitor
	contents = json.Unmarshal(jsonData, &contents)

	// Close file
	if err = jsonFile.Close(); err != nil {
		fmt.Println("error with closing file")
	}

}
