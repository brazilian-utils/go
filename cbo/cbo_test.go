package cbo_test

import (
	"testing"

	"github.com/brazilian-utils/go/cbo"
)

func TestGet(t *testing.T) {
	c := cbo.Get("212405")
	if c == nil || c.Description != "Analista de desenvolvimento de sistemas" {
		t.Errorf("Unexpected result for 212405: %+v", c)
	}

	masked := cbo.Get("2124-05")
	if masked == nil || masked.Code != "212405" {
		t.Errorf("Unexpected result for masked lookup: %+v", masked)
	}

	leadingZero := cbo.Get("0102-05")
	if leadingZero == nil || leadingZero.Code != "010205" || leadingZero.Description != "Oficial da aeronáutica" {
		t.Errorf("Unexpected result for leading-zero lookup: %+v", leadingZero)
	}

	if cbo.Get("223150") != nil {
		t.Error("Expected nil for nonexistent code")
	}
	if cbo.Get("000000") != nil {
		t.Error("Expected nil for 000000")
	}
	if cbo.Get("") != nil {
		t.Error("Expected nil for empty code")
	}
}

var isValidTests = []struct {
	input    interface{}
	expected bool
}{
	{"212405", true},
	{"2124-05", true},
	{"010205", true},
	{"223150", false},
	{"000000", false},
	{"21240", false},
	{"abcdef", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cbo.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"2124-05", "212405"},
	{"212405", "212405"},
	{"21?ABC24-05abc", "212405"},
	{"", ""},
	{"212405999", "212405"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cbo.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
