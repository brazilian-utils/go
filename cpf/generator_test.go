package cpf_test

import (
	"testing"

	"github.com/brazilian-utils/go/cpf"
)

func TestGenerate(t *testing.T) {
	generated := cpf.Generate()
	if len(generated) != 11 {
		t.Errorf("Expected generated CPF to have 11 characters, got %d", len(generated))
	}
	if !cpf.IsValid(generated) {
		t.Errorf("Expected generated CPF to be valid, got %v", generated)
	}
}

func TestGenerateMultiple(t *testing.T) {
	for i := 0; i < 100; i++ {
		generated := cpf.Generate()
		if !cpf.IsValid(generated) {
			t.Fatalf("Generated invalid CPF on iteration %d: %v", i, generated)
		}
	}
}
