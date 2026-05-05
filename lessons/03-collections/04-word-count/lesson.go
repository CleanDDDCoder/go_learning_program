package word_count

import "strings"

// CountWords returns how many times each whitespace-separated word appears.
func CountWords(text string) map[string]int {
	counts := map[string]int{}
	for _, word := range strings.Fields(text) {
		counts[word]++
	}
	return counts
}
