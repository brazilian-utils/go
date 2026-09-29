package legalnature

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

const codeSize = 4

// normalize strips non-digit characters and validates the code is 4 digits.
func normalize(code string) string {
	digits := helpers.OnlyNumbers(code)
	if len(digits) != codeSize {
		return ""
	}
	return digits
}

// IsValid checks if a string corresponds to a valid Natureza Jurídica code.
// Accepts formats "NNNN" or "NNN-N". The 92 codes in force are accepted,
// plus the retired codes this dataset knows about.
func IsValid(code string) bool {
	normalized := normalize(code)
	if normalized == "" {
		return false
	}
	_, ok := byCode[normalized]
	return ok
}

// GetDescription retrieves the description of a Natureza Jurídica code.
// Returns empty string if the code is not found.
func GetDescription(code string) string {
	normalized := normalize(code)
	if normalized == "" {
		return ""
	}
	return byCode[normalized].Description
}

// ListAll returns a copy of the in-force Natureza Jurídica table
// (code -> description).
//
// Deprecated: kept for backward compatibility; prefer List.
func ListAll() map[string]string {
	return List()
}

// ListOptions controls List.
type ListOptions struct {
	// IncludeRetired also includes the retired codes this dataset knows
	// about (best-effort, not necessarily exhaustive).
	IncludeRetired bool
}

// List returns the legal nature table as a map from code to description.
// Only the 92 codes in force by default; opts[0].IncludeRetired also
// includes the retired ones.
func List(opts ...ListOptions) map[string]string {
	includeRetired := len(opts) > 0 && opts[0].IncludeRetired

	result := make(map[string]string, len(inForce)+len(retired))
	for _, ln := range inForce {
		result[ln.Code] = ln.Description
	}
	if includeRetired {
		for _, ln := range retired {
			result[ln.Code] = ln.Description
		}
	}
	return result
}

// Get looks a legal nature code up; returns nil for an unknown code.
func Get(value string) *LegalNature {
	normalized := normalize(value)
	if normalized == "" {
		return nil
	}
	ln, ok := byCode[normalized]
	if !ok {
		return nil
	}
	copied := ln
	return &copied
}

// ListByCategoryOptions controls ListByCategory.
type ListByCategoryOptions struct {
	// IncludeRetired also includes the retired codes of the category.
	IncludeRetired bool
}

// ListByCategory returns every legal nature of a CONCLA category (the
// first digit of the code), sorted by code. category may be a numeric
// string or the digit itself; a full code (e.g. "2062") is not a category
// and returns an empty list.
func ListByCategory(category string, opts ...ListByCategoryOptions) []LegalNature {
	includeRetired := len(opts) > 0 && opts[0].IncludeRetired

	category = strings.TrimSpace(category)
	if len(category) != 1 || category[0] < '0' || category[0] > '9' {
		return []LegalNature{}
	}

	result := make([]LegalNature, 0)
	for _, ln := range inForce {
		if len(ln.Code) > 0 && string(ln.Code[0]) == category {
			result = append(result, ln)
		}
	}
	if includeRetired {
		for _, ln := range retired {
			if len(ln.Code) > 0 && string(ln.Code[0]) == category {
				result = append(result, ln)
			}
		}
	}
	return result
}

// ListByCategoryNumber is a convenience overload of ListByCategory for a
// numeric category (as the contract's "string | number" parameter allows).
func ListByCategoryNumber(category int, opts ...ListByCategoryOptions) []LegalNature {
	return ListByCategory(strconv.Itoa(category), opts...)
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 4 digits first.
	Pad bool
}

// Format formats a legal nature code as "NNN-N", masking as far as the
// digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if len(opts) > 0 && opts[0].Pad && digits != "" {
		if len(digits) < codeSize {
			digits = strings.Repeat("0", codeSize-len(digits)) + digits
		}
	}
	if len(digits) > codeSize {
		digits = digits[:codeSize]
	}
	if digits == "" {
		return ""
	}

	buf := bytes.Buffer{}
	for i, c := range digits {
		if i == 3 {
			buf.WriteString("-")
		}
		buf.WriteString(fmt.Sprintf("%c", c))
	}
	return buf.String()
}

// Parse removes legal nature formatting and keeps only digits, capped to
// 4 digits.
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if len(digits) > codeSize {
		digits = digits[:codeSize]
	}
	return digits
}

// Generate generates a random valid legal nature code, drawn only among
// the codes in force.
func Generate() string {
	return inForce[rand.Intn(len(inForce))].Code
}
