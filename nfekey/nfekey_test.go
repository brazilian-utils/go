package nfekey_test

import (
	"testing"

	"github.com/brazilian-utils/go/nfekey"
)

var validKey = "35170458716523000119550010000000121000123458"

var formatTests = []struct {
	input    string
	expected string
}{
	{validKey, "3517 0458 7165 2300 0119 5500 1000 0000 1210 0012 3458"},
	{"NFe" + validKey, "3517 0458 7165 2300 0119 5500 1000 0000 1210 0012 3458"},
	{"12345", "1234 5"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := nfekey.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{validKey, true},
	{"3517 0458 7165 2300 0119 5500 1000 0000 1210 0012 3458", true},
	{"NFe" + validKey, true},
	{"35170458716523000119550010000000001000123457", false},
	{"3517045871652300011955001000000012100012345", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := nfekey.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

// tpEmisTests cover every valid tpEmis digit (1..9, unchanged behavior)
// plus the invalid tpEmis=0 case; each is validKey with only the
// tpEmis digit and the trailing check digit changed.
var tpEmisTests = []struct {
	input    string
	expected bool
}{
	{"35170458716523000119550010000000120000123450", false}, // tpEmis=0
	{"35170458716523000119550010000000121000123458", true},  // tpEmis=1
	{"35170458716523000119550010000000122000123456", true},  // tpEmis=2
	{"35170458716523000119550010000000123000123454", true},  // tpEmis=3
	{"35170458716523000119550010000000124000123452", true},  // tpEmis=4
	{"35170458716523000119550010000000125000123450", true},  // tpEmis=5
	{"35170458716523000119550010000000126000123459", true},  // tpEmis=6
	{"35170458716523000119550010000000127000123457", true},  // tpEmis=7
	{"35170458716523000119550010000000128000123455", true},  // tpEmis=8
	{"35170458716523000119550010000000129000123453", true},  // tpEmis=9
}

func TestIsValidTpEmisRange(t *testing.T) {
	for _, table := range tpEmisTests {
		if res := nfekey.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

// cNFTests cover MOC rule B03-10: cNF (the 8-digit numeric code) must
// not be all-repeated, ascending/descending sequential, or equal to
// nNF, but only for models 55/65 -- a model-57 (CT-e) key is unaffected.
var cNFTests = []struct {
	name     string
	input    string
	expected bool
}{
	{"model 55, all-repeated cNF", "35170458716523000119550010000000121111111113", false},
	{"model 55, ascending cNF", "35170458716523000119550010000000121012345671", false},
	{"model 55, descending cNF", "35170458716523000119550010000000121876543215", false},
	{"model 55, cNF == nNF", "35170458716523000119550010000000121000000127", false},
	{"model 57 (CT-e), all-repeated cNF is fine", "35170458716523000119570010000000121111111110", true},
}

func TestIsValidCNFRule(t *testing.T) {
	for _, table := range cNFTests {
		if res := nfekey.IsValid(table.input); res != table.expected {
			t.Errorf("%s: Failing for %v \t Expected: %v | Received: %v", table.name, table.input, table.expected, res)
		}
	}
}

func TestGetInfo(t *testing.T) {
	info := nfekey.GetInfo(validKey)
	if info == nil {
		t.Fatal("Expected non-nil info")
	}
	if info.StateCode != "SP" || info.Year != 2017 || info.Month != 4 ||
		info.TaxID != "58716523000119" || info.Model != "55" || info.Series != 1 ||
		info.Number != 12 || info.EmissionType != 1 || info.Code != "00012345" || info.CheckDigit != 8 {
		t.Errorf("Unexpected info: %+v", info)
	}

	if nfekey.GetInfo("35170458716523000119550010000000121000123459") != nil {
		t.Error("Expected nil for wrong check digit")
	}
	if nfekey.GetInfo("") != nil {
		t.Error("Expected nil for empty input")
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"3517 0458 7165 2300 0119 5500 1000 0000 1210 0012 3458", validKey},
	{"NFe" + validKey, validKey},
	{"3517 0458", "35170458"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := nfekey.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
