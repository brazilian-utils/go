package pixkey_test

import (
	"testing"

	"github.com/brazilian-utils/go/pixkey"
)

func TestGetInfo(t *testing.T) {
	tests := []struct {
		input        string
		expectedType string
		expectedVal  string
	}{
		{"123.456.789-09", "cpf", "12345678909"},
		{"00.038.166/0001-05", "cnpj", "00038166000105"},
		{"fulano_da_silva.recebedor@example.com", "email", "fulano_da_silva.recebedor@example.com"},
		{"+5561912345678", "phone", "+5561912345678"},
		{"71c7d9be-4b85-4e43-9f1c-1f3b8b4e9a2d", "evp", "71c7d9be-4b85-4e43-9f1c-1f3b8b4e9a2d"},
	}
	for _, table := range tests {
		info := pixkey.GetInfo(table.input)
		if info == nil {
			t.Errorf("Failing for %v \t Expected non-nil info", table.input)
			continue
		}
		if info.Type != table.expectedType || info.Value != table.expectedVal {
			t.Errorf("Failing for %v \t Expected: %v/%v | Received: %v/%v", table.input, table.expectedType, table.expectedVal, info.Type, info.Value)
		}
	}

	if pixkey.GetInfo("11257245286") != nil {
		t.Error("Expected nil for invalid CPF")
	}
	if pixkey.GetInfo("") != nil {
		t.Error("Expected nil for empty input")
	}
}

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"123.456.789-09", true},
	{"00.038.166/0001-05", true},
	{"fulano_da_silva.recebedor@example.com", true},
	{"+5561912345678", true},
	{"71c7d9be-4b85-4e43-9f1c-1f3b8b4e9a2d", true},
	{"11257245286", false},
	{"fulano@example", false},
	{"1130000000", false},
	{"", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := pixkey.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
