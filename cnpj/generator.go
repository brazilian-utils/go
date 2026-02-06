package cnpj

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// Generate generates a random valid CNPJ string.
// The branch parameter specifies the branch number (1-9999); it defaults
// to 1 if 0 is provided.
func Generate(branch int) string {
	branch = branch % 10000
	if branch == 0 {
		branch = 1
	}

	base := fmt.Sprintf("%08d%04d", rand.Intn(100000000), branch)
	return base + checksum(base)
}

// checksum computes the two verifying checksum digits for a 12-digit CNPJ base.
func checksum(base string) string {
	first := hashDigit(base, 13)
	second := hashDigit(base+strconv.Itoa(first), 14)
	return fmt.Sprintf("%d%d", first, second)
}

// hashDigit calculates the checksum digit at the given position.
// Mirrors the Python _hashdigit function using the same weight sequence.
func hashDigit(cnpj string, position int) int {
	digits := strings.Split(cnpj[:position-1], "")
	weights := generateWeights(position)

	sum := 0
	for i, d := range digits {
		v, _ := strconv.Atoi(d)
		sum += v * weights[i]
	}

	mod := sum % 11
	if mod < 2 {
		return 0
	}
	return 11 - mod
}

// generateWeights produces the weight sequence used by the CNPJ checksum.
// For position 13: [5,4,3,2,9,8,7,6,5,4,3,2]
// For position 14: [6,5,4,3,2,9,8,7,6,5,4,3,2]
func generateWeights(position int) []int {
	var weights []int
	for i := position - 8; i >= 2; i-- {
		weights = append(weights, i)
	}
	for i := 9; i >= 2; i-- {
		weights = append(weights, i)
	}
	return weights
}
