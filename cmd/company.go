package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/GoToolSharing/passfinder/lib/generate"
	"github.com/GoToolSharing/passfinder/lib/utils"
	"github.com/spf13/cobra"
)

var (
	companyName            string
	city                   string
	yearSeparators         bool
	includeYear            bool
	includeAllPermutations bool
	startCaps              bool
	shortYear              bool
	includeSpecialChars    bool
	includeMixedCase       bool
	includeLeetCode        bool
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Generate a passlist based on company information",
	Long:  `Generate a passlist based on company name, current year, city, and other relevant information.`,
	Run: func(cmd *cobra.Command, args []string) {

		if (!includeYear && !shortYear) && yearSeparators {
			fmt.Println("You cannot use --year-separators without --year or --short-year")
			return
		}

		if includeYear && shortYear {
			fmt.Println("You cannot use both --year and --short-year")
			return
		}

		wordlist := generateCompanyPasslist(companyName, city)

		wordlist = utils.RemoveDuplicates(wordlist)
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
	companyCmd.Flags().BoolVar(&includeMixedCase, "mixed-case", false, "Include mixed case variations")
	companyCmd.Flags().BoolVar(&yearSeparators, "year-separators", false, "Special characters to separate the company name and the year")
	companyCmd.Flags().BoolVar(&includeSpecialChars, "end-special", false, "Include special characters at the end of the passwords")
	companyCmd.Flags().BoolVar(&includeAllPermutations, "all", false, "Run all permutations")
	companyCmd.Flags().BoolVar(&startCaps, "start-caps", false, "First letter in caps")
	companyCmd.Flags().BoolVar(&shortYear, "short-year", false, "Truncate the year to two digits")
	companyCmd.Flags().BoolVar(&includeLeetCode, "leet", false, "Add leet code")
	// companyCmd.Flags().BoolVar(&includeSpecialChars, "pass-pol", false, "Password Policy (remove bad passwords)")
}

func generateCompanyPasslist(name, city string) []string {
	var wordlist []string

	wordlist = append(wordlist, strings.ToLower(name)) // Init the wordlist

	if includeAllPermutations {
		includeMixedCase = true
		includeYear = true
		yearSeparators = true
		shortYear = false // We cannot have both year and shortYear
		includeSpecialChars = true
		includeLeetCode = true
	}

	if startCaps {
		wordlist = generate.WithStartCaps(wordlist)
	}

	if includeMixedCase {
		wordlist = generate.WithMixedCase(wordlist)
	}

	if includeLeetCode {
		wordlist = generate.WithLeetCode(wordlist)
	}

	if shortYear {
		year := time.Now().Year() % 100
		var separators string
		if yearSeparators {
			separators = "!@#$%+?=*"
		}
		wordlist = generate.WithYearAndSeparators(wordlist, year, separators)
	}

	if includeYear {
		year := time.Now().Year()
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
