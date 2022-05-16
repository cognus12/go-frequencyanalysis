package frequencyanalysis

import "regexp"

var rgxp *regexp.Regexp

var pattern string = `^-|-$|[,|!|:|;|"|\.{1,}]`

func init() {
	rgxp = regexp.MustCompile(pattern)
}
