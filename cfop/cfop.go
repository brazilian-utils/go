// Package cfop looks up a CFOP (Código Fiscal de Operações e
// Prestações) code in the consolidated Anexo II of Convênio SINIEF s/nº
// 1970 in force.
package cfop

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/brazilian-utils/go/helpers"
)

//go:embed data.json
var rawData []byte

// Cfop is one entry of the official CFOP table.
type Cfop struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type sourceFile struct {
	Data []Cfop `json:"data"`
}

var byCode map[string]Cfop

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("cfop: failed to parse embedded data.json: " + err.Error())
	}
	byCode = make(map[string]Cfop, len(src.Data))
	for _, c := range src.Data {
		byCode[c.Code] = c
	}
}

var maskRegex = regexp.MustCompile(`^(\d)[\s.\-/]?(\d{3})$`)

// normalize accepts the bare 4 digits, the N.NNN form with a single
// separator, or a non-negative integer, and returns the bare code.
func normalize(value interface{}) (string, bool) {
	var raw string
	switch v := value.(type) {
	case string:
		raw = v
	case int:
		raw = fmt.Sprintf("%d", v)
	case int64:
		raw = fmt.Sprintf("%d", v)
	default:
		return "", false
	}

	if len(raw) == 4 && helpers.OnlyNumbers(raw) == raw {
		return raw, true
	}
	if m := maskRegex.FindStringSubmatch(raw); m != nil {
		return m[1] + m[2], true
	}
	return "", false
}

// IsValid checks whether a CFOP code exists in the official table. Group
// and subgroup headings are not operable codes and are rejected.
func IsValid(value interface{}) bool {
	code, ok := normalize(value)
	if !ok {
		return false
	}
	_, found := byCode[code]
	return found
}

// Get looks a CFOP code up and returns its code and description; returns
// nil for a heading, an unknown code or a value not in an accepted form.
func Get(value interface{}) *Cfop {
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

// Parse removes CFOP formatting and keeps only digits, capped to 4
// digits. Nothing is padded.
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
	if len(digits) > 4 {
		digits = digits[:4]
	}
	return digits
}
