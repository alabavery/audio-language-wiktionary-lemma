package token

import (
	"audio-language/wiktionary/lemma/util"
)

// TokensItem is a single part of speech (or irrelevant datapoint)
type TokensItem struct {
	Name   string   `json:"name"`
	Tokens []string `json:"tokens"`
}

// TokensWrapper is struct wrapping the content of a word's tokens
// file (or wrapping nothing, if no such file exists)
type TokensWrapper struct {
	Word       string
	Content    []TokensItem
	FileExists bool
}

// NewTokensWrapper gives TokensWrapper pointer
func NewTokensWrapper(word string, tokensDir string) *TokensWrapper {
	content, exists := getFileContent(word, tokensDir)
	token := &TokensWrapper{
		Word:       word,
		Content:    content,
		FileExists: exists,
	}
	return token
}

func getFileContent(word string, tokensDir string) ([]TokensItem, bool) {
	var content []TokensItem
	exists := util.GetJSONWhenFileMayNotExist(
		tokensDir+"/"+word+".json",
		&content,
	)
	return content, exists
}

// GetItem - get an item of a given name
func (t TokensWrapper) GetItem(itemName string) TokensItem {
	// find the item by name in content
	for _, item := range t.Content {
		if item.Name == itemName {
			return item
		}
	}
	panic("No item found")
}
