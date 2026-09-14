package date

import (
	"regexp"
	"strings"
	"time"

	"github.com/brazilian-utils/go/helpers"
)

var dateRegex = regexp.MustCompile(`^\d{2}/\d{2}/\d{4}$`)

// Month names in lowercase, per Brazilian orthographic norm (months are
// common nouns, not proper nouns) and the LOA extenso convention (Lei
// 14.822/2024).
var monthNames = []string{
	"",
	"janeiro", "fevereiro", "março", "abril",
	"maio", "junho", "julho", "agosto",
	"setembro", "outubro", "novembro", "dezembro",
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
