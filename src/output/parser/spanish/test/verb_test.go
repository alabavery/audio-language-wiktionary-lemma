package test

import "testing"

func TestEstoy(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "verb",
		word:         "estoy",
		lemma:        "estar",
		token:        "# {{es-verb form of|mood=ind|tense=pres|num=s|pers=1|ending=ar|estar|nodot=1}}; [[am]]",
	})
}

func TestSe(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "verb",
		word:         "sé",
		lemma:        "ser",
		token:        "# {{es-verb form of|mood=imp|num=s|pers=2|formal=n|sense=+|ending=er|ser}}",
	})
}

func TestEs(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "verb",
		word:         "es",
		lemma:        "ser",
		token:        "# {{es-verb form of|mood=ind|tense=pres|num=s|pers=2|formal=y|ending=er|ser}}",
	})
}

func TestIr(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "verb",
		word:         "voy",
		lemma:        "ir",
		token:        "# {{es-verb form of|mood=ind|tense=pres|num=s|pers=1|ending=ir|ir}}",
	})
}

func TestPuedo(t *testing.T) {
	runTest(t, testData{
		partOfSpeech: "verb",
		word:         "puedo",
		lemma:        "poder",
		token:        "# {{es-verb form of|ending=er|mood=indicative|tense=present|pers=1|number=singular|poder}",
	})
}
