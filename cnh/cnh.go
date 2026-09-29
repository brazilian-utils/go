package cnh

import (
	"math/rand"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 11 digits first.
	Pad bool
}

// Format formats a CNH number as "000000000-00" (9 digits, hyphen, 2
// check digits), masking as far as the digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < cnhSize {
		digits = strings.Repeat("0", cnhSize-len(digits)) + digits
	}
	if len(digits) > cnhSize {
		digits = digits[:cnhSize]
	}
	if digits == "" {
		return ""
	}
	if len(digits) <= 9 {
		return digits
	}
	return digits[:9] + "-" + digits[9:]
}

// Parse removes CNH formatting and keeps only digits, capped to 11 digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > cnhSize {
		digits = digits[:cnhSize]
	}
	return digits
}

// Generate generates a random valid CNH number (11 digits, unformatted).
func Generate() string {
	for {
		base := make([]int, 9)
		for i := range base {
			base[i] = rand.Intn(10)
		}

		firstRemainder := 0
		for i := 0; i < 9; i++ {
			firstRemainder += base[i] * (9 - i)
		}
		firstRemainder %= 11
		firstDigit := firstRemainder
		if firstDigit > 9 {
			firstDigit = 0
		}

		secondSum := 0
		for i := 0; i < 9; i++ {
			secondSum += base[i] * (i + 1)
		}
		secondSum %= 11
		if firstRemainder >= 10 {
			if secondSum-2 < 0 {
				secondSum += 9
			} else {
				secondSum -= 2
			}
		}
		secondDigit := secondSum
		if secondDigit > 9 {
			secondDigit = 0
		}

		var sb strings.Builder
		for _, d := range base {
			sb.WriteByte(byte('0' + d))
		}
		sb.WriteByte(byte('0' + firstDigit))
		sb.WriteByte(byte('0' + secondDigit))

		candidate := sb.String()
		if !helpers.Contains(blacklist, candidate) {
			return candidate
		}
	}
}
