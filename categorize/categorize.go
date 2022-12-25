package categorize

import "github.com/Pav0l/nance/lib/json"

type result struct {
	Target         string
	ReviewManually string
}

type Categorize struct {
	categories map[string]string
}

// NewClassifier takes in file name of categories map which is used to catogirize inputs
func NewClassifier(categoriesFileName string) *Categorize {
	// read the categories JSON file only once and store it in memory
	categories := json.ReadFile(categoriesFileName)

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
