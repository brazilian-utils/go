package cep

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/agaragon/brutils-go/helpers"
)

// Every CEP has exactly 8 characters
const cepSize = 8

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
