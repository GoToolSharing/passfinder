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

func WithMask(wordlist []string, mask string) []string {
	specialChars := []rune("!@#$%^&*()_+-=[]{}|;:'\",.<>?/\\")
	digits := []rune("0123456789")
	lowercase := []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	allChars := append(append(append(lowercase, uppercase...), digits...), specialChars...)

	for _, word := range wordlist {
		var maskedWords []string
		applyMask(word, mask, "", &maskedWords, specialChars, digits, lowercase, uppercase, allChars)
		wordlist = append(wordlist, maskedWords...)
	}

	return wordlist
}

func applyMask(word, mask, current string, result *[]string, specialChars, digits, lowercase, uppercase, allChars []rune) {
	if len(mask) == 0 {
		*result = append(*result, current)
		return
	}

	switch mask[0] {
	case '%':
		if len(mask) > 1 {
			switch mask[1] {
			case 'w':
				applyMask(word, mask[2:], current+word, result, specialChars, digits, lowercase, uppercase, allChars)
			case 's':
				for _, char := range specialChars {
					applyMask(word, mask[2:], current+string(char), result, specialChars, digits, lowercase, uppercase, allChars)
				}
			case 'l':
				for _, char := range lowercase {
					applyMask(word, mask[2:], current+string(char), result, specialChars, digits, lowercase, uppercase, allChars)
				}
			case 'u':
				for _, char := range uppercase {
					applyMask(word, mask[2:], current+string(char), result, specialChars, digits, lowercase, uppercase, allChars)
				}
			case 'd':
				for _, char := range digits {
					applyMask(word, mask[2:], current+string(char), result, specialChars, digits, lowercase, uppercase, allChars)
				}
			case 'a':
				for _, char := range allChars {
					applyMask(word, mask[2:], current+string(char), result, specialChars, digits, lowercase, uppercase, allChars)
				}
			}
		}
	default:
		applyMask(word, mask[1:], current+string(mask[0]), result, specialChars, digits, lowercase, uppercase, allChars)
	}
}

func WithUppercase(wordlist []string) []string {
	for _, word := range wordlist {
		wordlist = append(wordlist, strings.ToUpper(word))
	}
	return wordlist
}
