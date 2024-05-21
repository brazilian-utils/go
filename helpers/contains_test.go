package helpers_test

import (
	"testing"

	"github.com/agaragon/brutils-go/helpers"
)

var sliceOfString []string = []string{"foo", "bar"}

func TestContaints(t *testing.T) {
	tables := []struct {
		input    string
		expected bool
	}{
		{"foo", true},
		{"bar", true},
		{"foobar", false},
		{"another string", false},
	}

	for _, table := range tables {
		if res := helpers.Contains(sliceOfString, table.input); res != table.expected {
			t.Fatalf("Failed for %s (expected: %t, received: %t)", table.input, table.expected, res)
		}
	}
}
