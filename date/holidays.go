package date

import "time"

// IsHoliday checks if the given date is a national or state holiday in Brazil.
// If uf is empty, only national holidays are checked.
// Returns (result, ok). ok is false if the uf is invalid.
func IsHoliday(date time.Time, uf string) (bool, bool) {
	if uf != "" && !isValidUF(uf) {
		return false, false
	}

	year := date.Year()
	month := date.Month()
	day := date.Day()

	// Check national holidays (fixed)
	for _, h := range fixedNationalHolidays {
		if month == h.month && day == h.day {
			return true, true
		}
	}

	// Dia da Consciência Negra became a national holiday in 2024
	if year >= 2024 && month == time.November && day == 20 {
		return true, true
	}

	// Check national holidays (movable, Easter-based)
	easter := computeEaster(year)
	for _, offset := range easterOffsets {
		h := easter.AddDate(0, 0, offset)
		if month == h.Month() && day == h.Day() {
			return true, true
		}
	}

	// Check state holidays if uf is provided
	if uf != "" {
		if holidays, ok := stateHolidays[uf]; ok {
			for _, h := range holidays {
				if month == h.month && day == h.day {
					return true, true
				}
			}
		}
	}

	return false, true
}

type fixedDate struct {
	month time.Month
	day   int
}

// Fixed national holidays
var fixedNationalHolidays = []fixedDate{
	{time.January, 1},   // Confraternização Universal
	{time.April, 21},    // Tiradentes
	{time.May, 1},       // Dia do Trabalho
	{time.September, 7}, // Independência do Brasil
	{time.October, 12},  // Nossa Senhora Aparecida
	{time.November, 2},  // Finados
	{time.November, 15}, // Proclamação da República
	{time.December, 25}, // Natal
}

// Easter-based movable national holidays (offsets from Easter Sunday)
var easterOffsets = []int{
	-48, // Carnival Monday
	-47, // Carnival Tuesday
	-2,  // Good Friday (Sexta-feira Santa)
	60,  // Corpus Christi
}

// State holidays by UF
var stateHolidays = map[string][]fixedDate{
	"AC": {
		{time.January, 23},  // Dia do Evangélico
		{time.June, 15},     // Aniversário do Acre
		{time.September, 5}, // Dia da Amazônia
		{time.November, 17}, // Tratado de Petrópolis
	},
	"AL": {
		{time.June, 24},      // São João
		{time.June, 29},      // São Pedro
		{time.September, 16}, // Emancipação Política
		{time.November, 20},  // Consciência Negra
	},
	"AP": {
		{time.March, 19},    // São José
		{time.July, 25},     // São Tiago
		{time.October, 5},   // Criação do Estado
		{time.November, 20}, // Consciência Negra
	},
	"AM": {
		{time.September, 5}, // Elevação do Amazonas à categoria de província
		{time.November, 20}, // Consciência Negra
		{time.December, 8},  // Nossa Senhora da Conceição
	},
	"BA": {
		{time.July, 2}, // Independência da Bahia
	},
	"CE": {
		{time.March, 19}, // São José
		{time.March, 25}, // Abolição da escravidão no Ceará
	},
	"DF": {
		{time.April, 21},    // Fundação de Brasília
		{time.November, 30}, // Dia do Evangélico
	},
	"ES": {
		{time.October, 28}, // Dia do Servidor Público
	},
	"GO": {
		{time.October, 24}, // Pedra Fundamental de Goiânia
		{time.October, 28}, // Dia do Servidor Público
	},
	"MA": {
		{time.July, 28}, // Adesão do Maranhão à Independência do Brasil
	},
	"MT": {
		{time.November, 20}, // Consciência Negra
	},
	"MS": {
		{time.October, 11}, // Criação do Estado
	},
	"MG": {
		{time.April, 21}, // Data Magna do Estado
	},
	"PA": {
		{time.August, 15}, // Adesão do Grão-Pará à Independência do Brasil
	},
	"PB": {
		{time.August, 5}, // Fundação do Estado
	},
	"PR": {
		{time.December, 19}, // Emancipação Política do Paraná
	},
	"PE": {
		{time.March, 6}, // Data Magna de Pernambuco
	},
	"PI": {
		{time.October, 19}, // Dia do Piauí
	},
	"RJ": {
		{time.April, 23},    // São Jorge
		{time.November, 20}, // Consciência Negra
	},
	"RN": {
		{time.June, 29},   // São Pedro
		{time.October, 3}, // Mártires de Cunhaú e Uruaçu
	},
	"RS": {
		{time.September, 20}, // Revolução Farroupilha
	},
	"RO": {
		{time.January, 4}, // Criação do Estado
		{time.June, 18},   // Dia do Evangélico
	},
	"RR": {
		{time.October, 5}, // Criação do Estado
	},
	"SC": {
		{time.August, 11}, // Data Magna do Estado
	},
	"SP": {
		{time.July, 9}, // Revolução Constitucionalista
	},
	"SE": {
		{time.July, 8}, // Emancipação Política de Sergipe
	},
	"TO": {
		{time.March, 18},    // Autonomia do Estado do Tocantins
		{time.September, 8}, // Nossa Senhora da Natividade
		{time.October, 5},   // Criação do Estado
	},
}

var validUFs = []string{
	"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO",
	"MA", "MT", "MS", "MG", "PA", "PB", "PR", "PE", "PI",
	"RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO",
}

func isValidUF(uf string) bool {
	for _, v := range validUFs {
		if v == uf {
			return true
		}
	}
	return false
}

// computeEaster calculates the date of Easter Sunday for a given year
// using the Anonymous Gregorian algorithm (Computus).
func computeEaster(year int) time.Time {
	a := year % 19
	b := year / 100
	c := year % 100
	d := b / 4
	e := b % 4
	f := (b + 8) / 25
	g := (b - f + 1) / 3
	h := (19*a + b - d - g + 15) % 30
	i := c / 4
	k := c % 4
	l := (32 + 2*e + 2*i - h - k) % 7
	m := (a + 11*h + 22*l) / 451
	month := (h + l - 7*m + 114) / 31
	day := ((h + l - 7*m + 114) % 31) + 1

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
