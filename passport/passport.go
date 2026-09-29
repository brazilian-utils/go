// Package passport formats, generates and validates a Brazilian passport
// number.
package passport

import (
	"math/rand"
	"regexp"
	"strings"
)

var passportRegex = regexp.MustCompile(`^[A-Z]{2}\d{6}$`)

const passportSize = 8

// sanitize removes every non-alphanumeric character and upper-cases the
// rest.
func sanitize(value string) string {
	var b strings.Builder
	for _, r := range value {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			b.WriteRune(r)
		}
	}
	return strings.ToUpper(b.String())
}

// Format formats a passport number: upper-cased, without symbols, capped
// to 8 characters (the same operation as Parse).
func Format(value string) string {
	return Parse(value)
}

// Parse removes every non-alphanumeric character from a passport number,
// upper-cases it and caps it to 8 characters.
func Parse(value string) string {
	cleaned := sanitize(value)
	if len(cleaned) > passportSize {
		cleaned = cleaned[:passportSize]
	}
	return cleaned
}

// RemoveSymbols removes the passport formatting symbols ('.', '-'),
// keeping everything else unchanged.
func RemoveSymbols(value string) string {
	r := strings.NewReplacer(".", "", "-", "")
	return r.Replace(value)
}

// IsValid validates a Brazilian passport number: 2 letters followed by 6
// digits, after removing non-alphanumeric characters and case-folding.
// There is no check digit. value may be a string or a number for
// symmetry with Format/Parse, but a number is never valid.
func IsValid(value interface{}) bool {
	s, ok := value.(string)
	if !ok {
		return false
	}
	cleaned := sanitize(s)
	return passportRegex.MatchString(cleaned)
}

// Generate generates a random valid passport number: 2 uppercase letters
// followed by 6 digits.
func Generate() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var b strings.Builder
	b.WriteByte(letters[rand.Intn(len(letters))])
	b.WriteByte(letters[rand.Intn(len(letters))])
	for i := 0; i < 6; i++ {
		b.WriteByte(byte('0' + rand.Intn(10)))
	}
	return b.String()
}
