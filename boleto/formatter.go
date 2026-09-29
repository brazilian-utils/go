package boleto

import (
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

// FormatOptions controls Format.
type FormatOptions struct {
	// Pad left-pads the value with zeros to the length of the detected
	// pattern (47 digits, or 48 when the value starts with '8') before
	// masking.
	Pad bool
}

// cobrancaSeparators maps the digit index (in the unmasked 47-digit
// cobrança bancária linha digitável) to the separator printed right
// before it.
var cobrancaSeparators = map[int]string{
	5: ".", 10: " ", 15: ".", 21: " ", 26: ".", 32: " ", 33: " ",
}

// Format formats a value as a boleto linha digitável.
//
// The 47-digit cobrança bancária linha digitável is grouped as
// "00000.00000 00000.000000 00000.000000 0 00000000000000". A 48-digit
// linha digitável starting with '8' (boleto de arrecadação) gets four
// blocks of 11 digits, each followed by its check digit. Masks as far as
// the digits go.
func Format(value string, opts ...FormatOptions) string {
	digits := helpers.OnlyNumbers(value)
	if digits == "" {
		return ""
	}

	targetLen := digitableLineLength
	if digits[0] == '8' {
		targetLen = arrecadacaoDigitableLineLength
	}

	if len(opts) > 0 && opts[0].Pad && len(digits) < targetLen {
		digits = strings.Repeat("0", targetLen-len(digits)) + digits
	}
	if len(digits) > targetLen {
		digits = digits[:targetLen]
	}

	if targetLen == arrecadacaoDigitableLineLength {
		return formatArrecadacao(digits)
	}
	return formatCobranca(digits)
}

func formatCobranca(digits string) string {
	var b strings.Builder
	for i := 0; i < len(digits); i++ {
		if sep, ok := cobrancaSeparators[i]; ok {
			b.WriteString(sep)
		}
		b.WriteByte(digits[i])
	}
	return b.String()
}

// formatArrecadacao groups the digits as four blocks of 11 digits, each
// followed by a hyphen and its check digit, separated by a space.
func formatArrecadacao(digits string) string {
	var b strings.Builder
	for i := 0; i < len(digits); i++ {
		pos := i % 12
		if i > 0 && pos == 0 {
			b.WriteString(" ")
		} else if pos == 11 {
			b.WriteString("-")
		}
		b.WriteByte(digits[i])
	}
	return b.String()
}

// Parse removes boleto formatting and keeps only digits, capped to 47
// digits (48 for a boleto de arrecadação, detected by a leading '8').
func Parse(value string) string {
	digits := helpers.OnlyNumbers(value)
	if digits == "" {
		return ""
	}
	targetLen := digitableLineLength
	if digits[0] == '8' {
		targetLen = arrecadacaoDigitableLineLength
	}
	if len(digits) > targetLen {
		digits = digits[:targetLen]
	}
	return digits
}
