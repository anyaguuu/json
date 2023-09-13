package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"log/slog"
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
	var data []byte
	var d interface{}

	// Take in cli arg for file containing JSON data
	flag.StringVar(&jsonDataFilename, "j", "", "File containing JSON data")
	flag.StringVar(&jsonSchemaFilename, "s", "", "File containing JSON schema")
	flag.Parse()

	// Create JSON Schema compiler
	compiler = jsonschema.NewCompiler()
	// Compile the JSON schema
	schema, err = compiler.Compile(jsonSchemaFilename)
	if err != nil {
		slog.Error("schema compilation error", "error", err)
		return
	}

	// Open JSON file
	jsonFile, err = os.Open(jsonDataFilename) // For read access.
	if err != nil {
		fmt.Println("error with opening file")
	}

	// Read contents of JSON file
	jsonData, err = os.ReadFile(jsonDataFilename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %q", jsonData)

	// Unmarshal json data, apply visitor
	if err := json.Unmarshal(data, &d); err != nil {
		slog.Error("unable to unmarshal data", "data", data, "error", err)
		return
	}

	// Validate json against schema
	if err := schema.Validate(d); err != nil {
		msg := fmt.Sprintf("%#v", err)
		slog.Error("data does not conform to the schema", "error", msg)
	}
	fmt.Println("JSON data conforms to the schema")

	// Call Accept(data, visitor) with visitor object

	// Close file
	if err = jsonFile.Close(); err != nil {
		fmt.Println("error with closing file")
	}

}
