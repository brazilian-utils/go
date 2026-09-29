package date_test

import (
	"testing"
	"time"

	"github.com/brazilian-utils/go/date"
)

// ConvertDateToText tests

var convertTests = []struct {
	input    string
	expected string
}{
	// Months lowercase: they're common nouns in Portuguese orthography,
	// matching the LOA extenso convention (Lei 14.822/2024).
	{"01/01/2024", "primeiro de janeiro de dois mil e vinte e quatro"},
	{"15/06/1990", "quinze de junho de mil novecentos e noventa"},
	{"25/12/2000", "vinte e cinco de dezembro de dois mil"},
	{"02/03/1985", "dois de março de mil novecentos e oitenta e cinco"},
	{"31/07/2023", "trinta e um de julho de dois mil e vinte e três"},
	{"10/10/2010", "dez de outubro de dois mil e dez"},
	{"2024-12-25", "vinte e cinco de dezembro de dois mil e vinte e quatro"},

	// Invalid formats
	{"", ""},
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
		result := date.IsHoliday(date.IsHolidayOptions{Date: table.date, UF: table.uf})
		if result != table.expected {
			t.Errorf("Failing for %v uf=%v \t Expected: %v | Received: %v", table.date.Format("2006-01-02"), table.uf, table.expected, result)
		}
	}
}

func TestIsHolidayInvalidUF(t *testing.T) {
	// An unrecognized UF is ignored (as if omitted), so a national holiday
	// still reports true.
	if !date.IsHoliday(date.IsHolidayOptions{Date: d(2024, 1, 1), UF: "XX"}) {
		t.Error("Expected true: national holiday regardless of an unknown UF")
	}
}

func TestIsHolidayNoOptions(t *testing.T) {
	if date.IsHoliday() {
		t.Error("Expected false when no options are given")
	}
}

func TestIsBusinessDay(t *testing.T) {
	if !date.IsBusinessDay(d(2024, 1, 2)) {
		t.Error("Expected 2024-01-02 (Tuesday) to be a business day")
	}
	if date.IsBusinessDay(d(2024, 1, 1)) {
		t.Error("Expected 2024-01-01 (New Year) to not be a business day")
	}
	if date.IsBusinessDay(d(2024, 1, 6)) {
		t.Error("Expected 2024-01-06 (Saturday) to not be a business day")
	}
	// Carnival Monday 2024-02-12 is optional: business day when excluded.
	no := false
	if !date.IsBusinessDay(d(2024, 2, 12), date.BusinessDayOptions{IncludeOptionalHolidays: &no}) {
		t.Error("Expected Carnival Monday to be a business day when optional holidays are excluded")
	}
}

func TestAddBusinessDays(t *testing.T) {
	// 2024-01-01 is a holiday; adding 1 business day should land on 01-02.
	result := date.AddBusinessDays(d(2024, 1, 1), 1)
	if result == nil || !result.Equal(d(2024, 1, 2)) {
		t.Errorf("Expected 2024-01-02, got %v", result)
	}

	zero := date.AddBusinessDays(d(2024, 1, 6), 0)
	if zero == nil || !zero.Equal(d(2024, 1, 6)) {
		t.Errorf("Expected amount=0 to return the same date, got %v", zero)
	}
}

func TestSubBusinessDays(t *testing.T) {
	result := date.SubBusinessDays(d(2024, 1, 2), 1)
	if result == nil || !result.Equal(d(2023, 12, 29)) {
		t.Errorf("Expected 2023-12-29, got %v", result)
	}
}

func TestDifferenceInBusinessDays(t *testing.T) {
	diff := date.DifferenceInBusinessDays(d(2024, 1, 8), d(2024, 1, 1))
	if diff == nil {
		t.Fatal("Expected non-nil difference")
	}
	// Business days from 2024-01-01 (holiday) to 2024-01-08 exclusive:
	// 01-02, 01-03, 01-04, 01-05 = 4
	if *diff != 4 {
		t.Errorf("Expected 4, got %v", *diff)
	}

	same := date.DifferenceInBusinessDays(d(2024, 1, 1), d(2024, 1, 1))
	if same == nil || *same != 0 {
		t.Errorf("Expected 0 for the same date, got %v", same)
	}
}

func TestGetHolidays(t *testing.T) {
	holidays := date.GetHolidays(2024)
	if len(holidays) == 0 {
		t.Fatal("Expected a non-empty holiday list for 2024")
	}
	for i := 1; i < len(holidays); i++ {
		if holidays[i].Date.Before(holidays[i-1].Date) {
			t.Errorf("Expected holidays sorted by date, got %v before %v", holidays[i].Date, holidays[i-1].Date)
		}
	}

	if len(date.GetHolidays(1899)) != 0 {
		t.Error("Expected empty list for a year before 1900")
	}
}
