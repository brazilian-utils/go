package cpf

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

var dotIndexes = []int{3, 6}
var hyphenIndexes = []int{9}

// Format returns the CPF with all the "." and "-"
func Format(cpf string) string {
	cpfNumbers := helpers.OnlyNumbers(cpf)

	// Truncate the CPF to its maximum size
	if len(cpfNumbers) > cpfSize {
		cpfNumbers = cpfNumbers[:cpfSize]
	}

	return format(cpfNumbers)
}

func format(nomalizedCpf string) string {
	buf := bytes.Buffer{}
	for index, character := range nomalizedCpf {
		if helpers.ContainsInt(dotIndexes, index) {
			buf.WriteString(".")
		}
		if helpers.ContainsInt(hyphenIndexes, index) {
			buf.WriteString("-")
		}
		buf.WriteString(fmt.Sprintf("%c", character))
	}

	return buf.String()
}

// RemoveSymbols removes the CPF formatting symbols ('.', '-'), keeping
// everything else unchanged.
func RemoveSymbols(value string) string {
	r := strings.NewReplacer(".", "", "-", "")
	return r.Replace(value)
}

// Parse removes CPF formatting and keeps only digits, capped to 11 digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > cpfSize {
		digits = digits[:cpfSize]
	}
	return digits
}
