package parser

import (
	"audio-language/wiktionary/lemma/output/parser/common"
	"audio-language/wiktionary/lemma/output/parser/spanish"
	"audio-language/wiktionary/lemma/token"
)

// GetParser gives parser that will get lemma from tokens
func GetParser(language string, word string, partOfSpeech string) func(token.TokensItem) ([]string, bool) {
	singleTokenParser := getSingleTokenParser(language, word, partOfSpeech)

	return func(t token.TokensItem) ([]string, bool) {
		if common.IsOwnLemma(t) {
			return []string{word}, true
		}

		allFound := make(map[string]bool)
		for _, t := range t.Tokens {
			found := singleTokenParser(t)
			for _, f := range found {
				allFound[f] = true
			}
		}
		asArray := []string{}
		for key := range allFound {
			asArray = append(asArray, key)
		}
		return asArray, len(allFound) > 0
	}
}

func getSingleTokenParser(language string, word string, partOfSpeech string) func(string) []string {
	switch language {
	case "spanish":
		return spanish.GetSpanishSingleTokenParser(word, partOfSpeech)
	default:
		panic("language is not supported")
	}
}
