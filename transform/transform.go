package transform

import (
	"log"

	"github.com/Pav0l/nance/lib/json"
)

type Transform struct {
	supportedHeaders map[string]string
}

func NewTransformer(headersFileName string) *Transform {
	supportedHeaders := json.ReadFile(headersFileName)

	return &Transform{
		supportedHeaders: supportedHeaders,
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
