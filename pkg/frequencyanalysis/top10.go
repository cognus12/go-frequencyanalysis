package frequencyanalysis

import (
	"fmt"
	"regexp"
	"strings"
)

var rgxp *regexp.Regexp

// var pattern string = `^,|,$|^\.|\.{1,}$|^-|-$|^!|!$|^:|:$|^"|"$`
var pattern string = `^[,|-|!|:|;|"|\.{1,}]|[,|-|!|:|;|"|\.{1,}]$`

func init() {
	rgxp = regexp.MustCompile(pattern)
}

func CleanWord(w string) string {
	result := rgxp.ReplaceAllString(w, "")

	return strings.ToLower(result)
}

func Top10(s string) []string {

	var res []string

	for _, w := range strings.Fields(s) {

		fmt.Println("origin: ", w, "cleaned: ", CleanWord(w))
		res = append(res, CleanWord(w))
	}

	return res
}
