package date

import (
	"regexp"
	"strings"
	"time"

	"github.com/brazilian-utils/brutils-go/helpers"
)

var dateRegex = regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)

var monthNames = []string{
	"",
	"Janeiro", "Fevereiro", "Março", "Abril",
	"Maio", "Junho", "Julho", "Agosto",
	"Setembro", "Outubro", "Novembro", "Dezembro",
}

// ConvertDateToText converts a date in Brazilian format (dd/mm/yyyy) to its
// Portuguese text representation.
// Returns empty string if the date is invalid.
func ConvertDateToText(date string) string {
	if !dateRegex.MatchString(date) {
		return ""
	}

	t, err := time.Parse("02/01/2006", date)
	if err != nil {
		return ""
	}

	day := t.Day()
	month := int(t.Month())
	year := t.Year()

	var dayStr string
	if day == 1 {
		dayStr = "Primeiro"
	} else {
		dayStr = helpers.NumberToPortuguese(int64(day))
		dayStr = strings.ToUpper(dayStr[:1]) + dayStr[1:]
	}

	yearStr := helpers.NumberToPortuguese(int64(year))

	return dayStr + " de " + monthNames[month] + " de " + yearStr
}
