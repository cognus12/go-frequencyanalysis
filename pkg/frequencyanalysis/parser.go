package frequencyanalysis

import (
	"strings"
)

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
