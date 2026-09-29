package date

import (
	"sort"
	"time"
)

// minYear and maxYear bound every date/business-day calculation in this
// package, per the shared brazilian-utils contract.
const minYear = 1900
const maxYear = 2099

// IsHolidayOptions carries the date to check and, optionally, a state
// (UF) whose holidays also count. An unknown UF is ignored (as if omitted).
type IsHolidayOptions struct {
	Date time.Time
	UF   string
}

// IsHoliday checks whether a given date is a Brazilian holiday.
// Returns false when the date is missing (zero value), invalid (year
// outside 1900-2099) or the options are omitted.
func IsHoliday(opts ...IsHolidayOptions) bool {
	if len(opts) == 0 {
		return false
	}
	o := opts[0]
	if o.Date.IsZero() {
		return false
	}
	if o.Date.Year() < minYear || o.Date.Year() > maxYear {
		return false
	}
	uf := o.UF
	if uf != "" && !isValidUF(uf) {
		uf = ""
	}
	holiday, _ := isHoliday(o.Date, uf)
	return holiday
}

// isHoliday is the internal (pre-contract) implementation: it reports
// whether the date is a holiday and whether the given uf was recognized.
func isHoliday(date time.Time, uf string) (bool, bool) {
	if uf != "" && !isValidUF(uf) {
		return false, false
	}

	year := date.Year()
	month := date.Month()
	day := date.Day()

	for _, h := range fixedNationalHolidays {
		if month == h.month && day == h.day {
			return true, true
		}
	}

	// Dia da Consciência Negra became a national holiday in 2024
	if year >= 2024 && month == time.November && day == 20 {
		return true, true
	}

	easter := computeEaster(year)
	for _, offset := range easterOffsets {
		h := easter.AddDate(0, 0, offset)
		if month == h.Month() && day == h.Day() {
			return true, true
		}
	}

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

// isHolidayForBusinessDay is like isHoliday but can exclude the optional
// (facultativo) holidays: Carnaval and Corpus Christi.
func isHolidayForBusinessDay(date time.Time, uf string, includeOptional bool) bool {
	year := date.Year()
	month := date.Month()
	day := date.Day()

	for _, h := range fixedNationalHolidays {
		if month == h.month && day == h.day {
			return true
		}
	}

	if year >= 2024 && month == time.November && day == 20 {
		return true
	}

	easter := computeEaster(year)
	for _, h := range movableNationalHolidays {
		d := easter.AddDate(0, 0, h.offset)
		if month != d.Month() || day != d.Day() {
			continue
		}
		if h.optional && !includeOptional {
			continue
		}
		return true
	}

	if uf != "" {
		if holidays, ok := stateHolidays[uf]; ok {
			for _, h := range holidays {
				if month == h.month && day == h.day {
					return true
				}
			}
		}
	}

	return false
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

var fixedNationalHolidayNames = map[fixedDate]string{
	{time.January, 1}:    "Confraternização Universal",
	{time.April, 21}:     "Tiradentes",
	{time.May, 1}:        "Dia do Trabalho",
	{time.September, 7}:  "Independência do Brasil",
	{time.October, 12}:   "Nossa Senhora Aparecida",
	{time.November, 2}:   "Finados",
	{time.November, 15}:  "Proclamação da República",
	{time.December, 25}:  "Natal",
}

// Easter-based movable national holidays (offsets from Easter Sunday)
var easterOffsets = []int{
	-48, // Carnival Monday
	-47, // Carnival Tuesday
	-2,  // Good Friday (Sexta-feira Santa)
	60,  // Corpus Christi
}

type movableHoliday struct {
	offset   int
	name     string
	optional bool
}

var movableNationalHolidays = []movableHoliday{
	{-48, "Carnaval (segunda-feira)", true},
	{-47, "Carnaval (terça-feira)", true},
	{-2, "Sexta-feira Santa", false},
	{60, "Corpus Christi", true},
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

// Holiday describes a single Brazilian holiday, as returned by GetHolidays.
type Holiday struct {
	Name string
	Date time.Time
	Type string // "national", "state", "optional" or "religious"
}

// GetHolidays returns the national Brazilian holidays of a year, sorted by
// date. Returns an empty list when year is outside 1900-2099.
func GetHolidays(year int) []Holiday {
	if year < minYear || year > maxYear {
		return []Holiday{}
	}

	holidays := make([]Holiday, 0, 12)

	for _, h := range fixedNationalHolidays {
		holidays = append(holidays, Holiday{
			Name: fixedNationalHolidayNames[h],
			Date: time.Date(year, h.month, h.day, 0, 0, 0, 0, time.UTC),
			Type: "national",
		})
	}

	if year >= 2024 {
		holidays = append(holidays, Holiday{
			Name: "Dia Nacional de Zumbi e da Consciência Negra",
			Date: time.Date(year, time.November, 20, 0, 0, 0, 0, time.UTC),
			Type: "national",
		})
	}

	easter := computeEaster(year)
	for _, h := range movableNationalHolidays {
		typ := "national"
		if h.optional {
			typ = "optional"
		}
		holidays = append(holidays, Holiday{
			Name: h.name,
			Date: easter.AddDate(0, 0, h.offset),
			Type: typ,
		})
	}

	sort.Slice(holidays, func(i, j int) bool {
		return holidays[i].Date.Before(holidays[j].Date)
	})

	return holidays
}

// BusinessDayOptions controls the business-day calculations. Optional
// holidays (Carnaval, Corpus Christi) are counted by default; set
// IncludeOptionalHolidays to a false pointer to exclude them. UF adds one
// state's holidays.
type BusinessDayOptions struct {
	IncludeOptionalHolidays *bool
	UF                      string
}

func (o *BusinessDayOptions) includeOptional() bool {
	if o == nil || o.IncludeOptionalHolidays == nil {
		return true
	}
	return *o.IncludeOptionalHolidays
}

func (o *BusinessDayOptions) uf() string {
	if o == nil {
		return ""
	}
	if o.UF != "" && !isValidUF(o.UF) {
		return ""
	}
	return o.UF
}

func validBusinessDate(t time.Time) bool {
	return !t.IsZero() && t.Year() >= minYear && t.Year() <= maxYear
}

// IsBusinessDay checks whether a date is a Brazilian business day (dia
// útil): not a Saturday, a Sunday or a holiday. Returns false for an
// invalid date or a year outside 1900-2099.
func IsBusinessDay(value time.Time, opts ...BusinessDayOptions) bool {
	if !validBusinessDate(value) {
		return false
	}
	var o *BusinessDayOptions
	if len(opts) > 0 {
		o = &opts[0]
	}
	weekday := value.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}
	return !isHolidayForBusinessDay(value, o.uf(), o.includeOptional())
}

// AddBusinessDays adds a number of Brazilian business days to a date,
// skipping weekends and holidays. Returns nil for an invalid date, a
// non-integer amount is not representable in Go (amount is already an
// int), or a result outside 1900-2099.
func AddBusinessDays(date time.Time, amount int, opts ...BusinessDayOptions) *time.Time {
	if !validBusinessDate(date) {
		return nil
	}
	var o *BusinessDayOptions
	if len(opts) > 0 {
		o = &opts[0]
	}

	result := date
	step := 1
	if amount < 0 {
		step = -1
	}
	remaining := amount
	if remaining < 0 {
		remaining = -remaining
	}

	for remaining > 0 {
		result = result.AddDate(0, 0, step)
		if result.Year() < minYear || result.Year() > maxYear {
			return nil
		}
		weekday := result.Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			continue
		}
		if isHolidayForBusinessDay(result, o.uf(), o.includeOptional()) {
			continue
		}
		remaining--
	}

	if result.Year() < minYear || result.Year() > maxYear {
		return nil
	}
	return &result
}

// SubBusinessDays subtracts a number of Brazilian business days from a
// date: the same as AddBusinessDays with the opposite amount.
func SubBusinessDays(date time.Time, amount int, opts ...BusinessDayOptions) *time.Time {
	return AddBusinessDays(date, -amount, opts...)
}

// DifferenceInBusinessDays counts the Brazilian business days between two
// dates: laterDate is never counted, earlierDate is counted when it is a
// business day. Returns nil when either date is invalid or outside
// 1900-2099.
func DifferenceInBusinessDays(laterDate, earlierDate time.Time, opts ...BusinessDayOptions) *int {
	if !validBusinessDate(laterDate) || !validBusinessDate(earlierDate) {
		return nil
	}
	var o *BusinessDayOptions
	if len(opts) > 0 {
		o = &opts[0]
	}

	laterDay := time.Date(laterDate.Year(), laterDate.Month(), laterDate.Day(), 0, 0, 0, 0, time.UTC)
	earlierDay := time.Date(earlierDate.Year(), earlierDate.Month(), earlierDate.Day(), 0, 0, 0, 0, time.UTC)

	if laterDay.Equal(earlierDay) {
		zero := 0
		return &zero
	}

	negative := laterDay.Before(earlierDay)
	from, to := earlierDay, laterDay
	if negative {
		from, to = laterDay, earlierDay
	}

	count := 0
	for d := from; d.Before(to); d = d.AddDate(0, 0, 1) {
		weekday := d.Weekday()
		if weekday == time.Saturday || weekday == time.Sunday {
			continue
		}
		if isHolidayForBusinessDay(d, o.uf(), o.includeOptional()) {
			continue
		}
		count++
	}

	if negative {
		count = -count
	}
	return &count
}
