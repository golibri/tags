package tags

func (t *text) process() *text {
	index := make(map[string]float64)
	sum := float64(len(t.StemIndex))
	for stem, count := range t.StemIndex {
		index[stem] = float64(count) / sum
	}
	results := pickHighestResults(index)
	t.Stems = results
	t.
		removeUnimportantStems().
		stemsToWords().
		removeUnimportantWords()
	return t
}

func (t *text) stemsToWords() *text {
	list := []string{}
	for _, words := range t.Dictionary {
		for _, word := range words {
			list = append(list, word)
		}
	}
	t.Words = list
	return t
}

func (t *text) removeUnimportantStems() *text {
	for stem := range t.StemIndex {
		if !contains(t.Stems, stem) {
			delete(t.StemIndex, stem)
			delete(t.Dictionary, stem)
		}
	}
	return t
}

func (t *text) removeUnimportantWords() *text {
	for _, word := range t.Words {
		if !contains(t.Words, word) {
			delete(t.WordIndex, word)
		}
	}
	return t
}

func pickHighestResults(index map[string]float64) []string {
	numMatches := 10
	results := []string{}
	if len(index) <= numMatches {
		for stem := range index {
			results = append(results, stem)
		}
		return results
	}
	var max float64
	var target string
	for i := 0; i < numMatches; i++ {
		max = 0
		target = ""
		for stem, val := range index {
			if val > max {
				max = val
				target = stem
			}
		}
		delete(index, target)
		results = append(results, target)
	}
	return results
}
