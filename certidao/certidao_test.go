package certidao_test

import (
	"testing"

	"github.com/brazilian-utils/go/certidao"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"10453901552013100012021000012321", "104539 01 55 2013 1 00012 021 0000123 21"},
	{"104539.01.55.2013.1.00012.021.0000123-21", "104539 01 55 2013 1 00012 021 0000123 21"},
	{"10453901", "104539 01"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := certidao.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"10453901552013100012021000012321", true},
	{"104539 01 55 2013 1 00012 021 0000123 21", true},
	{"094003 01 55 2011 1 00110 002 0051917 43", true},
	{"10453901552013100012021000012322", false},
	{"104539015520131000120210000123", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := certidao.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGetInfo(t *testing.T) {
	info := certidao.GetInfo("104539 01 55 2013 1 00012 021 0000123 21")
	if info == nil {
		t.Fatal("Expected non-nil info")
	}
	if info.RegistryCNS != "104539" || info.Acervo != "01" || info.Service != "55" ||
		info.Year != 2013 || info.Type != "birth" || info.TypeCode != 1 ||
		info.Book != "00012" || info.Page != "021" || info.Term != "0000123" || info.CheckDigits != "21" {
		t.Errorf("Unexpected info: %+v", info)
	}

	info2 := certidao.GetInfo("094300 01 55 2010 1 00020 112 0000120-87")
	if info2 == nil {
		t.Fatal("Expected non-nil info")
	}
	if info2.RegistryCNS != "094300" || info2.Book != "00020" || info2.Page != "112" || info2.Term != "0000120" || info2.CheckDigits != "87" {
		t.Errorf("Unexpected info: %+v", info2)
	}

	if certidao.GetInfo("10453901552013100012021000012322") != nil {
		t.Error("Expected nil for wrong check digits")
	}
	if certidao.GetInfo("not-a-matricula") != nil {
		t.Error("Expected nil for garbage input")
	}
	if certidao.GetInfo("") != nil {
		t.Error("Expected nil for empty input")
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"104539 01 55 2013 1 00012 021 0000123 21", "10453901552013100012021000012321"},
	{"104539.01.55.2013.1.00012.021.0000123-21abc", "10453901552013100012021000012321"},
	{"10453901552013100012021000012321999", "10453901552013100012021000012321"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := certidao.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
