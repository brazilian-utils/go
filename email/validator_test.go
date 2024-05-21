package email_test

import (
	"testing"

	"github.com/agaragon/brutils-go/email"
)

var table = []struct {
	input    string
	expected bool
}{
	{"test", false},
	{"test@test", false},
	{"wrong-invalid", false},
	{"wrong-invalid@gmail.com", false},
	{".invalid@gmail.com", false},
	{"test@gmail.com", true},
	{"valid.valid@hotmail.com", true},
}

func TestValidate(t *testing.T) {
	for _, row := range table {
		if res := email.IsValid(row.input); res != row.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", row.input, row.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 1000; i++ {
		generatedEmail := email.Generate()
		if !email.IsValid(generatedEmail) {
			t.Errorf("The generated email is not valid %s", generatedEmail)
		}
	}
}
