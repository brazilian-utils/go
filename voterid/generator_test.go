package voterid_test

import (
	"testing"

	"github.com/brazilian-utils/go/voterid"
)

func TestGenerate(t *testing.T) {
	// Test default generation (ZZ - foreigners)
	voterID := voterid.Generate()
	if voterID == "" {
		t.Error("Generate() returned empty string")
	}
	if !voterid.IsValid(voterID) {
		t.Errorf("Generate() produced invalid voter ID: %s", voterID)
	}
	if len(voterID) != 12 {
		t.Errorf("Generate() produced voter ID with wrong length: %d", len(voterID))
	}

	// Test generation for all valid states
	states := []string{
		"SP", "MG", "RJ", "RS", "BA", "PR", "CE", "PE", "SC", "GO",
		"MA", "PB", "PA", "ES", "PI", "RN", "AL", "MT", "MS", "DF",
		"SE", "AM", "RO", "AC", "AP", "RR", "TO", "ZZ",
	}

	for _, state := range states {
		voterID := voterid.Generate(state)
		if voterID == "" {
			t.Errorf("Generate(%q) returned empty string", state)
			continue
		}
		if !voterid.IsValid(voterID) {
			t.Errorf("Generate(%q) produced invalid voter ID: %s", state, voterID)
		}
		if len(voterID) != 12 {
			t.Errorf("Generate(%q) produced voter ID with wrong length: %d", state, len(voterID))
		}
	}

	// Test invalid state
	voterID = voterid.Generate("XX")
	if voterID != "" {
		t.Errorf("Generate(\"XX\") should return empty string for invalid state, got: %s", voterID)
	}

	// Test lowercase state (should be normalized)
	voterID = voterid.Generate("sp")
	if voterID == "" {
		t.Error("Generate(\"sp\") returned empty string")
	}
	if !voterid.IsValid(voterID) {
		t.Errorf("Generate(\"sp\") produced invalid voter ID: %s", voterID)
	}
}

func TestGenerateUniqueness(t *testing.T) {
	// Generate multiple voter IDs and ensure they're unique
	ids := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		voterID := voterid.Generate()
		if ids[voterID] {
			t.Errorf("Generate() produced duplicate voter ID: %s", voterID)
		}
		ids[voterID] = true
	}

	if len(ids) != iterations {
		t.Errorf("Expected %d unique voter IDs, got %d", iterations, len(ids))
	}
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		voterid.Generate()
	}
}

func BenchmarkGenerateWithState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		voterid.Generate("SP")
	}
}
