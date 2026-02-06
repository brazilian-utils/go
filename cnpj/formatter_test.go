package cnpj_test

import (
	"testing"

	"github.com/brazilian-utils/brutils-go/cnpj"
)

var formatTests = []struct {
	input    string
	expected string
}{
	{"03560714000142", "03.560.714/0001-42"},
	{"13723705000189", "13.723.705/0001-89"},
	{"60391947000100", "60.391.947/0001-00"},
	{"1372370500018", ""},
	{"137237050001890", ""},
	{"abcdefghijklmn", ""},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := cnpj.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
