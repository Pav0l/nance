package transform

import (
	"log"

	"github.com/Pav0l/nance/categorize"
	"github.com/Pav0l/nance/diacritics"
	"github.com/Pav0l/nance/lib/json"
)

type Transform struct {
	supportedHeaders map[string]string
	c                categorize.Categorize
	hadSpenderAlrdy  bool
}

func NewTransformer(headers string, categories categorize.Categorize) *Transform {
	supportedHeaders := json.ReadFile(headers)

	return &Transform{
		supportedHeaders: supportedHeaders,
		c:                categories,
		hadSpenderAlrdy:  false,
	}
}

func (t *Transform) RemoveUnnecessaryColumns(rows [][]string) [][]string {
	headers := rows[0]

	newHeaders, rowIndexesToKeep := t.rebuildHeaders(headers)
	rows[0] = newHeaders

	// start from index 1 to skip header row
	for i := 1; i < len(rows); i++ {
		rows[i] = t.rebuildRow(rows[i], rowIndexesToKeep)
	}

	return rows
}

func (t *Transform) AppendToRow(row []string, isHeader bool, spender string) []string {
	if isHeader {
		return t.appendToHeader(row)
	}

	// I don't like this - it infers header indexes to be specific value which we do not validate anywhere
	partner := row[2]
	category := row[4]

	categorized := t.c.Categorize(partner, category)

	if !t.hadSpenderAlrdy {
		row = append(row, spender)
	}
	return append(row, diacritics.Replace(categorized.Target), categorized.ReviewManually)
}

func (t *Transform) appendToHeader(header []string) []string {
	t.hadSpenderAlrdy = header[len(header)-1] == "Spender"

	if !t.hadSpenderAlrdy {
		header = append(header, "Spender")
	}

	header = append(header, "Category", "Review Manually")
	return header
}

func (t *Transform) rebuildRow(original []string, rowIndexesToKeep []int8) []string {
	newRow := []string{}

	for i := 0; i < len(rowIndexesToKeep); i++ {
		value := original[rowIndexesToKeep[i]]
		newRow = append(newRow, value)
	}

	return newRow
}

func (t *Transform) rebuildHeaders(headers []string) ([]string, []int8) {
	rowIndexesToKeep := []int8{}
	newHeaders := []string{}

	for i := 0; i < len(headers); i++ {
		header := headers[i]

		if t.supportedHeaders[header] != "" {
			newHeaders = append(newHeaders, t.supportedHeaders[header])
			log.Printf(" %s is renamed to %s", header, t.supportedHeaders[header])
			rowIndexesToKeep = append(rowIndexesToKeep, int8(i))
		}
	}

	return newHeaders, rowIndexesToKeep
}
