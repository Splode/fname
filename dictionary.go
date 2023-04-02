// Package fname contains functions for generating random, human-friendly names.
package fname

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed data/adjective
var _adjective string
var adjective = split(_adjective)

//go:embed data/adverb
var _adverb string
var adverb = split(_adverb)

//go:embed data/noun
var _noun string
var noun = split(_noun)

//go:embed data/verb
var _verb string
var verb = split(_verb)

// Dictionary is a collection of words.
type Dictionary struct {
	adjectives []string
	adverbs    []string
	nouns      []string
	verbs      []string
}

// NewDictionary creates a new dictionary.
func NewDictionary() *Dictionary {
	// TODO: allow for custom dictionary
	return &Dictionary{
		adjectives: adjective,
		adverbs:    adverb,
		nouns:      noun,
		verbs:      verb,
	}
}

// LengthAdjective returns the number of adjectives in the dictionary.
func (d *Dictionary) LengthAdjective() int {
	return len(d.adjectives)
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

func split(s string) []string {
	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
