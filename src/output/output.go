package output

import (
	"audio-language/wiktionary/lemma/constants"
	"audio-language/wiktionary/lemma/output/parser"
	"audio-language/wiktionary/lemma/token"
)

type item struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Lemmas       []string `json:"lemmas"`
	Exists       bool     `json:"exists"`
}

// LemmasWrapper wraps lemmas
type LemmasWrapper struct {
	Language      string
	Word          string
	Content       []item
	HasContent    bool
	tokensWrapper *token.TokensWrapper
}

// NewLemmasWrapper gives a LemmasWrapper
func NewLemmasWrapper(
	word string,
	language string,
	t *token.TokensWrapper,
) *LemmasWrapper {
	return &LemmasWrapper{
		Language:      language,
		Word:          word,
		HasContent:    false,
		tokensWrapper: t,
	}
}

// GetLemmas saves lemmas for each part of speech into the Content of the LemmasWrapper
func (w *LemmasWrapper) GetLemmas() {
	t := w.tokensWrapper
	language := w.Language

	if !t.FileExists {
		// fmt.Printf("\nSkipping %v -- no tokens file exists\n", w.Word)
		return
	}
	w.HasContent = true

	for _, section := range t.Content {
		if !isPartOfSpeechName(section.Name) {
			continue
		}
		partOfSpeech := section.Name
		l := item{
			PartOfSpeech: partOfSpeech,
			Exists:       false,
		}
		tokensItemForPos := t.GetItem(partOfSpeech)
		parser := parser.GetParser(language, w.Word, partOfSpeech)
		parsed, exists := parser(tokensItemForPos)
		l.Exists = exists
		l.Lemmas = parsed

		w.Content = append(w.Content, l)
	}
}

func isPartOfSpeechName(name string) bool {
	_, exists := constants.PartsOfSpeech[name]
	return exists
}
