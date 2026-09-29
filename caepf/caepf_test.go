package caepf_test

import (
	"testing"

	"github.com/brazilian-utils/go/caepf"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"29311861000184", "293.118.610/001-84"},
	{"293.118.610/001-84", "293.118.610/001-84"},
	{"2931", "293.1"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := caepf.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"293.118.610/001-84", true},
	{"29311861000184", true},
	{"41142260000101", true},
	{"29311861000185", false},
	{"00000000000000", false},
	{"1234567890", false},
	{"abc.118.610/001-84", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := caepf.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"293.118.610/001-84", "29311861000184"},
	{"29311861000184", "29311861000184"},
	{"293.?ABC118.610/001-84abc", "29311861000184"},
	{"", ""},
	{"29311861000184999", "29311861000184"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := caepf.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
