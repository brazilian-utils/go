package number_test

import (
	"testing"

	"github.com/brazilian-utils/go/number"
)

var convertTests = []struct {
	input    float64
	expected string
}{
	{0, "zero"},
	{1, "um"},
	{3, "três"},
	{100, "cem"},
	{101, "cento e um"},
	{123, "cento e vinte e três"},
	{1000, "mil"},
	{1000000, "um milhão"},
	{-3, "menos três"},
	{12.9, "doze"},
	{1235, "mil duzentos e trinta e cinco"},
}

func TestConvertToWords(t *testing.T) {
	for _, table := range convertTests {
		if res := number.ConvertToWords(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestConvertToWordsFeminine(t *testing.T) {
	feminineTests := []struct {
		input    float64
		expected string
	}{
		{1, "uma"},
		{2, "duas"},
		{200, "duzentas"},
		{201, "duzentas e uma"},
	}
	for _, table := range feminineTests {
		if res := number.ConvertToWords(table.input, number.ConvertToWordsOptions{Feminine: true}); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestConvertToWordsOutOfRange(t *testing.T) {
	if res := number.ConvertToWords(1_000_000_000_000_000); res != "" {
		t.Errorf("Expected empty string for value exceeding max, got %v", res)
	}
}
