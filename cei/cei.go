// Package cei formats, parses and validates a CEI (Cadastro Específico
// do INSS) number.
package cei

import (
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

const ceiSize = 12

var weights = []int{7, 4, 1, 8, 5, 2, 1, 6, 3, 7, 4}

// IsValid validates a CEI: 12 digits, 11 base digits and one check
// digit. A value whose digits are all the same is rejected.
func IsValid(value string) bool {
	digits := helpers.OnlyNumbers(value)
	if len(digits) != ceiSize || digits != cleanMask(value) {
		return false
	}
	if allSame(digits) {
		return false
	}
	return checkDigit(digits[:11]) == int(digits[11]-'0')
}

// cleanMask keeps only the digits and the usual grouping/mask
// characters (whitespace, '.', '-', '/'); anything else invalidates the
// value.
func cleanMask(value string) string {
	var b strings.Builder
	for _, r := range value {
		switch r {
		case ' ', '.', '-', '/':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b.WriteRune(r)
		default:
			return "" // force a mismatch against the digits-only value
		}
	}
	return b.String()
}

func checkDigit(base string) int {
	sum := 0
	for i := 0; i < 11; i++ {
		sum += int(base[i]-'0') * weights[i]
	}
	reduced := sum/10 + sum%10
	unit := reduced % 10
	return (10 - unit) % 10
}

func allSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 12 digits first.
	Pad bool
}

// Format formats a CEI with the usual mask "00.000.00000/00", masking as
// far as the digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < ceiSize {
		digits = strings.Repeat("0", ceiSize-len(digits)) + digits
	}
	if len(digits) > ceiSize {
		digits = digits[:ceiSize]
	}
	if digits == "" {
		return ""
	}

	seps := map[int]string{2: ".", 5: ".", 10: "/"}
	var b strings.Builder
	for i, c := range digits {
		if s, ok := seps[i]; ok {
			b.WriteString(s)
		}
		b.WriteRune(c)
	}
	return b.String()
}

// Parse removes CEI formatting and keeps only digits, capped to 12
// digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > ceiSize {
		digits = digits[:ceiSize]
	}
	return digits
}
