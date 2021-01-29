package spanish

// GetSpanishSingleTokenParser gets parser for particular part of speech
func GetSpanishSingleTokenParser(word string, partOfSpeech string) func(string) []string {
	switch partOfSpeech {
	case "verb":
		return getParseVerbLemma(word)
	default:
		return func(string) []string {
			return []string{}
		}
	}
}
