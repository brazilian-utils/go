package ncm_test

import (
	"testing"

	"github.com/brazilian-utils/go/ncm"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"84713012", "8471.30.12"},
	{"8471.30.12", "8471.30.12"},
	{"84713", "8471.3"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := ncm.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    interface{}
	expected bool
}{
	{"22030000", true},
	{"2203.00.00", true},
	{"01012100", true},
	{"12345678", false},
	{"2203000", false},
	{"abcdefgh", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := ncm.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"8471.30.12", "84713012"},
	{"84713012", "84713012"},
	{"84?ABC71.30.12abc", "84713012"},
	{"", ""},
	{"84713012999", "84713012"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := ncm.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
