package phone_test

import (
	"testing"

	"github.com/brazilian-utils/go/phone"
)

var isValidTests = []struct {
	input    string
	kinds    []string
	expected bool
}{
	// Mobile (11 digits, DDD + 9 + 8 digits)
	{"11994029275", nil, true},
	{"11994029275", []string{"mobile"}, true},
	{"11994029275", []string{"landline"}, false},

	// Landline (10 digits, DDD + 2-5 + 7 digits)
	{"1635014415", nil, true},
	{"1635014415", []string{"landline"}, true},
	{"1635014415", []string{"mobile"}, false},

	// Invalid
	{"123", nil, false},
	{"", nil, false},
	{"0035014415", nil, false}, // DDD starts with 0
	{"1005014415", nil, false}, // DDD second digit is 0
	{"1165014415", nil, true},  // Landline prefix 6 valid (Anatel Res. 749/2022 art. 11, I, "a")
	{"1175014415", nil, false}, // Landline prefix 7 invalid
	{"11994029275a", nil, false},

	// From cases/phone.json (a subset checked against reference JS/Python/Ruby)
	{"48976579784", nil, true},
	{"48976579785", nil, true},
	{"00000000000", nil, false},
	{"3147981464", nil, true},
	{"3147981465", nil, true},
	{"0000000000", nil, false},
	{"38942432195", nil, true},
	{"38942432196", nil, true},
	{"48976-5797", nil, false},
	{"31479-8146", nil, false},
	{"38942-4321", nil, false},
	{"   ", nil, false},
	{"abc", nil, false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		var res bool
		if table.kinds == nil {
			res = phone.IsValid(table.input)
		} else {
			res = phone.IsValid(table.input, phone.IsValidOptions{Kinds: table.kinds})
		}
		if res != table.expected {
			t.Errorf("Failing for %v kinds=%v \t Expected: %v | Received: %v", table.input, table.kinds, table.expected, res)
		}
	}
}

func TestIsValidLandline(t *testing.T) {
	landlineTests := []struct {
		input    string
		expected bool
	}{
		{"(11) 3000-0000", true},
		{"1130000000", true},
		{"+55 11 3000-0000", true},
		{"11987654321", false},
		{"113000000", false},
		{"", false},
	}
	for _, table := range landlineTests {
		if res := phone.IsValidLandline(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestIsValidMobile(t *testing.T) {
	mobileTests := []struct {
		input    string
		expected bool
	}{
		{"(11) 98765-4321", true},
		{"+55 11 98765-4321", true},
		{"5511987654321", true},
		{"1130000000", false},
		{"1198765432", false},
		{"", false},
	}
	for _, table := range mobileTests {
		if res := phone.IsValidMobile(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestIsValidService(t *testing.T) {
	serviceTests := []struct {
		input    string
		expected bool
	}{
		{"08001234567", true},
		{"0800 123 4567", true},
		{"40041234", true},
		{"11987654321", false},
		{"0800123456", false},
		{"911", false},
		{"", false},
	}
	for _, table := range serviceTests {
		if res := phone.IsValidService(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"988887777", "98888-7777"},
	{"98888777", "98888-777"},
	{"", ""},
	{"1130000000", "11300-0000"},
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
	{"+5511994029275", "11994029275"}, // "+" is part of the country code, not kept
}

func TestRemoveInternationalDialingCode(t *testing.T) {
	for _, table := range removeCodeTests {
		if res := phone.RemoveInternationalDialingCode(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"(11) 98888-7777", "11988887777"},
	{"98888-7777", "988887777"},
	{"+55 11 98888-7777", "11988887777"},
	{"005511988887777", "11988887777"},
	{"551130000000", "1130000000"},
	{"55988887777", "55988887777"},
	{"0800 123 4567", "08001234567"},
	{"11988887777123", "11988887777"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := phone.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerateMobile(t *testing.T) {
	generated := phone.Generate("mobile")
	if !phone.IsValid(generated, phone.IsValidOptions{Kinds: []string{"mobile"}}) {
		t.Errorf("Generated invalid mobile number: %v", generated)
	}
}

func TestGenerateLandline(t *testing.T) {
	generated := phone.Generate("landline")
	if !phone.IsValid(generated, phone.IsValidOptions{Kinds: []string{"landline"}}) {
		t.Errorf("Generated invalid landline number: %v", generated)
	}
}

func TestGenerateRandom(t *testing.T) {
	generated := phone.Generate("")
	if !phone.IsValid(generated) {
		t.Errorf("Generated invalid phone number: %v", generated)
	}
}
