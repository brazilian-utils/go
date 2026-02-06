package legalprocess

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const processSize = 20

// randIntn wraps rand.Intn for use across the package.
var randIntn = rand.Intn

// IsValid checks if a legal process ID is valid.
// Validates format (20 digits), checksum, and orgão/tribunal/foro codes.
func IsValid(legalProcessID string) bool {
	clean := removeSymbols(legalProcessID)

	if len(clean) != processSize || !isDigits(clean) {
		return false
	}

	dd := clean[7:9]
	j, _ := strconv.Atoi(clean[13:14])
	tr, _ := strconv.Atoi(clean[14:16])
	oooo, _ := strconv.Atoi(clean[16:20])

	orgao, ok := orgaos[j]
	if !ok {
		return false
	}

	if !orgao.tribunals[tr] || !orgao.foros[oooo] {
		return false
	}

	// Checksum: base = NNNNNNN + YYYYJTTOOOO (digits without DD)
	base := clean[0:7] + clean[9:20]
	return checksum(base) == dd
}

// Format formats a 20-digit legal process ID into the standard
// "NNNNNNN-DD.YYYY.J.TT.OOOO" format. Returns empty string if invalid.
func Format(legalProcessID string) string {
	if len(legalProcessID) != processSize || !isDigits(legalProcessID) {
		return ""
	}

	return fmt.Sprintf("%s-%s.%s.%s.%s.%s",
		legalProcessID[0:7],
		legalProcessID[7:9],
		legalProcessID[9:13],
		legalProcessID[13:14],
		legalProcessID[14:16],
		legalProcessID[16:20],
	)
}

// Generate generates a random valid legal process ID.
// year must not be in the past; orgao must be 1-9.
// Returns empty string if arguments are invalid.
func Generate(year int, orgao int) string {
	if year < time.Now().Year() || orgao < 1 || orgao > 9 {
		return ""
	}

	data, ok := orgaos[orgao]
	if !ok {
		return ""
	}

	tr := randomKey(data.tribunals)
	oooo := randomKey(data.foros)
	nnnnnnn := fmt.Sprintf("%07d", randIntn(10000000))

	base := fmt.Sprintf("%s%04d%d%02d%04d", nnnnnnn, year, orgao, tr, oooo)
	dd := checksum(base)

	return fmt.Sprintf("%s%s%04d%d%02d%04d", nnnnnnn, dd, year, orgao, tr, oooo)
}

// checksum computes the 2-digit verification code.
// Formula: 97 - ((basenum * 100) % 97), zero-padded.
func checksum(basenum string) string {
	// Compute basenum % 97 digit by digit to avoid overflow
	remainder := 0
	for _, ch := range basenum {
		remainder = (remainder*10 + int(ch-'0')) % 97
	}
	// (basenum * 100) % 97 = (remainder * 100) % 97
	result := 97 - (remainder*100)%97
	return fmt.Sprintf("%02d", result)
}

// removeSymbols strips '.' and '-' from a legal process ID string.
func removeSymbols(s string) string {
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, "-", "")
	return s
}

// isDigits checks if a string contains only digit characters.
func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}
