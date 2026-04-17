package utils

import (
	"regexp"
	"strings"
)

func IsValidName(name string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z]+$`, strings.ToLower(strings.TrimSpace(name)))
	return matched
}