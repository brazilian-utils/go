package currency

import "github.com/brazilian-utils/brutils-go/helpers"

// numberToPortuguese delegates to the shared helpers implementation.
func numberToPortuguese(n int64) string {
	return helpers.NumberToPortuguese(n)
}
