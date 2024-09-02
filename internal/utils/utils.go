package utils

import "strings"

func NormalizeInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
