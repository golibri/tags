package tags

func contains(arr []string, s string) bool {
	included := false
	for _, element := range arr {
		if element == s {
			included = true
			break
		}
	}
	return included
}
