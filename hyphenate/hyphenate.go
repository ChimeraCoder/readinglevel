package hyphenate

import (
	"io/ioutil"
	"regexp"
	"strings"
	"unicode"
)

// A Corpus is a body of text that supports certain features
// Currently the only required queries are Syllables (number of syllables)
// and Words (number of words)
type Corpus interface {
	Syllables(string) int
	Words() int
}

type cmuCorpus map[string][]string

var cmuCorpusCached cmuCorpus = map[string][]string{}

// CMUCorpus returns the CMU corpus
func CMUCorpus() (Corpus, error) {
	bts, err := ioutil.ReadFile("cmudict.corpus")
	if err != nil {
		return nil, err
	}

	cmu := string(bts)

	re := regexp.MustCompile(`^[A-Z]`)
	var tmpCorpus cmuCorpus = map[string][]string{}
	for _, line := range strings.Split(cmu, "\n") {
		line = strings.TrimSpace(line)

		//TODO account for the lines that represent alternate pronunciations

		// Ignore the lines that don't start with a character A-Z
		if len(line) == 0 || !re.MatchString(line[:1]) {
			continue
		}

		linesplit := strings.Split(line, " ")
		word := linesplit[0]
		tmpCorpus[strings.ToUpper(word)] = linesplit[1:]
	}

	return tmpCorpus, nil
}

// Syllables returns the number of syllables for the word, according to the corpus
// If the word is not in the corpus, it will return 0
func (c cmuCorpus) Syllables(word string) int {
	phonemes, ok := c[strings.ToUpper(word)]
	if !ok {
		return 0
	}

	count := 0
	for _, phoneme := range phonemes {
		for _, r := range phoneme {
			if unicode.IsNumber(r) {
				count++
			}
		}
	}
	return count
}

func (c cmuCorpus) Words() int {
	return len(c)
}
