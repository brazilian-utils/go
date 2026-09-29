// Package registroprofissional checks the structure of a professional
// council registration number (registro/inscrição profissional): digit
// count and UF only, never a check digit (there is no published one for
// OAB, CRM, CRO or CRP, and CREA is not covered).
package registroprofissional

import (
	"regexp"
	"strings"
)

var brazilianUFs = map[string]bool{
	"AC": true, "AL": true, "AP": true, "AM": true, "BA": true, "CE": true,
	"DF": true, "ES": true, "GO": true, "MA": true, "MT": true, "MS": true,
	"MG": true, "PA": true, "PB": true, "PR": true, "PE": true, "PI": true,
	"RJ": true, "RN": true, "RS": true, "RO": true, "RR": true, "SC": true,
	"SP": true, "SE": true, "TO": true,
}

// oabCrmRegex: 4 to 6 digits plus the UF ("123456/SP" or "123456-SP").
var oabCrmRegex = regexp.MustCompile(`^(\d{4,6})[/\-]([A-Za-z]{2})$`)

// croRegex: 3 to 6 digits plus the UF.
var croRegex = regexp.MustCompile(`^(\d{3,6})[/\-]([A-Za-z]{2})$`)

// crpRegex: a 2-digit regional code (01-24) plus 4 to 6 digits.
var crpRegex = regexp.MustCompile(`^(\d{2})/(\d{4,6})$`)

// crcRegex: UF, 6 digits, registration type (O or P) and a digit,
// optionally a transfer suffix (T-MG or S-MG).
var crcRegex = regexp.MustCompile(`^([A-Za-z]{2})-(\d{6})/([OPop])-(\d)(?:-([TSts])-([A-Za-z]{2}))?$`)

// IsValidParams carries the registration to check.
type IsValidParams struct {
	// Value is the registration number to check.
	Value string
	// Council is one of "OAB", "CRM", "CRO", "CRP" or "CRC"
	// (case-insensitive).
	Council string
	// State, when given, is matched against the registration's UF. For
	// CRP this is ignored (a CRP regional code is not a state code).
	State string
}

// IsValid checks the structure of a professional council registration
// number: digit count and UF only, never a check digit.
func IsValid(params IsValidParams) bool {
	value := strings.TrimSpace(params.Value)
	council := strings.ToUpper(strings.TrimSpace(params.Council))
	wantState := strings.ToUpper(strings.TrimSpace(params.State))

	switch council {
	case "OAB", "CRM":
		m := oabCrmRegex.FindStringSubmatch(value)
		if m == nil {
			return false
		}
		uf := strings.ToUpper(m[2])
		if !brazilianUFs[uf] {
			return false
		}
		return wantState == "" || wantState == uf

	case "CRO":
		m := croRegex.FindStringSubmatch(value)
		if m == nil {
			return false
		}
		uf := strings.ToUpper(m[2])
		if !brazilianUFs[uf] {
			return false
		}
		return wantState == "" || wantState == uf

	case "CRP":
		m := crpRegex.FindStringSubmatch(value)
		if m == nil {
			return false
		}
		region := m[1]
		return region >= "01" && region <= "24"

	case "CRC":
		m := crcRegex.FindStringSubmatch(value)
		if m == nil {
			return false
		}
		uf := strings.ToUpper(m[1])
		if !brazilianUFs[uf] {
			return false
		}
		if len(m) > 6 && m[6] != "" && !brazilianUFs[strings.ToUpper(m[6])] {
			return false
		}
		return wantState == "" || wantState == uf

	default:
		return false
	}
}
