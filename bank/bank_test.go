package bank_test

import (
	"testing"

	"github.com/brazilian-utils/go/bank"
)

func TestGetByCode(t *testing.T) {
	bb := bank.GetByCode("001")
	if bb == nil || bb.ISPB != "00000000" || bb.Name != "Banco do Brasil S.A." {
		t.Errorf("Unexpected result for code 001: %+v", bb)
	}

	unpadded := bank.GetByCode("1")
	if unpadded == nil || unpadded.Code != "001" {
		t.Errorf("Unexpected result for unpadded code 1: %+v", unpadded)
	}

	num := bank.GetByCode(1)
	if num == nil || num.Code != "001" {
		t.Errorf("Unexpected result for numeric code 1: %+v", num)
	}

	itau := bank.GetByCode("341")
	if itau == nil || itau.ISPB != "60701190" || itau.Name != "ITAÚ UNIBANCO S.A." {
		t.Errorf("Unexpected result for code 341: %+v", itau)
	}

	if bank.GetByCode("999") != nil {
		t.Error("Expected nil for nonexistent code 999")
	}
	if bank.GetByCode("abc") != nil {
		t.Error("Expected nil for abc")
	}
	if bank.GetByCode("") != nil {
		t.Error("Expected nil for empty code")
	}
}

func TestGetByIspb(t *testing.T) {
	bb := bank.GetByIspb("00000000")
	if bb == nil || bb.Code != "001" {
		t.Errorf("Unexpected result for ISPB 00000000: %+v", bb)
	}

	itau := bank.GetByIspb("60701190")
	if itau == nil || itau.Code != "341" {
		t.Errorf("Unexpected result for ISPB 60701190: %+v", itau)
	}

	if bank.GetByIspb("99999999") != nil {
		t.Error("Expected nil for nonexistent ISPB")
	}
	if bank.GetByIspb("abc") != nil {
		t.Error("Expected nil for abc")
	}
	if bank.GetByIspb("") != nil {
		t.Error("Expected nil for empty value")
	}
}

func TestList(t *testing.T) {
	all := bank.List()
	if len(all) != 71 {
		t.Errorf("Expected 71 banks, got %d", len(all))
	}
}
