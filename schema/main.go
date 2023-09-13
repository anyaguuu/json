package main

import (
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

	// Take in cli arg for file containing JSON data
	flag.StringVar(&jsonDataFilename, "f", "", "File containing JSON data")
	flag.Parse()

	// Read file and unmarshal contents into var type any
	jsonFile, err = os.Open(jsonDataFilename) // For read access.
	if err != nil {
		fmt.Println("error with opening file")
	}
	jsonData, err = os.ReadFile(jsonDataFilename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %q", jsonData)

	if err = jsonFile.Close(); err != nil {
		fmt.Println("error with closing file")
	}

}
