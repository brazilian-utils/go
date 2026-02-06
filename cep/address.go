package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/brazilian-utils/brutils-go/helpers"
)

// Address represents the address data returned by the ViaCEP API.
type Address struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

// viaCEPError is the response shape when a CEP is not found.
type viaCEPError struct {
	Erro bool `json:"erro"`
}

// baseAPIURL is the ViaCEP API base URL. It is a variable so tests can override it.
var baseAPIURL = "https://viacep.com.br/ws"

// validUFs contains all valid Brazilian state abbreviations.
var validUFs = []string{
	"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO",
	"MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI",
	"RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO",
}

// GetAddressFromCEP fetches address information for a given CEP using the ViaCEP API.
// Returns an error if the CEP is invalid or not found.
func GetAddressFromCEP(cep string) (*Address, error) {
	cleaned := helpers.OnlyNumbers(cep)
	if !IsValid(cleaned) {
		return nil, fmt.Errorf("invalid CEP: %s", cep)
	}

	reqURL := fmt.Sprintf("%s/%s/json/", baseAPIURL, cleaned)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch CEP: %w", err)
	}
	defer resp.Body.Close()

	// Check for API error response ({"erro": true})
	var raw json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	var errResp viaCEPError
	if json.Unmarshal(raw, &errResp) == nil && errResp.Erro {
		return nil, fmt.Errorf("CEP not found: %s", cep)
	}

	var addr Address
	if err := json.Unmarshal(raw, &addr); err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	return &addr, nil
}

// GetCEPFromAddress fetches CEP options for a given address using the ViaCEP API.
// federalUnit must be a valid 2-letter Brazilian state abbreviation (e.g. "SP", "RJ").
func GetCEPFromAddress(federalUnit, city, street string) ([]Address, error) {
	if !helpers.Contains(validUFs, federalUnit) {
		return nil, fmt.Errorf("invalid UF: %s", federalUnit)
	}

	if city == "" {
		return nil, errors.New("city must not be empty")
	}
	if street == "" {
		return nil, errors.New("street must not be empty")
	}

	reqURL := fmt.Sprintf("%s/%s/%s/%s/json/",
		baseAPIURL,
		url.PathEscape(federalUnit),
		url.PathEscape(city),
		url.PathEscape(street),
	)

	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch address: %w", err)
	}
	defer resp.Body.Close()

	var addresses []Address
	if err := json.NewDecoder(resp.Body).Decode(&addresses); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if len(addresses) == 0 {
		return nil, fmt.Errorf("no results found for: %s - %s - %s", federalUnit, city, street)
	}

	return addresses, nil
}
