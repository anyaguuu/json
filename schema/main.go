package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

func main() {
	var jsonDataFilename string
	var jsonSchemaFilename string
	var compiler *jsonschema.Compiler
	var schema *jsonschema.Schema
	var jsonData []byte
	var jsonFile *os.File
	var err error
	var data any

	// Take in cli arg for file containing JSON data
	flag.StringVar(&jsonDataFilename, "j", "", "File containing JSON data")
	flag.StringVar(&jsonSchemaFilename, "s", "", "File containing JSON schema")
	flag.Parse()

	// Create JSON Schema compiler
	compiler = jsonschema.NewCompiler()
	// Compile the JSON schema
	schema, err = compiler.Compile(jsonSchemaFilename)

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
