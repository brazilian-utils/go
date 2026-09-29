package helpers_test

import (
	"testing"

	"github.com/brazilian-utils/go/helpers"
)

var numberToPortugueseTests = []struct {
	input    int64
	expected string
}{
	{0, "zero"},
	{1, "um"},
	{14, "quatorze"},
	{100, "cem"},
	{1000, "mil"},
	{1234, "mil duzentos e trinta e quatro"},
	{-42, "menos quarenta e dois"}, // was a panic (index out of range [-1])
	{-1000, "menos mil"},
}

func TestNumberToPortuguese(t *testing.T) {
	for _, table := range numberToPortugueseTests {
		if res := helpers.NumberToPortuguese(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
