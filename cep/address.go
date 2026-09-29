package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/brazilian-utils/go/helpers"
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
// ViaCEP sends "erro" as the boolean true, but in practice has also been
// observed sending it as the string "true" — decoding straight into a bool
// silently fails the Unmarshal in that case, so this accepts either shape.
type viaCEPError struct {
	Erro flexBool `json:"erro"`
}

// flexBool unmarshals from a JSON bool or from a JSON string ("true"/"false").
type flexBool bool

func (b *flexBool) UnmarshalJSON(data []byte) error {
	var asBool bool
	if err := json.Unmarshal(data, &asBool); err == nil {
		*b = flexBool(asBool)
		return nil
	}

	var asString string
	if err := json.Unmarshal(data, &asString); err != nil {
		return err
	}
	*b = flexBool(asString == "true")
	return nil
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
	if json.Unmarshal(raw, &errResp) == nil && bool(errResp.Erro) {
		return nil, fmt.Errorf("CEP not found: %s", cep)
	}

	var addr Address
	if err := json.Unmarshal(raw, &addr); err != nil {
		return nil, fmt.Errorf("failed to parse address: %w", err)
	}

	return &addr, nil
}

// GetCEPFromAddressParams carries the state (federative unit), city and
// street to look up.
type GetCEPFromAddressParams struct {
	// UF is the 2-letter Brazilian state abbreviation (e.g. "SP", "RJ"),
	// case-insensitive.
	UF string
	// City is the city name.
	City string
	// Street is the street name.
	Street string
}

// GetCEPFromAddress fetches CEP options for a given address using the
// ViaCEP API.
func GetCEPFromAddress(params GetCEPFromAddressParams) ([]Address, error) {
	federalUnit := strings.ToUpper(strings.TrimSpace(params.UF))
	city := strings.TrimSpace(params.City)
	street := strings.TrimSpace(params.Street)

	if !helpers.Contains(validUFs, federalUnit) {
		return nil, fmt.Errorf("invalid UF: %s", params.UF)
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
