package transform

import (
	"log"

	"github.com/Pav0l/nance/categorize"
	"github.com/Pav0l/nance/diacritics"
	"github.com/Pav0l/nance/lib/json"
)

type Transform struct {
	c                categorize.Categorize
	supportedHeaders map[string]string
	// set to `true` if the original CSV has "Spender" column
	hadSpenderAlrdy bool
	// header row indexes we want to keep
	originalRowIndexesToKeep []int8
	// headerIndex stores a map of header name and its index after rebuilding the original header
	headerIndex headerIndex
}

type headerIndex map[string]int8

const SUM_HEADER = "Sum"
const PARTNER_HEADER = "Partner"
const OG_CATEGORY_HEADER = "Original Category"
const SPENDER_HEADER = "Spender"
const CATEGORY_HEADER = "Category"
const REVIEW_MANUALLY_HEADER = "Review Manually"

func NewTransformer(headers string, categories categorize.Categorize) *Transform {
	supportedHeaders := json.ReadFile(headers)

	return &Transform{
		supportedHeaders: supportedHeaders,
		c:                categories,
		headerIndex:      headerIndex{},
		hadSpenderAlrdy:  false,
	}
}

func (t *Transform) RemoveUnnecessaryColumns(rows [][]string) [][]string {
	headers := rows[0]

	newHeaders := t.rebuildHeaders(headers)
	rows[0] = newHeaders

	// start from index 1 to skip header row
	for i := 1; i < len(rows); i++ {
		rows[i] = t.rebuildRow(rows[i])
	}

	return rows
}

func (t *Transform) AppendToRow(row []string, isHeader bool, spender string) []string {
	if isHeader {
		return t.appendToHeader(row)
	}
	// remove zero sum rows
	sum := row[t.headerIndex[SUM_HEADER]]
	if sum == "0,00" {
		return []string{}
	}

	partner := row[t.headerIndex[PARTNER_HEADER]]
	category := row[t.headerIndex[OG_CATEGORY_HEADER]]
	categorized := t.c.Categorize(partner, category)

	if !t.hadSpenderAlrdy {
		row = append(row, spender)
	}
	return append(row, diacritics.Replace(categorized.Target), categorized.ReviewManually)
}

func (t *Transform) appendToHeader(header []string) []string {
	t.hadSpenderAlrdy = header[len(header)-1] == SPENDER_HEADER

	if !t.hadSpenderAlrdy {
		header = append(header, SPENDER_HEADER)
	}

	header = append(header, CATEGORY_HEADER, REVIEW_MANUALLY_HEADER)
	return header
}

func (t *Transform) rebuildRow(original []string) []string {
	newRow := []string{}

	for i := 0; i < len(t.originalRowIndexesToKeep); i++ {
		value := original[t.originalRowIndexesToKeep[i]]
		newRow = append(newRow, value)
	}

	return newRow
}

func (t *Transform) rebuildHeaders(headers []string) []string {
	newHeaders := []string{}

	for i := 0; i < len(headers); i++ {
		header := headers[i]

		if t.supportedHeaders[header] != "" {
			newHeaders = append(newHeaders, t.supportedHeaders[header])
			log.Printf("%s is renamed to %s", header, t.supportedHeaders[header])
			t.originalRowIndexesToKeep = append(t.originalRowIndexesToKeep, int8(i))
			// new header index is the last index (length -1) of indexes we want to keep
			t.headerIndex[t.supportedHeaders[header]] = int8(len(t.originalRowIndexesToKeep) - 1)
		}
	}

	return newHeaders
}
