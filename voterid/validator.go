package voterid

import (
	"strconv"

	"github.com/brazilian-utils/go/helpers"
)

// Valid federative union codes (01-28)
var validFederativeUnions = []string{
	"01", "02", "03", "04", "05", "06", "07", "08", "09", "10",
	"11", "12", "13", "14", "15", "16", "17", "18", "19", "20",
	"21", "22", "23", "24", "25", "26", "27", "28",
}

// São Paulo and Minas Gerais codes
var spMgCodes = []string{"01", "02"}

// IsValid validates if a given voter ID is valid
// It checks format, length, federative union codes, and verifying digits
func IsValid(voterID string) bool {
	// Get only numbers
	voterIDNumbers := helpers.OnlyNumbers(voterID)

	// Check if it has valid length
	if !isLengthValid(voterIDNumbers) {
		return false
	}

	// Extract components
	sequentialNumber := getSequentialNumber(voterIDNumbers)
	federativeUnion := getFederativeUnion(voterIDNumbers)
	verifyingDigits := getVerifyingDigits(voterIDNumbers)

	// Validate federative union
	if !isFederativeUnionValid(federativeUnion) {
		return false
	}

	// Calculate and validate first verifying digit
	vd1 := calculateVD1(sequentialNumber, federativeUnion)
	digit1, _ := strconv.Atoi(string(verifyingDigits[0]))
	if vd1 != digit1 {
		return false
	}

	// Calculate and validate second verifying digit
	vd2 := calculateVD2(federativeUnion, vd1)
	digit2, _ := strconv.Atoi(string(verifyingDigits[1]))
	if vd2 != digit2 {
		return false
	}

	return true
}

// isLengthValid checks if the voter ID has valid length
// Typically 12 digits, but SP and MG can have 13 (edge case with 9-digit sequential)
func isLengthValid(voterID string) bool {
	length := len(voterID)

	if length == 12 {
		return true
	}

	// Edge case: SP and MG with 13 digits
	if length == 13 {
		federativeUnion := getFederativeUnion(voterID)
		return helpers.Contains(spMgCodes, federativeUnion)
	}

	return false
}

// getSequentialNumber extracts the sequential number (first 8 digits)
func getSequentialNumber(voterID string) string {
	return voterID[:8]
}

// getFederativeUnion extracts the federative union (2 digits before last 2)
// Indexed backwards as sequential number can be 8 or 9 digits
func getFederativeUnion(voterID string) string {
	return voterID[len(voterID)-4 : len(voterID)-2]
}

// getVerifyingDigits extracts the verifying digits (last 2 digits)
func getVerifyingDigits(voterID string) string {
	return voterID[len(voterID)-2:]
}

// isFederativeUnionValid checks if the federative union code is valid (01-28)
func isFederativeUnionValid(federativeUnion string) bool {
	return helpers.Contains(validFederativeUnions, federativeUnion)
}

// calculateVD1 calculates the first verifying digit
func calculateVD1(sequentialNumber string, federativeUnion string) int {
	// Weights: 2, 3, 4, 5, 6, 7, 8, 9
	sum := 0
	for i := 0; i < 8; i++ {
		digit, _ := strconv.Atoi(string(sequentialNumber[i]))
		sum += digit * (i + 2)
	}

	rest := sum % 11
	vd1 := rest

	// Edge case: rest == 0 and federative union is SP or MG
	if rest == 0 && helpers.Contains(spMgCodes, federativeUnion) {
		vd1 = 1
	}

	// Edge case: rest == 10
	if rest == 10 {
		vd1 = 0
	}

	return vd1
}

// calculateVD2 calculates the second verifying digit
func calculateVD2(federativeUnion string, vd1 int) int {
	// Weights: 7, 8, 9
	digit0, _ := strconv.Atoi(string(federativeUnion[0]))
	digit1, _ := strconv.Atoi(string(federativeUnion[1]))

	sum := (digit0 * 7) + (digit1 * 8) + (vd1 * 9)

	rest := sum % 11
	vd2 := rest

	// Edge case: rest == 0 and federative union is SP or MG
	if rest == 0 && helpers.Contains(spMgCodes, federativeUnion) {
		vd2 = 1
	}

	// Edge case: rest == 10
	if rest == 10 {
		vd2 = 0
	}

	return vd2
}
