package cnpj

import (
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

var blacklist = []string{
	"00000000000000",
	"11111111111111",
	"22222222222222",
	"33333333333333",
	"44444444444444",
	"55555555555555",
	"66666666666666",
	"77777777777777",
	"88888888888888",
	"99999999999999",
}

// Position of verification digits
var verifierIndexes = []int{12, 13}

// Every CNPJ has exactly 14 characters
const cnpjSize = 14

// IsValidOptions controls IsValid.
type IsValidOptions struct {
	// Version selects the accepted format: 1 (default) numeric only,
	// 2 also accepts the alphanumeric CNPJ. Any other value is read as 1.
	Version int
}

// IsValid validates if a given CNPJ is valid. options[0].Version == 2
// also accepts the alphanumeric CNPJ (case-insensitive).
func IsValid(cnpj string, options ...IsValidOptions) bool {
	version := 1
	if len(options) > 0 && options[0].Version == 2 {
		version = 2
	}

	cleaned := removeMaskChars(cnpj)

	if version == 2 {
		return hasValidLength(cleaned) && isValidChecksumAlnum(strings.ToUpper(cleaned))
	}

	digits := helpers.OnlyNumbers(cnpj)
	return hasValidLength(digits) && !isBlacklisted(digits) && isValidChecksumAlnum(digits)
}

// removeMaskChars strips the usual CNPJ mask characters and whitespace,
// keeping digits and letters (for the alphanumeric CNPJ).
func removeMaskChars(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '.', '-', '/', ' ':
			continue
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}

// isValidChecksumAlnum validates the 2 check digits of a 14-character CNPJ,
// numeric or alphanumeric (value = ASCII code - 48, per RFB's alphanumeric
// CNPJ check-digit spec; the check digits themselves are always numeric).
func isValidChecksumAlnum(cnpj string) bool {
	if len(cnpj) != cnpjSize {
		return false
	}
	for _, verifier := range verifierIndexes {
		if cnpj[verifier] < '0' || cnpj[verifier] > '9' {
			return false
		}
		mod := computeModAlnum(cnpj[:verifier], verifier)
		if int(cnpj[verifier]-'0') != mod {
			return false
		}
	}
	return true
}

// Compute the mod for the current slice of characters
func computeModAlnum(chars string, verifier int) (res int) {
	weights := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	if verifier == verifierIndexes[len(verifierIndexes)-1] {
		weights = append([]int{6}, weights...)
	}

	var mod int
	for index := 0; index < len(chars); index++ {
		c := chars[index]
		if c < '0' || (c > '9' && c < 'A') || c > 'Z' {
			return -1
		}
		value := int(c) - 48
		mod += value * weights[index]
	}

	if mod = mod % 11; mod >= 2 {
		res = 11 - mod
	}

	return
}

// Validates the string length
func hasValidLength(cnpj string) bool {
	return len(cnpj) == cnpjSize
}

func isBlacklisted(cnpj string) bool {
	return helpers.Contains(blacklist, cnpj)
}
