// Package bank provides lookups over the Brazilian banks that carry a
// COMPE code, published by Banco Central do Brasil in the STR
// participants list.
package bank

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data.json
var rawData []byte

// Bank is one COMPE-bearing STR participant.
type Bank struct {
	Code string `json:"code"`
	ISPB string `json:"ispb"`
	Name string `json:"name"`
}

type sourceFile struct {
	Data []Bank `json:"data"`
}

var (
	all      []Bank
	byCode   map[string]Bank
	byISPB   map[string]Bank
)

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("bank: failed to parse embedded data.json: " + err.Error())
	}
	all = src.Data

	byCode = make(map[string]Bank, len(all))
	byISPB = make(map[string]Bank, len(all))
	for _, b := range all {
		byCode[b.Code] = b
		byISPB[b.ISPB] = b
	}
}

// List returns every bank with a COMPE code.
func List() []Bank {
	result := make([]Bank, len(all))
	copy(result, all)
	return result
}

// GetByCode looks a bank up by its 3-digit COMPE code. code may be a
// string or a number (zero-padded to 3 digits). Returns nil when no bank
// has that code.
func GetByCode(code interface{}) *Bank {
	s, ok := normalizeCode(code, 3)
	if !ok {
		return nil
	}
	b, ok := byCode[s]
	if !ok {
		return nil
	}
	copied := b
	return &copied
}

// GetByIspb looks a bank up by its 8-digit ISPB. value may be a string
// or a number, with or without leading zeros. Returns nil when no bank
// with a COMPE code has that ISPB.
func GetByIspb(value interface{}) *Bank {
	s, ok := normalizeCode(value, 8)
	if !ok {
		return nil
	}
	b, ok := byISPB[s]
	if !ok {
		return nil
	}
	copied := b
	return &copied
}

// IsValidCode reports whether code is a known COMPE code (used by
// bankaccount to validate a bank code).
func IsValidCode(code string) bool {
	_, ok := byCode[code]
	return ok
}

func normalizeCode(value interface{}, width int) (string, bool) {
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("%0*d", width, v), true
	case int64:
		return fmt.Sprintf("%0*d", width, v), true
	case float64:
		return fmt.Sprintf("%0*d", width, int64(v)), true
	case string:
		trimmed := strings.TrimSpace(v)
		if trimmed == "" {
			return "", false
		}
		n, err := strconv.Atoi(trimmed)
		if err != nil {
			return "", false
		}
		return fmt.Sprintf("%0*d", width, n), true
	default:
		return "", false
	}
}
