package fname

import (
	"math/rand"
	"time"
)

type RandomNameGenerator struct {
	dict      *Dictionary
	delimiter string
	seed      int64
}

type RandomNameGeneratorOption func(*RandomNameGenerator)

func WithDelimiter(delimiter string) RandomNameGeneratorOption {
	return func(r *RandomNameGenerator) {
		r.delimiter = delimiter
	}
}

func WithSeed(seed int64) RandomNameGeneratorOption {
	return func(r *RandomNameGenerator) {
		r.seed = seed
	}
}

func NewRandomNameGenerator(opts ...RandomNameGeneratorOption) *RandomNameGenerator {
	r := &RandomNameGenerator{
		dict:      NewDictionary(),
		delimiter: "-",
		seed:      time.Now().UnixNano(),
	}
	for _, opt := range opts {
		opt(r)
	}
	rand.Seed(r.seed)
	return r
}

func (r *RandomNameGenerator) Generate() string {
	adjective, err := r.dict.Adjective(rand.Intn(r.dict.LengthAdjective()))
	if err != nil {
		panic(err)
	}
	noun, err := r.dict.Noun(rand.Intn(r.dict.LengthNoun()))
	if err != nil {
		panic(err)
	}
	return adjective + r.delimiter + noun
}
