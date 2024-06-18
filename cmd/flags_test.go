package cmd

import (
	"strings"
	"testing"
)

func testCleanup() {
	includeYearSeparators = false
	includeYear = -1
	includeStartCaps = false
	includeShortYear = -1
	includeEndSpecial = false
	includeMixedCase = false
	includeLeetCode = false
	includeUppercase = false
	includeMask = ""
	includePostal = 0
	includeNumbers = 0
}

func TestWithoutFlag(t *testing.T) {
	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestEndSpecialFlag(t *testing.T) {
	includeEndSpecial = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo!\ndemo@\ndemo#\ndemo$\ndemo%\ndemo+\ndemo?\ndemo=\ndemo*"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestUppercaseFlag(t *testing.T) {
	includeUppercase = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\nDEMO"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestStartCapsFlag(t *testing.T) {
	includeStartCaps = true

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\nDemo"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}

func TestMaskFlag(t *testing.T) {
	includeMask = "%w%d"

	wordlist := generateCompanyPasslist("demo")

	got := strings.Join(wordlist, "\n")

	expected := "demo\ndemo0\ndemo1\ndemo2\ndemo3\ndemo4\ndemo5\ndemo6\ndemo7\ndemo8\ndemo9"

	if got != expected {
		t.Errorf("Expected %q, but got %q", expected, got)
	}

	t.Cleanup(testCleanup)
}
