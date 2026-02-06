package email

import "regexp"

// emailRegex validates email addresses following RFC 5322 conventions.
// The first character of the local part must not be a dot.
var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9_%+\-][a-zA-Z0-9._%+\-]*@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

// IsValid checks if a string corresponds to a valid email address.
func IsValid(email string) bool {
	return emailRegex.MatchString(email)
}
