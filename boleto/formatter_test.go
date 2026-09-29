package boleto_test

import (
	"testing"

	"github.com/brazilian-utils/go/boleto"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"10491443385511900000200000000141325230000093423", "10491.44338 55119.000002 00000.000141 3 25230000093423"},
	{"104914", "10491.4"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := boleto.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"10491.44338 55119.000002 00000.000141 3 25230000093423", "10491443385511900000200000000141325230000093423"},
	{"84610000000-5 24610029110-2 00546033900-4 69589506108-0", "846100000005246100291102005460339004695895061080"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := boleto.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 20; i++ {
		generated := boleto.Generate()
		if len(generated) != 47 {
			t.Fatalf("Expected 47 digits, got %d: %v", len(generated), generated)
		}
		if !boleto.IsValid(generated) {
			t.Fatalf("Generated invalid boleto on iteration %d: %v", i, generated)
		}
	}
}

func TestGetInfo(t *testing.T) {
	info := boleto.GetInfo("60757135008571297205909229083648919190488862573")
	if info == nil {
		t.Fatal("Expected non-nil info for a valid boleto")
	}
	if info.BankCode == "" {
		t.Error("Expected a non-empty bank code")
	}

	if boleto.GetInfo("invalid") != nil {
		t.Error("Expected nil for an invalid boleto")
	}
}
