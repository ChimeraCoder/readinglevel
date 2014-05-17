package readinglevel

import (
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
