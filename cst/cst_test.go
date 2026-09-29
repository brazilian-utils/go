package cst_test

import (
	"testing"

	"github.com/brazilian-utils/go/cst"
)

var isValidTests = []struct {
	input    interface{}
	tax      string
	expected bool
}{
	{"110", "", true},
	{"00", "", true},
	{"07", "", true},
	{"999", "", false},
	{"0-0", "", false},
	{"", "", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		var res bool
		if table.tax == "" {
			res = cst.IsValid(table.input)
		} else {
			res = cst.IsValid(table.input, cst.IsValidOptions{Tax: table.tax})
		}
		if res != table.expected {
			t.Errorf("Failing for %v tax=%v \t Expected: %v | Received: %v", table.input, table.tax, table.expected, res)
		}
	}
}

func TestIsValidWithTax(t *testing.T) {
	if !cst.IsValid("110", cst.IsValidOptions{Tax: "icms"}) {
		t.Error("Expected 110 to be valid for icms")
	}
	if cst.IsValid("110", cst.IsValidOptions{Tax: "ipi"}) {
		t.Error("Expected 110 (3 digits) to be invalid for ipi (2-digit table)")
	}
	if !cst.IsValid("07", cst.IsValidOptions{Tax: "pis"}) {
		t.Error("Expected 07 to be valid for pis")
	}
	if !cst.IsValid("07", cst.IsValidOptions{Tax: "cofins"}) {
		t.Error("Expected 07 to be valid for cofins")
	}
	if cst.IsValid("07", cst.IsValidOptions{Tax: "ipi"}) {
		t.Error("Expected 07 to be invalid for ipi")
	}
}
