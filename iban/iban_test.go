package iban_test

import (
	"testing"

	"github.com/brazilian-utils/go/iban"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"BR1500000000000010932840814P2", "BR15 0000 0000 0000 1093 2840 814P 2"},
	{"br1500000000000010932840814p2", "BR15 0000 0000 0000 1093 2840 814P 2"},
	{"BR150", "BR15 0"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := iban.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"BR1500000000000010932840814P2", true},
	{"BR15 0000 0000 0000 1093 2840 814P 2", true},
	{"br1500000000000010932840814p2", true},
	{"BR1500000000000010932840814P3", false},
	{"DE89370400440532013000", false},
	{"BR15000000000000109328408", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := iban.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGetInfo(t *testing.T) {
	info := iban.GetInfo("BR1500000000000010932840814P2")
	if info == nil {
		t.Fatal("Expected non-nil info")
	}
	if info.CountryCode != "BR" || info.CheckDigits != "15" || info.BankISPB != "00000000" ||
		info.Branch != "00001" || info.Account != "0932840814" || info.AccountType != "P" || info.Owner != "2" {
		t.Errorf("Unexpected info: %+v", info)
	}

	info2 := iban.GetInfo("BR3860701190000010000012345C1")
	if info2 == nil {
		t.Fatal("Expected non-nil info")
	}
	if info2.CheckDigits != "38" || info2.BankISPB != "60701190" || info2.Branch != "00001" ||
		info2.Account != "0000012345" || info2.AccountType != "C" || info2.Owner != "1" {
		t.Errorf("Unexpected info: %+v", info2)
	}

	if iban.GetInfo("BR1500000000000010932840814P3") != nil {
		t.Error("Expected nil for wrong check")
	}
	if iban.GetInfo("DE89370400440532013000") != nil {
		t.Error("Expected nil for non-Brazilian IBAN")
	}
	if iban.GetInfo("") != nil {
		t.Error("Expected nil for empty input")
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"BR15 0000 0000 0000 1093 2840 814P 2", "BR1500000000000010932840814P2"},
	{"br15-0000.0000/0000 1093 2840 814p-2", "BR1500000000000010932840814P2"},
	{"BR1500000000000010932840814P2EXTRA", "BR1500000000000010932840814P2"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := iban.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
