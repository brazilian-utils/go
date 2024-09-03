package cep

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/agaragon/brutils-go/helpers"
)

// ApiBaseURL is the base URL for the API endpoint, exposed for overriding in tests
var ApiBaseURL = "https://viacep.com.br/ws"

// Every CEP has exactly 8 characters
const cepSize = 8

type Address struct {
	Cep string `json:"cep"`
	Street string `json:"logradouro"`
	Complement string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City string `json:"localidade"`
	State string `json:"uf"`
	IBGE string `json:"ibge"`
	GIA string `json:"gia"`
	DDD string `json:"ddd"`
	SIAFI string `json:"siafi"`
}

// IsValid validates if a given CEP is valid
func IsValid(cep string) bool {
	digits := helpers.OnlyNumbers(cep)

	return hasValidLength(digits)
}

func Generate() string {
	rand.Seed(time.Now().UnixNano())
	var nums [8]int
	for i := range nums {
		nums[i] = rand.Intn(10) // generate a digit from 0-9
	}

	formattedString := fmt.Sprintf("%d%d%d%d%d-%d%d%d", nums[0], nums[1], nums[2], nums[3], nums[4], nums[5], nums[6], nums[7])
	return formattedString
}

// Validates the string length
func hasValidLength(cep string) bool {
	return len(cep) == cepSize
}

func Clean(cep string) string {
	return helpers.OnlyNumbers(cep)
}

// FetchAddress retrieves address data from the API based on a given CEP (ZIP code).
func FetchAddress(cep string) (Address, error) {
	// Construct the full URL with the CEP
	url := fmt.Sprintf("%s/%s/json/", ApiBaseURL, cep)

	// Send the GET request to the API
	resp, err := http.Get(url)
	if err != nil {
		return Address{}, fmt.Errorf("error making request to API: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return Address{}, fmt.Errorf("API returned non-OK status: %d", resp.StatusCode)
	}

	// Decode the JSON response into an Address struct
	var address Address
	if err := json.NewDecoder(resp.Body).Decode(&address); err != nil {
		return Address{}, fmt.Errorf("error decoding response from API: %v", err)
	}

	return address, nil
}
