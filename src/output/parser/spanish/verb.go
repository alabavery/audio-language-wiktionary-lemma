package spanish

import (
	"regexp"
)

func getParseVerbLemma(word string) func(string) []string {
	return basic
}

func basic(t string) []string {
	re := regexp.MustCompile("^[a-z]+r$")
	return onEachSection(t, func(section string, _ int, _ []string) []string {
		return re.FindAllString(section, -1)
	})
}
