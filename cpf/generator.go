package cpf

import (
	"fmt"
	"math/rand"
	"strings"
)

// Generate generates a random valid CPF string.
func Generate() string {
	base := fmt.Sprintf("%09d", rand.Intn(999999998)+1)

	first := computeMod(strings.Split(base, ""))
	withFirst := base + fmt.Sprintf("%d", first)
	second := computeMod(strings.Split(withFirst, ""))

	return fmt.Sprintf("%s%d%d", base, first, second)
}
