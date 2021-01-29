package spanish

import "testing"

type testData struct {
	word  string
	token string
	lemma string
}

func runTest(t *testing.T, td testData) {
	method := GetSpanishSingleTokenParser(td.word, "verb")
	result := method(td.token)
	if len(result) == 0 {
		t.Errorf("For %v, expected %v, but got no results", td.word, td.lemma)
		return
	}
	if result[0] != td.lemma {
		t.Errorf("For %v, expected '%v', but got '%v'", td.word, td.lemma, result[0])
	}
}

func TestEstoy(t *testing.T) {
	runTest(t, testData{
		word:  "estoy",
		lemma: "estar",
		token: "# {{es-verb form of|mood=ind|tense=pres|num=s|pers=1|ending=ar|estar|nodot=1}}; [[am]]",
	})
}

func TestSe(t *testing.T) {
	runTest(t, testData{
		word:  "sé",
		lemma: "ser",
		token: "# {{es-verb form of|mood=imp|num=s|pers=2|formal=n|sense=+|ending=er|ser}}",
	})
}

func TestEs(t *testing.T) {
	runTest(t, testData{
		word:  "es",
		lemma: "ser",
		token: "# {{es-verb form of|mood=ind|tense=pres|num=s|pers=2|formal=y|ending=er|ser}}",
	})
}

func TestIr(t *testing.T) {
	runTest(t, testData{
		word:  "voy",
		lemma: "ir",
		token: "# {{es-verb form of|mood=ind|tense=pres|num=s|pers=1|ending=ir|ir}}",
	})
}

func TestPuedo(t *testing.T) {
	runTest(t, testData{
		word:  "puedo",
		lemma: "poder",
		token: "# {{es-verb form of|ending=er|mood=indicative|tense=present|pers=1|number=singular|poder}",
	})
}
