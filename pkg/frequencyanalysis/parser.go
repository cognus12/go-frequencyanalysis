package frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	Word  string
	Count int
}

func cleanWord(w string) string {
	result := rgxp.ReplaceAllString(w, "")

	return strings.ToLower(result)
}

func split(text string) (slice []string) {

	for _, w := range strings.Fields(text) {

		slice = append(slice, cleanWord(w))
	}

	return slice
}

func countWords(words []string) (m map[string]int) {

	m = make(map[string]int)

	for _, w := range words {

		if w == "" {
			continue
		}

		count, exists := m[w]

		if exists {
			m[w] = count + 1
		} else {
			m[w] = 1
		}

	}

	return m
}

func mapByCount(m map[string]int) []wordCount {
	var data []wordCount

	for w, c := range m {
		data = append(data, wordCount{w, c})
	}

	return data
}

func sortByCount(data []wordCount) []wordCount {
	sort.Slice(data, func(i, j int) bool {
		if data[i].Count == data[j].Count {
			return data[i].Word < data[j].Word
		}

		return data[i].Count > data[j].Count
	})

	var limit int

	if len(data) > 10 {
		limit = 10
	} else {
		limit = len(data)
	}

	top := data[0:limit]

	return top
}

func createTop(data []wordCount) []string {
	var top []string

	for _, w := range data {
		top = append(top, w.Word)
	}

	return top
}
