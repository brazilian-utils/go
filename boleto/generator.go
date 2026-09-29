package boleto

import (
	"fmt"
	"math/rand"
	"strings"
)

// GenerateOptions controls Generate.
type GenerateOptions struct {
	// Type picks the layout: "" or "cobranca" (default, 47-digit cobrança
	// bancária) or "arrecadacao" (48-digit boleto de arrecadação, a
	// best-effort structural generation).
	Type string
}

// Generate generates a random valid boleto number. By default a 47-digit
// cobrança bancária linha digitável; Type "arrecadacao" generates a
// 48-digit boleto de arrecadação instead.
func Generate(opts ...GenerateOptions) string {
	if len(opts) > 0 && opts[0].Type == "arrecadacao" {
		return generateArrecadacao()
	}
	return generateCobranca()
}

func randDigits(n int) string {
	var b strings.Builder
	for i := 0; i < n; i++ {
		b.WriteByte(byte('0' + rand.Intn(10)))
	}
	return b.String()
}

// generateCobranca builds a 47-digit cobrança bancária linha digitável
// whose check digits (three mod10 field digits and one mod11 overall
// digit) are computed to match, so the result always satisfies IsValid.
func generateCobranca() string {
	field1 := randDigits(4) + randDigits(5) // dl[0:9], check digit appended below
	dv1 := getMod10(field1)

	field2 := randDigits(10) // dl[10:20]
	dv2 := getMod10(field2)

	field3 := randDigits(10) // dl[21:31]
	dv3 := getMod10(field3)

	field5 := randDigits(14) // dl[33:47]

	// The overall check digit (position 32) is a mod11 over the barcode
	// without its own check digit: dl[0:4] + dl[33:47] + dl[4:9] +
	// dl[10:20] + dl[21:31] (per the reordering validateMod11CheckDigit
	// uses).
	dvGeneral := getMod11(field1[0:4] + field5 + field1[4:9] + field2 + field3)

	return fmt.Sprintf("%s%d%s%d%s%d%d%s",
		field1, dv1, field2, dv2, field3, dv3, dvGeneral, field5)
}

// generateArrecadacao builds a structurally well-formed 48-digit boleto
// de arrecadação: four blocks of 11 digits, each followed by a mod10
// check digit. This is a best-effort layout (the FEBRABAN arrecadação
// layout picks mod10 or mod11 per block from a segment identifier
// digit, which this generator does not yet distinguish, and
// boleto.IsValid does not yet verify this layout).
func generateArrecadacao() string {
	var b strings.Builder

	firstBlock := "8" + randDigits(10)
	b.WriteString(firstBlock)
	b.WriteString(fmt.Sprintf("%d", getMod10(firstBlock)))

	for i := 0; i < 3; i++ {
		block := randDigits(11)
		b.WriteString(block)
		b.WriteString(fmt.Sprintf("%d", getMod10(block)))
	}

	return b.String()
}
