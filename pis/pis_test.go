package pis_test

import (
	"testing"

	"github.com/brazilian-utils/go/pis"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"12345678900", true},
	{"52555121685", true},

	// Invalid
	{"12345678901", false},
	{"1234567890", false},
	{"123456789012", false},
	{"", false},
	{"8217853746a", false},
	{"123.45678.90-9", false}, // formatted, not plain digits
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := pis.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"12345678900", "123.45678.90-0"},
	{"52555121685", "525.55121.68-5"},
	{"invalid", ""},
	{"1234567890", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := pis.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 100; i++ {
		generated := pis.Generate()
		if len(generated) != 11 {
			t.Fatalf("Expected 11 digits, got %d: %v", len(generated), generated)
		}
		if !pis.IsValid(generated) {
			t.Fatalf("Generated invalid PIS on iteration %d: %v", i, generated)
		}
	}
}
