package phone

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"

	"github.com/brazilian-utils/go/helpers"
)

var mobileRegex = regexp.MustCompile(`^[1-9][1-9]9\d{8}$`)

// Landline first digit 2-6, per Anatel Res. 749/2022 art. 11, I, "a".
var landlineRegex = regexp.MustCompile(`^[1-9][1-9][2-6]\d{7}$`)

var serviceNGCRegex = regexp.MustCompile(`^(0300|0303|0500|0800|0900)\d{7}$`)
var serviceAbbreviatedRegex = regexp.MustCompile(`^(300|400)\d{5}$`)

// publicUtilityCodes lists the 3-digit Anatel-designated public utility
// codes (Códigos de Acesso de Utilidade Pública). 112 and 911 are
// intentionally excluded per the shared contract.
var publicUtilityCodes = map[string]bool{
	"100": true, "128": true, "132": true, "133": true, "135": true,
	"136": true, "137": true, "138": true, "140": true, "141": true,
	"142": true, "143": true, "144": true, "145": true, "146": true,
	"147": true, "148": true, "150": true, "151": true, "153": true,
	"154": true, "155": true, "156": true, "157": true, "158": true,
	"159": true, "160": true, "161": true, "162": true, "163": true,
	"164": true, "180": true, "181": true, "185": true, "188": true,
	"190": true, "191": true, "192": true, "193": true, "194": true,
	"196": true, "197": true, "198": true, "199": true,
}

// IsValidOptions controls IsValid.
type IsValidOptions struct {
	// Kinds narrows or widens the accepted kinds: any of "mobile",
	// "landline", "service". Defaults to ["mobile", "landline"]. An empty
	// (but non-nil) slice rejects everything.
	Kinds []string
	// MobileVersion picks the numbering rule used for a mobile number
	// (see IsValidMobile).
	MobileVersion int
}

// IsValid validates a Brazilian phone number: mobile or landline by
// default. A country code (+55, 0055 or a bare 55) is accepted and
// removed first.
func IsValid(value string, opts ...IsValidOptions) bool {
	kinds := []string{"mobile", "landline"}
	mobileVersion := 0 // 0: no extra numbering-plan restriction (matches the base contract tests)
	if len(opts) > 0 {
		if opts[0].Kinds != nil {
			kinds = opts[0].Kinds
		}
		mobileVersion = opts[0].MobileVersion
	}

	for _, kind := range kinds {
		switch kind {
		case "mobile":
			cleaned := cleanForCheck(value)
			if !mobileRegex.MatchString(cleaned) {
				continue
			}
			if mobileVersion == 0 {
				return true
			}
			if IsValidMobile(value, IsValidMobileOptions{Version: mobileVersion}) {
				return true
			}
		case "landline":
			if IsValidLandline(value) {
				return true
			}
		case "service":
			if IsValidService(value) {
				return true
			}
		}
	}
	return false
}

// cleanForCheck removes the usual mask symbols and a leading country code,
// but (unlike Parse) does not strip stray non-digit characters: a regex
// match against the result will correctly reject them.
func cleanForCheck(value string) string {
	return RemoveSymbols(RemoveInternationalDialingCode(value))
}

// IsValidMobileOptions controls IsValidMobile.
type IsValidMobileOptions struct {
	// Version picks the numbering rule: 1 (default) accepts a first
	// subscriber digit of 6, 7, 8 or 9; 2 follows Resolução Anatel nº
	// 749/2022: only 7, 8 or 9, rejecting the 700 (satellite) series.
	Version int
}

// IsValidMobile validates a mobile number: DDD plus 9 digits. A country
// code is accepted and removed first.
func IsValidMobile(value string, opts ...IsValidMobileOptions) bool {
	version := 1
	if len(opts) > 0 && opts[0].Version == 2 {
		version = 2
	}

	cleaned := cleanForCheck(value)
	if !mobileRegex.MatchString(cleaned) {
		return false
	}

	subscriber := cleaned[3:]
	firstDigit := subscriber[0]

	if version == 2 {
		if firstDigit < '7' || firstDigit > '9' {
			return false
		}
		if strings.HasPrefix(subscriber, "700") {
			return false
		}
		return true
	}

	return firstDigit >= '6' && firstDigit <= '9'
}

