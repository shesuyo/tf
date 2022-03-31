package tf

import "regexp"

var rgxIsPhoneNumber = regexp.MustCompile(`^1\d{10}$`)

func IsPhoneNumber(phone string) bool {
	return rgxIsPhoneNumber.MatchString(phone)
}
