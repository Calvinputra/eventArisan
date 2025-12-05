package helper

import (
	"regexp"
	"strconv"
	"strings"
)

func ToSlug(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove all non-alphanumeric characters except spaces
	reg, _ := regexp.Compile("[^a-z0-9\\s]+")
	s = reg.ReplaceAllString(s, "")

	// Replace multiple spaces with a single hyphen
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, " ", "-")

	// Remove multiple consecutive hyphens
	reg, _ = regexp.Compile("-+")
	s = reg.ReplaceAllString(s, "-")

	return s
}

func EditSlug(s string, i int64) string {
	return strings.ToLower(s) + "-" + strconv.FormatInt(i+1, 10)
}
