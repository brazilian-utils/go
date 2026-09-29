package voterid

import (
	"fmt"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

// Format formats a voter ID for display with visual spaces
// Returns the formatted voter ID as "XXXX XXXX XX XX" or empty string if invalid
func Format(voterID string) string {
	voterIDNumbers := helpers.OnlyNumbers(voterID)

	// Validate the voter ID
	if !IsValid(voterIDNumbers) {
		return ""
	}

	// Truncate to 12 characters for formatting (standard length)
	if len(voterIDNumbers) > 12 {
		voterIDNumbers = voterIDNumbers[:12]
	}

	return fmt.Sprintf("%s %s %s %s",
		voterIDNumbers[:4],
		voterIDNumbers[4:8],
		voterIDNumbers[8:10],
		voterIDNumbers[10:12],
	)
}

// RemoveSymbols removes the voter ID formatting symbols (space, '.'),
// keeping everything else unchanged.
func RemoveSymbols(value string) string {
	r := strings.NewReplacer(" ", "", ".", "")
	return r.Replace(value)
}

// Parse removes voter ID formatting and keeps only digits, capped to 12
// digits (13 when the UF digits, at position 10-11 of a 13-digit value,
// are 01 - SP - or 02 - MG).
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)

	capLen := 12
	if len(digits) >= 13 {
		uf := digits[9:11]
		if uf == "01" || uf == "02" {
			capLen = 13
		}
	}
	if len(digits) > capLen {
		digits = digits[:capLen]
	}
	return digits
}
