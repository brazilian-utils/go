package phone_test

import (
	"testing"

	"github.com/brazilian-utils/go/phone"
)

var isValidTests = []struct {
	input    string
	pType    string
	expected bool
}{
	// Mobile (11 digits, DDD + 9 + 8 digits)
	{"11994029275", "", true},
	{"11994029275", "mobile", true},
	{"11994029275", "landline", false},

	// Landline (10 digits, DDD + 2-5 + 7 digits)
	{"1635014415", "", true},
	{"1635014415", "landline", true},
	{"1635014415", "mobile", false},

	// Invalid
	{"123", "", false},
	{"", "", false},
	{"0035014415", "", false},  // DDD starts with 0
	{"1005014415", "", false},  // DDD second digit is 0
	{"1165014415", "", false},  // Landline prefix 6 invalid
	{"11994029275a", "", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := phone.IsValid(table.input, table.pType); res != table.expected {
			t.Errorf("Failing for %v type=%v \t Expected: %v | Received: %v", table.input, table.pType, table.expected, res)
		}
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"11994029275", "(11)99402-9275"},
	{"1635014415", "(16)3501-4415"},
	{"333333", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := phone.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var removeSymbolsTests = []struct {
	input    string
	expected string
}{
	{"(11)99402-9275", "11994029275"},
	{"(11) 99402-9275", "11994029275"},
	{"+55 11 99402-9275", "5511994029275"},
}

func TestRemoveSymbols(t *testing.T) {
	for _, table := range removeSymbolsTests {
		if res := phone.RemoveSymbols(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var removeCodeTests = []struct {
	input    string
	expected string
}{
	{"5511994029275", "11994029275"},
	{"1635014415", "1635014415"},
	{"+5511994029275", "+11994029275"},
}

func TestRemoveInternationalDialingCode(t *testing.T) {
	for _, table := range removeCodeTests {
		if res := phone.RemoveInternationalDialingCode(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerateMobile(t *testing.T) {
	generated := phone.Generate("mobile")
	if !phone.IsValid(generated, "mobile") {
		t.Errorf("Generated invalid mobile number: %v", generated)
	}
}

func TestGenerateLandline(t *testing.T) {
	generated := phone.Generate("landline")
	if !phone.IsValid(generated, "landline") {
		t.Errorf("Generated invalid landline number: %v", generated)
	}
}

func TestGenerateRandom(t *testing.T) {
	generated := phone.Generate("")
	if !phone.IsValid(generated, "") {
		t.Errorf("Generated invalid phone number: %v", generated)
	}
}
