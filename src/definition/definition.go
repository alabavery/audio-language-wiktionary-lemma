package definition

import "github.com/ninetypercentlanguage/misc/files"

// PartOfSpeechDefinitions has definitions for single part of speech
type PartOfSpeechDefinitions struct {
	PartOfSpeech string   `json:"part_of_speech"`
	Definitions  []string `json:"definitions"`
}

// DefinitionsWrapper is struct wrapping the content of a word's definition
// file (or wrapping nothing, if no such file exists)
type DefinitionsWrapper struct {
	Word       string
	Content    []PartOfSpeechDefinitions
	FileExists bool
}

// NewDefinitionsWrapper gives a DefinitionsWrapper
func NewDefinitionsWrapper(word string, definitionsDir string) *DefinitionsWrapper {
	content, exists := getFileContent(word, definitionsDir)
	definition := &DefinitionsWrapper{
		Word:       word,
		Content:    content,
		FileExists: exists,
	}
	return definition
}

func getFileContent(word string, definitionsDir string) ([]PartOfSpeechDefinitions, bool) {
	var posDefs []PartOfSpeechDefinitions
	exists := files.GetJSONWhenFileMayNotExist(
		definitionsDir+"/"+word+".json",
		&posDefs,
	)
	return posDefs, exists
}

// IsDefined - does the part of speech have non-empty definitions
func (d PartOfSpeechDefinitions) IsDefined() bool {
	return len(d.Definitions) > 0
}
