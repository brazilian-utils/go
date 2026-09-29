// Package caepf formats, parses and validates a CAEPF (Cadastro de
// Atividade Econômica da Pessoa Física) number.
package caepf

import (
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

const caepfSize = 14

var cnpjWeights1 = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
var cnpjWeights2 = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// IsValid validates a CAEPF: 14 digits, the 9-digit CPF base of the
// holder, a 3-digit sequence and 2 check digits (CNPJ modulus 11,
// shifted by 12 mod 100). A base whose digits are all the same is
// rejected.
func IsValid(value string) bool {
	digits := helpers.OnlyNumbers(value)
	if len(digits) != caepfSize {
		return false
	}

	base12 := digits[:12]
	if allSame(base12[:9]) {
		return false
	}

	d1 := cnpjMod(base12, cnpjWeights1)
	d2 := cnpjMod(base12+itoa(d1), cnpjWeights2)

	dd := (d1*10 + d2 + 12) % 100
	actual := int(digits[12]-'0')*10 + int(digits[13]-'0')

	return dd == actual
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 14 digits first.
	Pad bool
}

// Format formats a CAEPF with the mask "000.000.000/000-00", masking as
// far as the digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < caepfSize {
		digits = strings.Repeat("0", caepfSize-len(digits)) + digits
	}
	if len(digits) > caepfSize {
		digits = digits[:caepfSize]
	}
	if digits == "" {
		return ""
	}

	positions := []struct {
		idx int
		sep string
	}{
		{3, "."}, {6, "."}, {9, "/"}, {12, "-"},
	}

	var b strings.Builder
	for i, c := range digits {
		for _, p := range positions {
			if i == p.idx {
				b.WriteString(p.sep)
			}
		}
		b.WriteRune(c)
	}
	return b.String()
}

// Parse removes CAEPF formatting and keeps only digits, capped to 14
// digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > caepfSize {
		digits = digits[:caepfSize]
	}
	return digits
}

func allSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}

func cnpjMod(chars string, weights []int) int {
	sum := 0
	for i := 0; i < len(chars); i++ {
		sum += int(chars[i]-'0') * weights[i]
	}
	mod := sum % 11
	if mod < 2 {
		return 0
	}
	return 11 - mod
}

func itoa(n int) string {
	return string(rune('0' + n))
}
