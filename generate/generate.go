package generate

import (
	"fmt"
	"strconv"
	"strings"
)

// WithSpecialChars adds special characters to each word in the wordlist
func WithSpecialChars(wordlist []string) []string {
	var result []string
	specialChars := []string{"!", "@", "#", "$", "%", "+", "?", "=", "*"}
	for _, word := range wordlist {
		for _, char := range specialChars {
			result = append(result, word+char)
		}
	}
	return result
}

// WithNumbers appends numbers to each word in the wordlist up to a specified range
func WithNumbers(wordlist []string, specificRange int) []string {
	var result []string
	for _, word := range wordlist {
		for i := 0; i <= specificRange; i++ {
			result = append(result, word+strconv.Itoa(i))
		}
	}
	return result
}

// WithMixedCase generates case variations for each word in the wordlist
func WithMixedCase(wordlist []string) []string {
	var result []string
	for _, word := range wordlist {
		result = append(result, generateCaseVariations(word)...)
	}
	return result
}

// generateCaseVariations generates all possible case variations of a word
func generateCaseVariations(word string) []string {
	var result []string
	helper(word, "", 0, &result)
	return result
}

// helper is a recursive function that generates case variations of a word
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

// WithYear appends the year to each word in the wordlist
func WithYear(wordlist []string, year int, yearRange int, separators string) []string {
	var result []string
	separatorsList := strings.Split(separators, "")

	addYearWithSeparators := func(word string, year int) {
		result = append(result, fmt.Sprintf("%s%d", word, year))
		for _, separator := range separatorsList {
			result = append(result, fmt.Sprintf("%s%s%d", word, separator, year))
		}
	}

	if yearRange > 0 {
		for i := year - yearRange; i <= year+yearRange; i++ {
			for _, word := range wordlist {
				addYearWithSeparators(word, i)
			}
		}
	} else {
		for _, word := range wordlist {
			addYearWithSeparators(word, year)
		}
	}

	return append(wordlist, result...)
}

// WithStartCaps capitalizes the first letter of each word in the wordlist
func WithStartCaps(wordlist []string) []string {
	var result []string
	for _, word := range wordlist {
		if len(word) > 0 {
			result = append(result, strings.ToUpper(string(word[0]))+word[1:])
		} else {
			result = append(result, word)
		}
	}
	return result
}

// WithLeetCode converts certain characters in each word to leet speak
func WithLeetCode(wordlist []string) []string {
	var result []string
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
		result = append(result, generateLeetVariations(word, leetMap)...)
	}

	return result
}

// generateLeetVariations generates all possible leet speak variations of a word
func generateLeetVariations(word string, leetMap map[rune][]string) []string {
	var result []string
	helperLeet(word, "", 0, &result, leetMap)
	return result
}

// helperLeet is a recursive function that generates leet speak variations of a word
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

// WithMask applies a custom mask to each word in the wordlist
func WithMask(wordlist []string, mask string) []string {
	var result []string
	specialChars := []rune("!@#$%^&*()_+-=[]{}|;:'\",.<>?/\\")
	digits := []rune("0123456789")
	lowercase := []rune("abcdefghijklmnopqrstuvwxyz")
	uppercase := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	allChars := append(append(append(lowercase, uppercase...), digits...), specialChars...)

	for _, word := range wordlist {
		var maskedWords []string
		applyMask(word, mask, "", &maskedWords, specialChars, digits, lowercase, uppercase, allChars)
		result = append(result, maskedWords...)
	}

	return result
}

// applyMask is a recursive function that applies a custom mask to a word
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

// WithUppercase converts each word in the wordlist to uppercase
func WithUppercase(wordlist []string) []string {
	var result []string
	for _, word := range wordlist {
		result = append(result, strings.ToUpper(word))
	}
	return result
}

// WithPostal appends the postal code to each word in the wordlist
func WithPostal(wordlist []string, includePostal int) []string {
	var result []string
	for _, word := range wordlist {
		result = append(result, word+strconv.Itoa(includePostal))
	}
	return result
}

func WithCity(wordlist []string, includeCity string) []string {
	var result []string

	for _, word := range wordlist {
		result = append(wordlist, word+includeCity)
	}

	return result
}
