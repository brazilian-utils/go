package cpf

import (
	"strconv"
	"strings"

	"github.com/agaragon/brutils-go/helpers"
)

// Valid CPFs but blacklisted they are reserved
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

// Position of verification digits
var verifierIndexes = []int{9, 10}

// Every CPF has exactly 11 characters
const cpfSize = 11

// IsValid validates if a given CPF is valid
func IsValid(cpf string) bool {
	cpfNumbers := helpers.OnlyNumbers(cpf)

	return hasValidLength(cpfNumbers) && !isBlacklisted(cpf) && isValidChecksum(cpfNumbers)
}

// Perform checksum validation
func isValidChecksum(cpf string) bool {
	validity := true

	for _, verifier := range verifierIndexes {
		mod := computeMod(strings.Split(cpf[:verifier], ""))

		valid, _ := strconv.Atoi(string(cpf[verifier]))
		validity = validity && (valid == mod)
	}

	return validity
}

// Compute the mod for the current slice of strings
func computeMod(digits []string) (res int) {
	weight := len(digits) + 1

	var mod int
	for _, digitStr := range digits {
		digit, _ := strconv.Atoi(digitStr)
		mod += digit * weight
		weight--
	}

	if mod = mod % 11; mod >= 2 {
		res = 11 - mod
	}

	return
}

// Validates the string length
func hasValidLength(cpf string) bool {
	return len(cpf) == cpfSize
}

func isBlacklisted(cpf string) bool {
	return helpers.Contains(blacklist, cpf)
}
