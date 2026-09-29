package pis

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

const pisSize = 11

var weights = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// IsValid checks if a PIS (Programa de Integração Social) number is valid.
// A number whose digits are all the same is rejected even if the check
// digit happens to match (same rule as CPF/CNPJ).
func IsValid(pis string) bool {
	cleaned := helpers.OnlyNumbers(pis)
	if len(cleaned) != pisSize || cleaned != pis {
		return false
	}

	if allSameDigit(cleaned) {
		return false
	}

	expected := checksum(cleaned[:10])
	return int(cleaned[10]-'0') == expected
}

func allSameDigit(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

// Format formats a valid PIS into "NNN.NNNNN.NN-N".
// Returns empty string if invalid.
func Format(pis string) string {
	if !IsValid(pis) {
		return ""
	}
	return fmt.Sprintf("%s.%s.%s-%s", pis[:3], pis[3:8], pis[8:10], pis[10:11])
}

// Generate generates a random valid PIS number.
func Generate() string {
	base := fmt.Sprintf("%010d", rand.Intn(10000000000))
	return base + fmt.Sprintf("%d", checksum(base))
}

// RemoveSymbols removes the PIS formatting symbols ('.', '-'), keeping
// everything else unchanged.
func RemoveSymbols(value string) string {
	r := strings.NewReplacer(".", "", "-", "")
	return r.Replace(value)
}

// Parse removes PIS formatting and keeps only digits, capped to 11 digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > pisSize {
		digits = digits[:pisSize]
	}
	return digits
}

// checksum computes the check digit for the first 10 digits of a PIS.
func checksum(base string) int {
	sum := 0
	for i, w := range weights {
		sum += int(base[i]-'0') * w
	}
	digit := 11 - (sum % 11)
	if digit >= 10 {
		return 0
	}
	return digit
}
