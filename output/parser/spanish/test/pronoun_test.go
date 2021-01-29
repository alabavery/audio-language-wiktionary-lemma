package test

import "testing"

func TestLos(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "pronoun",
		word:         "los",
		lemma:        "ellos",
		token:        "# {{inflection of|es|ellos||acc}} and [[ustedes]] (when referring to more than one man); [[them]], [[you all]] (formal)",
	})
}
