package fname

import (
	"bufio"
	"embed"
	"fmt"
)

const (
	adjectiveFilePath = "data/adjective"
	adverbFilePath    = "data/adverb"
	nounFilePath      = "data/noun"
	verbFilePath      = "data/verb"
)

//go:embed data/*
var dataFS embed.FS

type Dictionary struct {
	adectives []string
	adverbs   []string
	nouns     []string
	verbs     []string
}

func NewDictionary() *Dictionary {
	// TODO: allow for custom dictionary
	a, err := loadFile(adjectiveFilePath)
	if err != nil {
		panic(err)
	}
	av, err := loadFile(adverbFilePath)
	if err != nil {
		panic(err)
	}
	n, err := loadFile(nounFilePath)
	if err != nil {
		panic(err)
	}
	v, err := loadFile(verbFilePath)
	if err != nil {
		panic(err)
	}

	return &Dictionary{
		adectives: a,
		adverbs:   av,
		nouns:     n,
		verbs:     v,
	}
}

func (d *Dictionary) Adjective(idx int) (string, error) {
	if idx < 0 || idx >= len(d.adectives) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.adectives[idx], nil
}

func (d *Dictionary) Adverb(idx int) (string, error) {
	if idx < 0 || idx >= len(d.adverbs) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.adverbs[idx], nil
}

func (d *Dictionary) Noun(idx int) (string, error) {
	if idx < 0 || idx >= len(d.nouns) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.nouns[idx], nil
}

func (d *Dictionary) Verb(idx int) (string, error) {
	if idx < 0 || idx >= len(d.verbs) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.verbs[idx], nil
}

func (d *Dictionary) LengthAdjective() int {
	return len(d.adectives)
}

func (d *Dictionary) LengthAdverb() int {
	return len(d.adverbs)
}

func (d *Dictionary) LengthNoun() int {
	return len(d.nouns)
}

func (d *Dictionary) LengthVerb() int {
	return len(d.verbs)
}

func loadFile(path string) ([]string, error) {
	f, err := dataFS.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var words []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w := scanner.Text()
		if w != "" {
			words = append(words, scanner.Text())
		}
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return words, nil
}
