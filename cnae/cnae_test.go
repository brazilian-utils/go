package cnae_test

import (
	"testing"

	"github.com/brazilian-utils/go/cnae"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"6201501", "6201-5/01"},
	{"6201-5/01", "6201-5/01"},
	{"62015", "6201-5"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cnae.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGet(t *testing.T) {
	c := cnae.Get("6201501")
	if c == nil || c.Description != "DESENVOLVIMENTO DE PROGRAMAS DE COMPUTADOR SOB ENCOMENDA" {
		t.Errorf("Unexpected result for 6201501: %+v", c)
	}

	masked := cnae.Get("6201-5/01")
	if masked == nil || masked.Code != "6201501" {
		t.Errorf("Unexpected result for masked lookup: %+v", masked)
	}

	leadingZero := cnae.Get("0111-3/01")
	if leadingZero == nil || leadingZero.Code != "0111301" || leadingZero.Description != "CULTIVO DE ARROZ" {
		t.Errorf("Unexpected result for leading-zero lookup: %+v", leadingZero)
	}

	if cnae.Get("0000000") != nil {
		t.Error("Expected nil for 0000000")
	}
	if cnae.Get("620150") != nil {
		t.Error("Expected nil for too-short code")
	}
	if cnae.Get("") != nil {
		t.Error("Expected nil for empty code")
	}
}

var isValidTests = []struct {
	input    interface{}
	expected bool
}{
	{"6201501", true},
	{"6201-5/01", true},
	{"0111301", true},
	{"0000000", false},
	{"620150", false},
	{"62015011", false},
	{"abcdefg", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cnae.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"6201-5/01", "6201501"},
	{"6201501", "6201501"},
	{"62?ABC01-5/01abc", "6201501"},
	{"", ""},
	{"6201501999", "6201501"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cnae.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
