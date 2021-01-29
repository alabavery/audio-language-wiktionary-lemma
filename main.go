package main

import (
	"audio-language/wiktionary/lemma/output"
	"audio-language/wiktionary/lemma/token"
	"audio-language/wiktionary/lemma/word"
	"flag"
	"fmt"
)

func main() {
	wordsFile, tokensDir, language := getFlags()
	words := word.GetWords(wordsFile)

	for _, word := range words {
		t := token.NewTokensWrapper(word, tokensDir)
		l := output.NewLemmasWrapper(word, language, t)
		l.GetLemmas()
		fmt.Println("\n\n", word, ":")
		for _, pos := range l.Content {
			fmt.Printf("\tpart of speech: %v; has-lemma: %v; lemma: %v\n", pos.PartOfSpeech, pos.Exists, pos.Lemma)
		}
	}
}

func getFlags() (string, string, string) {
	wordsFilePtr := flag.String("words", "", "the path of the words file")
	tokensDirPtr := flag.String("tokens", "", "the path of the tokens directory")
	languagePtr := flag.String("language", "", "the subject language")
	flag.Parse()

	if *wordsFilePtr == "" {
		panic("need a -words flag")
	}
	if *tokensDirPtr == "" {
		panic("need a -tokens flag")
	}
	if *languagePtr == "" {
		panic("need a -language flag")
	}

	return *wordsFilePtr, *tokensDirPtr, *languagePtr
}
