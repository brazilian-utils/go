package cns_test

import (
	"testing"

	"github.com/brazilian-utils/go/cns"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"123456789010001", "123 4567 8901 0001"},
	{"898 0000 0004 3208", "898 0000 0004 3208"},
	{"1234", "123 4"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cns.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"123456789010000", true},
	{"123 4567 8901 0000", true},
	{"898000000043208", true},
	{"700000000000005", true},
	{"123456789010001", false},
	{"312345678901234", false},
	{"12345678901", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cns.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"123 4567 8901 0000", "123456789010000"},
	{"123456789010000", "123456789010000"},
	{"123.?ABC4567 8901-0000abc", "123456789010000"},
	{"", ""},
	{"123456789010000999", "123456789010000"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cns.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
