package help

import (
	"regexp"
	"strings"
)

func Slugify(value string) string {
	invalidChars := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	whiteSpace := regexp.MustCompile(`\s+`)
	dashes := regexp.MustCompile(`-+`)
	value = invalidChars.ReplaceAllString(value, "-")
	value = whiteSpace.ReplaceAllString(value, "-")
	value = dashes.ReplaceAllString(value, "-")
	return strings.ToLower(value)
}
