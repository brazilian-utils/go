// Package cns formats, parses and validates a CNS (Cartão Nacional de
// Saúde) number.
package cns

import (
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

const cnsSize = 15

// cleanMask keeps digits and the usual grouping separators (whitespace,
// '.', '-', '/'); any other character makes the value invalid.
func cleanMask(value string) string {
	var b strings.Builder
	for _, r := range value {
		switch r {
		case ' ', '.', '-', '/':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b.WriteRune(r)
		default:
			return ""
		}
	}
	return b.String()
}

// IsValid validates a CNS number: 15 digits. Definitive cards (starting
// 1 or 2) and provisional ones (starting 7, 8 or 9) each use their own
// modulus 11 rule; a number starting with 5 (or anything else) is
// rejected.
func IsValid(value string) bool {
	digits := helpers.OnlyNumbers(value)
	if len(digits) != cnsSize || digits != cleanMask(value) {
		return false
	}

	switch digits[0] {
	case '1', '2':
		return isValidDefinitive(digits)
	case '7', '8', '9':
		return isValidProvisional(digits)
	default:
		return false
	}
}

func isValidDefinitive(digits string) bool {
	pis := digits[:11]
	sum := 0
	for i := 0; i < 11; i++ {
		sum += int(pis[i]-'0') * (15 - i)
	}
	resto := sum % 11
	dv := 11 - resto

	var suffix string
	switch {
	case dv == 11:
		suffix = "0000"
	case dv == 10:
		sum2 := sum + 2
		resto2 := sum2 % 11
		dv2 := 11 - resto2
		suffix = "001" + string(rune('0'+dv2))
	default:
		suffix = "000" + string(rune('0'+dv))
	}

	return pis+suffix == digits
}

func isValidProvisional(digits string) bool {
	sum := 0
	for i := 0; i < 15; i++ {
		sum += int(digits[i]-'0') * (15 - i)
	}
	return sum%11 == 0
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 15 digits first.
	Pad bool
}

// Format formats a CNS number into groups of 3-4-4-4 digits separated by
// spaces, masking as far as the digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < cnsSize {
		digits = strings.Repeat("0", cnsSize-len(digits)) + digits
	}
	if len(digits) > cnsSize {
		digits = digits[:cnsSize]
	}
	if digits == "" {
		return ""
	}

	groupSizes := []int{3, 4, 4, 4}
	var b strings.Builder
	pos := 0
	for gi, size := range groupSizes {
		if pos >= len(digits) {
			break
		}
		if gi > 0 {
			b.WriteString(" ")
		}
		end := pos + size
		if end > len(digits) {
			end = len(digits)
		}
		b.WriteString(digits[pos:end])
		pos = end
	}
	return b.String()
}

// Parse removes CNS formatting and keeps only digits, capped to 15
// digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > cnsSize {
		digits = digits[:cnsSize]
	}
	return digits
}
