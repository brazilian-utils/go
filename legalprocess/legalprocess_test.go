package legalprocess_test

import (
	"testing"
	"time"

	"github.com/brazilian-utils/brutils-go/legalprocess"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	// Valid IDs from Python docstring
	{"68476506020233030000", true},
	{"51808233620233030000", true},

	// Formatted valid IDs
	{"6847650-60.2023.3.03.0000", true},

	// Invalid
	{"123", false},
	{"", false},
	{"12345678901234567890", false}, // valid length but likely bad checksum/orgao
	{"0000000000000000000a", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := legalprocess.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"12345678901234567890", "1234567-89.0123.4.56.7890"},
	{"98765432109876543210", "9876543-21.0987.6.54.3210"},
	{"68476506020233030000", "6847650-60.2023.3.03.0000"},
	{"123", ""},
	{"1234567890123456789a", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := legalprocess.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	year := time.Now().Year()
	for orgao := 1; orgao <= 9; orgao++ {
		generated := legalprocess.Generate(year, orgao)
		if generated == "" {
			t.Errorf("Expected non-empty result for orgao %d", orgao)
			continue
		}
		if len(generated) != 20 {
			t.Errorf("Expected 20 digits for orgao %d, got %d: %v", orgao, len(generated), generated)
			continue
		}
		if !legalprocess.IsValid(generated) {
			t.Errorf("Generated invalid ID for orgao %d: %v", orgao, generated)
		}
	}
}

func TestGenerateInvalidArgs(t *testing.T) {
	if res := legalprocess.Generate(2020, 5); res != "" {
		t.Errorf("Expected empty for past year, got %v", res)
	}
	if res := legalprocess.Generate(time.Now().Year(), 0); res != "" {
		t.Errorf("Expected empty for orgao 0, got %v", res)
	}
	if res := legalprocess.Generate(time.Now().Year(), 10); res != "" {
		t.Errorf("Expected empty for orgao 10, got %v", res)
	}
}
