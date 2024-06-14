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
	companyName           string
	includeYearSeparators bool
	includeYear           bool
	includeStartCaps      bool
	includeShortYear      bool
	includeEndSpecial     bool
	includeMixedCase      bool
	includeLeetCode       bool
	includeUppercase      bool
	includeMask           string
	includeYearRange      int
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

		if (!includeYear && !includeShortYear) && includeYearRange != 0 {
			fmt.Println("You cannot use --year-range without --year or --short-year")
			return
		}
		// End Temp

		wordlist := generateCompanyPasslist(companyName)

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
	companyCmd.Flags().BoolVar(&includeYearSeparators, "year-separators", false, "Add special characters between the company name and the year")
	companyCmd.Flags().IntVar(&includeYearRange, "year-range", 0, "Include a range of years around the current year")
	companyCmd.Flags().BoolVar(&includeMixedCase, "mixed-case", false, "Include variations with mixed case")
	companyCmd.Flags().BoolVarP(&includeEndSpecial, "end-special", "s", false, "Add special characters at the end of the passwords")
	companyCmd.Flags().BoolVar(&includeStartCaps, "start-caps", false, "Capitalize the first letter of the passwords")
	companyCmd.Flags().BoolVar(&includeShortYear, "short-year", false, "Truncate the year to the last two digits")
	companyCmd.Flags().BoolVarP(&includeLeetCode, "leet", "l", false, "Convert characters to leet speak")
	companyCmd.Flags().BoolVarP(&includeUppercase, "uppercase", "u", false, "Capitalize all letters of the passwords")
	companyCmd.Flags().StringVarP(&includeMask, "mask", "m", "", "Apply a custom mask to the passwords")
}

func generateCompanyPasslist(name string) []string {
	var wordlist []string

	wordlist = append(wordlist, strings.ToLower(name))

	if includeStartCaps {
		wordlist = generate.WithStartCaps(wordlist)
	}

	if includeMixedCase {
		wordlist = generate.WithMixedCase(wordlist)
	}

	if includeLeetCode {
		wordlist = generate.WithLeetCode(wordlist)
	}

	if includeUppercase {
		wordlist = generate.WithUppercase(wordlist)
	}

	year := time.Now().Year()
	if includeShortYear {
		year = year % 100
	}

	if includeYear || includeShortYear {
		var separators string
		if includeYearSeparators {
			separators = "!@#$%+?=*"
		}
		wordlist = generate.WithYearAndSeparators(wordlist, year, separators, includeYearRange)
	}

	if includeEndSpecial {
		wordlist = generate.WithSpecialChars(wordlist)
	}

	if includeMask != "" {
		wordlist = generate.WithMask(wordlist, includeMask)
	}

	return utils.RemoveDuplicates(wordlist)
}
