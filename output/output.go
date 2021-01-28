package output

import (
	"audio-language/wiktionary/lemma/definition"
	"audio-language/wiktionary/lemma/output/parser"
	"audio-language/wiktionary/lemma/token"
	"fmt"
)

type item struct {
	PartOfSpeech string
	Lemma        string
	Exists       bool
}

// LemmasWrapper wraps lemmas
type LemmasWrapper struct {
	Language           string
	Word               string
	Content            []item
	HasContent         bool
	definitionsWrapper *definition.DefinitionsWrapper
	tokensWrapper      *token.TokensWrapper
}

// NewLemmasWrapper gives a LemmasWrapper
func NewLemmasWrapper(
	word string,
	language string,
	d *definition.DefinitionsWrapper,
	t *token.TokensWrapper,
) *LemmasWrapper {
	return &LemmasWrapper{
		Language:           language,
		Word:               word,
		HasContent:         false,
		definitionsWrapper: d,
		tokensWrapper:      t,
	}
}

// GetLemmas saves lemmas for each part of speech into the Content of the LemmasWrapper
func (w *LemmasWrapper) GetLemmas() {
	d := w.definitionsWrapper
	t := w.tokensWrapper
	language := w.Language

	if !d.FileExists {
		// currently, we are relying on definitions to tell us which items in
		// tokens are parts of speech and which are irrelevant (e.g. Pronunciation)
		return
	}
	if !t.FileExists {
		// unexpected state -- tokens should always exist if definitions exist
		panic("definition file exists, but tokens file does not")
	}
	w.HasContent = true

	for _, partOfSpeechItem := range d.Content {
		l := item{
			PartOfSpeech: partOfSpeechItem.PartOfSpeech,
		}
		if partOfSpeechItem.IsDefined() {
			l.Lemma = d.Word
			l.Exists = true
		} else {
			partOfSpeech := partOfSpeechItem.PartOfSpeech
			tokensItemForPos := t.GetItem(partOfSpeech)
			parser := parser.GetParser(language, partOfSpeech)
			parsed, exists := parser(tokensItemForPos)
			l.Exists = exists
			if exists {
				if len(parsed) > 1 {
					fmt.Println("\n", w.Word, partOfSpeech, exists, parsed)
					panic("multiple lemmas??")
				}
				l.Lemma = parsed[0]
			}
		}
		w.Content = append(w.Content, l)
	}
}
