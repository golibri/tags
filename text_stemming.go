package tags

import (
	"github.com/dchest/stemmer/german"
	"github.com/dchest/stemmer/porter2"
)

func (t *text) buildStems() *text {
	t.Dictionary = make(map[string][]string)
	if t.Language == "de" {
		t.buildStemsGerman()
	} else {
		t.buildStemsEnglish()
	}
	return t
}

func (t *text) buildStemsGerman() {
	stemmer := german.Stemmer
	index := make(map[string]int)
	for word, count := range t.WordIndex {
		stem := stemmer.Stem(word)
		index[stem] = index[stem] + count
		t.addToDictionary(stem, word)
	}
	t.StemIndex = index
}

func (t *text) buildStemsEnglish() {
	stemmer := porter2.Stemmer
	index := make(map[string]int)
	for word, count := range t.WordIndex {
		stem := stemmer.Stem(word)
		index[stem] = index[stem] + count
		t.addToDictionary(stem, word)
	}
	t.StemIndex = index
}

func (t *text) addToDictionary(stem string, word string) *text {
	list := t.Dictionary[stem]
	if !contains(list, word) {
		t.Dictionary[stem] = append(list, word)
	}
	return t
}
