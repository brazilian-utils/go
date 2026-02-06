package licenseplate

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var oldFormatRegex = regexp.MustCompile(`^[A-Za-z]{3}[0-9]{4}$`)
var mercosulRegex = regexp.MustCompile(`^[A-Za-z]{3}[0-9][A-Za-z][0-9]{2}$`)

// IsValid checks if a Brazilian license plate is valid.
// plateType can be "old_format", "mercosul", or "" (accepts either).
func IsValid(plate string, plateType string) bool {
	plate = strings.TrimSpace(plate)
	switch plateType {
	case "old_format":
		return isValidOldFormat(plate)
	case "mercosul":
		return isValidMercosul(plate)
	default:
		return isValidOldFormat(plate) || isValidMercosul(plate)
	}
}

// GetFormat returns the format of a license plate:
// "LLLNNNN" for old format, "LLLNLNN" for Mercosul, or empty string if invalid.
func GetFormat(plate string) string {
	plate = strings.TrimSpace(plate)
	if isValidOldFormat(plate) {
		return "LLLNNNN"
	}
	if isValidMercosul(plate) {
		return "LLLNLNN"
	}
	return ""
}

// Format formats a license plate into the correct display pattern.
// Old format gets a dash: "ABC-1234". Mercosul gets uppercased: "ABC1E34".
// Returns empty string if invalid.
func Format(plate string) string {
	plate = strings.TrimSpace(plate)
	upper := strings.ToUpper(plate)
	if isValidOldFormat(plate) {
		return upper[:3] + "-" + upper[3:]
	}
	if isValidMercosul(plate) {
		return upper
	}
	return ""
}

// ConvertToMercosul converts an old format plate (LLLNNNN) to Mercosul (LLLNLNN).
// The 5th character (second digit) is converted to a letter (0→A, 1→B, ..., 9→J).
// Returns empty string if the input is not a valid old format plate.
func ConvertToMercosul(plate string) string {
	plate = strings.TrimSpace(plate)
	if !isValidOldFormat(plate) {
		return ""
	}

	upper := strings.ToUpper(plate)
	fifthChar := rune(upper[4])
	letter := rune('A') + (fifthChar - '0')

	return fmt.Sprintf("%s%c%s", upper[:4], letter, upper[5:])
}

// Generate generates a random license plate in the given format.
// format must be "LLLNLNN" (Mercosul, default) or "LLLNNNN" (old format).
// Returns empty string if the format is invalid.
func Generate(format string) string {
	format = strings.ToUpper(format)
	if format != "LLLNLNN" && format != "LLLNNNN" {
		return ""
	}

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var buf strings.Builder
	for _, ch := range format {
		if ch == 'L' {
			buf.WriteByte(letters[rand.Intn(len(letters))])
		} else {
			buf.WriteByte(byte('0' + rand.Intn(10)))
		}
	}
	return buf.String()
}

func isValidOldFormat(plate string) bool {
	return oldFormatRegex.MatchString(plate)
}

func isValidMercosul(plate string) bool {
	return mercosulRegex.MatchString(plate)
}
