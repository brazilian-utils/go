package voterid_test

import (
	"testing"

	"github.com/brazilian-utils/go/voterid"
)

var validationTests = []struct {
	input    string
	expected bool
}{
	// Valid voter IDs (from Python tests)
	{"217633460930", true},
	{"690847092828", true},
	{"163204010922", true},
	{"858224120973", true},
	{"149426030183", true},
	{"033568860230", true},

	// Valid voter IDs with formatting
	{"6908 4709 28 28", true},
	{"1632 0401 09 22", true},

	// Invalid: wrong length
	{"12345", false},
	{"123456789012345", false},
	{"", false},

	// Invalid: wrong check digits (from Python tests)
	{"123456789011", false},
	{"427503840223", false},
	{"427503840214", false},
	{"690847092827", false},
	{"163204010921", false},

	// Invalid: invalid federative union (00)
	{"123456780012", false},

	// Invalid: invalid federative union (29)
	{"123456782912", false},

	// Invalid: non-numeric
	{"abcd1234efgh", false},

	// Edge case: 13 digits for SP (01) and MG (02)
	{"3244567800167", true},
	{"7865793030175", true},
	{"2195408310272", true},

	// Invalid: 13 digits for non-SP/MG state
	{"0123456780312", false},
}

func TestIsValid(t *testing.T) {
	for _, test := range validationTests {
		result := voterid.IsValid(test.input)
		if result != test.expected {
			t.Errorf("IsValid(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func BenchmarkIsValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		voterid.IsValid("690847092828")
	}
}
