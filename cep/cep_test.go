package cep_test

import (
	"testing"

	"github.com/brazilian-utils/go/cep"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"01001000", true},
	{"12345678", true},
	{"00000000", true},
	{"99999999", true},
	{"1234567", false},
	{"123456789", false},
	{"1234567a", false},
	{"abcdefgh", false},
	{"", false},
	{"01001-000", false},
	{"01.001-000", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := cep.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"01001000", "01001-000"},
	{"12345678", "12345-678"},
	{"99999999", "99999-999"},
	{"1234567", ""},
	{"123456789", ""},
	{"abcdefgh", ""},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cep.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var removeSymbolsTests = []struct {
	input    string
	expected string
}{
	{"91906292", "91906292"},
	{"91906-292", "91906292"},
	{"", ""},
	{"   ", "   "},
	{"abc", "abc"},
}

func TestRemoveSymbols(t *testing.T) {
	for _, table := range removeSymbolsTests {
		if res := cep.RemoveSymbols(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"01001-000", "01001000"},
	{"01001000", "01001000"},
	{"a0.10cr01?00#ab0", "01001000"},
	{"", ""},
	{"01001000123", "01001000"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cep.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	generated := cep.Generate()
	if len(generated) != 8 {
		t.Errorf("Expected generated CEP to have 8 characters, got %d", len(generated))
	}
	if !cep.IsValid(generated) {
		t.Errorf("Expected generated CEP to be valid, got %v", generated)
	}
}
