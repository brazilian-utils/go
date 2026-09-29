package cnpj_test

import (
	"testing"

	"github.com/brazilian-utils/go/cnpj"
)

func TestGenerate(t *testing.T) {
	generated := cnpj.Generate()
	if len(generated) != 14 {
		t.Errorf("Expected generated CNPJ to have 14 characters, got %d", len(generated))
	}
	if !cnpj.IsValid(generated) {
		t.Errorf("Expected generated CNPJ to be valid, got %v", generated)
	}
}

func TestGenerateWithBranch(t *testing.T) {
	generated := cnpj.Generate(cnpj.GenerateOptions{Branch: 1234})
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
	generated := cnpj.Generate(cnpj.GenerateOptions{Branch: 0})
	if !cnpj.IsValid(generated) {
		t.Errorf("Expected generated CNPJ to be valid, got %v", generated)
	}
}

func TestGenerateAlphanumeric(t *testing.T) {
	for i := 0; i < 10; i++ {
		generated := cnpj.Generate(cnpj.GenerateOptions{Version: 2})
		if len(generated) != 14 {
			t.Errorf("Expected generated CNPJ to have 14 characters, got %d", len(generated))
		}
		if !cnpj.IsValid(generated, cnpj.IsValidOptions{Version: 2}) {
			t.Errorf("Expected generated alphanumeric CNPJ to be valid, got %v", generated)
		}
	}
}
