// Determine relevant Tags (aka Keywords) from a given text string
package tags

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type corpus struct {
	input     string
	text      string
	sentences []string
	words     []string
	tags      []string
}

func Calculate(str string) []string {
	c := corpus{input: str}
	c.calculateText()
	c.calculateSentences()
	c.calculateWords()
	// fmt.Println(c)
	// return c.tags
	return c.words
}

func (c *corpus) calculateText() {
	spaces := regexp.MustCompile(`\s{2,}`)
	c.text = spaces.ReplaceAllString(c.input, " ")
}

func (c *corpus) calculateSentences() {
	piper := regexp.MustCompile(`|`)
	t := piper.ReplaceAllString(c.text, "")
	splitter := regexp.MustCompile(`[!\?\.]\s`)
	t = splitter.ReplaceAllString(t, "|")
	c.sentences = strings.Split(t, "|")
}

func (c *corpus) calculateWords() {
	cleaner := regexp.MustCompile(`[!\?\.]`)
	t := cleaner.ReplaceAllString(c.text, " ")
	t = strings.ToLower(t)
	words := strings.Fields(t)
	var result []string
	characters := regexp.MustCompile(`^[\p{L}\p{M}\w]+$`)
	for _, e := range words {
		if len(e) > 2 && characters.FindString(e) != "" {
			result = append(result, e)
		}
	}
	c.words = result
}

func (c *corpus) calculateTags() {
	var count map[string]int
	for _, e := range c.words {
		count[e]++
	}

	sum := len(count)
	min := sum * 10 / 100
	max := sum * 50 / 100
	for k, v := range count {
		fmt.Println(k)
		if v < min || v > max {
			delete(count, k)
		}
	}

	var tags []string
	limit := int(math.Sqrt(float64(len(c.words))))
	for k, _ := range count {
		if len(tags) < limit {
			tags = append(tags, k)
		}
	}

	c.tags = tags
}
