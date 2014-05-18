package readinglevel

import (
	"github.com/ChimeraCoder/textcorpora/cmu"
	"regexp"
	"strings"
)

var whitespaceRe = regexp.MustCompile(`\W+`)

func lettersPerHWord(corpus string) float64 {

	words := NumWords(corpus)
	letters := len(corpus) - (words - 1) // Whitespace doesn't count as letters

	return 100 * (float64(letters) / float64(words))
}

func NumWords(corpus string) int {
	corpus = whitespaceRe.ReplaceAllLiteralString(corpus, " ")
	corpus = strings.TrimSpace(corpus)
	return len(strings.Split(corpus, " "))
}

func NumSentences(corpus string) int {
	//TODO replace with a more sophisticated means of separation
	re := regexp.MustCompile(`\s+`)
	corpus = re.ReplaceAllLiteralString(corpus, " ")
	corpus = strings.TrimSpace(corpus)
	return len(strings.Split(corpus, ".")) - 1 // Subtract one, assuming the last character is a period
}

func ColemanLiau(corpus string) float64 {
	// CLI = 0.0588{L} - 0.296{S} - 15.8\,\!
	// https://en.wikipedia.org/wiki/Coleman%E2%80%93Liau_index
	L := lettersPerHWord(corpus)
	S := float64(NumSentences(corpus)) / float64(100*NumWords(corpus))

	return 0.0588*L - 0.296*S - 15.8
}

func NumSyllables(text string) (int, error) {
	text = whitespaceRe.ReplaceAllLiteralString(text, " ")
	text = strings.TrimSpace(text)
	words := strings.Split(text, " ")

	var syllables int
	var notFound int

	corpus, err := cmu.CMUCorpus()
	if err != nil {
		return 0, err
	}
	for _, word := range words {
		// TODO check for words in which the lookup failed
		s := corpus.Syllables(word)
		syllables += s

		if s == 0 {
			notFound++
		}
	}
	return syllables, nil
}

// FleschKincaidGrade returns the grade level of the given body of text
// according to the Flesch-Kincaid grade level test
// It underestimates the grade level slightly, as unknown words
// are treated as having 0 syllables
func FleschKincaidGrade(corpus string) (float64, error) {
	words := NumWords(corpus)
	sentences := NumSentences(corpus)
	syllables, err := NumSyllables(corpus)
	if err != nil {
		return 0, err
	}
	wordsPerSentence := float64(words) / float64(sentences)
	syllablesPerWord := float64(syllables) / float64(words)
	return .39*wordsPerSentence + 11.8*syllablesPerWord - 15.59, nil
}

// FleschKincaidEase returns the reading ease score of the given body of text
// according to the Flesch-Kincaid reading ease test
// It overestimates the reading ease score slightly, as unknown words
// are treated as having 0 syllables
func FleschKincaidEase(corpus string) (float64, error) {
	words := NumWords(corpus)
	sentences := NumSentences(corpus)
	syllables, err := NumSyllables(corpus)
	if err != nil {
		return 0, err
	}
	wordsPerSentence := float64(words) / float64(sentences)
	syllablesPerWord := float64(syllables) / float64(words)
	return 206.835 - 1.015*(wordsPerSentence) - 84.6*syllablesPerWord, nil
}
