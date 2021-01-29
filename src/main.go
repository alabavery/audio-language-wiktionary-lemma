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

	total := 0
	lemmad := 0
	for i, word := range words {
		t := token.NewTokensWrapper(word, tokensDir)
		l := output.NewLemmasWrapper(word, language, t)
		l.GetLemmas()

		tot, lem := getStats(l, i+1)
		total += tot
		lemmad += lem
	}
	fmt.Printf("\nTotal items: %v; Total lemmas: %v (%v)", total, lemmad, float32(lemmad)/float32(total))
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

func getStats(l *output.LemmasWrapper, wordRank int) (int, int) {
	total := 0
	lemmad := 0
	for _, pos := range l.Content {
		total++
		if pos.Exists {
			lemmad++
		} else {
			fmt.Printf("\nrank: %v; word: %v; part of speech: %v\n", wordRank, l.Word, pos.PartOfSpeech)
		}
	}
	return total, lemmad
}
