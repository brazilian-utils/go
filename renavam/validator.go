package renavam

const renavamSize = 11

var weights = []int{2, 3, 4, 5, 6, 7, 8, 9, 2, 3}

// IsValid validates a Brazilian vehicle registration number (RENAVAM).
// A valid RENAVAM is exactly 11 digits, not all the same digit, with a
// correct check digit.
func IsValid(renavam string) bool {
	if len(renavam) != renavamSize || !isDigits(renavam) || allSame(renavam) {
		return false
	}

	return checkDigit(renavam) == int(renavam[10]-'0')
}

// checkDigit computes the verification digit from the first 10 digits (reversed).
func checkDigit(renavam string) int {
	sum := 0
	for i := 0; i < 10; i++ {
		// Reverse: digit at position 9-i gets weight at position i
		sum += int(renavam[9-i]-'0') * weights[i]
	}
	dv := 11 - (sum % 11)
	if dv >= 10 {
		return 0
	}
	return dv
}

func isDigits(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return len(s) > 0
}

func allSame(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] != s[0] {
			return false
		}
	}
	return true
}
