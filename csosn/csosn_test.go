package csosn_test

import (
	"testing"

	"github.com/brazilian-utils/go/csosn"
)

var isValidTests = []struct {
	input    interface{}
	expected bool
}{
	{"101", true},
	{"102", true},
	{"999", false},
	{"10", false},
	{"1-01", false},
	{"abc", false},
	{"", false},
	{101, true},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := csosn.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
