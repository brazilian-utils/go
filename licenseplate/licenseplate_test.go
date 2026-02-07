package licenseplate_test

import (
	"testing"

	"github.com/brazilian-utils/go/licenseplate"
)

var isValidTests = []struct {
	input    string
	pType    string
	expected bool
}{
	// Old format
	{"ABC1234", "", true},
	{"abc1234", "", true},
	{"ABC1234", "old_format", true},
	{"ABC1234", "mercosul", false},

	// Mercosul
	{"ABC1D23", "", true},
	{"abc1d23", "", true},
	{"ABC1D23", "mercosul", true},
	{"ABC1D23", "old_format", false},

	// Invalid
	{"ABCD123", "", false},
	{"AB12345", "", false},
	{"ABC123", "", false},
	{"ABC12345", "", false},
	{"", "", false},
	{"ABC-1234", "", false}, // dash not stripped by IsValid
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := licenseplate.IsValid(table.input, table.pType); res != table.expected {
			t.Errorf("Failing for %v type=%v \t Expected: %v | Received: %v", table.input, table.pType, table.expected, res)
		}
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"ABC1234", "ABC-1234"},
	{"abc1234", "ABC-1234"},
	{"abc1e34", "ABC1E34"},
	{"ABC1E34", "ABC1E34"},
	{"ABC123", ""},
	{"ABCD123", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := licenseplate.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var getFormatTests = []struct {
	input    string
	expected string
}{
	{"ABC1234", "LLLNNNN"},
	{"abc1234", "LLLNNNN"},
	{"ABC1D23", "LLLNLNN"},
	{"abc1d23", "LLLNLNN"},
	{"ABCD123", ""},
}

func TestGetFormat(t *testing.T) {
	for _, table := range getFormatTests {
		if res := licenseplate.GetFormat(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var convertTests = []struct {
	input    string
	expected string
}{
	{"ABC4567", "ABC4F67"},
	{"ABC4067", "ABC4A67"},
	{"ABC4967", "ABC4J67"},
	{"abc1234", "ABC1C34"},
	{"ABC1D23", ""}, // already Mercosul, not old format
	{"invalid", ""},
}

func TestConvertToMercosul(t *testing.T) {
	for _, table := range convertTests {
		if res := licenseplate.ConvertToMercosul(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerateMercosul(t *testing.T) {
	plate := licenseplate.Generate("LLLNLNN")
	if !licenseplate.IsValid(plate, "mercosul") {
		t.Errorf("Generated invalid Mercosul plate: %v", plate)
	}
}

func TestGenerateOldFormat(t *testing.T) {
	plate := licenseplate.Generate("LLLNNNN")
	if !licenseplate.IsValid(plate, "old_format") {
		t.Errorf("Generated invalid old format plate: %v", plate)
	}
}

func TestGenerateInvalid(t *testing.T) {
	if res := licenseplate.Generate("invalid"); res != "" {
		t.Errorf("Expected empty for invalid format, got %v", res)
	}
}
