package test

import "testing"

func TestUn(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "adjective",
		word:         "un",
		lemma:        "uno",
		token:        "# {{lb|es|before the noun}} {{apocopic form of|es|uno}} [[one]]",
	})
}
