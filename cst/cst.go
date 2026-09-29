// Package cst validates a CST (Código de Situação Tributária) code
// against the ICMS (Tabela A + Tabela B), IPI and PIS/COFINS tables.
package cst

import (
	"fmt"
	"regexp"
)

var icmsOrigins = map[string]bool{
	"0": true, "1": true, "2": true, "3": true, "4": true,
	"5": true, "6": true, "7": true, "8": true,
}

var icmsCstCodes = map[string]bool{
	"00": true, "10": true, "20": true, "30": true, "40": true, "41": true,
	"50": true, "51": true, "60": true, "70": true, "90": true,
}

var ipiCstCodes = map[string]bool{
	"00": true, "01": true, "02": true, "03": true, "04": true, "05": true,
	"49": true, "50": true, "51": true, "52": true, "53": true, "54": true,
	"55": true, "99": true,
}

var pisCofinsCstCodes = map[string]bool{
	"01": true, "02": true, "03": true, "04": true, "05": true, "06": true,
	"07": true, "08": true, "09": true, "49": true, "50": true, "51": true,
	"52": true, "53": true, "54": true, "55": true, "56": true, "60": true,
	"61": true, "62": true, "63": true, "64": true, "65": true, "66": true,
	"67": true, "70": true, "71": true, "72": true, "73": true, "74": true,
	"75": true, "98": true, "99": true,
}

var icmsMaskRegex = regexp.MustCompile(`^(\d)[\s.\-/]?(\d{2})$`)
var twoDigitsRegex = regexp.MustCompile(`^\d{2}$`)
var oneDigitRegex = regexp.MustCompile(`^\d$`)

// IsValidOptions controls IsValid.
type IsValidOptions struct {
	// Tax picks the table: "icms", "ipi", "pis" or "cofins". Omitted or
	// unknown, every table is accepted.
	Tax string
}

// IsValid checks whether value is valid for a tax's CST table. Accepts a
// string or a non-negative integer; the 3-digit ICMS form may have a
// single separator after the origin digit. A single digit is padded to
// the 3-digit ICMS form; a 2-digit string is a Tabela B code (IPI or
// PIS/COFINS).
func IsValid(value interface{}, opts ...IsValidOptions) bool {
	tax := ""
	if len(opts) > 0 {
		tax = opts[0].Tax
	}

	raw, ok := toRawString(value)
	if !ok {
		return false
	}

	checkICMS := tax == "" || tax == "icms"
	checkIPI := tax == "" || tax == "ipi"
	checkPISCofins := tax == "" || tax == "pis" || tax == "cofins"

	if oneDigitRegex.MatchString(raw) {
		raw = "00" + raw
	}

	if checkICMS {
		if m := icmsMaskRegex.FindStringSubmatch(raw); m != nil {
			if icmsOrigins[m[1]] && icmsCstCodes[m[2]] {
				return true
			}
		}
	}

	if twoDigitsRegex.MatchString(raw) {
		if checkIPI && ipiCstCodes[raw] {
			return true
		}
		if checkPISCofins && pisCofinsCstCodes[raw] {
			return true
		}
	}

	return false
}

func toRawString(value interface{}) (string, bool) {
	switch v := value.(type) {
	case string:
		if v == "" {
			return "", false
		}
		return v, true
	case int:
		return fmt.Sprintf("%d", v), true
	case int64:
		return fmt.Sprintf("%d", v), true
	default:
		return "", false
	}
}
