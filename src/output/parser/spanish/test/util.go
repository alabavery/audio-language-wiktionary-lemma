package test

import (
	"audio-language/wiktionary/lemma/output/parser/spanish"
	"testing"
)

type testData struct {
	word         string
	token        string
	lemma        string
	partOfSpeech string
}

func runTest(t *testing.T, td testData) {
	method := spanish.GetSpanishSingleTokenParser(td.word, td.partOfSpeech)
	result := method(td.token)
	if len(result) == 0 {
		t.Errorf("For %v, expected %v, but got no results", td.word, td.lemma)
		return
	}
	if result[0] != td.lemma {
		t.Errorf("For %v, expected '%v', but got '%v'", td.word, td.lemma, result[0])
	}
}
