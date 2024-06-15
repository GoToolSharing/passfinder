package cmd

import (
	"fmt"

	"github.com/GoToolSharing/passfinder/lib/username"
	"github.com/GoToolSharing/passfinder/lib/utils"
	"github.com/spf13/cobra"
)

var (
	firstName string
	lastName  string
	patterns  []string
)

var personCmd = &cobra.Command{
	Use:   "person",
	Short: "Generate a list of passwords based on a person's information",
	Long:  `Generate a list of probable passwords based on a person's first and last names, with customizable patterns.`,
	Run: func(cmd *cobra.Command, args []string) {
		passwords := generatePasswords(firstName, lastName, patterns)
		for _, password := range passwords {
			fmt.Println(password)
		}
	},
}

func init() {
	rootCmd.AddCommand(personCmd)
	personCmd.Flags().StringVarP(&firstName, "firstname", "f", "", "First Name")
	err := personCmd.MarkFlagRequired("firstname")
	if err != nil {
		return
	}
	personCmd.Flags().StringVarP(&lastName, "lastname", "l", "", "Last name")
	err = personCmd.MarkFlagRequired("lastname")
	if err != nil {
		return
	}
	personCmd.Flags().StringSliceVar(&patterns, "pattern", []string{}, "Specify patterns for password generation (e.g., 'flast, first, last')")
}

func generatePasswords(firstName, lastName string, patterns []string) []string {
	var passwords []string

	patternFuncs := map[string]func(string, string) []string{
		"first":     username.PatternFirst,
		"firstlast": username.PatternFirstLastNoSpace,
		"flast":     username.PatternFLastNoDot,
		"sfirst":    username.PatternSFirst,
		"lastf":     username.PatternLastFirstInit,
		"last":      username.PatternLast,
		"fl":        username.PatternFirstInitLastInit,
	}

	if len(patterns) == 0 {
		patterns = []string{
			"first",
			"firstlast",
			"flast",
			"sfirst",
			"lastf",
			"last",
			"fl",
		}
	}

	for _, pattern := range patterns {
		if patternFunc, exists := patternFuncs[pattern]; exists {
			passwords = append(passwords, patternFunc(firstName, lastName)...)
		}
	}

	return utils.RemoveDuplicates(passwords)
}
