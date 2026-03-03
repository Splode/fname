// Package fname contains functions for generating random, human-friendly names.
package fname

import (
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

// NewDictionary creates a new Dictionary backed by the default embedded word lists.
// To use custom word lists, use NewCustomDictionary and pass it via WithDictionary.
func NewDictionary() *Dictionary {
	return &Dictionary{
		adjectives: adjective,
		adverbs:    adverb,
		nouns:      noun,
		verbs:      verb,
	}
}

// NewCustomDictionary creates a Dictionary with caller-supplied word lists.
// Any nil slice falls back to the corresponding default embedded word list.
func NewCustomDictionary(adjectives, adverbs, nouns, verbs []string) *Dictionary {
	d := NewDictionary()
	if adjectives != nil {
		d.adjectives = adjectives
	}
	if adverbs != nil {
		d.adverbs = adverbs
	}
	if nouns != nil {
		d.nouns = nouns
	}
	if verbs != nil {
		d.verbs = verbs
	}
	return d
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
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
}
