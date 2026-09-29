// Package cbo looks up a CBO (Classificação Brasileira de Ocupações)
// code in the official CBO 2002 occupation table.
package cbo

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

//go:embed data.json
var rawData []byte

// Cbo is one entry of the official CBO 2002 occupation table.
type Cbo struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type sourceFile struct {
	Data []Cbo `json:"data"`
}

var byCode map[string]Cbo

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("cbo: failed to parse embedded data.json: " + err.Error())
	}
	byCode = make(map[string]Cbo, len(src.Data))
	for _, c := range src.Data {
		byCode[c.Code] = c
	}
}

var maskedRegex = regexp.MustCompile(`^(\d{4})[\s.\-/](\d{2})$`)
var bareDigitsRegex = regexp.MustCompile(`^\d{1,6}$`)

// normalize accepts the 6 digits, the "NNNN-NN" mask (a single separator
// between the groups) or a non-negative integer; bare digits are
// left-padded to 6, a masked value is read as written.
func normalize(value interface{}) (string, bool) {
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("%06d", v), v >= 0
	case int64:
		return fmt.Sprintf("%06d", v), v >= 0
	case string:
		trimmed := strings.TrimSpace(v)
		if m := maskedRegex.FindStringSubmatch(trimmed); m != nil {
			return m[1] + m[2], true
		}
		if bareDigitsRegex.MatchString(trimmed) {
			return fmt.Sprintf("%06s", trimmed), true
		}
		return "", false
	default:
		return "", false
	}
}

// IsValid checks whether a CBO code exists in the official table.
func IsValid(value interface{}) bool {
	code, ok := normalize(value)
	if !ok {
		return false
	}
	_, found := byCode[code]
	return found
}

// Get looks a CBO code up and returns its code and official title;
// returns nil for an unknown code or a value not in an accepted form.
func Get(value interface{}) *Cbo {
	code, ok := normalize(value)
	if !ok {
		return nil
	}
	c, found := byCode[code]
	if !found {
		return nil
	}
	copied := c
	return &copied
}

// Parse removes CBO formatting and keeps only digits, capped to 6
// digits. Nothing is left-padded.
func Parse(value interface{}) string {
	var raw string
	switch v := value.(type) {
	case string:
		raw = v
	case int:
		raw = fmt.Sprintf("%d", v)
	case int64:
		raw = fmt.Sprintf("%d", v)
	default:
		return ""
	}
	digits := helpers.OnlyNumbers(raw)
	if len(digits) > 6 {
		digits = digits[:6]
	}
	return digits
}
