package email

import (
	"math/rand"
	"regexp"
	"strings"
)

func IsValid(email string) bool {
	/*
	   Check if a string corresponds to a valid email address.

	   Args:
	       email(str): The input string to be checked.

	   Returns:
	       bool: True if email is a valid email address, False otherwise.

	   Example:
	       >>> isVaild("brutils@brutils.com")
	       True
	       >>> isVaild("invalid-email@brutils")
	       False

	   .. note::
	       The rules for validating an email address generally follow the
	       specifications defined by RFC 5322 (updated by RFC 5322bis),
	       which is the widely accepted standard for email address formats.
	*/
	regexPattern := `^[a-zA-Z0-9][a-zA-Z0-9.]+@[a-zA-Z]+\.[a-zA-Z]{2,63}$`
	emailRegex := regexp.MustCompile(regexPattern)

	return emailRegex.MatchString(email)
}
func randomStringFromSet(length int, charSet string) string {
	var output strings.Builder
	for i := 0; i < length; i++ {
		randomChar := charSet[rand.Intn(len(charSet))]
		output.WriteByte(randomChar)
	}
	return output.String()
}

func Generate() string {
	/*
			Generates a random email

			Args:
				No args
			Returns:
				string: A valid email
			Example:
				>>> Generate()
				example@example.com
			.. note::
		        The rules for generate an email address generally follow the
		        specifications defined by RFC 5322 (updated by RFC 5322bis),
		        which is the widely accepted standard for email address formats.
	*/

	localPart := randomStringFromSet(1, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") +
		randomStringFromSet(rand.Intn(10)+1, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789.")

	// Domain name: Only letters
	domain := randomStringFromSet(rand.Intn(10)+1, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	// TLD: 2 to 63 letters
	tld := randomStringFromSet(rand.Intn(62)+2, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	return localPart + "@" + domain + "." + tld
}
