package cnh

import (
	"strconv"

	"github.com/brazilian-utils/go/helpers"
)

// Every CNH has exactly 11 digits
const cnhSize = 11

// blacklist contains reserved sequences that are technically valid checksums
// but not real CNH numbers.
var blacklist = []string{
	"00000000000",
	"11111111111",
	"22222222222",
	"33333333333",
	"44444444444",
	"55555555555",
	"66666666666",
	"77777777777",
	"88888888888",
	"99999999999",
}

// IsValid validates if a given CNH (Carteira Nacional de Habilitação) is valid.
// It checks the format (11 digits) and verifies both check digits.
func IsValid(cnh string) bool {
	cnhNumbers := helpers.OnlyNumbers(cnh)

	if len(cnhNumbers) != cnhSize {
		return false
	}

	if helpers.Contains(blacklist, cnhNumbers) {
		return false
	}

	digits := toDigits(cnhNumbers)
	firstVerifier := digits[9]
	secondVerifier := digits[10]

	if !checkFirstVerifier(digits, firstVerifier) {
		return false
	}

	return checkSecondVerifier(digits, secondVerifier, firstVerifier)
}

// checkFirstVerifier validates the 10th digit (first check digit).
// Sum of first 9 digits weighted by (9-i), mod 11; result is 0 if > 9.
func checkFirstVerifier(digits []int, firstVerifier int) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (9 - i)
	}

	result := sum % 11
	if result > 9 {
		result = 0
	}

	return result == firstVerifier
}

// checkSecondVerifier validates the 11th digit (second check digit).
// Sum of first 9 digits weighted by (i+1), mod 11; adjusted if first verifier > 9.
func checkSecondVerifier(digits []int, secondVerifier int, firstVerifier int) bool {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (i + 1)
	}

	result := sum % 11

	if firstVerifier > 9 {
		if (result - 2) < 0 {
			result = result + 9
		} else {
			result = result - 2
		}
	}

	if result > 9 {
		result = 0
	}

	return result == secondVerifier
}

// toDigits converts a numeric string into a slice of ints.
func toDigits(s string) []int {
	digits := make([]int, len(s))
	for i, ch := range s {
		digits[i], _ = strconv.Atoi(string(ch))
	}
	return digits
}
