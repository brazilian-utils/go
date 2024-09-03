package cep_test

import (
	"testing"

	"github.com/agaragon/brutils-go/cep"
)

var tables = []struct {
	input    string
	expected bool
}{
	{"000111", false},
	{"13723705000189", false},
	{"60.391.947/0001-0", false},
	{"abcdefgh", false},
	{"99999999999999", false},
	{"12345678", true},
	{"12345-678", true},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := cep.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := cep.Generate()
		if !cep.IsValid(res) {
			t.Errorf("An invalid cep was generated: %s", res)
		}
	}
}
