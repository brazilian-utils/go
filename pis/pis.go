package pis

import (
	"fmt"
	"math/rand"

	"github.com/brazilian-utils/brutils-go/helpers"
)

const pisSize = 11

var weights = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// IsValid checks if a PIS (Programa de Integração Social) number is valid.
func IsValid(pis string) bool {
	cleaned := helpers.OnlyNumbers(pis)
	if len(cleaned) != pisSize || cleaned != pis {
		return false
	}

	expected := checksum(cleaned[:10])
	return int(cleaned[10]-'0') == expected
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
