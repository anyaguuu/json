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
	var schemaFilename string
	var jsonData []byte
	var jsonFile *os.File
	var err error
	var data any

	// Take in cli arg for file containing JSON data
	flag.StringVar(&jsonDataFilename, "j", "", "File containing JSON data")
	flag.StringVar(&schemaFilename, "s", "", "File containing JSON schema")
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

	// Unmarshal json data, apply visitor
	data = json.Unmarshal(jsonData, &data)

	// Call Accept(data, visitor) with visitor object

	// Close file
	if err = jsonFile.Close(); err != nil {
		fmt.Println("error with closing file")
	}

}
