package generate

func WithSpecialChars(wordlist []string) []string {
	specialChars := []string{"!", "@", "#", "$", "%", "+", "?", "="} // TODO: Configurable charset
	for _, word := range wordlist {
		for _, char := range specialChars {
			wordlist = append(wordlist, word+char)
		}
	}

	return wordlist
}
