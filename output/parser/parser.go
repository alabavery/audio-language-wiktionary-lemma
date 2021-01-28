package parser

import (
	"audio-language/wiktionary/lemma/output/parser/spanish"
	"audio-language/wiktionary/lemma/token"
)

// GetParser gives parser that will get lemma from tokens
func GetParser(language string, partOfSpeech string) func(token.TokensItem) ([]string, bool) {
	var singleTokenParser func(string) []string
	switch language {
	case "spanish":
		singleTokenParser = spanish.GetSpanishSingleTokenParser(partOfSpeech)
	default:
		panic("language is not supported")
	}

	return func(t token.TokensItem) ([]string, bool) {
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
