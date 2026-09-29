package areacode_test

import (
	"reflect"
	"testing"

	"github.com/brazilian-utils/go/areacode"
)

func TestGetInfo(t *testing.T) {
	sp := areacode.GetInfo("11")
	if sp == nil || sp.StateCode != "SP" || sp.RegionCode != "SE" || !reflect.DeepEqual(sp.StateCodes, []string{"SP"}) {
		t.Errorf("Unexpected result for DDD 11: %+v", sp)
	}

	spNum := areacode.GetInfo(11)
	if spNum == nil || spNum.StateCode != "SP" {
		t.Errorf("Unexpected result for numeric DDD 11: %+v", spNum)
	}

	df := areacode.GetInfo("61")
	if df == nil || df.StateCode != "DF" || !reflect.DeepEqual(df.StateCodes, []string{"DF", "GO"}) {
		t.Errorf("Unexpected result for DDD 61: %+v", df)
	}

	pr := areacode.GetInfo(42)
	if pr == nil || pr.StateCode != "PR" || !reflect.DeepEqual(pr.StateCodes, []string{"PR", "SC"}) {
		t.Errorf("Unexpected result for DDD 42: %+v", pr)
	}

	if areacode.GetInfo("00") != nil {
		t.Error("Expected nil for DDD 00")
	}
	if areacode.GetInfo("20") != nil {
		t.Error("Expected nil for unassigned DDD 20")
	}
	if areacode.GetInfo("") != nil {
		t.Error("Expected nil for empty input")
	}
}

func TestListByState(t *testing.T) {
	sp := areacode.ListByState("SP")
	if !reflect.DeepEqual(sp, []int{11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Errorf("Unexpected result for SP: %v", sp)
	}

	spLower := areacode.ListByState("sp")
	if !reflect.DeepEqual(spLower, []int{11, 12, 13, 14, 15, 16, 17, 18, 19}) {
		t.Errorf("Unexpected result for sp (lowercase): %v", spLower)
	}

	ac := areacode.ListByState("AC")
	if !reflect.DeepEqual(ac, []int{68}) {
		t.Errorf("Unexpected result for AC: %v", ac)
	}

	pe := areacode.ListByState("PE")
	if !reflect.DeepEqual(pe, []int{81, 87}) {
		t.Errorf("Unexpected result for PE: %v", pe)
	}

	sc := areacode.ListByState("SC")
	if !reflect.DeepEqual(sc, []int{42, 47, 48, 49}) {
		t.Errorf("Unexpected result for SC: %v", sc)
	}

	if len(areacode.ListByState("XX")) != 0 {
		t.Error("Expected empty for XX")
	}
	if len(areacode.ListByState("")) != 0 {
		t.Error("Expected empty for empty input")
	}
}
