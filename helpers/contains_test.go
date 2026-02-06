package helpers_test

import (
	"testing"

	"github.com/brazilian-utils/brutils-go/helpers"
)

var sliceOfString []string = []string{"foo", "bar"}
var sliceOfInt []int = []int{1, 999}

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

func TestContaintsInt(t *testing.T) {
	tables := []struct {
		input    int
		expected bool
	}{
		{1, true},
		{999, true},
		{10, false},
		{99, false},
	}

	for _, table := range tables {
		if res := helpers.ContainsInt(sliceOfInt, table.input); res != table.expected {
			t.Fatalf("Failed for %v (expected: %t, received: %t)", table.input, table.expected, res)
		}
	}
}
