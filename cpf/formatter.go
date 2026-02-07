package cpf

import (
	"bytes"
	"fmt"

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
