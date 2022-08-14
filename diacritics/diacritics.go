package diacritics

var d = map[string]string{
	"á":  "a",
	"ä":  "a",
	"č":  "c",
	"ď":  "d",
	"dž": "dz",
	"é":  "e",
	"í":  "i",
	"ĺ":  "l",
	"ľ":  "l",
	"ň":  "n",
	"ó":  "o",
	"ô":  "o",
	"ŕ":  "r",
	"š":  "s",
	"ť":  "t",
	"ú":  "t",
	"ý":  "y",
	"ž":  "z",
	"Á":  "A",
	"Ä":  "A",
	"Č":  "C",
	"Ď":  "D",
	"DŽ": "DZ",
	"É":  "E",
	"Í":  "I",
	"Ĺ":  "L",
	"Ľ":  "L",
	"Ň":  "N",
	"Ó":  "O",
	"Ô":  "O",
	"Ŕ":  "R",
	"Š":  "S",
	"Ť":  "T",
	"Ú":  "T",
	"Ý":  "Y",
	"Ž":  "Z",
}

// Replace replaces character with diacritics with it's non-diacritic alternative or returns the character.
// e.g. `á` is replaced with `a`, but `b` just returns `b`
func Replace(s string) string {
	replaced := d[s]

	if len(replaced) == 0 {
		return s
	}

	return replaced
}
