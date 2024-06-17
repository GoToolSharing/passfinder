package cmd

import (
	"reflect"
	"strings"
	"testing"
)

func cleanup() {
	includeYearSeparators = false
	includeYear = false
	includeStartCaps = false
	includeShortYear = false
	includeEndSpecial = false
	includeMixedCase = false
	includeLeetCode = false
	includeUppercase = false
	includeMask = ""
	includeYearRange = 0
}

func TestWithoutFlag(t *testing.T) {
	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestYearFlag(t *testing.T) {
	includeYear = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo2024"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestYearRangeFlag(t *testing.T) {
	includeYear = true
	includeYearRange = 3

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo2021\ndemo2022\ndemo2023\ndemo2024\ndemo2025\ndemo2026\ndemo2027"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestEndSpecialFlag(t *testing.T) {
	includeEndSpecial = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo!\ndemo@\ndemo#\ndemo$\ndemo%\ndemo+\ndemo?\ndemo=\ndemo*"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestUppercaseFlag(t *testing.T) {
	includeUppercase = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\nDEMO"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestLeetFlag(t *testing.T) {
	includeLeetCode = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\nd3m0\nd3mo\ndem0"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestStartCapsFlag(t *testing.T) {
	includeStartCaps = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\nDemo"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestShortYearFlag(t *testing.T) {
	includeShortYear = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo24"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestMixedCaseFlag(t *testing.T) {
	includeMixedCase = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemO\ndeMo\ndeMO\ndEmo\ndEmO\ndEMo\ndEMO\nDemo\nDemO\nDeMo\nDeMO\nDEmo\nDEmO\nDEMo\nDEMO"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestYearSeparatorsFlag(t *testing.T) {
	includeYear = true
	includeYearSeparators = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo2024\ndemo!2024\ndemo@2024\ndemo#2024\ndemo$2024\ndemo%2024\ndemo+2024\ndemo?2024\ndemo=2024\ndemo*2024"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestMaskFlag(t *testing.T) {
	includeMask = "%w%d"

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo0\ndemo1\ndemo2\ndemo3\ndemo4\ndemo5\ndemo6\ndemo7\ndemo8\ndemo9"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(cleanup)
}

func TestNumbersFlag(t *testing.T) {
	includeNumbers = true

	// Generate 20 sets of random numbers
	var wordlists [][]string
	for i := 0; i < 20; i++ {
		wordlist := generateCompanyPasslist("demo")
		wordlists = append(wordlists, wordlist)
	}

	// Check that all sets of random numbers are different
	for i := 0; i < len(wordlists)-1; i++ {
		for j := i + 1; j < len(wordlists); j++ {
			if reflect.DeepEqual(wordlists[i], wordlists[j]) {
				t.Errorf("Expected different sets of random numbers, but got %q and %q", wordlists[i], wordlists[j])
			}
		}
	}

	t.Cleanup(cleanup)
}

func TestNumbersRangeFlag(t *testing.T) {
	includeNumbers = true
	includeNumbersRange = 10

	// Generate 20 sets of random numbers
	var wordlists [][]string
	for i := 0; i < 20; i++ {
		wordlist := generateCompanyPasslist("demo")
		wordlists = append(wordlists, wordlist)
	}

	// -1 because we don't count the first demo pass
	if len(wordlists[0])-1 != 10 {
		t.Errorf("Expected 10 words, got %d", len(wordlists[0]))
	}

	// Check that all sets of random numbers are different
	for i := 0; i < len(wordlists)-1; i++ {
		for j := i + 1; j < len(wordlists); j++ {
			if reflect.DeepEqual(wordlists[i], wordlists[j]) {
				t.Errorf("Expected different sets of random numbers, but got %q and %q", wordlists[i], wordlists[j])
			}
		}
	}

	t.Cleanup(cleanup)
}
