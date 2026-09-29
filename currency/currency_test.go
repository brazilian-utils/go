package currency_test

import (
	"math"
	"testing"

	"github.com/brazilian-utils/go/currency"
)

var formatTests = []struct {
	input    float64
	expected string
}{
	{1234.56, "R$ 1.234,56"},
	{0, "R$ 0,00"},
	{-9876.54, "R$ -9.876,54"},
	{1000000.00, "R$ 1.000.000,00"},
	{0.99, "R$ 0,99"},
	{1.00, "R$ 1,00"},
	{123456789.01, "R$ 123.456.789,01"},
	{-0.01, "R$ -0,01"},
}

func TestFormatCurrency(t *testing.T) {
	for _, table := range formatTests {
		if res := currency.FormatCurrency(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestFormatCurrencyNaN(t *testing.T) {
	if res := currency.FormatCurrency(math.NaN()); res != "" {
		t.Errorf("Expected empty string for NaN, got %v", res)
	}
}

func TestFormatCurrencyInf(t *testing.T) {
	if res := currency.FormatCurrency(math.Inf(1)); res != "" {
		t.Errorf("Expected empty string for +Inf, got %v", res)
	}
}

var convertTests = []struct {
	input    float64
	expected string
}{
	// Examples from the shared brazilian-utils contract (cases/currency.json):
	// all lower case, no commas between groups.
	{1523.45, "mil quinhentos e vinte e três reais e quarenta e cinco centavos"},
	{1.00, "um real"},
	{0.50, "cinquenta centavos"},
	{0.00, "zero reais"},

	// Additional cases
	{0.01, "um centavo"},
	{2.50, "dois reais e cinquenta centavos"},
	{100, "cem reais"},
	{101, "cento e um reais"},
	{1000, "mil reais"},
	{1001, "mil e um reais"},
	{1200, "mil e duzentos reais"},
	{2000, "dois mil reais"},
	{1000000, "um milhão de reais"},
	{2000000, "dois milhões de reais"},
	{1000000000, "um bilhão de reais"},
	{1000000000000, "um trilhão de reais"},
	{-5.25, "menos cinco reais e vinte e cinco centavos"},
	{1500000, "um milhão e quinhentos mil reais"},
	{1000230, "um milhão duzentos e trinta reais"},
	{1.999, "um real e noventa e nove centavos"},
}

func TestConvertRealToText(t *testing.T) {
	for _, table := range convertTests {
		if res := currency.ConvertRealToText(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestConvertRealToTextNaN(t *testing.T) {
	if res := currency.ConvertRealToText(math.NaN()); res != "" {
		t.Errorf("Expected empty string for NaN, got %v", res)
	}
}

func TestConvertRealToTextExceedsMax(t *testing.T) {
	if res := currency.ConvertRealToText(1_000_000_000_000_001); res != "" {
		t.Errorf("Expected empty string for value exceeding max, got %v", res)
	}
}

var formatValueTests = []struct {
	input    interface{}
	expected string
}{
	{1000.01, "1.000,01"},
	{0.01, "0,01"},
	{1.0, "1,00"},
	{1000000.01, "1.000.000,01"},
	{-10.1, "-10,10"},
	{"1234.56", "1.234,56"},
	{"1.234,56", "1.234,56"},
	{"R$ 1.234,56", "1.234,56"},
}

func TestFormat(t *testing.T) {
	for _, table := range formatValueTests {
		if res := currency.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected float64
}{
	{"R$ 1.234,56", 1234.56},
	{"1234,56", 1234.56},
	{"R$ 0,50", 0.5},
	{"1.000.000,50", 1000000.5},
	{"-1.234,56", -1234.56},
	{"0,00", 0},
	{"1234", 12.34},
	{"", 0},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := currency.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
