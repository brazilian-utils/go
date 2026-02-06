package cep

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAddressFromCEP_Valid(t *testing.T) {
	expected := Address{
		CEP:         "01001-000",
		Logradouro:  "Praça da Sé",
		Complemento: "lado ímpar",
		Bairro:      "Sé",
		Localidade:  "São Paulo",
		UF:          "SP",
		IBGE:        "3550308",
		GIA:         "1004",
		DDD:         "11",
		SIAFI:       "7107",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	oldURL := baseAPIURL
	baseAPIURL = server.URL
	defer func() { baseAPIURL = oldURL }()

	addr, err := GetAddressFromCEP("01001000")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if addr.CEP != expected.CEP {
		t.Errorf("Expected CEP %v, got %v", expected.CEP, addr.CEP)
	}
	if addr.Localidade != expected.Localidade {
		t.Errorf("Expected Localidade %v, got %v", expected.Localidade, addr.Localidade)
	}
	if addr.UF != expected.UF {
		t.Errorf("Expected UF %v, got %v", expected.UF, addr.UF)
	}
}

func TestGetAddressFromCEP_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"erro": true}`))
	}))
	defer server.Close()

	oldURL := baseAPIURL
	baseAPIURL = server.URL
	defer func() { baseAPIURL = oldURL }()

	addr, err := GetAddressFromCEP("00000000")
	if err == nil {
		t.Fatal("Expected error for non-existent CEP, got nil")
	}
	if addr != nil {
		t.Errorf("Expected nil address, got %v", addr)
	}
}

func TestGetAddressFromCEP_InvalidCEP(t *testing.T) {
	_, err := GetAddressFromCEP("123")
	if err == nil {
		t.Fatal("Expected error for invalid CEP, got nil")
	}
}

func TestGetCEPFromAddress_Valid(t *testing.T) {
	expected := []Address{
		{
			CEP:        "01001-000",
			Logradouro: "Praça da Sé",
			Bairro:     "Sé",
			Localidade: "São Paulo",
			UF:         "SP",
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(expected)
	}))
	defer server.Close()

	oldURL := baseAPIURL
	baseAPIURL = server.URL
	defer func() { baseAPIURL = oldURL }()

	addresses, err := GetCEPFromAddress("SP", "São Paulo", "Praça da Sé")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(addresses) != 1 {
		t.Fatalf("Expected 1 address, got %d", len(addresses))
	}
	if addresses[0].UF != "SP" {
		t.Errorf("Expected UF SP, got %v", addresses[0].UF)
	}
}

func TestGetCEPFromAddress_InvalidUF(t *testing.T) {
	_, err := GetCEPFromAddress("XX", "São Paulo", "Praça da Sé")
	if err == nil {
		t.Fatal("Expected error for invalid UF, got nil")
	}
}

func TestGetCEPFromAddress_EmptyCity(t *testing.T) {
	_, err := GetCEPFromAddress("SP", "", "Praça da Sé")
	if err == nil {
		t.Fatal("Expected error for empty city, got nil")
	}
}

func TestGetCEPFromAddress_EmptyStreet(t *testing.T) {
	_, err := GetCEPFromAddress("SP", "São Paulo", "")
	if err == nil {
		t.Fatal("Expected error for empty street, got nil")
	}
}

func TestGetCEPFromAddress_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`[]`))
	}))
	defer server.Close()

	oldURL := baseAPIURL
	baseAPIURL = server.URL
	defer func() { baseAPIURL = oldURL }()

	_, err := GetCEPFromAddress("SP", "Nonexistent", "Nonexistent")
	if err == nil {
		t.Fatal("Expected error for empty results, got nil")
	}
}
