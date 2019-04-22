package utils

import (
	"regexp"
)

// ValidateColorString validstes that the string that the user
// provided is actually a hex color code
func ValidateColorString(cs string) bool {
	if result, err := regexp.MatchString(`^#[A-Fa-f0-9]{6}$`, cs); err == nil {
		return result
	}
	return false
}
