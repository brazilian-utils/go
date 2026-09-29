package licenseplate

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var oldFormatRegex = regexp.MustCompile(`^[A-Za-z]{3}[0-9]{4}$`)
var mercosulRegex = regexp.MustCompile(`^[A-Za-z]{3}[0-9][A-Za-z][0-9]{2}$`)

// normalizeForValidation trims whitespace and removes the usual separator
// (a hyphen or an internal space) so a masked plate can be recognized.
func normalizeForValidation(plate string) string {
	plate = strings.TrimSpace(plate)
	plate = strings.ReplaceAll(plate, "-", "")
	plate = strings.ReplaceAll(plate, " ", "")
	return plate
}

// IsValid checks if a Brazilian license plate is valid: the old format
// (ABC1234) or the Mercosul format (ABC1D23), in any case, with or
// without a hyphen or space.
func IsValid(plate string) bool {
	normalized := normalizeForValidation(plate)
	return isValidOldFormat(normalized) || isValidMercosul(normalized)
}

// GetFormat returns the format of a license plate:
// "LLLNNNN" for old format, "LLLNLNN" for Mercosul, or empty string if invalid.
func GetFormat(plate string) string {
	normalized := normalizeForValidation(plate)
	if isValidOldFormat(normalized) {
		return "LLLNNNN"
	}
	if isValidMercosul(normalized) {
		return "LLLNLNN"
	}
	return ""
}

// Format formats a license plate into the correct display pattern.
// Old format gets a dash: "ABC-1234". Mercosul gets uppercased: "ABC1E34".
// Returns empty string if the input is not a valid plate.
func Format(plate string) string {
	normalized := normalizeForValidation(plate)
	upper := strings.ToUpper(normalized)
	if isValidOldFormat(normalized) {
		return upper[:3] + "-" + upper[3:]
	}
	if isValidMercosul(normalized) {
		return upper
	}
	return ""
}

// ConvertToMercosul converts an old format plate (LLLNNNN) to Mercosul (LLLNLNN).
// The 5th character (second digit) is converted to a letter (0→A, 1→B, ..., 9→J).
// Accepts the plate bare or with a hyphen mask. Returns empty string if the
// input is not a valid old format plate.
func ConvertToMercosul(plate string) string {
	normalized := normalizeForValidation(plate)
	if !isValidOldFormat(normalized) {
		return ""
	}

	upper := strings.ToUpper(normalized)
	fifthChar := rune(upper[4])
	letter := 'A' + (fifthChar - '0')

	return fmt.Sprintf("%s%c%s", upper[:4], letter, upper[5:])
}

// RemoveSymbols removes the formatting symbol ('-'), keeping everything
// else (including whitespace) unchanged.
func RemoveSymbols(value string) string {
	return strings.ReplaceAll(value, "-", "")
}

// Parse removes separators from a license plate, upper-cases it and caps
// it to 7 characters.
func Parse(value string) string {
	cleaned := RemoveSymbols(value)
	upper := strings.ToUpper(cleaned)
	if len(upper) > 7 {
		upper = upper[:7]
	}
	return upper
}

// Generate generates a random license plate in the given format.
// format must be "LLLNLNN" (Mercosul, the default) or "LLLNNNN" (old
// format); any other value falls back to the default.
func Generate(format ...string) string {
	f := "LLLNLNN"
	if len(format) > 0 {
		requested := strings.ToUpper(format[0])
		if requested == "LLLNLNN" || requested == "LLLNNNN" {
			f = requested
		}
	}

	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var buf strings.Builder
	for _, ch := range f {
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
