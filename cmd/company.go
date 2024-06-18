package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/GoToolSharing/passfinder/config"
	"github.com/GoToolSharing/passfinder/generate"
	"github.com/GoToolSharing/passfinder/utils"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var (
	companyName           string
	includeYearSeparators bool
	includeYear           int
	includeStartCaps      bool
	includeShortYear      int
	includeEndSpecial     bool
	includeMixedCase      bool
	includeLeetCode       bool
	includeUppercase      bool
	includeMask           string
	includeNumbers        int
	includePostal         int
	includeAll            bool
)

var companyCmd = &cobra.Command{
	Use:   "company",
	Short: "Generate a passlist based on company information",
	Long:  `Generate a passlist based on company name, current year, city, and other relevant information.`,
	Run:   runCompanyCmd,
}

func init() {
	rootCmd.AddCommand(companyCmd)

	companyCmd.Flags().StringVarP(&companyName, "name", "n", "", "Company name")
	_ = companyCmd.MarkFlagRequired("name")
	companyCmd.Flags().IntVarP(&includeYear, "year", "y", -1, "Include the current year in the passwords")
	companyCmd.Flags().BoolVar(&includeYearSeparators, "year-separators", false, "Add special characters between the company name and the year")
	companyCmd.Flags().BoolVar(&includeMixedCase, "mixed-case", false, "Include variations with mixed case")
	companyCmd.Flags().BoolVarP(&includeEndSpecial, "end-special", "s", false, "Add special characters at the end of the passwords")
	companyCmd.Flags().BoolVar(&includeStartCaps, "start-caps", false, "Capitalize the first letter of the passwords")
	companyCmd.Flags().IntVar(&includeShortYear, "short-year", -1, "Truncate the year to the last two digits")
	companyCmd.Flags().BoolVarP(&includeLeetCode, "leet", "l", false, "Convert characters to leet speak")
	companyCmd.Flags().BoolVarP(&includeUppercase, "uppercase", "u", false, "Capitalize all letters of the passwords")
	companyCmd.Flags().StringVarP(&includeMask, "mask", "m", "", "Apply a custom mask to the passwords")
	companyCmd.Flags().IntVar(&includeNumbers, "numbers", 0, "Include numbers to the passwords")
	companyCmd.Flags().IntVarP(&includePostal, "postal", "p", 0, "Include postal code to the passwords")
	companyCmd.Flags().BoolVarP(&includeAll, "all", "a", false, "Include all variations")
}

// runCompanyCmd executes the main logic for generating the password list
func runCompanyCmd(cmd *cobra.Command, args []string) {
	if err := validateFlags(); err != nil {
		fmt.Println(err)
		return
	}

	wordlist := generateCompanyPasslist(companyName)

	if !config.GlobalConfig.BatchParam {
		passwordLength := len(wordlist)
		isConfirmed := utils.AskConfirmation(fmt.Sprintf("%d passwords will be generated, do you want to continue?", passwordLength))
		if !isConfirmed {
			return
		}
	}

	if config.GlobalConfig.OutputFile != "" {
		writePasswordsToFile(wordlist, config.GlobalConfig.OutputFile)
	} else {
		for _, password := range wordlist {
			fmt.Println(password)
		}
	}
}

// validateFlags ensures that the provided flags are logically consistent
func validateFlags() error {
	if (includeYear == -1 && includeShortYear == -1) && includeYearSeparators {
		return fmt.Errorf("You cannot use --year-separators without --year or --short-year")
	}

	return nil
}

// writePasswordsToFile writes the generated passwords to the specified file
func writePasswordsToFile(wordlist []string, filename string) {
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		spin.Stop()
		os.Exit(0)
	}()

	spin.Start()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, password := range wordlist {
		if _, err := file.WriteString(password + "\n"); err != nil {
			log.Fatal(err)
		}
	}
	spin.Stop()
}

// generateCompanyPasslist generates a list of passwords based on the provided company name and flags
func generateCompanyPasslist(name string) []string {
	spin := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		spin.Stop()
		os.Exit(0)
	}()

	spin.Start()
	baseWordlist := []string{strings.ToLower(name)}
	wordlist := baseWordlist

	if includeAll {
		includeYearSeparators = true
		includeYear = 0
		includeStartCaps = true
		includeShortYear = 0
		includeEndSpecial = true
		includeMixedCase = true
		includeLeetCode = true
		includeUppercase = true
		includeNumbers = 20
	}

	if includeYear != -1 || includeShortYear != -1 {
		var separators string
		if includeYearSeparators {
			separators = "!@#$%+?=*"
		}
		var yearWordlist []string
		year := time.Now().Year()
		if includeYear != -1 {
			yearWordlist = append(yearWordlist, generate.WithYear(baseWordlist, year, includeYear, separators)...)
		}
		if includeShortYear != -1 {
			shortYear := year % 100
			yearWordlist = append(yearWordlist, generate.WithYear(baseWordlist, shortYear, includeShortYear, separators)...)
		}
		wordlist = append(wordlist, yearWordlist...)
	}

	if includeNumbers != 0 {
		wordlist = append(wordlist, generate.WithNumbers(baseWordlist, includeNumbers)...)
	}

	if includePostal != 0 {
		wordlist = append(wordlist, generate.WithPostal(baseWordlist, includePostal)...)
	}

	wordlist = append(wordlist, baseWordlist...)

	transformedWordlist := wordlist
	if includeEndSpecial {
		transformedWordlist = append(transformedWordlist, generate.WithSpecialChars(wordlist)...)
	}
	if includeMixedCase {
		transformedWordlist = append(transformedWordlist, generate.WithMixedCase(wordlist)...)
	}
	if includeLeetCode {
		transformedWordlist = append(transformedWordlist, generate.WithLeetCode(transformedWordlist)...)
	}
	if includeStartCaps {
		transformedWordlist = append(transformedWordlist, generate.WithStartCaps(transformedWordlist)...)
	}
	if includeUppercase {
		transformedWordlist = append(transformedWordlist, generate.WithUppercase(transformedWordlist)...)
	}
	if includeMask != "" {
		transformedWordlist = append(transformedWordlist, generate.WithMask(transformedWordlist, includeMask)...)
	}

	transformedWordlist = append(wordlist, transformedWordlist...)

	wordlist = utils.RemoveDuplicates(transformedWordlist)
	spin.Stop()

	return wordlist
}
