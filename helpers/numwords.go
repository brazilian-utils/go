package helpers

import "strings"

var onesWords = []string{
	"zero", "um", "dois", "três", "quatro", "cinco",
	"seis", "sete", "oito", "nove", "dez",
	"onze", "doze", "treze", "quatorze", "quinze",
	"dezesseis", "dezessete", "dezoito", "dezenove",
}

var tensWords = []string{
	"", "", "vinte", "trinta", "quarenta", "cinquenta",
	"sessenta", "setenta", "oitenta", "noventa",
}

var hundredsWords = []string{
	"", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos",
	"seiscentos", "setecentos", "oitocentos", "novecentos",
}

type scaleUnit struct {
	singular string
	plural   string
	value    int64
}

var scales = []scaleUnit{
	{"trilhão", "trilhões", 1_000_000_000_000},
	{"bilhão", "bilhões", 1_000_000_000},
	{"milhão", "milhões", 1_000_000},
	{"mil", "mil", 1_000},
}

// convertGroup converts a number 0-999 to Portuguese words.
func convertGroup(n int) string {
	if n == 0 {
		return ""
	}
	if n == 100 {
		return "cem"
	}

	var parts []string

	if n >= 100 {
		parts = append(parts, hundredsWords[n/100])
		n %= 100
	}

	if n >= 20 {
		parts = append(parts, tensWords[n/10])
		n %= 10
		if n > 0 {
			parts = append(parts, onesWords[n])
		}
	} else if n > 0 {
		parts = append(parts, onesWords[n])
	}

	return strings.Join(parts, " e ")
}

// NumberToPortuguese converts a non-negative integer to its
// Brazilian Portuguese text representation.
func NumberToPortuguese(n int64) string {
	if n == 0 {
		return "zero"
	}

	type group struct {
		text  string
		value int // the 3-digit group count (0-999)
	}

	var groups []group

	for _, s := range scales {
		if n >= s.value {
			count := int(n / s.value)
			n %= s.value

			scaleName := s.plural
			if count == 1 {
				scaleName = s.singular
			}

			// "mil" has no "um" prefix in Portuguese
			if s.value == 1000 && count == 1 {
				groups = append(groups, group{"mil", count})
			} else {
				groups = append(groups, group{convertGroup(count) + " " + scaleName, count})
			}
		}
	}

	// Units group (0-999)
	if n > 0 {
		groups = append(groups, group{convertGroup(int(n)), int(n)})
	}

	if len(groups) == 1 {
		return groups[0].text
	}

	// Join all groups: use ", " between groups, except the last
	// connector uses " e " when the last group is < 100 or a round hundred.
	lastIdx := len(groups) - 1
	lastValue := groups[lastIdx].value
	useE := lastValue < 100 || lastValue%100 == 0

	var sb strings.Builder
	for i := 0; i < lastIdx; i++ {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(groups[i].text)
	}

	if useE {
		sb.WriteString(" e ")
	} else {
		sb.WriteString(", ")
	}
	sb.WriteString(groups[lastIdx].text)

	return sb.String()
}
