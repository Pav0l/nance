package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"

	"github.com/Pav0l/nance/categorize"
	"github.com/Pav0l/nance/diacritics"
	"github.com/Pav0l/nance/transform"
)

func main() {
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

	transformer := transform.NewTransformer("headers.json")
	rows = transformer.RemoveUnnecessaryColumns(rows)

	// Prepare Header row
	header := append(rows[0], "Category", "Review Manually")
	w.Write(header)
	defer w.Flush()

	// Iterate over every data row (excluding header) and categorize, sanitize it and write it to target file
	c := categorize.NewClassifier("categories.json")
	for i := 1; i < len(rows); i++ {
		row := rows[i]

		// I don't like this - it infers header indexes to be specific value which we do not validate anywhere
		partner := row[2]
		category := row[4]

		categorized := c.Categorize(partner, category)

		w.Write(append(row, diacritics.Replace(categorized.Target), categorized.ReviewManually))
	}
}
