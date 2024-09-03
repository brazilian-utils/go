package cnpj_test

import (
	"testing"

	"github.com/agaragon/brutils-go/cnpj"
)

var tables = []struct {
	input    string
	expected bool
}{
	{"000111", false},
	{"00000000000000", false},
	{"11111111111111", false},
	{"22222222222222", false},
	{"33333333333333", false},
	{"44444444444444", false},
	{"55555555555555", false},
	{"66666666666666", false},
	{"77777777777777", false},
	{"88888888888888", false},
	{"99999999999999", false},
	{"13723705000189", true},
	{"60.391.947/0001-00", true},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := cnpj.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
