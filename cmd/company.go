package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/GoToolSharing/passfinder/lib/generate"
	"github.com/spf13/cobra"
)

var (
	companyName string
	city        string
	// includeYear         bool
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
		year := time.Now().Year()

		wordlist := generateCompanyPasslist(companyName, city, year)
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
	// companyCmd.Flags().BoolVar(&includeYear, "year", false, "Include the current year in passwords")
	// companyCmd.Flags().BoolVar(&includeCity, "with-city", false, "Include the city in passwords")
	// companyCmd.Flags().BoolVar(&includeAcronym, "acronym", false, "Include acronym of the company name")
	// companyCmd.Flags().BoolVar(&includeNumericSeq, "numeric", false, "Include numeric sequences in passwords")
	// companyCmd.Flags().BoolVar(&includeUpperCase, "uppercase", false, "Include uppercase variations")
	// companyCmd.Flags().BoolVar(&includeLowerCase, "lowercase", false, "Include lowercase variations")
	companyCmd.Flags().BoolVar(&includeMixedCase, "mixedcase", false, "Include mixed case variations")
	companyCmd.Flags().BoolVar(&includeSpecialChars, "end-special", false, "Include special characters at the end of the passwords")
}

func generateCompanyPasslist(name, city string, year int) []string {
	var wordlist []string

	wordlist = append(wordlist, strings.ToLower(name)) // Init the wordlist

	// if includeUpperCase {
	// 	wordlist = append(wordlist, strings.ToUpper(name))
	// }

	// if includeLowerCase {
	// 	wordlist = append(wordlist, strings.ToLower(name))
	// }

	// if includeYear {
	// 	wordlist = append(wordlist, fmt.Sprintf("%s%d", name, year))
	// 	wordlist = append(wordlist, fmt.Sprintf("%s%d", strings.ToLower(name), year))
	// 	wordlist = append(wordlist, fmt.Sprintf("%s%d", strings.ToUpper(name), year))
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
