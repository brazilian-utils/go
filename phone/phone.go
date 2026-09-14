package phone

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

var mobileRegex = regexp.MustCompile(`^[1-9][1-9]9\d{8}$`)
// Landline first digit 2-6, per Anatel Res. 749/2022 art. 11, I, "a".
var landlineRegex = regexp.MustCompile(`^[1-9][1-9][2-6]\d{7}$`)

// IsValid checks if a Brazilian phone number is valid.
// phoneType can be "mobile", "landline", or "" (accepts either).
// The number must be digits only, without country code, including the 2-digit DDD.
func IsValid(phoneNumber string, phoneType string) bool {
	switch phoneType {
	case "mobile":
		return isValidMobile(phoneNumber)
	case "landline":
		return isValidLandline(phoneNumber)
	default:
		return isValidMobile(phoneNumber) || isValidLandline(phoneNumber)
	}
}

// Format formats a phone number into the standard display pattern.
// Mobile: "(DD)NNNNN-NNNN", Landline: "(DD)NNNN-NNNN".
// Returns empty string if invalid.
func Format(phoneNumber string) string {
	if !IsValid(phoneNumber, "") {
		return ""
	}

	ddd := phoneNumber[:2]
	number := phoneNumber[2:]

	return fmt.Sprintf("(%s)%s-%s", ddd, number[:len(number)-4], number[len(number)-4:])
}

// RemoveSymbols removes common symbols from a phone number string: ()+-  and spaces.
func RemoveSymbols(phoneNumber string) string {
	r := strings.NewReplacer("(", "", ")", "", "-", "", "+", "", " ", "")
	return r.Replace(phoneNumber)
}

// RemoveInternationalDialingCode removes the Brazilian country code ("+55",
// "0055" or "55") from a phone number if present and the number is long
// enough. Unlike a plain digit-strip, this also drops a leading "+" instead
// of leaving it behind.
func RemoveInternationalDialingCode(phoneNumber string) string {
	cleaned := strings.ReplaceAll(phoneNumber, " ", "")
	digits := strings.TrimPrefix(cleaned, "+")

	switch {
	case strings.HasPrefix(digits, "0055") && len(digits) > 13:
		return digits[4:]
	case strings.HasPrefix(digits, "55") && len(digits) > 11:
		return digits[2:]
	}

	return phoneNumber
}

// Generate generates a random valid Brazilian phone number.
// phoneType can be "mobile", "landline", or "" (random choice).
func Generate(phoneType string) string {
	switch phoneType {
	case "mobile":
		return generateMobile()
	case "landline":
		return generateLandline()
	default:
		if rand.Intn(2) == 0 {
			return generateMobile()
		}
		return generateLandline()
	}
}

func isValidMobile(phoneNumber string) bool {
	return mobileRegex.MatchString(phoneNumber)
}

func isValidLandline(phoneNumber string) bool {
	return landlineRegex.MatchString(phoneNumber)
}

func generateDDD() string {
	return fmt.Sprintf("%d%d", rand.Intn(9)+1, rand.Intn(9)+1)
}

func generateMobile() string {
	ddd := generateDDD()
	var buf strings.Builder
	buf.WriteString(ddd)
	buf.WriteByte('9')
	for i := 0; i < 8; i++ {
		buf.WriteByte(byte('0' + rand.Intn(10)))
	}
	return buf.String()
}

func generateLandline() string {
	ddd := generateDDD()
	prefix := rand.Intn(5) + 2 // 2-6
	return fmt.Sprintf("%s%d%07d", ddd, prefix, rand.Intn(10000000))
}
