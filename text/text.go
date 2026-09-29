// Package text implements Brazilian-Portuguese-aware text helpers:
// capitalizing names/addresses and removing diacritics.
package text

import (
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// RemoveAccents removes diacritical marks (accents, tildes, cedillas)
// from a string: every character is decomposed (Unicode NFD) and every
// combining mark (general category M) is dropped, so accents from any
// script go.
func RemoveAccents(value string) string {
	decomposed := norm.NFD.String(value)
	var b strings.Builder
	b.Grow(len(decomposed))
	for _, r := range decomposed {
		if unicode.Is(unicode.M, r) {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// defaultPrepositions stay lower case between two words.
var defaultPrepositions = map[string]bool{
	"de": true, "da": true, "do": true, "das": true, "dos": true,
	"e": true, "em": true, "a": true, "o": true, "as": true, "os": true,
	"no": true, "na": true, "nos": true, "nas": true,
}

// defaultDesignations are upper-cased in full. "SA" without dots is
// intentionally absent: it is left alone (the surname "Sá").
var defaultDesignations = map[string]bool{
	"LTDA": true, "S.A.": true, "ME": true, "EPP": true, "EIRELI": true,
	"CNPJ": true, "CPF": true, "MEI": true,
}

var brazilianUFs = map[string]bool{
	"AC": true, "AL": true, "AP": true, "AM": true, "BA": true, "CE": true,
	"DF": true, "ES": true, "GO": true, "MA": true, "MT": true, "MS": true,
	"MG": true, "PA": true, "PB": true, "PR": true, "PE": true, "PI": true,
	"RJ": true, "RN": true, "RS": true, "RO": true, "RR": true, "SC": true,
	"SP": true, "SE": true, "TO": true,
}

// CapitalizeOptions controls Capitalize.
type CapitalizeOptions struct {
	// Prepositions, when non-nil, replaces the default list of words that
	// stay lower case between two other words.
	Prepositions map[string]bool
	// Designations, when non-nil, replaces the default list of words that
	// are always upper-cased in full.
	Designations map[string]bool
}

func isDelimiter(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r' ||
		r == '-' || r == '/' || r == '\''
}

func isRomanNumeral(word string) bool {
	if len(word) < 2 {
		return false
	}
	for _, r := range strings.ToUpper(word) {
		switch r {
		case 'I', 'V', 'X', 'L', 'C', 'D', 'M':
		default:
			return false
		}
	}
	return true
}

// Capitalize capitalizes a string the way a Brazilian name, company name
// or address is written: the first letter of each word, with
// prepositions/articles kept lower case between two words, recognized
// company designations and roman numerals upper-cased, and a state code
// right after a '/' upper-cased.
func Capitalize(value string, opts ...CapitalizeOptions) string {
	if value == "" {
		return ""
	}

	prepositions := defaultPrepositions
	designations := defaultDesignations
	if len(opts) > 0 {
		if opts[0].Prepositions != nil {
			prepositions = opts[0].Prepositions
		}
		if opts[0].Designations != nil {
			designations = opts[0].Designations
		}
	}

	runes := []rune(value)

	// Tokenize into words (runs of non-delimiter chars) and delimiters,
	// remembering, for each word, the delimiter immediately before it and
	// whether it's the first/last word.
	type token struct {
		text      string
		isWord    bool
		prevDelim rune
	}

	var tokens []token
	var cur strings.Builder
	var prevDelim rune
	flush := func() {
		if cur.Len() > 0 {
			tokens = append(tokens, token{text: cur.String(), isWord: true, prevDelim: prevDelim})
			cur.Reset()
		}
	}

	i := 0
	for i < len(runes) {
		r := runes[i]
		if isDelimiter(r) {
			flush()
			if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
				// collapse whitespace runs into a single space
				for i < len(runes) && (runes[i] == ' ' || runes[i] == '\t' || runes[i] == '\n' || runes[i] == '\r') {
					i++
				}
				tokens = append(tokens, token{text: " "})
				prevDelim = ' '
			} else {
				tokens = append(tokens, token{text: string(r)})
				prevDelim = r
				i++
			}
			continue
		}
		cur.WriteRune(r)
		i++
	}
	flush()

	// Index of first/last word tokens.
	firstWord, lastWord := -1, -1
	for idx, t := range tokens {
		if t.isWord {
			if firstWord == -1 {
				firstWord = idx
			}
			lastWord = idx
		}
	}

	var out strings.Builder
	for idx, t := range tokens {
		if !t.isWord {
			out.WriteString(t.text)
			continue
		}

		word := t.text
		upper := strings.ToUpper(word)

		switch {
		case designations[upper]:
			out.WriteString(upper)
		case t.prevDelim == '/' && brazilianUFs[upper]:
			out.WriteString(upper)
		case isRomanNumeral(word):
			out.WriteString(upper)
		case prepositions[strings.ToLower(word)] && idx != firstWord && idx != lastWord:
			out.WriteString(strings.ToLower(word))
		default:
			out.WriteString(titleCaseWord(word))
		}
	}

	return out.String()
}

// titleCaseWord upper-cases the first rune (if it is a letter) and
// lower-cases the rest.
func titleCaseWord(word string) string {
	runes := []rune(word)
	if len(runes) == 0 {
		return word
	}

	var b strings.Builder
	b.WriteRune(unicode.ToUpper(runes[0]))
	for _, r := range runes[1:] {
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
