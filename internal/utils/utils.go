package utils

import "strings"

func NormalizeInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
