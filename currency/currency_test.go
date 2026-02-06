package currency_test

import (
	"math"
	"testing"

	"github.com/brazilian-utils/brutils-go/currency"
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
	// Examples from the Python docstring
	{1523.45, "Mil, quinhentos e vinte e três reais e quarenta e cinco centavos"},
	{1.00, "Um real"},
	{0.50, "Cinquenta centavos"},
	{0.00, "Zero reais"},

	// Additional cases
	{0.01, "Um centavo"},
	{2.50, "Dois reais e cinquenta centavos"},
	{100, "Cem reais"},
	{101, "Cento e um reais"},
	{1000, "Mil reais"},
	{1001, "Mil e um reais"},
	{1200, "Mil e duzentos reais"},
	{2000, "Dois mil reais"},
	{1000000, "Um milhão de reais"},
	{2000000, "Dois milhões de reais"},
	{1000000000, "Um bilhão de reais"},
	{1000000000000, "Um trilhão de reais"},
	{-5.25, "Menos cinco reais e vinte e cinco centavos"},
	{1500000, "Um milhão e quinhentos mil reais"},
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
