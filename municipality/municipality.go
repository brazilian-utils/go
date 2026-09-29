// Package municipality provides lookups over the Brazilian municipalities
// published by the IBGE.
package municipality

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/brazilian-utils/go/helpers"
)

//go:embed data.json
var rawData []byte

const codeSize = 7

// Municipality is one Brazilian municipality.
type Municipality struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	StateCode string `json:"stateCode"`
}

type sourceFile struct {
	Data []Municipality `json:"data"`
}

var (
	all    []Municipality
	byCode map[string]Municipality
)

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("municipality: failed to parse embedded data.json: " + err.Error())
	}
	all = src.Data

	byCode = make(map[string]Municipality, len(all))
	for _, m := range all {
		byCode[m.Code] = m
	}
}

// GetByCode looks a municipality up by its 7-digit IBGE code. code may
// be a string or a non-negative integer. Returns nil when the code is
// not 7 digits long or matches no municipality.
func GetByCode(code interface{}) *Municipality {
	var digits string
	switch v := code.(type) {
	case int:
		digits = fmt.Sprintf("%d", v)
	case int64:
		digits = fmt.Sprintf("%d", v)
	case string:
		digits = helpers.OnlyNumbers(v)
	default:
		return nil
	}
	if len(digits) != codeSize {
		return nil
	}
	m, ok := byCode[digits]
	if !ok {
		return nil
	}
	copied := m
	return &copied
}

// List returns the Brazilian municipalities, sorted by name: every one
// of them when stateCode is omitted, or those of one state (matched
// case-sensitively) otherwise. An empty or unknown stateCode returns an
// empty list.
func List(stateCode ...string) []Municipality {
	if len(stateCode) == 0 {
		result := make([]Municipality, len(all))
		copy(result, all)
		return result
	}

	code := stateCode[0]
	if code == "" {
		return []Municipality{}
	}

	result := make([]Municipality, 0)
	for _, m := range all {
		if m.StateCode == code {
			result = append(result, m)
		}
	}
	return result
}
