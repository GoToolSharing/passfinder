package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/GoToolSharing/passfinder/config"
	"github.com/GoToolSharing/passfinder/lib/generate"
	"github.com/GoToolSharing/passfinder/lib/utils"
	"github.com/spf13/cobra"
)

var (
	companyName            string
	city                   string
	includeYearSeparators  bool
	includeYear            bool
	includeAllPermutations bool
	includeStartCaps       bool
	includeShortYear       bool
	includeSpecialChars    bool
	includeMixedCase       bool
	includeLeetCode        bool
	includeMask            string
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Generate a passlist based on company information",
	Long:  `Generate a passlist based on company name, current year, city, and other relevant information.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Temp
		if (!includeYear && !includeShortYear) && includeYearSeparators {
			fmt.Println("You cannot use --year-separators without --year or --short-year")
			return
		}

		if includeYear && includeShortYear {
			fmt.Println("You cannot use both --year and --short-year")
			return
		}
		// End Temp

		wordlist := generateCompanyPasslist(companyName, city)

		wordlist = utils.RemoveDuplicates(wordlist)
		if config.GlobalConfig.OutputFile != "" {
			file, err := os.OpenFile(config.GlobalConfig.OutputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			for _, password := range wordlist {
				_, err := file.WriteString(password + "\n")
				if err != nil {
					log.Fatal(err)
					break
				}
			}
		} else {
			for _, password := range wordlist {
				fmt.Println(password)
			}
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
	companyCmd.Flags().BoolVarP(&includeYear, "year", "y", false, "Include the current year in the passwords")
	companyCmd.Flags().BoolVar(&includeMixedCase, "mixed-case", false, "Include variations with mixed case")
	companyCmd.Flags().BoolVar(&includeYearSeparators, "year-separators", false, "Add special characters between the company name and the year")
	companyCmd.Flags().BoolVarP(&includeSpecialChars, "end-special", "s", false, "Add special characters at the end of the passwords")
	companyCmd.Flags().BoolVarP(&includeAllPermutations, "all", "a", false, "Generate all possible permutations of the password")
	companyCmd.Flags().BoolVar(&includeStartCaps, "start-caps", false, "Capitalize the first letter of the passwords")
	companyCmd.Flags().BoolVar(&includeShortYear, "short-year", false, "Truncate the year to the last two digits")
	companyCmd.Flags().BoolVarP(&includeLeetCode, "leet", "l", false, "Convert characters to leet speak")
	companyCmd.Flags().StringVarP(&includeMask, "mask", "m", "", "Apply a custom mask to the passwords")
}

func generateCompanyPasslist(name, city string) []string {
	var wordlist []string

	wordlist = append(wordlist, strings.ToLower(name))

	if includeAllPermutations {
		includeMixedCase = true
		includeYear = true
		includeYearSeparators = true
		includeShortYear = false
		includeSpecialChars = true
		includeLeetCode = true
		includeStartCaps = true
	}

	if includeStartCaps {
		wordlist = generate.WithStartCaps(wordlist)
	}

	if includeMixedCase {
		wordlist = generate.WithMixedCase(wordlist)
	}

	if includeLeetCode {
		wordlist = generate.WithLeetCode(wordlist)
	}

	if includeShortYear {
		year := time.Now().Year() % 100
		var separators string
		if includeYearSeparators {
			separators = "!@#$%+?=*"
		}
		wordlist = generate.WithYearAndSeparators(wordlist, year, separators)
	}

	if includeYear {
		year := time.Now().Year()
		var separators string
		if includeYearSeparators {
			separators = "!@#$%+?=*"
		}
		wordlist = generate.WithYearAndSeparators(wordlist, year, separators)
	}

	if includeSpecialChars {
		wordlist = generate.WithSpecialChars(wordlist)
	}

	if includeMask != "" {
		wordlist = generate.WithMask(wordlist, includeMask)
	}

	return wordlist
}
