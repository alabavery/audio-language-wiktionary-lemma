package spanish

import (
	"strings"
)

// type patternComponent struct {
// 	name    string
// 	pattern string
// }

// var pattern1Components = []patternComponent{
// 	{
// 		name:    "pipe-lookbehind",
// 		pattern: "(?<=|)",
// 	},
// 	{
// 		name:    "lemma",
// 		pattern: "[a-z]+",
// 	},
// 	{
// 		name:    "end-brackets",
// 		pattern: "(?=}{1,2}$)",
// 	},
// }

// func makeRegex(patternComponents []patternComponent) *regexp.Regexp {
// 	pattern := ""
// 	for _, component := range patternComponents {
// 		pattern += component.pattern
// 	}
// 	return regexp.MustCompile(pattern)
// }

func parseVerbLemma(tok string) []string {
	return basic(tok)
}

func basic(t string) []string {
	found := []string{}

	sections := strings.Split(t, "|")
	if len(sections) == 0 {
		return found
	}
	firstSection := sections[0]
	if !strings.HasSuffix(firstSection, "form of") {
		return found
	}

	last := sections[len(sections)-1]
	bracketStarts := strings.Index(last, "}")
	if bracketStarts == -1 {
		return found
	}
	found = append(found, last[:bracketStarts])
	return found
}
