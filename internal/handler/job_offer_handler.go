package handler

import (
	"regexp"
	"strings"
)

func extractKeywords(htmlPage string) []string {
	text := strings.ToLower(htmlPage)
	text = regexp.MustCompile(`[^\w\s]`).ReplaceAllString(text, " ")
	keywords := strings.Split(text, " ")
	return keywords
}
