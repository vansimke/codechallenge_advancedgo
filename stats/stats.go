package stats

import "strings"

type Stats struct{}

func (s Stats) TotalLines(lines []string) int {
	return len(lines)
}

func (s Stats) TotalWords(lines []string) int {
	count := 0
	for _, line := range lines {
		count += len(strings.Fields(line))
	}
	return count
}

func (s Stats) WordFrequency(lines []string) map[string]int {
	var result = make(map[string]int)
	for _, line := range lines {
		for _, word := range strings.Fields(line) {
			result[word]++
		}
	}
	return result
}
