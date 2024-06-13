package generate

import (
	"fmt"
	"strings"
)

func WithSpecialChars(wordlist []string) []string {
	specialChars := []string{"!", "@", "#", "$", "%", "+", "?", "=", "*"} // TODO: Configurable charset
	for _, word := range wordlist {
		for _, char := range specialChars {
			wordlist = append(wordlist, word+char)
		}
	}

	return wordlist
}

func WithMixedCase(wordlist []string) []string {
	for _, word := range wordlist {
		caseVariations := generateCaseVariations(word)
		wordlist = append(wordlist, caseVariations...)
	}
	return wordlist
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

func WithYearAndSeparators(wordlist []string, year int, separators string) []string {
	separatorsList := strings.Split(separators, "")

	for _, word := range wordlist {
		wordlist = append(wordlist, fmt.Sprintf("%s%d", word, year))
		for _, separator := range separatorsList {
			wordlist = append(wordlist, fmt.Sprintf("%s%s%d", word, separator, year))
		}
	}
	return wordlist
}

func WithStartCaps(wordlist []string) []string {
	for _, word := range wordlist {
		if len(word) > 0 {
			wordlist = append(wordlist, strings.ToUpper(string(word[0]))+word[1:])
		} else {
			wordlist = append(wordlist, word)
		}
	}
	return wordlist
}

func WithLeetCode(wordlist []string) []string {
	leetMap := map[rune][]string{
		'a': {"4", "@"}, 'A': {"4", "@"},
		'e': {"3"}, 'E': {"3"},
		'i': {"1", "!"}, 'I': {"1", "!"},
		'o': {"0"}, 'O': {"0"},
		's': {"5", "$"}, 'S': {"5", "$"},
		't': {"7"}, 'T': {"7"},
		'l': {"1"}, 'L': {"1"},
	}

	for _, word := range wordlist {
		leetVariations := generateLeetVariations(word, leetMap)
		wordlist = append(wordlist, leetVariations...)
	}

	return wordlist
}

func generateLeetVariations(word string, leetMap map[rune][]string) []string {
	var result []string
	helperLeet(word, "", 0, &result, leetMap)
	return result
}

func helperLeet(word, current string, index int, result *[]string, leetMap map[rune][]string) {
	if index == len(word) {
		*result = append(*result, current)
		return
	}

	char := rune(word[index])
	if leetChars, ok := leetMap[char]; ok {
		for _, leetChar := range leetChars {
			helperLeet(word, current+leetChar, index+1, result, leetMap)
		}
	}
	helperLeet(word, current+string(char), index+1, result, leetMap)
}
