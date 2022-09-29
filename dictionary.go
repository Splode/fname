// Package fname contains functions for generating random, human-friendly names.
package fname

import (
	_ "embed"
	"strings"
)

//go:embed data/adjective
var _adjective string
var adjective = strings.Split(_adjective, "\n")

//go:embed data/adverb
var _adverb string
var adverb = strings.Split(_adverb, "\n")

//go:embed data/noun
var _noun string
var noun = strings.Split(_noun, "\n")

//go:embed data/verb
var _verb string
var verb = strings.Split(_verb, "\n")

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
	return &Dictionary{
		adectives: adjective,
		adverbs:   adverb,
		nouns:     noun,
		verbs:     verb,
	}
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
