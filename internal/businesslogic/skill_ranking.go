package businesslogic

import (
	"fmt"
	"sort"
)

func rankKeywords(keywords []string) []string {
	keywordCounts := make(map[string]int)
	for _, keyword := range keywords {
		if keyword == "" {
			continue
		}
		keywordCounts[keyword]++
	}
	rankedKeywords := make([]string, 0, len(keywordCounts))
	for keyword, count := range keywordCounts {
		rankedKeywords = append(rankedKeywords, fmt.Sprintf("%s: %d", keyword, count))
	}
	sort.Slice(rankedKeywords, func(i, j int) bool {
		return rankedKeywords[i] < rankedKeywords[j]
	})
	return rankedKeywords
}
