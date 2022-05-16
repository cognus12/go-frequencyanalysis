package frequencyanalysis

import "fmt"

func Top10(text string) []string {
	words := split(text)
	wordsMap := countWords(words)
	counts := mapByCount(wordsMap)
	sorted := sortByCount(counts)

	fmt.Println(sorted)

	return createTop(sorted)
}
