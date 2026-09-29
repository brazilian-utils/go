package currency

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// FormatCurrencyOptions controls currency.Format.
type FormatCurrencyOptions struct {
	// Precision sets the number of decimal places (default 2, clamped 0-20).
	Precision int
	// HasPrecision distinguishes an explicit Precision: 0 from "unset".
	HasPrecision bool
	// Symbol prepends "R$ " to the result when true.
	Symbol bool
}

// FormatCurrency formats a float64 value as Brazilian currency "X.XXX,XX"
// (no currency symbol by default). Returns empty string for NaN or Inf values.
//
// Deprecated: kept for backward compatibility; prefer Format.
func FormatCurrency(value float64) string {
	return formatFloat(value, 2, true)
}

// Format formats a value (string or number) in the BRL pattern "1.234,56".
// A string is parsed the same way Parse reads it. Returns an empty string
// for a non-finite or unparsable value.
func Format(value interface{}, opts ...FormatCurrencyOptions) string {
	precision := 2
	symbol := false
	if len(opts) > 0 {
		symbol = opts[0].Symbol
		if opts[0].HasPrecision {
			precision = opts[0].Precision
		}
	}
	if precision < 0 {
		precision = 0
	}
	if precision > 20 {
		precision = 20
	}

	var f float64
	switch v := value.(type) {
	case float64:
		f = v
	case float32:
		f = float64(v)
	case int:
		f = float64(v)
	case int64:
		f = float64(v)
	case string:
		parsed, ok := parseAmount(v, precision)
		if !ok {
			return ""
		}
		f = parsed
	default:
		return ""
	}

	return formatFloat(f, precision, symbol)
}

func formatFloat(value float64, precision int, symbol bool) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return ""
	}

	s := strconv.FormatFloat(value, 'f', precision, 64)

	var intPart, decPart string
	if dotIdx := strings.Index(s, "."); dotIdx >= 0 {
		intPart = s[:dotIdx]
		decPart = s[dotIdx+1:]
	} else {
		intPart = s
	}

	negative := strings.HasPrefix(intPart, "-")
	if negative {
		intPart = intPart[1:]
	}

	// Insert thousand separators (dots in Brazilian format)
	var buf strings.Builder
	for i, c := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			buf.WriteByte('.')
		}
		buf.WriteRune(c)
	}

	prefix := ""
	if symbol {
		prefix = "R$ "
	}
	if negative {
		prefix += "-"
	}

	result := prefix + buf.String()
	if decPart != "" {
		result += "," + decPart
	}
	return result
}

// ParseCurrencyOptions controls currency.Parse.
type ParseCurrencyOptions struct {
	// Precision is the max number of decimal digits considered (default 2).
	Precision int
	// HasPrecision distinguishes an explicit Precision: 0 from "unset".
	HasPrecision bool
}

// Parse parses a BRL amount string (e.g. "R$ 1.234,56") into a float64.
// A value with no separator is read as cents (divided by 10^precision).
// Returns 0 for an empty string.
func Parse(value string, opts ...ParseCurrencyOptions) float64 {
	precision := 2
	if len(opts) > 0 && opts[0].HasPrecision {
		precision = opts[0].Precision
	}
	f, _ := parseAmount(value, precision)
	return f
}

// parseAmount implements the currency.parse contract.
func parseAmount(value string, precision int) (float64, bool) {
	s := strings.TrimSpace(value)
	if s == "" {
		return 0, true
	}

	// Strip everything except digits, '.', ',' and a leading '-'.
	var cleaned strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' || r == '.' || r == ',' {
			cleaned.WriteRune(r)
		}
	}
	s = cleaned.String()

	negative := strings.Contains(value, "-")

	maxDecDigits := precision
	if maxDecDigits < 2 {
		maxDecDigits = 2
	}

	// Find the last '.' or ',' that plausibly acts as the decimal separator:
	// followed only by 1..maxDecDigits digits until the end of the string.
	decimalSepIdx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' || s[i] == ',' {
			trailing := len(s) - i - 1
			if trailing >= 1 && trailing <= maxDecDigits {
				decimalSepIdx = i
			}
			break
		}
	}

	var intStr, decStr string
	if decimalSepIdx >= 0 {
		intStr = s[:decimalSepIdx]
		decStr = s[decimalSepIdx+1:]
	} else {
		intStr = s
		decStr = ""
	}

	// Remove remaining separators (thousands grouping) from the integer part.
	intStr = strings.ReplaceAll(intStr, ".", "")
	intStr = strings.ReplaceAll(intStr, ",", "")

	if intStr == "" && decStr == "" {
		return 0, true
	}

	if decimalSepIdx < 0 {
		// No separator at all: the whole digit string is read as cents.
		n, err := strconv.ParseFloat(intStr, 64)
		if err != nil {
			return 0, false
		}
		result := n / math.Pow10(precision)
		if negative {
			result = -result
		}
		return result, true
	}

	combined := intStr + "." + decStr
	n, err := strconv.ParseFloat(combined, 64)
	if err != nil {
		return 0, false
	}
	if negative {
		n = -n
	}
	return n, true
}

// ConvertRealToText converts a monetary value in Brazilian Reais to its
// Portuguese text representation. Values are truncated to 2 decimal places.
// Returns empty string for NaN, Inf, or values exceeding 1 quadrillion.
func ConvertRealToText(value float64) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return ""
	}

	if math.Abs(value) > 1_000_000_000_000_000 {
		return ""
	}

	negative := value < 0
	abs := math.Abs(value)

	// Truncate to 2 decimal places using string formatting to avoid
	// floating-point precision issues (mirrors Python's ROUND_DOWN).
	s := fmt.Sprintf("%.4f", abs)
	dotIdx := strings.Index(s, ".")
	intStr := s[:dotIdx]
	decStr := s[dotIdx+1 : dotIdx+3]

	reais, _ := strconv.ParseInt(intStr, 10, 64)
	centavos, _ := strconv.Atoi(decStr)

	var parts []string

	if reais > 0 {
		reaisText := numberToPortuguese(reais)
		currencyWord := "reais"
		if reais == 1 {
			currencyWord = "real"
		}
		connector := ""
		if strings.HasSuffix(reaisText, "lhão") || strings.HasSuffix(reaisText, "lhões") {
			connector = "de "
		}
		parts = append(parts, reaisText+" "+connector+currencyWord)
	}

	if centavos > 0 {
		centavosText := numberToPortuguese(int64(centavos))
		centavoWord := "centavos"
		if centavos == 1 {
			centavoWord = "centavo"
		}
		if reais > 0 {
			parts = append(parts, "e "+centavosText+" "+centavoWord)
		} else {
			parts = append(parts, centavosText+" "+centavoWord)
		}
	}

	if reais == 0 && centavos == 0 {
		parts = append(parts, "zero reais")
	}

	result := strings.Join(parts, " ")
	if negative {
		result = "menos " + result
	}

	return result
}
