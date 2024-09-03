package cpf_test

import (
	"testing"

	"github.com/agaragon/brutils-go/cpf"
)

var tables = []struct {
	input    string
	expected bool
}{
	{"000111", false},
	{"00000000000", false},
	{"11111111111", false},
	{"22222222222", false},
	{"33333333333", false},
	{"44444444444", false},
	{"55555555555", false},
	{"66666666666", false},
	{"77777777777", false},
	{"88888888888", false},
	{"99999999999", false},
	{"403644788290", false},
	{"40364478829", true},
	{"962.718.458-60", true},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := cpf.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
