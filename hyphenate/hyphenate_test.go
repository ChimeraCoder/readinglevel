package hyphenate

import (
	"testing"
)

func Test_CMUDict(t *testing.T) {
	err := loadCMUCorpus()
	if err != nil {
		t.Error(err)
	}

	const expectedLength = 133265
	if cmuCorpusCached.Words() != expectedLength {
		t.Errorf("Expected %d words in corpus and found %d", expectedLength, cmuCorpusCached.Words())
	}
}

func Test_CMUSyllables(t *testing.T) {
	err := loadCMUCorpus()
	if err != nil {
		t.Error(err)
	}

	TestWords := map[string]int{
		"Zombie":     2,
		"antibiotic": 5,
		"RECORDED":   3,
		"recaptured": 3,
		"blame":      1,
		"discussion": 3,
		"whims":      1,
		"tilting":    2,
		"gong":       1,
	}

	for word, expected := range TestWords {
		if syllables := cmuCorpusCached.Syllables(word); syllables != expected {
			t.Errorf("expected %d syllables for \"%s\" and received %d", expected, word, syllables)
		}
	}

}
