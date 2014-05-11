package hyphenate

import (
    "strings"
    "io/ioutil"
    "regexp"
    "unicode"
)


type Corpus map[string][]string

var cmuCorpus = Corpus(map[string][]string{})

func loadCMUCorpus() error{
    bts, err := ioutil.ReadFile("cmudict.corpus")
    if err != nil{
        return err
    }

    cmu := string(bts)

    re := regexp.MustCompile(`^[A-Z]`)
    for _, line := range strings.Split(cmu, "\n"){
        line = strings.TrimSpace(line)

        //TODO account for the lines that represent alternate pronunciations

        // Ignore the lines that don't start with a character A-Z
        if len(line) == 0 || !re.MatchString(line[:1]){
            continue
        }

        linesplit := strings.Split(line, " ")
        word := linesplit[0]
        cmuCorpus[strings.ToUpper(word)] = linesplit[1:]
    }

    return nil
}


// Syllables returns the number of syllables for the word, according to the corpus
// If the word is not in the corpus, it will return 0
func (c Corpus) Syllables(word string) int {
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
