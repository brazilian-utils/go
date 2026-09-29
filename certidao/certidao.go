// Package certidao formats, parses and validates the matrícula of a
// certidão de registro civil (art. 473 of the Código Nacional de Normas
// da Corregedoria Nacional de Justiça).
package certidao

import (
	"strconv"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

const certidaoSize = 32

// bookTypeNames maps the raw book type code (1-9) to its name.
var bookTypeNames = map[int]string{
	1: "birth",
	2: "marriage",
	3: "religiousMarriage",
	4: "death",
	5: "stillbirth",
	6: "banns",
	7: "other",
	8: "emancipation",
	9: "interdiction",
}

func cleanMask(value string) string {
	var b strings.Builder
	for _, r := range value {
		switch r {
		case ' ', '.', '-', '/':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b.WriteRune(r)
		default:
			return ""
		}
	}
	return b.String()
}

// dv computes the ISO-7064-like modulus 11 check digit used by the
// matrícula: weights start at 2 for the rightmost digit and increase by
// 1 leftward (no wraparound).
func dv(base string) int {
	peso := 2
	soma := 0
	for i := len(base) - 1; i >= 0; i-- {
		soma += int(base[i]-'0') * peso
		peso++
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func checkDigits(base30 string) string {
	d1 := dv(base30)
	d2 := dv(base30 + string(rune('0'+d1)))
	return string(rune('0'+d1)) + string(rune('0'+d2))
}

// IsValidOptions controls IsValid.
type IsValidOptions struct {
	// Accept narrows the valid book types to the listed raw codes
	// (1-9). Default: every type.
	Accept []int
}

// IsValid validates the 32-digit matrícula of a certidão de registro
// civil.
func IsValid(value string, opts ...IsValidOptions) bool {
	digits := helpers.OnlyNumbers(value)
	if len(digits) != certidaoSize || digits != cleanMask(value) {
		return false
	}

	service := digits[8:10]
	if service != "55" {
		return false
	}

	typeCode := int(digits[14] - '0')
	if typeCode < 1 || typeCode > 9 {
		return false
	}

	if len(opts) > 0 && opts[0].Accept != nil {
		if !helpers.ContainsInt(opts[0].Accept, typeCode) {
			return false
		}
	}

	base30 := digits[:30]
	return checkDigits(base30) == digits[30:32]
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 32 digits first.
	Pad bool
}

var groupSizes = []int{6, 2, 2, 4, 1, 5, 3, 7, 2}

// Format formats the matrícula into the printed mask of art. 473: the 32
// digits grouped as 6 2 2 4 1 5 3 7 2, separated by spaces. Masks as far
// as the digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < certidaoSize {
		digits = strings.Repeat("0", certidaoSize-len(digits)) + digits
	}
	if len(digits) > certidaoSize {
		digits = digits[:certidaoSize]
	}
	if digits == "" {
		return ""
	}

	var b strings.Builder
	pos := 0
	for gi, size := range groupSizes {
		if pos >= len(digits) {
			break
		}
		if gi > 0 {
			b.WriteString(" ")
		}
		end := pos + size
		if end > len(digits) {
			end = len(digits)
		}
		b.WriteString(digits[pos:end])
		pos = end
	}
	return b.String()
}

// Parse removes the certidão matrícula formatting and keeps only digits,
// capped to 32 digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > certidaoSize {
		digits = digits[:certidaoSize]
	}
	return digits
}

// Info is the parsed matrícula of a certidão de registro civil.
type Info struct {
	RegistryCNS  string
	Acervo       string
	Service      string
	Year         int
	Type         string
	TypeCode     int
	Book         string
	Page         string
	Term         string
	CheckDigits  string
}

// GetInfo parses the matrícula into its fields; returns nil when it is
// not valid (same rules as IsValid).
func GetInfo(value string) *Info {
	if !IsValid(value) {
		return nil
	}
	digits := helpers.OnlyNumbers(value)

	year, _ := strconv.Atoi(digits[10:14])
	typeCode := int(digits[14] - '0')

	return &Info{
		RegistryCNS: digits[0:6],
		Acervo:      digits[6:8],
		Service:     digits[8:10],
		Year:        year,
		Type:        bookTypeNames[typeCode],
		TypeCode:    typeCode,
		Book:        digits[15:20],
		Page:        digits[20:23],
		Term:        digits[23:30],
		CheckDigits: digits[30:32],
	}
}
