package municipality_test

import (
	"testing"

	"github.com/brazilian-utils/go/municipality"
)

func TestGetByCode(t *testing.T) {
	sp := municipality.GetByCode("3550308")
	if sp == nil || sp.Name != "São Paulo" || sp.StateCode != "SP" {
		t.Errorf("Unexpected result for 3550308: %+v", sp)
	}

	spNum := municipality.GetByCode(3550308)
	if spNum == nil || spNum.Name != "São Paulo" {
		t.Errorf("Unexpected result for numeric code: %+v", spNum)
	}

	mt := municipality.GetByCode("5101837")
	if mt == nil || mt.Name != "Boa Esperança do Norte" || mt.StateCode != "MT" {
		t.Errorf("Unexpected result for 5101837: %+v", mt)
	}

	masked := municipality.GetByCode("355-030-8")
	if masked == nil || masked.Code != "3550308" {
		t.Errorf("Unexpected result for masked code: %+v", masked)
	}

	if municipality.GetByCode("0000000") != nil {
		t.Error("Expected nil for 0000000")
	}
	if municipality.GetByCode("123") != nil {
		t.Error("Expected nil for too-short code")
	}
	if municipality.GetByCode("12345678") != nil {
		t.Error("Expected nil for too-long code")
	}
	if municipality.GetByCode("") != nil {
		t.Error("Expected nil for empty code")
	}
}

func TestList(t *testing.T) {
	all := municipality.List()
	if len(all) != 5571 {
		t.Errorf("Expected 5571 municipalities, got %d", len(all))
	}

	df := municipality.List("DF")
	if len(df) != 1 || df[0].Name != "Brasília" {
		t.Errorf("Expected only Brasília for DF, got %+v", df)
	}

	if len(municipality.List("ZZ")) != 0 {
		t.Error("Expected empty for unknown state ZZ")
	}
	if len(municipality.List("")) != 0 {
		t.Error("Expected empty for empty stateCode")
	}
}
