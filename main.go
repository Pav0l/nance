package main

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"

	"github.com/Pav0l/nance/categorize"
	"github.com/Pav0l/nance/diacritics"
)

func main() {
	// Read the source CSV file that we want to sanitize and categorize
	data, err := os.ReadFile("source.csv")
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bytes.NewReader(data))
	record, err := reader.ReadAll()
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

	// Prepare Header row
	targetHeader := append(record[0], "category", "reviewManually")
	w.Write(targetHeader)
	defer w.Flush()

	// Iterate over every data row (excluding header) and categorize, sanitize it and write it to target file
	rowCount := len(record)
	c := categorize.NewClassifier()
	for i := 1; i < rowCount; i++ {
		row := record[i]

		partner := row[2]
		category := row[3]

		categorized := c.Categorize(partner, category)

		// date, sum, partner, category(original), category(new), reviewManually
		w.Write([]string{row[0], row[1], diacritics.Replace(partner), category, diacritics.Replace(categorized.Target), categorized.ReviewManually})
	}
}
