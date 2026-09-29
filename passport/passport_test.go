package passport_test

import (
	"testing"

	"github.com/brazilian-utils/go/passport"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"AB123456", "AB123456"},
	{"acd12736", "ACD12736"},
	{"AB-123.456", "AB123456"},
	{"AB12", "AB12"},
	{"", ""},
	{"AB123456789", "AB123456"},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := passport.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var isValidTests = []struct {
	input    interface{}
	expected bool
}{
	{"AA111111", true},
	{"CL125167", true},
	{"ab123456", true},
	{"AB-123456", true},
	{"AB.123.456", true},
	{"1", false},
	{"1112223334-", false},
	{1, false},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := passport.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"Ab123456", "AB123456"},
	{" AB 123 456 ", "AB123456"},
	{"-AB1-23-4-56-", "AB123456"},
	{"AB123456789", "AB123456"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := passport.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 20; i++ {
		generated := passport.Generate()
		if !passport.IsValid(generated) {
			t.Errorf("Generated invalid passport number: %v", generated)
		}
	}
}
