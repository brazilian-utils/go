package cnpj_test

import (
	"testing"

	"github.com/brazilian-utils/go/cnpj"
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

var removeSymbolsTests = []struct {
	input    string
	expected string
}{
	{"10799163989271", "10799163989271"},
	{"10.799.163/9892-71", "10799163989271"},
	{"", ""},
	{"   ", "   "},
	{"abc", "abc"},
}

func TestRemoveSymbols(t *testing.T) {
	for _, table := range removeSymbolsTests {
		if res := cnpj.RemoveSymbols(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"46.843.485/0001-86", "46843485000186"},
	{"46843485000186", "46843485000186"},
	{"46.?ABC843.485/0001-86abc", "46843485000186"},
	{"", ""},
	{"46843485000186123", "46843485000186"},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := cnpj.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
