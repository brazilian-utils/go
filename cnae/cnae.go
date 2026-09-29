// Package cnae looks up a CNAE (Classificação Nacional de Atividades
// Econômicas) subclass code in the CNAE-Subclasses 2.3 table (the
// current subclass revision of CNAE 2.0).
package cnae

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

const cnaeSize = 7

// Cnae is one entry of the official CNAE-Subclasses 2.3 table.
type Cnae struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type sourceFile struct {
	Data []Cnae `json:"data"`
}

var byCode map[string]Cnae

func init() {
	var src sourceFile
	if err := json.Unmarshal(rawData, &src); err != nil {
		panic("cnae: failed to parse embedded data.json: " + err.Error())
	}
	byCode = make(map[string]Cnae, len(src.Data))
	for _, c := range src.Data {
		byCode[c.Code] = c
	}
}

var maskedRegex = regexp.MustCompile(`^\d{4}-\d/\d{2}$`)
var bareDigitsRegex = regexp.MustCompile(`^\d{1,7}$`)

// normalize accepts a masked "NNNN-N/NN" value (read as written) or bare
// digits (a string or a number, left-padded to 7); any other value is
// rejected.
func normalize(value interface{}) (string, bool) {
	switch v := value.(type) {
	case int:
		return fmt.Sprintf("%07d", v), v >= 0
	case int64:
		return fmt.Sprintf("%07d", v), v >= 0
	case string:
		trimmed := strings.TrimSpace(v)
		if maskedRegex.MatchString(trimmed) {
			return helpers.OnlyNumbers(trimmed), true
		}
		if bareDigitsRegex.MatchString(trimmed) {
			return fmt.Sprintf("%07s", trimmed), true
		}
		return "", false
	default:
		return "", false
	}
}

// IsValid checks whether a CNAE subclass code exists in the current
// table. Bare digits are left-padded to 7; a masked value is read as
// written.
func IsValid(value interface{}) bool {
	code, ok := normalize(value)
	if !ok {
		return false
	}
	_, found := byCode[code]
	return found
}

// Get looks a CNAE subclass code up and returns its code (7 bare digits)
// and official description; returns nil for an unknown code or a value
// not in an accepted form.
func Get(value interface{}) *Cnae {
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

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to 7 digits first.
	Pad bool
}

// Format formats a CNAE subclass code with the mask "NNNN-N/NN"; only
// the structure changes (use IsValid to check the code). Masks as far as
// the digits go. Returns an empty string when there is no digit at all.
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
	if len(opts) > 0 && opts[0].Pad && digits != "" && len(digits) < cnaeSize {
		digits = strings.Repeat("0", cnaeSize-len(digits)) + digits
	}
	if len(digits) > cnaeSize {
		digits = digits[:cnaeSize]
	}
	if digits == "" {
		return ""
	}

	seps := map[int]string{4: "-", 5: "/"}
	var b strings.Builder
	for i, c := range digits {
		if s, ok := seps[i]; ok {
			b.WriteString(s)
		}
		b.WriteRune(c)
	}
	return b.String()
}

// Parse removes CNAE formatting and keeps only digits, capped to 7
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
	if len(digits) > cnaeSize {
		digits = digits[:cnaeSize]
	}
	return digits
}
