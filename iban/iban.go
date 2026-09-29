// Package iban formats, parses and validates a Brazilian IBAN
// (International Bank Account Number).
package iban

import (
	"regexp"
	"strings"
)

const ibanSize = 29

// bareRegex matches the compact 29-character Brazilian IBAN.
var bareRegex = regexp.MustCompile(`^BR\d{2}\d{8}\d{5}\d{10}[A-Z][0-9A-Z]$`)

// groupedRegex matches groups of 4 characters separated by a single
// whitespace, '.', '-' or '/'.
var groupedSepRegex = regexp.MustCompile(`^[\sA-Z0-9.\-/]+$`)

// cleanMask removes the usual grouping separators (whitespace, '.', '-',
// '/'), keeping letters and digits; a separator inside a 4-character
// group makes the value invalid, so only whole-group separators are
// stripped by construction of the caller.
func stripSeparators(value string) string {
	r := strings.NewReplacer(" ", "", ".", "", "-", "", "/", "")
	return r.Replace(value)
}

// Parse removes IBAN formatting, keeps letters and digits upper-cased,
// capped to 29 characters.
func Parse(value string) string {
	cleaned := onlyAlnum(strings.ToUpper(value))
	if len(cleaned) > ibanSize {
		cleaned = cleaned[:ibanSize]
	}
	return cleaned
}

func onlyAlnum(value string) string {
	var b strings.Builder
	for _, r := range value {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// FormatOptions is reserved for future options (none defined yet).
type FormatOptions struct{}

// Format formats an IBAN in the ISO 13616 print grouping: blocks of 4
// characters separated by spaces, upper-cased. Does not validate. The
// result is capped at the 29 characters of a Brazilian IBAN.
func Format(value string) string {
	cleaned := Parse(value)
	if cleaned == "" {
		return ""
	}

	var b strings.Builder
	for i := 0; i < len(cleaned); i += 4 {
		if i > 0 {
			b.WriteString(" ")
		}
		end := i + 4
		if end > len(cleaned) {
			end = len(cleaned)
		}
		b.WriteString(cleaned[i:end])
	}
	return b.String()
}

// IsValid validates a Brazilian IBAN; any other country is invalid.
// Accepts the compact form or groups of 4 split by one whitespace, '.',
// '-' or '/', in any case; a separator inside a group makes it invalid.
func IsValid(value string) bool {
	upper := strings.ToUpper(strings.TrimSpace(value))
	if !groupedSepRegex.MatchString(upper) {
		return false
	}
	cleaned := stripSeparators(upper)
	if !bareRegex.MatchString(cleaned) {
		return false
	}
	return mod97(cleaned) == 1
}

// mod97 computes the ISO 7064 MOD 97-10 remainder over the IBAN: the
// first 4 characters move to the end, then each letter is replaced with
// its numeric value (A=10, ..., Z=35), and the resulting decimal string
// is reduced mod 97 digit by digit.
func mod97(iban string) int {
	rearranged := iban[4:] + iban[:4]
	var numeric strings.Builder
	for _, r := range rearranged {
		if r >= '0' && r <= '9' {
			numeric.WriteRune(r)
		} else {
			numeric.WriteString(itoa(int(r) - 'A' + 10))
		}
	}

	remainder := 0
	for _, r := range numeric.String() {
		remainder = (remainder*10 + int(r-'0')) % 97
	}
	return remainder
}

func itoa(n int) string {
	if n < 10 {
		return string(rune('0' + n))
	}
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}

// Info is the parsed Brazilian IBAN.
type Info struct {
	CountryCode string
	CheckDigits string
	BankISPB    string
	Branch      string
	Account     string
	AccountType string
	Owner       string
}

// GetInfo parses a Brazilian IBAN into its fields; returns nil whenever
// IsValid would return false.
func GetInfo(value string) *Info {
	if !IsValid(value) {
		return nil
	}
	cleaned := stripSeparators(strings.ToUpper(strings.TrimSpace(value)))

	return &Info{
		CountryCode: cleaned[0:2],
		CheckDigits: cleaned[2:4],
		BankISPB:    cleaned[4:12],
		Branch:      cleaned[12:17],
		Account:     cleaned[17:27],
		AccountType: cleaned[27:28],
		Owner:       cleaned[28:29],
	}
}
