package cpf_test

import (
	"testing"

	"github.com/brazilian-utils/go/cpf"
)

var formatTestValues = []struct {
	input    string
	expected string
}{
	{"", ""},
	{"9", "9"},
	{"94", "94"},
	{"943", "943"},
	{"9438", "943.8"},
	{"94389", "943.89"},
	{"943895", "943.895"},
	{"9438957", "943.895.7"},
	{"94389575", "943.895.75"},
	{"943895751", "943.895.751"},
	{"9438957510", "943.895.751-0"},
	{"94389575104", "943.895.751-04"},
	{"94389575104000000", "943.895.751-04"},
	{"943.?ABC895.751-04abc", "943.895.751-04"},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTestValues {
		if res := cpf.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
