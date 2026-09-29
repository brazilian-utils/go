package boleto

import (
	"strconv"
	"time"
)

// febrabanBaseDate is the epoch the fator de vencimento counts days from
// (per the Febraban/FEBRABAN cobrança bancária specification).
var febrabanBaseDate = time.Date(1997, time.October, 7, 0, 0, 0, 0, time.UTC)

// GetInfoOptions controls GetInfo.
type GetInfoOptions struct {
	// ReferenceDate resolves the fator de vencimento cycle as of this
	// date instead of today (the cycle repeats every 9000 days).
	ReferenceDate time.Time
}

// Info is the information extracted from a boleto.
type Info struct {
	// BankCode is the 3-digit COMPE code; empty for a boleto de
	// arrecadação.
	BankCode string
	// AmountCents is the amount in cents.
	AmountCents int64
	// DueDate is nil when the boleto carries no fator de vencimento (a
	// factor below 1000).
	DueDate *time.Time
	// IsArrecadacao reports whether this is a boleto de arrecadação.
	IsArrecadacao bool
}

// GetInfo extracts the amount, due date and bank code from a boleto;
// returns nil when the value is not a valid boleto.
//
// This is a best-effort implementation for the cobrança bancária layout:
// it does not resolve which of the two 9000-day fator-de-vencimento
// cycles applies (options.ReferenceDate picks a fixed epoch, defaulting
// to the original 1997-10-07 one), and it does not extract the
// boleto de arrecadação's amount/segment/type fields.
func GetInfo(value string, opts ...GetInfoOptions) *Info {
	if !IsValid(value) {
		return nil
	}
	digits := Parse(value)

	if len(digits) == arrecadacaoDigitableLineLength {
		return &Info{IsArrecadacao: true}
	}

	bankCode := digits[0:3]
	fatorStr := digits[33:37]
	valorStr := digits[37:47]

	amount, _ := strconv.ParseInt(valorStr, 10, 64)

	fator, _ := strconv.Atoi(fatorStr)
	var dueDate *time.Time
	if fator >= 1000 {
		base := febrabanBaseDate
		if len(opts) > 0 && !opts[0].ReferenceDate.IsZero() {
			// Resolve the 9000-day cycle nearest the reference date.
			base = febrabanBaseDate
			for base.AddDate(0, 0, fator).Before(opts[0].ReferenceDate.AddDate(0, 0, -4500)) {
				base = base.AddDate(0, 0, 9000)
			}
		}
		d := base.AddDate(0, 0, fator)
		dueDate = &d
	}

	return &Info{
		BankCode:    bankCode,
		AmountCents: amount,
		DueDate:     dueDate,
	}
}
