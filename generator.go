// Package fname contains functions for generating random, human-friendly names.
package fname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Generator is a random name generator.
type Generator struct {
	dict      *Dictionary
	delimiter string
	rand      *rand.Rand
	size      uint
}

// GeneratorOption is a function that configures a Generator.
type GeneratorOption func(*Generator)

// WithDelimiter sets the delimiter used to join words.
func WithDelimiter(delimiter string) GeneratorOption {
	return func(r *Generator) {
		r.delimiter = delimiter
	}
}

// WithSeed sets the seed used to generate random numbers.
func WithSeed(seed int64) GeneratorOption {
	return func(r *Generator) {
		r.rand.Seed(seed)
	}
}

// WithSize sets the number of words in the generated name.
func WithSize(size uint) GeneratorOption {
	return func(r *Generator) {
		r.size = size
	}
}

// NewGenerator creates a new Generator.
func NewGenerator(opts ...GeneratorOption) *Generator {
	r := &Generator{
		dict:      NewDictionary(),
		delimiter: "-",
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		size:      2,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Generate generates a random name.
func (r *Generator) Generate() (string, error) {
	// TODO: address case where adjective and noun are the same, such as "orange-orange" or "sound-sound"
	adjective := r.dict.adectives[r.rand.Intn(r.dict.LengthAdjective())]
	noun := r.dict.nouns[r.rand.Intn(r.dict.LengthNoun())]
	words := []string{adjective, noun}

	switch r.size {
	case 2:
		return strings.Join(words, r.delimiter), nil
	case 3:
		verb := r.dict.verbs[r.rand.Intn(r.dict.LengthVerb())]
		words = append(words, verb)
	case 4:
		verb := r.dict.verbs[r.rand.Intn(r.dict.LengthVerb())]
		words = append(words, verb)
		adverb := r.dict.adverbs[r.rand.Intn(r.dict.LengthAdverb())]
		words = append(words, adverb)
	default:
		return "", fmt.Errorf("invalid size: %d", r.size)
	}
	return strings.Join(words, r.delimiter), nil
}
