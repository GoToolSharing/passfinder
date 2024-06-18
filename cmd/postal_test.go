package cmd

import (
	"strings"
	"testing"
)

func TestPostalFlag(t *testing.T) {
	includePostal = 75000

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo75000"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}
