// Package csosn validates a CSOSN (Código de Situação da Operação no
// Simples Nacional) code against the official 10-code table (Ajuste
// SINIEF 07/2005 / NF-e MOC Tabela 5.3 Tabela B).
package csosn

import "strconv"

var validCodes = map[string]bool{
	"101": true, "102": true, "103": true,
	"201": true, "202": true, "203": true,
	"300": true, "400": true, "500": true, "900": true,
}

// IsValid checks whether value is one of the 10 official CSOSN codes.
// Accepts the bare 3 digits as a string, or a non-negative integer; a
// CSOSN has no printed grouping, so a masked value is rejected.
func IsValid(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return validCodes[v]
	case int:
		if v < 0 {
			return false
		}
		return validCodes[strconv.Itoa(v)]
	case int64:
		if v < 0 {
			return false
		}
		return validCodes[strconv.FormatInt(v, 10)]
	default:
		return false
	}
}
