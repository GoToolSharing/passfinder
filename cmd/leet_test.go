package cmd

import (
	"strings"
	"testing"
)

func TestLeetFlag(t *testing.T) {
	includeLeetCode = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\nd3m0\nd3mo\ndem0"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}
