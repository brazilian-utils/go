package legalnature_test

import (
	"testing"

	"github.com/brazilian-utils/brutils-go/legalnature"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"2062", true},
	{"206-2", true},
	{"1015", true},
	{"101-5", true},
	{"5002", true},
	{"9999", false},
	{"0000", false},
	{"206", false},
	{"20622", false},
	{"", false},
	{"abcd", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := legalnature.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var getDescriptionTests = []struct {
	input    string
	expected string
}{
	{"2062", "Sociedade Empresária Limitada"},
	{"101-5", "Órgão Público do Poder Executivo Federal"},
	{"5002", "Organização Internacional e Outras Instituições Extraterritoriais"},
	{"0000", ""},
	{"9999", ""},
	{"", ""},
}

func TestGetDescription(t *testing.T) {
	for _, table := range getDescriptionTests {
		if res := legalnature.GetDescription(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestListAll(t *testing.T) {
	all := legalnature.ListAll()
	if len(all) == 0 {
		t.Error("Expected non-empty map from ListAll")
	}
	// Verify it's a copy by modifying it
	all["9999"] = "test"
	if legalnature.IsValid("9999") {
		t.Error("ListAll should return a copy, not the original map")
	}
}
