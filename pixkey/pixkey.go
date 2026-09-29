// Package pixkey identifies a Pix key (chave Pix) and normalizes it to
// the canonical form the DICT expects inside a BR Code.
package pixkey

import (
	"regexp"
	"strings"

	"github.com/brazilian-utils/go/cnpj"
	"github.com/brazilian-utils/go/cpf"
	"github.com/brazilian-utils/go/helpers"
	"github.com/brazilian-utils/go/phone"
)

var emailRegex = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
var uuidRegex = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// mobileRegex checks the structural shape of a Brazilian mobile number
// (DDD + 9-indicator + 8 digits) without the newer numbering-plan
// restriction on the subscriber's leading digit, matching how a Pix key
// is recognized in practice (see phone.IsValid's own base structural
// check for the same reasoning).
var mobileRegex = regexp.MustCompile(`^[1-9][1-9]9\d{8}$`)

// isMobileNumber reports whether value, once parsed like phone.Parse
// does (mask/country-code stripped), is structurally a Brazilian mobile
// number.
func isMobileNumber(value string) bool {
	return mobileRegex.MatchString(phone.Parse(value))
}

const maxEmailLength = 77

// Info is the identified Pix key.
type Info struct {
	// Type is one of "cpf", "cnpj", "email", "phone" or "evp".
	Type string
	// Value is the canonical form of the key.
	Value string
}

// looksLikePhone reports whether value carries an explicit phone marker:
// a leading "+55" or a DDD written in parentheses.
func looksLikePhone(value string) bool {
	trimmed := strings.TrimSpace(value)
	return strings.HasPrefix(trimmed, "+55") || strings.Contains(trimmed, "(")
}

// GetInfo identifies a Pix key and normalizes it to the canonical form;
// returns nil when the value is not a valid Pix key.
func GetInfo(value string) *Info {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return nil
	}

	if strings.Contains(trimmed, "@") {
		if len(trimmed) > maxEmailLength || !emailRegex.MatchString(trimmed) {
			return nil
		}
		return &Info{Type: "email", Value: strings.ToLower(trimmed)}
	}

	if uuidRegex.MatchString(trimmed) {
		return &Info{Type: "evp", Value: strings.ToLower(trimmed)}
	}

	if looksLikePhone(trimmed) {
		if isMobileNumber(trimmed) {
			return &Info{Type: "phone", Value: "+55" + phone.Parse(trimmed)}
		}
		return nil
	}

	digits := helpers.OnlyNumbers(trimmed)

	if len(digits) == 14 && cnpj.IsValid(trimmed) {
		return &Info{Type: "cnpj", Value: digits}
	}

	if len(digits) == 11 {
		if cpf.IsValid(trimmed) {
			return &Info{Type: "cpf", Value: digits}
		}
		if isMobileNumber(trimmed) {
			return &Info{Type: "phone", Value: "+55" + phone.Parse(trimmed)}
		}
	}

	return nil
}

// IsValidOptions controls IsValid.
type IsValidOptions struct {
	// Types restricts the accepted key types to this list ("cpf",
	// "cnpj", "email", "phone", "evp"). An empty (non-nil) slice rejects
	// everything. Default: every type.
	Types []string
}

// IsValid checks whether value is a valid Pix key: a CPF, a CNPJ, an
// email, a Brazilian mobile phone or a random EVP key.
func IsValid(value string, opts ...IsValidOptions) bool {
	info := GetInfo(value)
	if info == nil {
		return false
	}
	if len(opts) > 0 && opts[0].Types != nil {
		return helpers.Contains(opts[0].Types, info.Type)
	}
	return true
}
