package tags

import (
	"github.com/bbalet/stopwords"
	"github.com/dchest/stemmer/german"
	"github.com/dchest/stemmer/porter2"
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
	if language != "en" { // most languages use english phrases online, too
		content = stopwords.CleanString(content, "en", false)
	}
	t := newText(content, language)
	r := Result{}
	r.Words = t.Words
	r.Stems = t.Stems
	r.Dictionary = t.Dictionary
	return r
}

// Stem exposes access to the used stemmer directly
func Stem(word string, language string) string {
	var stem string
	if language == "de" {
		stemmer := german.Stemmer
		stem = stemmer.Stem(word)
	} else {
		stemmer := porter2.Stemmer
		stem = stemmer.Stem(word)
	}
	return stem
}
