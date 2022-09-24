// package fname contains functions for generating random, human-friendly names.
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

// Dictionary is a collection of words.
type Dictionary struct {
	adectives []string
	adverbs   []string
	nouns     []string
	verbs     []string
}

// NewDictionary creates a new dictionary.
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

// Adjective returns a random adjective.
func (d *Dictionary) Adjective(idx int) (string, error) {
	if idx < 0 || idx >= len(d.adectives) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.adectives[idx], nil
}

// Adverb returns a random adverb.
func (d *Dictionary) Adverb(idx int) (string, error) {
	if idx < 0 || idx >= len(d.adverbs) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.adverbs[idx], nil
}

// Noun returns a random noun.
func (d *Dictionary) Noun(idx int) (string, error) {
	if idx < 0 || idx >= len(d.nouns) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.nouns[idx], nil
}

// Verb returns a random verb.
func (d *Dictionary) Verb(idx int) (string, error) {
	if idx < 0 || idx >= len(d.verbs) {
		return "", fmt.Errorf("index out of range: %d", idx)
	}
	return d.verbs[idx], nil
}

// LengthAdjective returns the number of adjectives in the dictionary.
func (d *Dictionary) LengthAdjective() int {
	return len(d.adectives)
}

// LengthAdverb returns the number of adverbs in the dictionary.
func (d *Dictionary) LengthAdverb() int {
	return len(d.adverbs)
}

// LengthNoun returns the number of nouns in the dictionary.
func (d *Dictionary) LengthNoun() int {
	return len(d.nouns)
}

// LengthVerb returns the number of verbs in the dictionary.
func (d *Dictionary) LengthVerb() int {
	return len(d.verbs)
}

// loadFile loads a file from the embedded filesystem, and returns a slice of strings containing each line.
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
