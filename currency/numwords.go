package currency

import "github.com/brazilian-utils/go/helpers"

// numberToPortuguese delegates to the shared helpers implementation.
func numberToPortuguese(n int64) string {
	return helpers.NumberToPortuguese(n)
}
