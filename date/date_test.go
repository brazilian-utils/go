package date_test

import (
	"testing"
	"time"

	"github.com/brazilian-utils/brutils-go/date"
)

// ConvertDateToText tests

var convertTests = []struct {
	input    string
	expected string
}{
	{"01/01/2024", "Primeiro de Janeiro de dois mil e vinte e quatro"},
	{"15/06/1990", "Quinze de Junho de mil, novecentos e noventa"},
	{"25/12/2000", "Vinte e cinco de Dezembro de dois mil"},
	{"02/03/1985", "Dois de Março de mil, novecentos e oitenta e cinco"},
	{"31/07/2023", "Trinta e um de Julho de dois mil e vinte e três"},
	{"10/10/2010", "Dez de Outubro de dois mil e dez"},

	// Invalid formats
	{"", ""},
	{"2024-01-01", ""},
	{"01-01-2024", ""},
	{"32/01/2024", ""},
	{"01/13/2024", ""},
	{"abc", ""},
	{"00/01/2024", ""},
}

func TestConvertDateToText(t *testing.T) {
	for _, table := range convertTests {
		if res := date.ConvertDateToText(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

// IsHoliday tests

func d(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

var holidayTests = []struct {
	date     time.Time
	uf       string
	expected bool
}{
	// National fixed holidays
	{d(2024, 1, 1), "", true},   // New Year
	{d(2024, 4, 21), "", true},  // Tiradentes
	{d(2024, 5, 1), "", true},   // Labor Day
	{d(2024, 9, 7), "", true},   // Independence
	{d(2024, 10, 12), "", true}, // Nossa Senhora Aparecida
	{d(2024, 11, 2), "", true},  // Finados
	{d(2024, 11, 15), "", true}, // Republic
	{d(2024, 12, 25), "", true}, // Christmas

	// Not a holiday
	{d(2024, 1, 2), "", false},
	{d(2024, 6, 10), "", false},

	// Easter-based movable holidays for 2024 (Easter = March 31)
	{d(2024, 2, 12), "", true}, // Carnival Monday
	{d(2024, 2, 13), "", true}, // Carnival Tuesday
	{d(2024, 3, 29), "", true}, // Good Friday
	{d(2024, 5, 30), "", true}, // Corpus Christi

	// Consciência Negra as national (>= 2024)
	{d(2024, 11, 20), "", true},
	{d(2023, 11, 20), "", false}, // Before 2024, not national

	// State holidays
	{d(2024, 7, 9), "SP", true},   // SP: Revolução Constitucionalista
	{d(2024, 7, 2), "BA", true},   // BA: Independência da Bahia
	{d(2024, 9, 20), "RS", true},  // RS: Revolução Farroupilha
	{d(2024, 4, 23), "RJ", true},  // RJ: São Jorge
	{d(2024, 12, 25), "RJ", true}, // National holiday checked with UF

	// State holiday not in another state
	{d(2024, 7, 9), "RJ", false},
	{d(2024, 3, 2), "SP", false},
}

func TestIsHoliday(t *testing.T) {
	for _, table := range holidayTests {
		result, ok := date.IsHoliday(table.date, table.uf)
		if !ok {
			t.Errorf("Failing for %v uf=%v \t Got ok=false, expected ok=true", table.date.Format("2006-01-02"), table.uf)
			continue
		}
		if result != table.expected {
			t.Errorf("Failing for %v uf=%v \t Expected: %v | Received: %v", table.date.Format("2006-01-02"), table.uf, table.expected, result)
		}
	}
}

func TestIsHolidayInvalidUF(t *testing.T) {
	_, ok := date.IsHoliday(d(2024, 1, 1), "XX")
	if ok {
		t.Error("Expected ok=false for invalid UF")
	}
}
