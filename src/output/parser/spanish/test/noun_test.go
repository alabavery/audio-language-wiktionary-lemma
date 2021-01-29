package test

import "testing"

func TestAnimales(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "noun",
		word:         "animales",
		lemma:        "animal",
		token:        "# {{plural of|es|animal}",
	})
}

func TestAnos(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "noun",
		word:         "años",
		lemma:        "año",
		token:        "# {{plural of|es|a\u00f1o}}",
	})
}
