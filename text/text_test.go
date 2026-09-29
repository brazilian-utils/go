package text_test

import (
	"testing"

	"github.com/brazilian-utils/go/text"
)

var capitalizeTests = []struct {
	input    string
	expected string
}{
	{"esponja vegetal", "Esponja Vegetal"},
	{"JOAQUIM JOSÉ", "Joaquim José"},
	{"fulano de tal", "Fulano de Tal"},
	{"esponja de    aço 60G", "Esponja de Aço 60g"},
	{"de", "De"},
	{"", ""},
	// Additional, description-derived cases (not in the shared contract's
	// concrete case list, but documented behavior).
	{"empresa ltda", "Empresa LTDA"},
	{"santana/rs", "Santana/RS"},
}

func TestCapitalize(t *testing.T) {
	for _, table := range capitalizeTests {
		if res := text.Capitalize(table.input); res != table.expected {
			t.Errorf("Failing for %q \t Expected: %q | Received: %q", table.input, table.expected, res)
		}
	}
}

var removeAccentsTests = []struct {
	input    string
	expected string
}{
	{"São Paulo", "Sao Paulo"},
	{"Açaí", "Acai"},
	{"Piauí", "Piaui"},
	{"Brasil", "Brasil"},
	{"São Paulo, SP - 2024!", "Sao Paulo, SP - 2024!"},
	{"", ""},
}

func TestRemoveAccents(t *testing.T) {
	for _, table := range removeAccentsTests {
		if res := text.RemoveAccents(table.input); res != table.expected {
			t.Errorf("Failing for %q \t Expected: %q | Received: %q", table.input, table.expected, res)
		}
	}
}
