// Package creditcard validates a generic payment card number (Luhn).
package creditcard

import (
	"strings"
)

// IsValid validates a payment card number (credit or debit): 12 to 19
// digits and a Luhn check digit. No brand detection, issuer range
// lookup, expiry or CVV check. A number whose digits are all the same is
// rejected even though it passes Luhn.
//
// value may contain whitespace, '.', '-' and '/' between the digits;
// any other character makes it invalid.
func IsValid(value string) bool {
	cleaned := removeSeparators(value)
	if len(cleaned) < 12 || len(cleaned) > 19 {
		return false
	}
	for _, c := range cleaned {
		if c < '0' || c > '9' {
			return false
		}
	}
	if allSameDigit(cleaned) {
		return false
	}
	return luhnValid(cleaned)
}

func removeSeparators(value string) string {
	r := strings.NewReplacer(" ", "", ".", "", "-", "", "/", "")
	return r.Replace(value)
}

func allSameDigit(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func luhnValid(digits string) bool {
	sum := 0
	double := false
	for i := len(digits) - 1; i >= 0; i-- {
		d := int(digits[i] - '0')
		if double {
			d *= 2
			if d > 9 {
				d -= 9
			}
		}
		sum += d
		double = !double
	}
	return sum%10 == 0
}
