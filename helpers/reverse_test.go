package helpers_test

import (
	"testing"

	"github.com/brazilian-utils/brutils-go/helpers"
)

var reverseTestValues = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"Arara", "ararA"},
	{"The quick brown fox jumped over the lazy dog", "god yzal eht revo depmuj xof nworb kciuq ehT"},
}

func TestFormat(t *testing.T) {
	for _, table := range reverseTestValues {
		if res := helpers.Reverse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
