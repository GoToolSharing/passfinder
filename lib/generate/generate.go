package generate

import (
	"fmt"
	"time"

	"github.com/GoToolSharing/passfinder/lib/utils"
)

func WithSpecialChars(wordlist []string) []string {
	specialChars := []string{"!", "@", "#", "$", "%", "+", "?", "="} // TODO: Configurable charset
	for _, word := range wordlist {
		for _, char := range specialChars {
			wordlist = append(wordlist, word+char)
		}
	}

	return wordlist
}

func WithMixedCase(wordlist []string) []string {
	var newWordlist []string
	for _, word := range wordlist {
		caseVariations := generateCaseVariations(word)
		newWordlist = append(newWordlist, caseVariations...)
	}
	return utils.RemoveDuplicates(newWordlist)
}

func generateCaseVariations(word string) []string {
	var result []string
	helper(word, "", 0, &result)
	return result
}

func helper(word, current string, index int, result *[]string) {
	if index == len(word) {
		*result = append(*result, current)
		return
	}
	helper(word, current+string(word[index]), index+1, result)
	if word[index] >= 'a' && word[index] <= 'z' {
		helper(word, current+string(word[index]-'a'+'A'), index+1, result)
	} else if word[index] >= 'A' && word[index] <= 'Z' {
		helper(word, current+string(word[index]-'A'+'a'), index+1, result)
	}
}

func WithYear(wordlist []string) []string {
	currentYear := time.Now().Year()
	var newWordlist []string
	for _, word := range wordlist {
		newWordlist = append(newWordlist, fmt.Sprintf("%s%d", word, currentYear))
	}
	return newWordlist
}
