package spanish

import (
	"regexp"
	"strings"
)

func getParseVerbLemma(word string) func(string) []string {
	return basic
}

func basic(t string) []string {
	re := regexp.MustCompile("^[a-z]+r$")
	return onEachSection(t, func(section string) []string {
		return re.FindAllString(section, -1)
	})
}

func onEachSection(t string, parse func(string) []string) []string {
	openBracketSections := strings.Split(t, "{")
	bracketLessSections := []string{}
	for _, s := range openBracketSections {
		bracketLessSections = append(bracketLessSections, strings.Split(s, "}")...)
	}
	sections := []string{}
	for _, s := range bracketLessSections {
		sections = append(sections, strings.Split(s, "|")...)
	}

	results := []string{}
	for _, section := range sections {
		for _, r := range parse(section) {
			results = append(results, r)
		}
	}
	return results
}
