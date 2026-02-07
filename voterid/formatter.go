package voterid

import (
	"fmt"

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
