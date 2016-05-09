package tags

import (
	"regexp"
	"strings"
)

func (t *text) clean() *text {
	rxJunk := regexp.MustCompile(`[\.\,\?!;\s]\s`)
	t.Content = rxJunk.ReplaceAllString(t.Content, " ")
	rxSpace := regexp.MustCompile(`\s{2,}`)
	t.Content = rxSpace.ReplaceAllString(t.Content, " ")
	t.Content = strings.TrimSpace(t.Content)
	return t
}

func (t *text) splitWords() *text {
	list := strings.Fields(t.Content)
	index := make(map[string]int)
	for _, word := range list {
		if len(word) > 2 {
			index[word] = index[word] + 1
		}
	}
	t.WordIndex = index
	return t
}

func (t *text) uniqueWords() *text {
	words := []string{}
	for word := range t.WordIndex {
		words = append(words, word)
	}
	t.Words = words
	return t
}
