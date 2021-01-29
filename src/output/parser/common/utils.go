package common

import (
	"audio-language/wiktionary/lemma/token"
	"strings"
)

// IsOwnLemma tells you if a word is its own lemma
// we are looking for a token like  "{{head|es|verb form}}",
// if such a line is present, this is NOT its own lemma
// note that not all entries have a token with "head".  For those
// that don't we will consider them their own lemma
func IsOwnLemma(item token.TokensItem) bool {
	for _, t := range item.Tokens {
		if tokenIsHead(t) {
			return !strings.Contains(t, "form")
		}
	}
	return true
}

func tokenIsHead(singleToken string) bool {
	return strings.Contains(singleToken, "{head")
}
