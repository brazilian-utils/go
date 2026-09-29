package cno_test

import (
	"testing"

	"github.com/brazilian-utils/go/cno"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"111130137368", "11.113.01373/68"},
	{"11.084.01680/62", "11.084.01680/62"},
	{"11113", "11.113"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cno.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"110840168062", true},
	{"11.084.01680/62", true},
	{"401800097960", true},
	{"110840168063", false},
	{"000000000000", false},
	{"1234567890", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cno.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"11.113.01373/68", "111130137368"},
	{"111130137368", "111130137368"},
	{"11.?ABC113.01373/68abc", "111130137368"},
	{"", ""},
	{"111130137368999", "111130137368"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cno.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
