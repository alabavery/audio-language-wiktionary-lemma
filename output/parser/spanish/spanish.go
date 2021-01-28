package spanish

// GetSpanishSingleTokenParser gets parser for particular part of speech
func GetSpanishSingleTokenParser(partOfSpeech string) func(string) []string {
	switch partOfSpeech {
	case "verb":
		return parseVerbLemma
	default:
		return func(string) []string {
			return []string{}
		}
	}
}
