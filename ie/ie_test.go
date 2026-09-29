package ie_test

import (
	"testing"

	"github.com/brazilian-utils/go/ie"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		value    string
		state    string
		expected bool
	}{
		{"123456789012", "SP", true},
		{"123.456.789.012", "SP", true},
		{"P123456789012", "SP", true},
		{"12345678", "RJ", true},
		{"1234567", "RJ", false},
		{"123456789", "MG", false}, // MG needs 13
		{"", "SP", false},
		{"123456789012", "ZZ", false},
	}
	for _, table := range tests {
		res := ie.IsValid(ie.IsValidParams{Value: table.value, State: table.state})
		if res != table.expected {
			t.Errorf("Failing for %v/%v \t Expected: %v | Received: %v", table.value, table.state, table.expected, res)
		}
	}
}
