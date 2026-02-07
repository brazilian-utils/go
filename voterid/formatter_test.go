package voterid_test

import (
	"testing"

	"github.com/brazilian-utils/go/voterid"
)

var formatTests = []struct {
	input    string
	expected string
}{
	// Valid voter IDs
	{"690847092828", "6908 4709 28 28"},
	{"163204010922", "1632 0401 09 22"},
	{"217633460930", "2176 3346 09 30"},
	{"858224120973", "8582 2412 09 73"},
	{"149426030183", "1494 2603 01 83"},

	// Already formatted
	{"6908 4709 28 28", "6908 4709 28 28"},

	// Invalid voter IDs should return empty string
	{"12345", ""},
	{"690847092827", ""},
	{"", ""},
	{"abcdefghijkl", ""},
	{"123456789011", ""},

	// Edge case: 13 digits should be truncated to 12 for display
	{"3244567800167", "3244 5678 00 16"},
	{"7865793030175", "7865 7930 30 17"},
}

func TestFormat(t *testing.T) {
	for _, test := range formatTests {
		result := voterid.Format(test.input)
		if result != test.expected {
			t.Errorf("Format(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		voterid.Format("690847092828")
	}
}
