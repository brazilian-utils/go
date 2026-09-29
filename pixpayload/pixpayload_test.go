package pixpayload_test

import (
	"testing"

	"github.com/brazilian-utils/go/pixpayload"
)

const staticExample = "00020126580014br.gov.bcb.pix0136123e4567-e12b-12d1-a456-4266554400005204000053039865802BR5913Fulano de Tal6008BRASILIA62070503***63041D3D"
const withAmountExample = "00020126580014br.gov.bcb.pix0136bee05743-4291-4f3c-9259-595df1307ba1520400005303986540510.005802BR5914Alexandre Lima6019Presidente Prudente62180514Um-Id-Qualquer6304D475"
const wrongCRCExample = "00020126580014br.gov.bcb.pix0136123e4567-e12b-12d1-a456-4266554400005204000053039865802BR5913Fulano de Tal6008BRASILIA62070503***63041D3E"

func TestIsValid(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{staticExample, true},
		{withAmountExample, true},
		{wrongCRCExample, false},
		{"", false},
	}
	for _, table := range tests {
		if res := pixpayload.IsValid(table.input); res != table.expected {
			t.Errorf("Failing for %v \t Expected: %v | Received: %v", table.input, table.expected, res)
		}
	}
}

func TestGetInfo(t *testing.T) {
	info := pixpayload.GetInfo(staticExample)
	if info == nil {
		t.Fatal("Expected non-nil info")
	}
	if info.Key != "123e4567-e12b-12d1-a456-426655440000" || info.MerchantName != "Fulano de Tal" ||
		info.MerchantCity != "BRASILIA" || info.PointOfInitiation != "static" || info.Amount != nil || info.TxID != "" {
		t.Errorf("Unexpected info: %+v", info)
	}

	info2 := pixpayload.GetInfo(withAmountExample)
	if info2 == nil {
		t.Fatal("Expected non-nil info")
	}
	if info2.Key != "bee05743-4291-4f3c-9259-595df1307ba1" || info2.MerchantName != "Alexandre Lima" ||
		info2.MerchantCity != "Presidente Prudente" || info2.TxID != "Um-Id-Qualquer" ||
		info2.Amount == nil || *info2.Amount != 10 {
		t.Errorf("Unexpected info: %+v", info2)
	}

	if pixpayload.GetInfo(wrongCRCExample) != nil {
		t.Error("Expected nil for wrong CRC")
	}
	if pixpayload.GetInfo("") != nil {
		t.Error("Expected nil for empty input")
	}
}

func TestGenerateRoundTrip(t *testing.T) {
	payload := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{
		Key:          "123e4567-e12b-12d1-a456-426655440000",
		MerchantName: "Fulano de Tal",
		MerchantCity: "Brasilia",
	})
	if payload == "" {
		t.Fatal("Expected non-empty payload")
	}
	if !pixpayload.IsValid(payload) {
		t.Fatalf("Generated payload is not valid: %v", payload)
	}
	info := pixpayload.GetInfo(payload)
	if info == nil || info.Key != "123e4567-e12b-12d1-a456-426655440000" {
		t.Errorf("Unexpected round-trip info: %+v", info)
	}
}

func TestGenerateValidatesKeyFormat(t *testing.T) {
	validKeys := []string{
		"123.456.789-09",             // cpf
		"fulano@example.com",         // email
		"+5561912345678",             // phone
		"123e4567-e12b-12d1-a456-426655440000", // evp
	}
	for _, key := range validKeys {
		payload := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{
			Key:          key,
			MerchantName: "Fulano de Tal",
			MerchantCity: "Brasilia",
		})
		if payload == "" {
			t.Errorf("Expected non-empty payload for valid key %v", key)
		}
	}

	invalidKeys := []string{"pixkeyalias001", "not-a-key", "12345"}
	for _, key := range invalidKeys {
		payload := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{
			Key:          key,
			MerchantName: "Fulano de Tal",
			MerchantCity: "Brasilia",
		})
		if payload != "" {
			t.Errorf("Expected empty payload for invalid key %v, got %v", key, payload)
		}
	}
}

func TestGenerateRequiresExactlyOneOf(t *testing.T) {
	if res := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{}); res != "" {
		t.Errorf("Expected empty when neither key nor URL is given, got %v", res)
	}
	if res := pixpayload.Generate(pixpayload.GeneratePixPayloadParams{Key: "a", URL: "b"}); res != "" {
		t.Errorf("Expected empty when both key and URL are given, got %v", res)
	}
}
