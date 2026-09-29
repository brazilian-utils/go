package state_test

import (
	"testing"

	"github.com/brazilian-utils/go/state"
)

func TestGetByIbgeCode(t *testing.T) {
	sp := state.GetByIbgeCode("35")
	if sp == nil || sp.Code != "SP" || sp.Name != "São Paulo" || sp.RegionCode != "SE" {
		t.Errorf("Unexpected result for IBGE code 35: %+v", sp)
	}

	spNum := state.GetByIbgeCode(35)
	if spNum == nil || spNum.Code != "SP" {
		t.Errorf("Unexpected result for numeric IBGE code 35: %+v", spNum)
	}

	df := state.GetByIbgeCode("53")
	if df == nil || df.Code != "DF" {
		t.Errorf("Unexpected result for IBGE code 53: %+v", df)
	}

	if state.GetByIbgeCode("00") != nil {
		t.Error("Expected nil for IBGE code 00")
	}
	if state.GetByIbgeCode("99") != nil {
		t.Error("Expected nil for IBGE code 99")
	}
	if state.GetByIbgeCode("") != nil {
		t.Error("Expected nil for empty code")
	}
}

func TestGetCodeByName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"São Paulo", "SP"},
		{"Sao Paulo", "SP"},
		{"sao paulo", "SP"},
		{"  São Paulo  ", "SP"},
		{"Rio Grande do Norte", "RN"},
		{"Distrito Federal", "DF"},
	}
	for _, table := range tests {
		res := state.GetCodeByName(table.input)
		if res == nil || *res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}

	if state.GetCodeByName("Neverland") != nil {
		t.Error("Expected nil for Neverland")
	}
	if state.GetCodeByName("") != nil {
		t.Error("Expected nil for empty name")
	}
}

func TestGetNameByCode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SP", "São Paulo"},
		{"sp", "São Paulo"},
		{"  RJ  ", "Rio de Janeiro"},
		{"DF", "Distrito Federal"},
	}
	for _, table := range tests {
		res := state.GetNameByCode(table.input)
		if res == nil || *res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}

	if state.GetNameByCode("ZZ") != nil {
		t.Error("Expected nil for ZZ")
	}
	if state.GetNameByCode("") != nil {
		t.Error("Expected nil for empty code")
	}
}

func TestGetTimezone(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"SP", "America/Sao_Paulo"},
		{"AM", "America/Manaus"},
		{"AC", "America/Rio_Branco"},
		{"BA", "America/Bahia"},
		{"sp", "America/Sao_Paulo"},
		{"  SP  ", "America/Sao_Paulo"},
	}
	for _, table := range tests {
		res := state.GetTimezone(table.input)
		if res == nil || *res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}

	if state.GetTimezone("ZZ") != nil {
		t.Error("Expected nil for ZZ")
	}
	if state.GetTimezone("") != nil {
		t.Error("Expected nil for empty code")
	}
}

func TestList(t *testing.T) {
	all := state.List()
	if len(all) != 27 {
		t.Errorf("Expected 27 states, got %d", len(all))
	}
	for i := 1; i < len(all); i++ {
		if all[i].Name < all[i-1].Name {
			t.Errorf("Expected states sorted by name, got %v before %v", all[i].Name, all[i-1].Name)
		}
	}
}
