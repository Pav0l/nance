package categorize

import (
	"encoding/json"
	"log"
	"os"
)

type result struct {
	Target         string
	ReviewManually string
}

type Categorize struct {
	categories map[string]string
}

func NewClassifier() *Categorize {
	// read the categories JSON file only once and store it in memory
	categories := readCategories()

	return &Categorize{
		categories: categories,
	}
}

func (c *Categorize) Categorize(partner, originalCategory string) result {
	categorized := c.categories[originalCategory]

	if len(categorized) == 0 {
		categorized = c.categories[partner]

		if len(categorized) == 0 {
			return result{
				Target:         originalCategory,
				ReviewManually: "1",
			}
		}
	}

	return result{
		Target:         categorized,
		ReviewManually: "0",
	}
}

func readCategories() map[string]string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	fp := wd + "/categorize/categories.json"

	data, err := os.ReadFile(fp)
	if err != nil {
		log.Fatalln(err)
	}

	categories := map[string]string{}
	err = json.Unmarshal(data, &categories)
	if err != nil {
		log.Fatalln(err)
	}

	return categories
}
