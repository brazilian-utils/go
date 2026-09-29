package registroprofissional_test

import (
	"testing"

	"github.com/brazilian-utils/go/registroprofissional"
)

var tests = []struct {
	value    string
	council  string
	state    string
	expected bool
}{
	{"123456/SP", "OAB", "", true},
	{"123456-SP", "OAB", "", true},
	{"123456/SP", "OAB", "SP", true},
	{"123456/SP", "OAB", "RJ", false},
	{"12345/RJ", "CRM", "", true},
	{"12/RJ", "CRM", "", false}, // too short
	{"123/SC", "CRO", "", true},
	{"06/12345", "CRP", "", true},
	{"25/12345", "CRP", "", false}, // region out of range
	{"00/12345", "CRP", "", false},
	{"SP-123456/O-3", "CRC", "", true},
	{"SP-123456/P-3", "CRC", "", true},
	{"sp-123456/o-3", "CRC", "", true},
	{"SP-123456/O-3-T-MG", "CRC", "", true},
	{"SP-123456/O-3", "CRC", "SP", true},
	{"SP-123456/O-3", "CRC", "RJ", false},
	{"invalid", "OAB", "", false},
	{"123456/SP", "CREA", "", false},
	{"", "OAB", "", false},
}

func TestIsValid(t *testing.T) {
	for _, table := range tests {
		res := registroprofissional.IsValid(registroprofissional.IsValidParams{
			Value: table.value, Council: table.council, State: table.state,
		})
		if res != table.expected {
			t.Errorf("Failing for %v/%v state=%v \t Expected: %v | Received: %v", table.value, table.council, table.state, table.expected, res)
		}
	}
}
