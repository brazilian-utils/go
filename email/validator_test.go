package email_test

import (
	"testing"

	"github.com/brazilian-utils/brutils-go/email"
)

var tables = []struct {
	input    string
	expected bool
}{
	{"brutils@brutils.com", true},
	{"user@example.org", true},
	{"user.name@domain.com", true},
	{"user+tag@domain.co.uk", true},
	{"user%special@domain.com", true},
	{"user-name@domain.com", true},
	{"user_name@domain.com", true},
	{"a@b.co", true},

	// Invalid
	{"invalid-email@brutils", false},
	{".startswithdot@domain.com", false},
	{"@domain.com", false},
	{"user@", false},
	{"user@.com", false},
	{"", false},
	{"plaintext", false},
	{"user@@domain.com", false},
	{"user@domain", false},
}

func TestValidate(t *testing.T) {
	for _, table := range tables {
		if res := email.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}
