package cei_test

import (
	"testing"

	"github.com/brazilian-utils/go/cei"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"277297118187", "27.729.71181/87"},
	{"11.583.00249/85", "11.583.00249/85"},
	{"27729", "27.729"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cei.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"11.583.00249/85", true},
	{"115830024985", true},
	{"277297118187", true},
	{"115830024984", false},
	{"000000000000", false},
	{"1234567890", false},
	{"aa.583.00249/85", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cei.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"27.729.71181/87", "277297118187"},
	{"277297118187", "277297118187"},
	{"27.?ABC729.71181/87abc", "277297118187"},
	{"", ""},
	{"277297118187999", "277297118187"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cei.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
