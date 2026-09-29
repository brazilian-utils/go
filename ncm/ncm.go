// Package ncm formats, parses and validates a NCM (Nomenclatura Comum do
// Mercosul) code against the current table published by Siscomex.
package ncm

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

const ncmSize = 8

type entry struct {
	Code string `json:"code"`
}

type sourceFile struct {
	Data []entry `json:"data"`
}

var validCodes map[string]bool

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("ncm: failed to parse embedded data.json: " + err.Error())
	}
	validCodes = make(map[string]bool, len(src.Data))
	for _, e := range src.Data {
		validCodes[e.Code] = true
	}
}

var maskedRegex = regexp.MustCompile(`^\d{4}\.\d{2}\.\d{2}$`)
var bareDigitsRegex = regexp.MustCompile(`^\d{1,8}$`)

func normalize(value interface{}) (string, bool) {
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("%08d", v), v >= 0
	case int64:
		return fmt.Sprintf("%08d", v), v >= 0
	case string:
		trimmed := strings.TrimSpace(v)
		if maskedRegex.MatchString(trimmed) {
			return helpers.OnlyNumbers(trimmed), true
		}
		if bareDigitsRegex.MatchString(trimmed) {
			return fmt.Sprintf("%08s", trimmed), true
		}
		return "", false
	default:
		return "", false
	}
}

// IsValid checks whether an NCM code exists in the current table. Bare
// digits are left-padded to 8; a masked value is read as written.
func IsValid(value interface{}) bool {
	code, ok := normalize(value)
	if !ok {
		return false
	}
	return validCodes[code]
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 8 digits first.
	Pad bool
}

// Format formats an NCM code with the mask "NNNN.NN.NN"; only the
// structure changes (use IsValid to check the code). Masks as far as the
// digits go.
func Format(value interface{}, opts ...FormatOptions) string {
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
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < ncmSize {
		digits = strings.Repeat("0", ncmSize-len(digits)) + digits
	}
	if len(digits) > ncmSize {
		digits = digits[:ncmSize]
	}
	if digits == "" {
		return ""
	}

	seps := map[int]string{4: ".", 6: "."}
	var b strings.Builder
	for i, c := range digits {
		if s, ok := seps[i]; ok {
			b.WriteString(s)
		}
		b.WriteRune(c)
	}
	return b.String()
}

// Parse removes NCM formatting and keeps only digits, capped to 8
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
	if len(digits) > ncmSize {
		digits = digits[:ncmSize]
	}
	return digits
}
