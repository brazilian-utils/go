package cnpj

import (
	"bytes"
	"fmt"

	"github.com/brazilian-utils/go/helpers"
)

// Positions where formatting symbols are inserted: XX.XXX.XXX/XXXX-XX
var dotIndexes = []int{2, 5}
var slashIndexes = []int{8}
var hyphenIndexes = []int{12}

// Format returns the CNPJ with standard formatting: "XX.XXX.XXX/XXXX-XX".
// Returns an empty string if the input does not contain exactly 14 digits.
func Format(cnpj string) string {
	cnpjNumbers := helpers.OnlyNumbers(cnpj)

	if len(cnpjNumbers) != cnpjSize {
		return ""
	}

	return format(cnpjNumbers)
}

func format(normalized string) string {
	buf := bytes.Buffer{}
	for index, character := range normalized {
		if helpers.ContainsInt(dotIndexes, index) {
			buf.WriteString(".")
		}
		if helpers.ContainsInt(slashIndexes, index) {
			buf.WriteString("/")
		}
		if helpers.ContainsInt(hyphenIndexes, index) {
			buf.WriteString("-")
		}
		buf.WriteString(fmt.Sprintf("%c", character))
	}

	return buf.String()
}
