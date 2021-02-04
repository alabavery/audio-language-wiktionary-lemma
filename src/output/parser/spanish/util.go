package spanish

import (
	"strings"
)

func getSectionsOfSingleToken(t string) []string {
	openBracketSections := strings.Split(t, "{")
	bracketLessSections := []string{}
	for _, s := range openBracketSections {
		bracketLessSections = append(bracketLessSections, strings.Split(s, "}")...)
	}
	sections := []string{}
	for _, s := range bracketLessSections {
		sections = append(sections, strings.Split(s, "|")...)
	}
	return sections
}

func onEachSection(t string, parse func(string, int, []string) []string) []string {
	sections := getSectionsOfSingleToken(t)
	results := []string{}
	for i, section := range sections {
		results = append(results, parse(section, i, sections)...)
	}
	return results
}

// address those that are <something> of|es|<lemma>
func parseOfEs(t string) []string {
	return onEachSection(t, func(section string, i int, allSections []string) []string {
		if strings.Contains(section, "plural of") ||
			strings.Contains(section, "singular of") ||
			strings.Contains(section, "form of") ||
			strings.Contains(section, "inflection of") {
			if i+2 < len(allSections) {
				if strings.Contains(allSections[i+1], "es") {
					return []string{allSections[i+2]}
				}
			}
		}
		return []string{}
	})
}
