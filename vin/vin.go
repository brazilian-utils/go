// Package vin validates a Vehicle Identification Number (chassi).
package vin

import "strings"

const vinSize = 17

var transliteration = map[byte]int{
	'A': 1, 'B': 2, 'C': 3, 'D': 4, 'E': 5, 'F': 6, 'G': 7, 'H': 8,
	'J': 1, 'K': 2, 'L': 3, 'M': 4, 'N': 5, 'P': 7, 'R': 9,
	'S': 2, 'T': 3, 'U': 4, 'V': 5, 'W': 6, 'X': 7, 'Y': 8, 'Z': 9,
}

var weights = [vinSize]int{8, 7, 6, 5, 4, 3, 2, 10, 0, 9, 8, 7, 6, 5, 4, 3, 2}

// IsValid validates a VIN (chassi) structurally: 17 characters, none of
// the excluded letters I, O, Q, and the check digit at position 9 (North
// American rule), case-insensitive. A VIN whose characters are all the
// same is rejected even when the check digit matches.
func IsValid(value string) bool {
	upper := strings.ToUpper(value)
	if len(upper) != vinSize {
		return false
	}
	if allSame(upper) {
		return false
	}

	sum := 0
	for i := 0; i < vinSize; i++ {
		v, ok := charValue(upper[i])
		if !ok {
			return false
		}
		sum += v * weights[i]
	}

	remainder := sum % 11
	expected := byte('0' + remainder)
	if remainder == 10 {
		expected = 'X'
	}

	return upper[8] == expected
}

func charValue(c byte) (int, bool) {
	if c >= '0' && c <= '9' {
		return int(c - '0'), true
	}
	if v, ok := transliteration[c]; ok {
		return v, true
	}
	return 0, false
}

func allSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}
