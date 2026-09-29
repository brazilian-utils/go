// Package number writes an integer out in Brazilian Portuguese words
// ("por extenso").
package number

import (
	"math"
	"strings"
)

var onesMasculine = []string{
	"zero", "um", "dois", "três", "quatro", "cinco",
	"seis", "sete", "oito", "nove", "dez",
	"onze", "doze", "treze", "quatorze", "quinze",
	"dezesseis", "dezessete", "dezoito", "dezenove",
}

var onesFeminine = []string{
	"zero", "uma", "duas", "três", "quatro", "cinco",
	"seis", "sete", "oito", "nove", "dez",
	"onze", "doze", "treze", "quatorze", "quinze",
	"dezesseis", "dezessete", "dezoito", "dezenove",
}

var tensWords = []string{
	"", "", "vinte", "trinta", "quarenta", "cinquenta",
	"sessenta", "setenta", "oitenta", "noventa",
}

var hundredsMasculine = []string{
	"", "cento", "duzentos", "trezentos", "quatrocentos", "quinhentos",
	"seiscentos", "setecentos", "oitocentos", "novecentos",
}

var hundredsFeminine = []string{
	"", "cento", "duzentas", "trezentas", "quatrocentas", "quinhentas",
	"seiscentas", "setecentas", "oitocentas", "novecentas",
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

const maxValue = 999_999_999_999_999
const minValue = -999_999_999_999_999

// ConvertToWordsOptions controls ConvertToWords.
type ConvertToWordsOptions struct {
	// Feminine selects the feminine grammatical gender, which agrees
	// "um/uma", "dois/duas" and the hundreds. Default: masculine.
	Feminine bool
}

// ConvertToWords writes an integer out in Brazilian Portuguese cardinal
// words ("por extenso"), e.g. 1235 is "mil duzentos e trinta e cinco".
// A non-integer is truncated toward zero. Returns an empty string for a
// value outside -999999999999999..999999999999999 or not finite.
func ConvertToWords(value float64, opts ...ConvertToWordsOptions) string {
	if math.IsNaN(value) || math.IsInf(value, 0) {
		return ""
	}
	truncated := math.Trunc(value)
	if truncated < minValue || truncated > maxValue {
		return ""
	}

	feminine := len(opts) > 0 && opts[0].Feminine

	n := int64(truncated)
	if n == 0 {
		return "zero"
	}
	if n < 0 {
		return "menos " + convert(-n, feminine)
	}
	return convert(n, feminine)
}

func convert(n int64, feminine bool) string {
	if n == 0 {
		return "zero"
	}

	type group struct {
		text  string
		value int
	}
	var groups []group

	for _, s := range scales {
		if n >= s.value {
			count := n / s.value
			n %= s.value

			scaleName := s.plural
			if count == 1 {
				scaleName = s.singular
			}

			if s.value == 1000 && count == 1 {
				groups = append(groups, group{"mil", int(count)})
			} else {
				groups = append(groups, group{convertGroup(int(count), feminine) + " " + scaleName, int(count)})
			}
		}
	}

	if n > 0 {
		groups = append(groups, group{convertGroup(int(n), feminine), int(n)})
	}

	if len(groups) == 1 {
		return groups[0].text
	}

	lastIdx := len(groups) - 1
	lastValue := groups[lastIdx].value
	useE := lastValue < 100 || lastValue%100 == 0

	var sb strings.Builder
	for i := 0; i < lastIdx; i++ {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(groups[i].text)
	}
	if useE {
		sb.WriteString(" e ")
	} else {
		sb.WriteString(" ")
	}
	sb.WriteString(groups[lastIdx].text)

	return sb.String()
}

func convertGroup(n int, feminine bool) string {
	if n == 0 {
		return ""
	}
	if n == 100 {
		return "cem"
	}

	ones := onesMasculine
	hundreds := hundredsMasculine
	if feminine {
		ones = onesFeminine
		hundreds = hundredsFeminine
	}

	var parts []string

	if n >= 100 {
		parts = append(parts, hundreds[n/100])
		n %= 100
	}

	if n >= 20 {
		parts = append(parts, tensWords[n/10])
		n %= 10
		if n > 0 {
			parts = append(parts, ones[n])
		}
	} else if n > 0 {
		parts = append(parts, ones[n])
	}

	return strings.Join(parts, " e ")
}
