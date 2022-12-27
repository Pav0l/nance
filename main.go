package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/Pav0l/nance/categorize"
	"github.com/Pav0l/nance/transform"
)

func main() {

	spender := ""
	fmt.Print("Enter spender (P/L): ")
	_, err := fmt.Scanln(&spender)
	if err != nil {
		log.Fatal(err)
	}

	if spender != "P" && spender != "L" {
		log.Fatalf("Invalid spender. Expected: P/L. Received: %s", spender)
	}

	// Read the source CSV file that we want to sanitize and categorize
	data, err := os.ReadFile("source.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bytes.NewReader(data))
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Create the final file where we write sanitized and categorized data from source
	file, err := os.Create("target.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	transformer := transform.NewTransformer("headers.json", *categorize.NewClassifier("categories.json"))
	rows = transformer.RemoveUnnecessaryColumns(rows)

	// Iterate over every data row and categorize, sanitize it and write it to target file
	for i := 0; i < len(rows); i++ {
		w.Write(transformer.AppendToRow(rows[i], i == 0, spender))
	}
}
