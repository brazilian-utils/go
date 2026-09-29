package cfop_test

import (
	"testing"

	"github.com/brazilian-utils/go/cfop"
)

func TestGet(t *testing.T) {
	c := cfop.Get("5102")
	if c == nil || c.Description == "" {
		t.Fatalf("Expected non-nil result for 5102, got %+v", c)
	}

	masked := cfop.Get("5.102")
	if masked == nil || masked.Code != "5102" {
		t.Errorf("Expected masked lookup to resolve to 5102, got %+v", masked)
	}

	if cfop.Get("1100") != nil {
		t.Error("Expected nil for group heading 1100")
	}
	if cfop.Get("510") != nil {
		t.Error("Expected nil for too-short code")
	}
	if cfop.Get("") != nil {
		t.Error("Expected nil for empty code")
	}
}

var isValidTests = []struct {
	input    interface{}
	expected bool
}{
	{"5102", true},
	{"5.102", true},
	{"7504", true},
	{"0000", false},
	{"1100", false},
	{"510", false},
	{"abcd", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cfop.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"5.102", "5102"},
	{"5102", "5102"},
	{"5?ABC.102abc", "5102"},
	{"", ""},
	{"5102999", "5102"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cfop.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
