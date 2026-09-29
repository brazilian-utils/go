package cnpj

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

// GenerateOptions controls Generate.
type GenerateOptions struct {
	// Version picks the format: 1 (default, numeric) or 2 (alphanumeric).
	Version int
	// Branch is the número de ordem (1-9999); random when 0 or invalid.
	Branch int
}

// Generate generates a random valid CNPJ string. With no options, a
// numeric (version 1) CNPJ with a random branch is returned. The branch
// block stays numeric in both versions.
func Generate(opts ...GenerateOptions) string {
	var o GenerateOptions
	if len(opts) > 0 {
		o = opts[0]
	}

	branch := o.Branch % 10000
	if branch <= 0 {
		branch = rand.Intn(9999) + 1
	}

	if o.Version == 2 {
		return generateAlphanumeric(branch)
	}
	return generateNumeric(branch)
}

func generateNumeric(branch int) string {
	base := fmt.Sprintf("%08d%04d", rand.Intn(100000000), branch)
	return base + checksum(base)
}

const alnumChars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func generateAlphanumeric(branch int) string {
	var buf strings.Builder
	for i := 0; i < 8; i++ {
		buf.WriteByte(alnumChars[rand.Intn(len(alnumChars))])
	}
	base := fmt.Sprintf("%s%04d", buf.String(), branch)
	return base + checksumAlnum(base)
}

// checksum computes the two verifying checksum digits for a 12-digit
// numeric CNPJ base.
func checksum(base string) string {
	first := hashDigit(base, 13)
	second := hashDigit(base+strconv.Itoa(first), 14)
	return fmt.Sprintf("%d%d", first, second)
}

// checksumAlnum computes the two verifying checksum digits for a
// 12-character alphanumeric CNPJ base.
func checksumAlnum(base string) string {
	first := computeModAlnum(base, 12)
	second := computeModAlnum(base+strconv.Itoa(first), 13)
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
