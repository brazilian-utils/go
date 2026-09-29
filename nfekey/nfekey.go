// Package nfekey formats, parses and validates a DF-e (Documento Fiscal
// eletrônico) access key (chave de acesso): NF-e, NFC-e, CT-e, MDF-e,
// CT-e OS, GTV-e, BP-e, NF3e and NFCom.
package nfekey

import (
	"strconv"
	"strings"

	"github.com/brazilian-utils/go/helpers"
	"github.com/brazilian-utils/go/state"
)

const keySize = 44

// validModels: NF-e (55), NFC-e (65), CT-e (57), MDF-e (58), CT-e OS
// (67), GTV-e (64), BP-e (63), NF3e (66) and NFCom (62). CF-e-SAT (59)
// is out.
var validModels = map[string]bool{
	"55": true, "65": true, "57": true, "58": true,
	"67": true, "64": true, "63": true, "66": true, "62": true,
}

// idPrefixes are the XML Id prefixes stripped before parsing.
var idPrefixes = []string{"NFe", "CTe", "MDFe", "BPe", "NF3e", "NFCom"}

func stripPrefix(value string) string {
	for _, p := range idPrefixes {
		if strings.HasPrefix(value, p) {
			return value[len(p):]
		}
	}
	return value
}

// Parse removes the formatting of a DF-e access key and keeps only
// digits, capped to 44 digits. The XML Id prefixes are stripped first.
func Parse(value string) string {
	cleaned := stripPrefix(value)
	digits := helpers.OnlyNumbers(cleaned)
	if len(digits) > keySize {
		digits = digits[:keySize]
	}
	return digits
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 44 digits first.
	Pad bool
}

// Format formats a DF-e access key into groups of 4 digits separated by
// spaces, as the DANFE and the other auxiliary documents print it. Does
// not validate. Masks as far as the digits go.
func Format(value string, opts ...FormatOptions) string {
	cleaned := stripPrefix(value)
	digits := helpers.OnlyNumbers(cleaned)
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < keySize {
		digits = strings.Repeat("0", keySize-len(digits)) + digits
	}
	if len(digits) > keySize {
		digits = digits[:keySize]
	}
	if digits == "" {
		return ""
	}

	var b strings.Builder
	for i := 0; i < len(digits); i += 4 {
		if i > 0 {
			b.WriteString(" ")
		}
		end := i + 4
		if end > len(digits) {
			end = len(digits)
		}
		b.WriteString(digits[i:end])
	}
	return b.String()
}

// checkDigit computes the modulus 11 check digit over the first 43
// digits: weights cycle 2..9 starting from the rightmost digit.
func checkDigit(base43 string) int {
	peso := 2
	soma := 0
	for i := len(base43) - 1; i >= 0; i-- {
		soma += int(base43[i]-'0') * peso
		if peso == 9 {
			peso = 2
		} else {
			peso++
		}
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

// IsValid validates a 44-digit DF-e access key: the cUF must be a
// state, the model one of the supported ones, the document number not
// all zeros, the emission type (tpEmis) in 1..9, the numeric code
// (cNF, models 55/65 only per MOC rule B03-10) not trivially guessable,
// and the check digit (modulus 11 over the first 43 digits) must match.
func IsValid(value string) bool {
	digits := Parse(value)
	rawCleaned := helpers.OnlyNumbers(stripPrefix(value))
	if len(digits) != keySize || rawCleaned != digits {
		return false
	}

	cUF := digits[0:2]
	if state.GetByIbgeCode(cUF) == nil {
		return false
	}

	model := digits[20:22]
	if !validModels[model] {
		return false
	}

	number := digits[25:34]
	if allZero(number) {
		return false
	}

	tpEmis := digits[34:35]
	if tpEmis < "1" || tpEmis > "9" {
		return false
	}

	cNF := digits[35:43]
	if (model == "55" || model == "65") && !cnfOK(cNF, number) {
		return false
	}

	base43 := digits[:43]
	expected := checkDigit(base43)
	actual := int(digits[43] - '0')
	return expected == actual
}

func allZero(s string) bool {
	for _, c := range s {
		if c != '0' {
			return false
		}
	}
	return true
}

// cnfOK implements MOC rule B03-10: the numeric code (cNF) must not be
// all-repeated digits, strictly ascending/descending sequential (mod
// 10, so it wraps, e.g. "90123456" also counts as ascending), or equal
// to the document number (nNF).
func cnfOK(cnf, nnf string) bool {
	cnfInt, err1 := strconv.Atoi(cnf)
	nnfInt, err2 := strconv.Atoi(nnf)
	if err1 == nil && err2 == nil && cnfInt == nnfInt {
		return false
	}

	allSame := true
	for i := 1; i < len(cnf); i++ {
		if cnf[i] != cnf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	ascending := true
	descending := true
	for i := 1; i < len(cnf); i++ {
		a := int(cnf[i-1] - '0')
		b := int(cnf[i] - '0')
		if b != (a+1)%10 {
			ascending = false
		}
		if b != (a-1+10)%10 {
			descending = false
		}
	}

	return !(ascending || descending)
}

// Info is the parsed DF-e access key.
type Info struct {
	StateCode    string
	Year         int
	Month        int
	TaxID        string
	Model        string
	Series       int
	Number       int
	EmissionType int
	Code         string
	CheckDigit   int
}

// GetInfo parses the access key into its fields; returns nil when the
// key is not valid (same rules as IsValid).
func GetInfo(value string) *Info {
	if !IsValid(value) {
		return nil
	}
	digits := Parse(value)

	cUF := digits[0:2]
	st := state.GetByIbgeCode(cUF)
	stateCode := ""
	if st != nil {
		stateCode = st.Code
	}

	year, _ := strconv.Atoi(digits[2:4])
	month, _ := strconv.Atoi(digits[4:6])
	series, _ := strconv.Atoi(digits[22:25])
	number, _ := strconv.Atoi(digits[25:34])
	emissionType, _ := strconv.Atoi(digits[34:35])
	checkDigitVal := int(digits[43] - '0')

	return &Info{
		StateCode:    stateCode,
		Year:         2000 + year,
		Month:        month,
		TaxID:        digits[6:20],
		Model:        digits[20:22],
		Series:       series,
		Number:       number,
		EmissionType: emissionType,
		Code:         digits[35:43],
		CheckDigit:   checkDigitVal,
	}
}
