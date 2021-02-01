package main

import (
	"audio-language/wiktionary/lemma/output"
	"audio-language/wiktionary/lemma/token"
	"audio-language/wiktionary/lemma/word"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	wordsFile, tokensDir, language, targetDirectory := getFlags()
	words := word.GetWords(wordsFile)

	lemmas := []*output.LemmasWrapper{}
	total := 0
	lemmad := 0
	for i, word := range words {
		t := token.NewTokensWrapper(word, tokensDir)
		l := output.NewLemmasWrapper(word, language, t)
		lemmas = append(lemmas, l)

		l.GetLemmas()
		partsOfSpeech, lemmas := getSingleWordStats(l, i+1)
		total += partsOfSpeech
		lemmad += lemmas
	}
	save(targetDirectory, lemmas)
	fmt.Printf("\nTotal items: %v; Total lemmas: %v (%v)", total, lemmad, float32(lemmad)/float32(total))
}

func getFlags() (string, string, string, string) {
	wordsFilePtr := flag.String("words", "", "the path of the words file")
	tokensDirPtr := flag.String("tokens", "", "the path of the tokens directory")
	languagePtr := flag.String("language", "", "the subject language")
	targetDirectoryPointer := flag.String("target", "", "the path of the directory to save to")
	flag.Parse()

	if *wordsFilePtr == "" {
		panic("need a -words flag")
	}
	if *tokensDirPtr == "" {
		panic("need a -tokens flag")
	}
	language := *languagePtr
	if language == "" {
		fromEnv := os.Getenv("TARGET_LANGUAGE")
		if fromEnv == "" {
			panic("need a -language flag or a TARGET_LANGUAGE env var")
		} else {
			language = fromEnv
		}
	}

	return *wordsFilePtr, *tokensDirPtr, language, *targetDirectoryPointer
}

func getSingleWordStats(l *output.LemmasWrapper, wordRank int) (int, int) {
	total := 0
	lemmad := 0
	for _, pos := range l.Content {
		total++
		if pos.Exists {
			lemmad++
		} else {
			// fmt.Printf("\nrank: %v; word: %v; part of speech: %v\n", wordRank, l.Word, pos.PartOfSpeech)
		}
	}
	return total, lemmad
}

func save(targetDirectory string, lemmas []*output.LemmasWrapper) {
	for _, l := range lemmas {
		if l.HasContent {
			out, err := json.Marshal(l.Content)
			if err != nil {
				panic("Could not marshal json")
			}
			err = ioutil.WriteFile(fmt.Sprintf("%v/%v.json", targetDirectory, l.Word), out, 0644)
			if err != nil {
				panic("could not save file")
			}
		}
	}
}
