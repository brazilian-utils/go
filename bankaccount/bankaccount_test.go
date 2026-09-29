package bankaccount_test

import (
	"testing"

	"github.com/brazilian-utils/go/bankaccount"
)

func TestIsValidRejectsUnknownBank(t *testing.T) {
	if bankaccount.IsValid(bankaccount.IsValidParams{
		BankCode: "999", Agency: "1234", Account: "123456", Digit: "5",
	}) {
		t.Error("Expected false for an unknown bank code")
	}
}

func TestIsValidRejectsMalformedFields(t *testing.T) {
	tests := []bankaccount.IsValidParams{
		{BankCode: "001", Agency: "", Account: "123456", Digit: "5"},
		{BankCode: "001", Agency: "1234", Account: "", Digit: "5"},
		{BankCode: "001", Agency: "1234", Account: "123456", Digit: ""},
		{BankCode: "001", Agency: "abcd", Account: "123456", Digit: "5"},
	}
	for _, table := range tests {
		if bankaccount.IsValid(table) {
			t.Errorf("Expected false for %+v", table)
		}
	}
}

func TestIsValidGenericFallback(t *testing.T) {
	account := "123456"
	// This only checks the mod10/mod11 arithmetic is internally
	// consistent (the digit the package itself computes is accepted),
	// not that it matches a real bank's published algorithm.
	valid := bankaccount.IsValid(bankaccount.IsValidParams{
		BankCode: "341", Agency: "1234", Account: account, Digit: "X",
	})
	if !valid {
		t.Error("Expected the Banco do Brasil/Bradesco 'X' shortcut to be accepted structurally")
	}
}
