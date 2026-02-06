package cnpj_test

import (
	"testing"

	"github.com/brazilian-utils/brutils-go/cnpj"
)

func TestGenerate(t *testing.T) {
	generated := cnpj.Generate(1)
	if len(generated) != 14 {
		t.Errorf("Expected generated CNPJ to have 14 characters, got %d", len(generated))
	}
	if !cnpj.IsValid(generated) {
		t.Errorf("Expected generated CNPJ to be valid, got %v", generated)
	}
}

func TestGenerateWithBranch(t *testing.T) {
	generated := cnpj.Generate(1234)
	if len(generated) != 14 {
		t.Errorf("Expected generated CNPJ to have 14 characters, got %d", len(generated))
	}
	if !cnpj.IsValid(generated) {
		t.Errorf("Expected generated CNPJ to be valid, got %v", generated)
	}
	// Branch 1234 should appear at positions 8-11
	branch := generated[8:12]
	if branch != "1234" {
		t.Errorf("Expected branch 1234, got %v", branch)
	}
}

func TestGenerateWithZeroBranch(t *testing.T) {
	generated := cnpj.Generate(0)
	if !cnpj.IsValid(generated) {
		t.Errorf("Expected generated CNPJ to be valid, got %v", generated)
	}
	// Branch 0 should default to 1 → "0001"
	branch := generated[8:12]
	if branch != "0001" {
		t.Errorf("Expected branch 0001 for input 0, got %v", branch)
	}
}
