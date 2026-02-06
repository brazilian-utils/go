package boleto

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/brazilian-utils/brutils-go/helpers"
)

// Every Digitable Line from Boleto has exactly 47 characters
const digitableLineLength = 47

var digitableLineToBoletoConvertPositions = []struct {
	start int
	end   int
}{
	{0, 4},
	{32, 47},
	{4, 9},
	{10, 20},
	{21, 31},
}

var partialsToVerifyMod10 = []struct {
	start      int
	end        int
	digitIndex int
}{
	{0, 9, 9},
	{10, 20, 20},
	{21, 31, 31},
}

var mod10Weights = []int{2, 1}

var checkDigitMod11Position = 4

var mod11Weights = struct {
	initial   int
	end       int
	increment int
}{2, 9, 1}

// IsValid validates if a given Digitable Line is valid
func IsValid(digitableLine string) bool {
	digitableLineNumbers := helpers.OnlyNumbers(digitableLine)

	if !isValidLength(digitableLineNumbers) {
		return false
	}
	if !validateDigitableLinePartials(digitableLineNumbers) {
		return false
	}
	return validateMod11CheckDigit(digitableLineNumbers)
}

// validates the string length
func isValidLength(digitableLine string) bool {
	return len(digitableLine) == digitableLineLength
}

func validateDigitableLinePartials(digitableLine string) bool {
	isValid := true
	for _, partial := range partialsToVerifyMod10 {
		mod10 := getMod10(digitableLine[partial.start:partial.end])
		digitCharacter := fmt.Sprintf("%c", digitableLine[partial.digitIndex])
		digit, _ := strconv.Atoi(digitCharacter)
		if digit != mod10 {
			isValid = false
		}
	}

	return isValid
}

func getMod10(partial string) int {
	var sum int
	partialReversed := helpers.Reverse(partial)

	for index, partialRune := range partialReversed {
		partial := fmt.Sprintf("%c", partialRune)
		partialValue, _ := strconv.Atoi(partial)
		weight := mod10Weights[index%2]
		multiplier := partialValue * weight
		if multiplier > 9 {
			sum += 1 + (multiplier % 10)
		} else {
			sum += multiplier
		}
	}

	mod10 := sum % 10
	if mod10 > 0 {
		return 10 - mod10
	}

	return 0
}

func validateMod11CheckDigit(digitableLine string) bool {
	parsedDigitableLine := parseDigitableLine(digitableLine)
	mod11 := getMod11(parsedDigitableLine[0:checkDigitMod11Position] + parsedDigitableLine[checkDigitMod11Position+1:])
	mod11Value, _ := strconv.Atoi(fmt.Sprintf("%c", parsedDigitableLine[checkDigitMod11Position]))
	return mod11Value == mod11
}

func parseDigitableLine(digitableLine string) string {
	buf := bytes.Buffer{}
	for _, position := range digitableLineToBoletoConvertPositions {
		buf.WriteString(digitableLine[position.start:position.end])
	}
	return buf.String()
}

func getMod11(value string) int {
	weight := mod11Weights.initial

	var sum int
	valueReversed := helpers.Reverse(value)

	for _, valueRune := range valueReversed {
		value := fmt.Sprintf("%c", valueRune)
		valueValue, _ := strconv.Atoi(value)
		multiplier := valueValue * weight

		if weight < mod11Weights.end {
			weight++
		} else {
			weight = mod11Weights.initial
		}

		sum += multiplier
	}

	mod11 := sum % 11
	if mod11 != 0 && mod11 != 1 {
		return 11 - mod11
	}

	return 1
}
