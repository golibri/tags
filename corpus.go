package tags

import (
	"github.com/dchest/stemmer/german"
	"github.com/dchest/stemmer/porter2"
	"github.com/golibri/list"
	"math"
	"regexp"
	"strings"
)

type corpus struct {
	text      string                // initial text
	language  string                // language of the text, "de" or fallback "en"
	sentences []string              // list of sentences of the text
	words     list.List             // list of all words
	wordCount map[string]int        // <word>:<number of occurences>
	stems     map[string]*list.List // <stem>:<list of words for stem>
	stemCount map[string]int        // <stem>:<number of occurences>
	phrases   map[string]int        // relevant phrases + number of occurences
	tags      list.List             // final result list
}

func NewCorpus(s string, lang string) []string {
	c := corpus{language: lang}
	c.cleanText()
	c.calculateSentences()
	c.calculateWords()
	c.calculateStems()
	c.calculatePhrases()
	c.calculateTags()
	return c.tags.Contents()
}

// remove unneccessary junk from text.
func (c *corpus) cleanText() {
	c.text = regexp.MustCompile(`\s{2,}`).ReplaceAllString(c.text, " ")
}

// split the text into sentences (naive)
func (c *corpus) calculateSentences() {
	t := strings.Replace(c.text, "|", "", -1)
	t = regexp.MustCompile(`[!\?\.]\s`).ReplaceAllString(t, "|")
	t = regexp.MustCompile(`[\\n]{1,}`).ReplaceAllString(t, "|")
	t = regexp.MustCompile(`\|{2,}`).ReplaceAllString(t, "|")
	c.sentences = strings.Split(t, "|")
}

// split the sentences into a list of words
func (c *corpus) calculateWords() {
	words := list.New()
	for _, sentence := range c.sentences {
		t := strings.ToLower(sentence)
		for _, w := range strings.Fields(t) {
			if regexp.MustCompile(`^[\p{L}\p{M}\w]{3,}$`).FindString(w) != "" {
				words.Push(w)
			}

		}
	}
	c.words = words
}

// reduce words into stems
func (c *corpus) calculateStems() {
	eng := porter2.Stemmer
	ger := german.Stemmer
	c.words.Each(func(w string) {
		var s string
		if c.language == "de" {
			s = ger.Stem(w)
		} else {
			s = eng.Stem(w)
		}
		c.stemCount[s]++
		c.stems[s].Push(w)
	})
}

// count phrases (n-grams)
func (c *corpus) calculatePhrases() {
	// TODO: bigrams, trigrams. not useful for now, needs some work!
	c.phrases = c.stemCount
}

// select the best matching phrases as tags
func (c *corpus) calculateTags() {
	length := len(c.stems)
	min := length * 100 / 10
	max := length * 100 / 50
	for stem, _ := range c.stems {
		x := c.stemCount[stem]
		if x < max && x >= min {
			c.tags.Push(stem)
		}
	}
}

// transform the found tags (still stems!) to a sorted slice
func (c *corpus) tagsToSlice() []string {
	hits := make(map[string]int)
	c.tags.Each(func(s string) {
		l := c.stems[s].Length()
		e := c.stems[s].Commonest()
		hits[e] = l
	})
	max_tags := int(math.Sqrt(float64(20)))
	slice := make([]string, c.tags.Length())
	var max int
	var cur string
	for i := 0; i < max_tags; i++ {
		max = 0
		picked := make(map[string]bool)
		for k, v := range hits {
			if v > max {
				max = v
				cur = k
			}
		}
		picked[cur] = true
		slice = append(slice, cur)
	}
	return slice
}
