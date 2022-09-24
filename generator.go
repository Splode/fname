package fname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Generator struct {
	dict      *Dictionary
	delimiter string
	seed      int64
	size      uint
}

type GeneratorOption func(*Generator)

func WithDelimiter(delimiter string) GeneratorOption {
	return func(r *Generator) {
		r.delimiter = delimiter
	}
}

func WithSeed(seed int64) GeneratorOption {
	return func(r *Generator) {
		r.seed = seed
	}
}

func WithSize(size uint) GeneratorOption {
	return func(r *Generator) {
		r.size = size
	}
}

func NewGenerator(opts ...GeneratorOption) *Generator {
	r := &Generator{
		dict:      NewDictionary(),
		delimiter: "-",
		seed:      time.Now().UnixNano(),
		size:      2,
	}
	for _, opt := range opts {
		opt(r)
	}
	rand.Seed(r.seed)
	return r
}

func (r *Generator) Generate() (string, error) {
	// TODO: address case where adjective and noun are the same, such as "orange-orange" or "sound-sound"
	adjective, err := r.dict.Adjective(rand.Intn(r.dict.LengthAdjective()))
	if err != nil {
		return "", err
	}
	noun, err := r.dict.Noun(rand.Intn(r.dict.LengthNoun()))
	if err != nil {
		return "", err
	}
	words := []string{adjective, noun}

	switch r.size {
	case 2:
		return strings.Join(words, r.delimiter), nil
	case 3:
		verb, err := r.dict.Verb(rand.Intn(r.dict.LengthVerb()))
		if err != nil {
			return "", err
		}
		words = append(words, verb)
	case 4:
		verb, err := r.dict.Verb(rand.Intn(r.dict.LengthVerb()))
		if err != nil {
			return "", err
		}
		words = append(words, verb)
		adverb, err := r.dict.Adverb(rand.Intn(r.dict.LengthAdverb()))
		if err != nil {
			return "", err
		}
		words = append(words, adverb)
	default:
		return "", fmt.Errorf("invalid size: %d", r.size)
	}
	return strings.Join(words, r.delimiter), nil
}
