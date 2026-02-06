package currency

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// FormatCurrency formats a float64 value as Brazilian currency "R$ X.XXX,XX".
// Returns empty string for NaN or Inf values.
func FormatCurrency(value float64) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return ""
	}

	s := fmt.Sprintf("%.2f", value)

	dotIdx := strings.Index(s, ".")
	intPart := s[:dotIdx]
	decPart := s[dotIdx+1:]

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

	prefix := "R$ "
	if negative {
		prefix = "R$ -"
	}

	return prefix + buf.String() + "," + decPart
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

	// Capitalize first letter
	return strings.ToUpper(result[:1]) + result[1:]
}
