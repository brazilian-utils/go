package cep_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agaragon/brutils-go/cep"
	"github.com/stretchr/testify/assert"
)

var tablesValidate = []struct {
	input    string
	expected bool
}{
	{"000111", false},
	{"13723705000189", false},
	{"60.391.947/0001-0", false},
	{"abcdefgh", false},
	{"99999999999999", false},
	{"12345678", true},
	{"12345-678", true},
}

var tablesClean = []struct {
	input    string
	expected string
}{
	{"12345-678", "12345678"},
	{"00000-000", "00000000"},
	{"11111-111", "11111111"},
}

func TestValidate(t *testing.T) {
	for _, table := range tablesValidate {
		if res := cep.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGenerate(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := cep.Generate()
		if !cep.IsValid(res) {
			t.Errorf("An invalid cep was generated: %s", res)
		}
	}
}

func TestClean(t *testing.T) {
	for _, row := range tablesClean {
		if res := cep.Clean(row.input); res != row.expected {
			t.Errorf("Clean failed for %s \t Expected: %s | Received: %s", row.input, row.expected, res)
		}
	}
}

func TestFetchAddress(t *testing.T) {
	// Create a sample response that the API would return
	apiResponse := `{
		"cep": "01001-000",
		"logradouro": "Praça da Sé",
		"complemento": "lado ímpar",
		"bairro": "Sé",
		"localidade": "São Paulo",
		"uf": "SP",
		"ibge": "3550308",
		"gia": "1004",
		"ddd": "11",
		"siafi": "7107"
	}`
	
	// Setup a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(apiResponse))
	}))
	defer server.Close()

	// Assuming the FetchAddress method can be adjusted to accept a URL, you pass the server URL.
	// If not, you'll need to adjust FetchAddress or abstract the base URL so it can be changed during tests.
	oldURL := cep.ApiBaseURL // Save old URL
	cep.ApiBaseURL = server.URL // Temporarily set to test server URL
	defer func() { cep.ApiBaseURL = oldURL }() // Restore after test

	expectedAddress := cep.Address{
		Cep:          "01001-000",
		Street:       "Praça da Sé",
		Complement:   "lado ímpar",
		Neighborhood: "Sé",
		City:         "São Paulo",
		State:        "SP",
		IBGE:         "3550308",
		GIA:          "1004",
		DDD:          "11",
		SIAFI:        "7107",
	}

	// Call the function under test
	actualAddress, err := cep.FetchAddress("01001-000")
	assert.NoError(t, err)
	assert.Equal(t, expectedAddress, actualAddress)
}
