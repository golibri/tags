// Determine relevant Tags (aka Keywords) from a given text string
package tags

import (
	"github.com/endeveit/guesslanguage"
)

func Calculate(str string) []string {
	lang := detectLanguage(str)
	t := NewCorpus(str, lang)
	return t
}

// stemmer are only available for german and english currently
func detectLanguage(str string) string {
	lang, err := guesslanguage.Guess(str)
	if err != nil {
		return "en"
	}
	if lang == "de" {
		return "de"
	}
	return "en"
}
