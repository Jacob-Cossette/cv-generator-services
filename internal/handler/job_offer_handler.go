package handler

func extractKeywords(htmlPage string) []string {
    """Extracts keywords from an HTML page."""
    text := strings.ToLower(htmlPage)
    text = regexp.MustCompile(`[^\w\s]`).ReplaceAllString(text, " ")
    keywords := strings.Split(text, " ")
    return keywords
}