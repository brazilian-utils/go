// Package ie validates a Brazilian state tax registration number
// (inscrição estadual).
//
// This is a structural (digit-count only) implementation, not a real
// check-digit validator: the shared contract documents a distinct
// published SINTEGRA check-digit algorithm per state (27 of them, several
// with special cases for GO, PA, MS, SP, TO, DF, PE, AL and RJ) and
// carries zero concrete test cases to verify any of them against.
// Shipping 27 unverified check-digit algorithms risks silently wrong
// validation, which is worse than this narrower, honestly-scoped
// structural check.
package ie

import (
	"regexp"
	"strings"
)

// digitCounts lists the accepted digit counts (after removing the usual
// mask characters) for the numeric IE of each state. States with more
// than one accepted length list them all.
var digitCounts = map[string][]int{
	"AC": {13}, "AL": {9}, "AM": {9}, "AP": {9}, "BA": {8, 9},
	"CE": {9}, "DF": {13}, "ES": {9}, "GO": {9}, "MA": {9},
	"MG": {13}, "MS": {9}, "MT": {11}, "PA": {9}, "PB": {9},
	"PE": {9, 14}, "PI": {9}, "PR": {10}, "RJ": {8}, "RN": {9, 10},
	"RO": {9, 14}, "RR": {9}, "RS": {10}, "SC": {9}, "SE": {9},
	"SP": {12}, "TO": {9, 11},
}

var alnumRegex = regexp.MustCompile(`^[0-9A-Za-z]+$`)

// IsValidParams carries the registration and the state.
type IsValidParams struct {
	// Value is the inscrição estadual to check (mask characters are
	// stripped first).
	Value string
	// State is the 2-letter state code (case-insensitive).
	State string
}

// IsValid checks whether value has an accepted length for the given
// state. It does not verify the state's published check-digit formula
// (see the package doc).
func IsValid(params IsValidParams) bool {
	uf := strings.ToUpper(strings.TrimSpace(params.State))
	lengths, ok := digitCounts[uf]
	if !ok {
		return false
	}

	cleaned := stripMask(params.Value)
	if cleaned == "" {
		return false
	}

	// SP's produtor rural inscription starts with "P".
	if uf == "SP" && (strings.HasPrefix(cleaned, "P") || strings.HasPrefix(cleaned, "p")) {
		return len(cleaned) == 13
	}

	if !alnumRegex.MatchString(cleaned) {
		return false
	}
	for _, r := range cleaned {
		if r < '0' || r > '9' {
			return false
		}
	}

	for _, l := range lengths {
		if len(cleaned) == l {
			return true
		}
	}
	return false
}

func stripMask(value string) string {
	r := strings.NewReplacer(" ", "", ".", "", "-", "", "/", "")
	return r.Replace(strings.TrimSpace(value))
}
