package renavam

import (
	"math/rand"
	"strconv"
	"strings"
)

// Generate generates a random valid RENAVAM (11 digits, unformatted).
func Generate() string {
	for {
		var sb strings.Builder
		for i := 0; i < 10; i++ {
			sb.WriteString(strconv.Itoa(rand.Intn(10)))
		}
		base := sb.String()
		if allSame(base) {
			continue
		}

		candidate := base + strconv.Itoa(checkDigit(base))
		if IsValid(candidate) {
			return candidate
		}
	}
}
