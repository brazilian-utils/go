package renavam_test

import (
	"testing"

	"github.com/brazilian-utils/go/renavam"
)

var tables = []struct {
	input    string
	expected bool
}{
	{"86769597308", true},

	// Invalid
	{"12345678901", false},
	{"1234567890a", false},
	{"12345678 901", false},
	{"12345678", false},
	{"", false},
	{"00000000000", false},
	{"11111111111", false},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := renavam.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 50; i++ {
		generated := renavam.Generate()
		if len(generated) != 11 {
			t.Fatalf("Expected 11 digits, got %d: %v", len(generated), generated)
		}
		if !renavam.IsValid(generated) {
			t.Fatalf("Generated invalid RENAVAM on iteration %d: %v", i, generated)
		}
	}
}
