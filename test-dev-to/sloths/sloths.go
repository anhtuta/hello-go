package sloths

import (
	"strings"
)

// Sloths: con lười
// Slothful: lười biếng

// A string is considered slothful by this function if it contains the word "sloth",
// or it contains the hibiscus emoji, but not the race car emoji
func IsSlothful(s string) bool {
	s = strings.ToLower(s)
	slothsLikeThis := strings.Contains(s, "🌺") && !strings.Contains(s, "🏎️")
	if strings.Contains(s, "sloth") {
		return true
	} else if slothsLikeThis {
		return true
	}
	return false
}
