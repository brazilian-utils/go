package creditcard_test

import (
	"testing"

	"github.com/brazilian-utils/go/creditcard"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"4111111111111111", true},
	{"5555555555554444", true},
	{"378282246310005", true},
	{"4111 1111 1111 1111", true},
	{"4111111111111112", false},
	{"0000000000000000", false},
	{"60110000000", false},
	{"abcdabcdabcd", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := creditcard.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
