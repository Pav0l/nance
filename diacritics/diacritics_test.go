package diacritics

import "testing"

func TestReplace(t *testing.T) {
	testcases := []struct {
		receive, expect string
	}{
		// with diacritics
		{"á", "a"},
		{"ä", "a"},
		{"č", "c"},
		{"ď", "d"},
		{"dž", "dz"},
		{"é", "e"},
		{"í", "i"},
		{"ĺ", "l"},
		{"ľ", "l"},
		{"ň", "n"},
		{"ó", "o"},
		{"ô", "o"},
		{"ŕ", "r"},
		{"š", "s"},
		{"ť", "t"},
		{"ú", "t"},
		{"ý", "y"},
		{"ž", "z"},
		{"Á", "A"},
		{"Ä", "A"},
		{"Č", "C"},
		{"Ď", "D"},
		{"DŽ", "DZ"},
		{"É", "E"},
		{"Í", "I"},
		{"Ĺ", "L"},
		{"Ľ", "L"},
		{"Ň", "N"},
		{"Ó", "O"},
		{"Ô", "O"},
		{"Ŕ", "R"},
		{"Š", "S"},
		{"Ť", "T"},
		{"Ú", "T"},
		{"Ý", "Y"},
		{"Ž", "Z"},
		{"áŽÍksX", "aZIksX"},
		{"TŤ Ík,sX.com", "TT Ik,sX.com"},
		// no diacritics
		{"b", "b"},
		{" ", " "},
		{"", ""},
		{"foo", "foo"},
	}

	for _, testcase := range testcases {
		res := Replace(testcase.receive)
		if res != testcase.expect {
			t.Log("failed:", testcase.receive, "expected", testcase.expect, "received", res)
			t.Fail()
		}
	}
}
