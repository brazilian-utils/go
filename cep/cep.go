package cep

import (
	"bytes"
	"fmt"
	"math/rand"

	"github.com/brazilian-utils/brutils-go/helpers"
)

// Every CEP has exactly 8 digits
const cepSize = 8

// Position where the hyphen is inserted during formatting
var hyphenIndexes = []int{5}

// IsValid checks if a given CEP (Postal Code) is valid.
// A valid CEP is a string containing exactly 8 digits.
func IsValid(cep string) bool {
	cleaned := helpers.OnlyNumbers(cep)
	return len(cleaned) == cepSize && cleaned == cep
}

// Format formats a CEP string into the standard "XXXXX-XXX" format.
// Returns an empty string if the input is not a valid 8-digit CEP.
func Format(cep string) string {
	cleaned := helpers.OnlyNumbers(cep)
	if len(cleaned) != cepSize {
		return ""
	}

	buf := bytes.Buffer{}
	for index, character := range cleaned {
		if helpers.ContainsInt(hyphenIndexes, index) {
			buf.WriteString("-")
		}
		buf.WriteString(fmt.Sprintf("%c", character))
	}

	return buf.String()
}

// Generate generates a random 8-digit CEP string.
func Generate() string {
	buf := bytes.Buffer{}
	for i := 0; i < cepSize; i++ {
		buf.WriteString(fmt.Sprintf("%d", rand.Intn(10)))
	}
	return buf.String()
}
