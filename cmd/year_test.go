package cmd

import (
	"strings"
	"testing"
)

func TestYearFlag(t *testing.T) {
	includeYear = 0

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo2024"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestYearRangeFlag(t *testing.T) {
	includeYear = 3

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo2021\ndemo2022\ndemo2023\ndemo2024\ndemo2025\ndemo2026\ndemo2027"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestShortYearFlag(t *testing.T) {
	includeShortYear = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo24"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestYearSeparatorsFlag(t *testing.T) {
	includeYear = 0
	includeYearSeparators = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo2024\ndemo!2024\ndemo@2024\ndemo#2024\ndemo$2024\ndemo%2024\ndemo+2024\ndemo?2024\ndemo=2024\ndemo*2024"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}
