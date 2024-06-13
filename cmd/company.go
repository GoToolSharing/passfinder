package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/GoToolSharing/passfinder/lib/generate"
	"github.com/spf13/cobra"
)

var (
	companyName            string
	city                   string
	yearSeparators         bool
	includeYear            bool
	includeAllPermutations bool
	startCaps              bool
	// includeCity         bool
	// includeAcronym      bool
	includeSpecialChars bool
	// includeNumericSeq   bool
	// includeUpperCase    bool
	// includeLowerCase    bool
	includeMixedCase bool
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Generate a passlist based on company information",
	Long:  `Generate a passlist based on company name, current year, city, and other relevant information.`,
	Run: func(cmd *cobra.Command, args []string) {

		if !includeYear && yearSeparators {
			fmt.Println("You cannot use --year-separators without --year")
			return
		}

		year := time.Now().Year()

		wordlist := generateCompanyPasslist(companyName, city, year)

		// TODO: Remove duplicates
		for _, password := range wordlist {
			fmt.Println(password)
		}
	},
}

func init() {
	rootCmd.AddCommand(companyCmd)

	companyCmd.Flags().StringVarP(&companyName, "name", "n", "", "Company name")
	err := companyCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}
	// companyCmd.Flags().StringVarP(&city, "city", "c", "", "City of the company")
	companyCmd.Flags().BoolVar(&includeYear, "year", false, "Include the current year in passwords")
	// companyCmd.Flags().BoolVar(&includeNumericSeq, "numeric", false, "Include numeric sequences in passwords")
	companyCmd.Flags().BoolVar(&includeMixedCase, "mixed-case", false, "Include mixed case variations")
	companyCmd.Flags().BoolVar(&yearSeparators, "year-separators", false, "Special characters to separate the company name and the year")
	companyCmd.Flags().BoolVar(&includeSpecialChars, "end-special", false, "Include special characters at the end of the passwords")
	companyCmd.Flags().BoolVar(&includeAllPermutations, "all", false, "Run all permutations")
	companyCmd.Flags().BoolVar(&startCaps, "start-caps", false, "First letter in caps")
	// companyCmd.Flags().BoolVar(&includeSpecialChars, "pass-pol", false, "Password Policy (remove bad passwords)")
}

func generateCompanyPasslist(name, city string, year int) []string {
	var wordlist []string

	wordlist = append(wordlist, strings.ToLower(name)) // Init the wordlist

	if includeAllPermutations {
		includeMixedCase = true
		includeYear = true
		yearSeparators = true
		includeSpecialChars = true
	}

	if startCaps {
		wordlist = generate.WithStartCaps(wordlist)
	}

	// if includeUpperCase {
	// 	wordlist = append(wordlist, strings.ToUpper(name))
	// }

	// if includeLowerCase {
	// 	wordlist = append(wordlist, strings.ToLower(name))
	// }

	// if includeCity && city != "" {
	// 	wordlist = append(wordlist, fmt.Sprintf("%s%s", name, city))
	// 	wordlist = append(wordlist, fmt.Sprintf("%s%s", city, name))
	// }

	// if includeAcronym {
	// 	acronym := getAcronym(name)
	// 	wordlist = append(wordlist, acronym)
	// }

	// if includeNumericSeq {
	// 	for i := 1; i <= 3; i++ {
	// 		wordlist = append(wordlist, fmt.Sprintf("%s%d", name, i))
	// 		wordlist = append(wordlist, fmt.Sprintf("%s%d", strings.ToLower(name), i))
	// 		wordlist = append(wordlist, fmt.Sprintf("%s%d", strings.ToUpper(name), i))
	// 	}
	// }

	if includeMixedCase {
		wordlist = generate.WithMixedCase(wordlist)
	}

	if includeYear {
		var separators string
		if yearSeparators {
			separators = "!@#$%+?=*"
		}
		wordlist = generate.WithYearAndSeparators(wordlist, year, separators)
	}

	if includeSpecialChars {
		wordlist = generate.WithSpecialChars(wordlist)
	}

	return wordlist
}

// func getAcronym(name string) string {
// 	words := strings.Fields(name)
// 	acronym := ""
// 	for _, word := range words {
// 		acronym += strings.ToUpper(string(word[0]))
// 	}
// 	return acronym
// }
