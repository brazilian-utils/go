package cnh_test

import (
	"testing"

	"github.com/brazilian-utils/go/cnh"
)

var tables = []struct {
	input    string
	expected bool
}{
	// Valid CNHs
	{"98765432100", true},
	{"987654321-00", true},

	// Blacklisted sequences
	{"00000000000", false},
	{"11111111111", false},
	{"22222222222", false},
	{"33333333333", false},
	{"44444444444", false},
	{"55555555555", false},
	{"66666666666", false},
	{"77777777777", false},
	{"88888888888", false},
	{"99999999999", false},

	// Invalid length
	{"1234567890", false},
	{"123456789012", false},
	{"", false},

	// Invalid characters
	{"1234567890a", false},
	{"A2C45678901", false},

	// Invalid check digits
	{"12345678901", false},
	{"12345678999", false},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := cnh.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
