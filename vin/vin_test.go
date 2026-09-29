package vin_test

import (
	"testing"

	"github.com/brazilian-utils/go/vin"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"1HGCM82633A004352", true},
	{"1M8GDM9AXKP042788", true},
	{"1m8gdm9axkp042788", true},
	{"1HGCM82633A004353", false},
	{"1HGCM8263IA004352", false},
	{"1HGCM82633A00435", false},
	{"00000000000000000", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := vin.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
