package utils

import (
	"regexp"
)

func validateColorString(cs string) bool {
	if result, err := regexp.MatchString(`^#[A-Fa-f0-9]{6}$`, cs); err == nil {
		return result
	}
	return false
}
