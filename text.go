package tags

type text struct {
	Content    string
	Language   string
	Words      []string
	WordIndex  map[string]int
	Stems      []string
	StemIndex  map[string]int
	Dictionary map[string][]string
}

func newText(str string, language string) text {
	t := text{}
	t.Content = str
	t.Language = language
	t.
		clean().
		splitWords().
		uniqueWords().
		buildStems().
		process()
	return t
}
