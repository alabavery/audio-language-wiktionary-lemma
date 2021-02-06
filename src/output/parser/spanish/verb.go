package spanish

import (
	"regexp"
	"strings"
)

func getParseVerbLemma(word string) func(string) []string {
	return func(singleToken string) []string {
		compoundResults := compound(singleToken)
		if len(compoundResults) > 0 {
			return compoundResults
		}
		return basic(singleToken)
	}
}

func basic(t string) []string {
	re := regexp.MustCompile("^[a-z]+r$")
	return onEachSection(t, func(section string, _ int, _ []string) []string {
		return re.FindAllString(section, -1)
	})
}

func compound(t string) []string {
	return onEachSection(t, func(section string, idx int, allSections []string) []string {
		if strings.Contains(section, "compound") && idx+2 < len(allSections) {
			return []string{allSections[idx+1] + allSections[idx+2]}
		}
		return []string{}
	})
}
