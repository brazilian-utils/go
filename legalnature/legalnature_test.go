package legalnature_test

import (
	"testing"

	"github.com/brazilian-utils/go/legalnature"
)

var isValidTests = []struct {
	input    string
	expected bool
}{
	{"3329", false},
	{"0000", false},
	{"2240", true},
	{"2241", false},
	{"3311", false},
	{"224-0", true},
	{"", false},
	{"   ", false},
	{"abc", false},
	// Retired codes are still accepted by IsValid.
	{"2208", true},
}

func TestIsValid(t *testing.T) {
	for _, table := range isValidTests {
		if res := legalnature.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var getDescriptionTests = []struct {
	input    string
	expected string
}{
	{"2240", "Sociedade Simples Limitada"},
	{"224-0", "Sociedade Simples Limitada"},
}

func TestGetDescription(t *testing.T) {
	for _, table := range getDescriptionTests {
		if res := legalnature.GetDescription(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestList(t *testing.T) {
	all := legalnature.List()
	if len(all) != 92 {
		t.Errorf("Expected 92 in-force codes, got %d", len(all))
	}
	if all["2062"] != "Sociedade Empresária Limitada" {
		t.Errorf("Expected 2062 to be Sociedade Empresária Limitada, got %v", all["2062"])
	}
	if _, ok := all["2208"]; ok {
		t.Error("Expected retired code 2208 to be excluded by default")
	}

	withRetired := legalnature.List(legalnature.ListOptions{IncludeRetired: true})
	if len(withRetired) <= 92 {
		t.Errorf("Expected more than 92 entries with retired codes included, got %d", len(withRetired))
	}
	if _, ok := withRetired["2208"]; !ok {
		t.Error("Expected retired code 2208 to be included")
	}

	// Verify List returns a copy, not a shared map.
	all["9999"] = "test"
	if legalnature.IsValid("9999") {
		t.Error("List should return a copy, not the original map")
	}
}

func TestGet(t *testing.T) {
	ln := legalnature.Get("2062")
	if ln == nil {
		t.Fatal("Expected non-nil result for 2062")
	}
	if ln.Description != "Sociedade Empresária Limitada" || ln.Category.Code != "2" || ln.Legacy {
		t.Errorf("Unexpected result for 2062: %+v", ln)
	}

	masked := legalnature.Get("206-2")
	if masked == nil || masked.Code != "2062" {
		t.Errorf("Expected masked lookup to resolve to 2062, got %+v", masked)
	}

	legacy := legalnature.Get("2208")
	if legacy == nil || !legacy.Legacy || legacy.CurrentCode == nil || *legacy.CurrentCode != "2275" {
		t.Errorf("Expected legacy code 2208 with currentCode 2275, got %+v", legacy)
	}

	if legalnature.Get("0000") != nil {
		t.Error("Expected nil for unknown code 0000")
	}
	if legalnature.Get("206") != nil {
		t.Error("Expected nil for too-short code")
	}
	if legalnature.Get("") != nil {
		t.Error("Expected nil for empty code")
	}
}

func TestListByCategory(t *testing.T) {
	cat5 := legalnature.ListByCategory("5")
	if len(cat5) != 3 {
		t.Fatalf("Expected 3 entries for category 5, got %d", len(cat5))
	}
	if cat5[0].Code != "5010" || cat5[1].Code != "5029" || cat5[2].Code != "5037" {
		t.Errorf("Unexpected order/content for category 5: %+v", cat5)
	}

	cat4 := legalnature.ListByCategory("4")
	if len(cat4) != 6 {
		t.Errorf("Expected 6 entries for category 4, got %d", len(cat4))
	}

	if len(legalnature.ListByCategory("0")) != 0 {
		t.Error("Expected empty for category 0")
	}
	if len(legalnature.ListByCategory("9")) != 0 {
		t.Error("Expected empty for category 9")
	}
	if len(legalnature.ListByCategory("2062")) != 0 {
		t.Error("Expected empty: a full code is not a category")
	}
	if len(legalnature.ListByCategory("")) != 0 {
		t.Error("Expected empty for empty category")
	}
}

var formatTests = []struct {
	input    string
	expected string
}{
	{"2062", "206-2"},
	{"206-2", "206-2"},
	{"206", "206"},
	{"", ""},
}

func TestFormat(t *testing.T) {
	for _, table := range formatTests {
		if res := legalnature.Format(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

var parseTests = []struct {
	input    string
	expected string
}{
	{"206-2", "2062"},
	{"2062", "2062"},
	{"206299", "2062"},
	{"", ""},
}

func TestParse(t *testing.T) {
	for _, table := range parseTests {
		if res := legalnature.Parse(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 20; i++ {
		generated := legalnature.Generate()
		if !legalnature.IsValid(generated) {
			t.Errorf("Generated invalid legal nature code: %v", generated)
		}
	}
}
