package cnh_test

import (
	"testing"

	"github.com/brazilian-utils/go/cnh"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"00000000119", "000000001-19"},
	{"000.000.001-19", "000000001-19"},
	{"0000000011", "000000001-1"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cnh.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"000000001-19", "00000000119"},
	{"00000000119", "00000000119"},
	{"000.abc000001-19", "00000000119"},
	{"", ""},
	{"00000000119123", "00000000119"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cnh.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 50; i++ {
		generated := cnh.Generate()
		if len(generated) != 11 {
			t.Fatalf("Expected 11 digits, got %d: %v", len(generated), generated)
		}
		if !cnh.IsValid(generated) {
			t.Fatalf("Generated invalid CNH on iteration %d: %v", i, generated)
		}
	}
}
