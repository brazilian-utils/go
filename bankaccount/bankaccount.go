// Package bankaccount validates a Brazilian bank account (bank code,
// agency, account and check digit).
//
// This is a partial, best-effort implementation: the shared contract
// documents nine banks with their own published check-digit algorithm
// (Banco do Brasil, Santander, Banrisul, Caixa, Bradesco, Nubank -
// Verhoeff, Itaú, HSBC/Kirton and Citibank) plus a list of banks
// checked by structure only, and carries zero concrete test cases to
// verify any of that against. Only the documented generic fallback
// (mod 10 / mod 11 over the account number, bankCode membership and
// basic field-shape checks) is implemented; the nine bank-specific
// algorithms and the structure-only bank list are not.
package bankaccount

import (
	"regexp"

	"github.com/brazilian-utils/go/bank"
)

var digitsRegex = regexp.MustCompile(`^\d+$`)
var checkDigitRegex = regexp.MustCompile(`^[0-9]{1,2}$|^[Xx]$|^[Pp]$`)

// IsValidParams carries the fields of a bank account.
type IsValidParams struct {
	// BankCode is the 3-digit COMPE code.
	BankCode string
	// Agency is 1 to 5 digits.
	Agency string
	// Account is 1 to 13 digits.
	Account string
	// Digit is 1 or 2 characters, or "X" (Banco do Brasil) or "P"
	// (Bradesco).
	Digit string
}

// IsValid validates a Brazilian bank account. bankCode must be a Banco
// Central STR participant; every other listed bank uses the generic
// fallback (digit must match mod 10 or mod 11 over the account; a
// 2-character digit chains mod 10 then mod 11).
func IsValid(params IsValidParams) bool {
	if !bank.IsValidCode(params.BankCode) {
		return false
	}
	if params.Agency == "" || len(params.Agency) > 5 || !digitsRegex.MatchString(params.Agency) {
		return false
	}
	if params.Account == "" || len(params.Account) > 13 || !digitsRegex.MatchString(params.Account) {
		return false
	}
	if params.Digit == "" || !checkDigitRegex.MatchString(params.Digit) {
		return false
	}

	if len(params.Digit) == 1 {
		ch := params.Digit[0]
		if ch == 'X' || ch == 'x' || ch == 'P' || ch == 'p' {
			// Banco do Brasil (X) / Bradesco (P): structural only, not
			// re-derived here (see package doc).
			return true
		}
		d := int(ch - '0')
		return d == mod10(params.Account) || d == mod11(params.Account)
	}

	d1 := int(params.Digit[0] - '0')
	d2 := int(params.Digit[1] - '0')
	m10 := mod10(params.Account)
	if d1 != m10 {
		return false
	}
	m11 := mod11(params.Account + string(rune('0'+m10)))
	return d2 == m11
}

func mod10(digits string) int {
	sum := 0
	weight := 2
	for i := len(digits) - 1; i >= 0; i-- {
		v := int(digits[i]-'0') * weight
		if v > 9 {
			v = 1 + v%10
		}
		sum += v
		if weight == 2 {
			weight = 1
		} else {
			weight = 2
		}
	}
	mod := sum % 10
	if mod == 0 {
		return 0
	}
	return 10 - mod
}

func mod11(digits string) int {
	sum := 0
	weight := 2
	for i := len(digits) - 1; i >= 0; i-- {
		sum += int(digits[i]-'0') * weight
		if weight == 9 {
			weight = 2
		} else {
			weight++
		}
	}
	mod := sum % 11
	if mod == 0 || mod == 1 {
		return 0
	}
	return 11 - mod
}
