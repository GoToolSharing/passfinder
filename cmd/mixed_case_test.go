package cmd

import (
	"strings"
	"testing"
)

func TestMixedCaseFlag(t *testing.T) {
	includeMixedCase = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemO\ndeMo\ndeMO\ndEmo\ndEmO\ndEMo\ndEMO\nDemo\nDemO\nDeMo\nDeMO\nDEmo\nDEmO\nDEMo\nDEMO"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}
