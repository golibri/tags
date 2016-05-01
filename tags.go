package tags

import (
	"github.com/bbalet/stopwords"
)

// Result holds a data structure with all relevant metadata of the calculation
type Result struct {
	Words      []string
	Stems      []string
	Dictionary map[string][]string // stem -> [words]
}

// Calculate takes a text and a language hint, either "en" or "de" for now.
func Calculate(str string, language string) Result {
	content := stopwords.CleanString(str, language, true)
	t := newText(content, language)
	r := Result{}
	r.Words = t.Words
	r.Stems = t.Stems
	r.Dictionary = t.Dictionary
	return r
}
