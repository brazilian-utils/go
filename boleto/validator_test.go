package boleto_test

import (
	"testing"

	"github.com/brazilian-utils/go/boleto"
)

var tables = []struct {
	input    string
	expected bool
}{
	{"", false},
	{"000111", false},
	{"00190000020114971860168524522114675860000102656", false},
	{"00190000090114971860168524522114975860000102656", false},
	{"00190000090114971860168524522114675860000102656", true},
	{"0019000009 01149.718601 68524.522114 6 75860000102656", true},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := boleto.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