// IsValidLandline validates a landline number: DDD plus 8 digits. A
// country code is accepted and removed first.
func IsValidLandline(value string) bool {
	cleaned := cleanForCheck(value)
	return landlineRegex.MatchString(cleaned)
}

// IsValidService validates a Brazilian service number, dialed without a
// DDD: the Código Não Geográfico (0300/0303/0500/0800/0900 + 7 digits),
// the abbreviated 300X/400X form (8 digits), or a 3-digit public utility
// code (112 and 911 excluded).
func IsValidService(value string) bool {
	cleaned := cleanForCheck(value)
	switch len(cleaned) {
	case 11:
		return serviceNGCRegex.MatchString(cleaned)
	case 8:
		return serviceAbbreviatedRegex.MatchString(cleaned)
	case 3:
		return publicUtilityCodes[cleaned]
	default:
		return false
	}
}

// FormatOptions controls Format.
type FormatOptions struct {
	// Mask picks the display pattern: "sn" (subscriber number only,
	// default), "ddd" (DDD + number), "e164", "international", "service"
	// or "auto" (chosen from the structure of the number).
	Mask string
}

// Format formats a phone number according to Brazilian patterns. Without
// options, the default mask is the subscriber number only (no DDD),
// masked as far as the digits go.
func Format(value string, opts ...FormatOptions) string {
	mask := "sn"
	if len(opts) > 0 && opts[0].Mask != "" {
		mask = opts[0].Mask
	}

	switch mask {
	case "ddd":
		return formatDDDAndNumber(value)
	case "e164":
		digits := Parse(value)
		if len(digits) != 10 && len(digits) != 11 {
			return ""
		}
		return "+55" + digits
	case "international":
		digits := Parse(value)
		if len(digits) != 10 && len(digits) != 11 {
			return ""
		}
		ddd := digits[:2]
		number := digits[2:]
		return "+55 " + ddd + " " + snMask(number)
	case "service":
		digits := helpers.OnlyNumbers(value)
		if len(digits) == 11 {
			return digits[:4] + " " + digits[4:7] + " " + digits[7:]
		}
		if len(digits) == 8 {
			return digits[:4] + "-" + digits[4:]
		}
		return ""
	case "auto":
		if IsValidService(value) {
			return Format(value, FormatOptions{Mask: "service"})
		}
		if IsValidMobile(value) || IsValidLandline(value) {
			return formatDDDAndNumber(value)
		}
		return Format(value, FormatOptions{Mask: "sn"})
	default:
		digits := helpers.OnlyNumbers(value)
		if len(digits) > 9 {
			digits = digits[:9]
		}
		return snMask(digits)
	}
}

// snMask masks a digit string as "NNNNN-NNNN" (or as far as it goes).
func snMask(digits string) string {
	if len(digits) < 5 {
		return digits
	}
	return digits[:5] + "-" + digits[5:]
}

func formatDDDAndNumber(value string) string {
	digits := Parse(value)
	if len(digits) != 10 && len(digits) != 11 {
		return ""
	}
	ddd := digits[:2]
	number := digits[2:]
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

// Parse removes phone formatting and keeps only digits, capped to 11
// digits. A country code (+55, 0055 or a bare 55) is stripped first, only
// when 10 or 11 digits are left, so a DDD of 55 is not mistaken for it.
func Parse(value string) string {
	cleaned := RemoveInternationalDialingCode(value)
	digits := helpers.OnlyNumbers(cleaned)
	if len(digits) > 11 {
		digits = digits[:11]
	}
	return digits
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
