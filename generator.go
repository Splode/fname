package fname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Casing string

const (
	Lower Casing = "lower"
	Upper Casing = "upper"
	Title Casing = "title"
)

// Generator is a random name generator.
type Generator struct {
	casing    Casing
	dict      *Dictionary
	delimiter string
	rand      *rand.Rand
	size      uint
}

// GeneratorOption is a function that configures a Generator.
type GeneratorOption func(*Generator)

// WithCasing sets the casing used to format the generated name.
func WithCasing(casing Casing) GeneratorOption {
	return func(g *Generator) {
		g.casing = casing
	}
}

// WithDelimiter sets the delimiter used to join words.
func WithDelimiter(delimiter string) GeneratorOption {
	return func(g *Generator) {
		g.delimiter = delimiter
	}
}

// WithSeed sets the seed used to generate random numbers.
func WithSeed(seed int64) GeneratorOption {
	return func(g *Generator) {
		g.rand.Seed(seed)
	}
}

// WithSize sets the number of words in the generated name.
func WithSize(size uint) GeneratorOption {
	return func(g *Generator) {
		g.size = size
	}
}

// NewGenerator creates a new Generator.
func NewGenerator(opts ...GeneratorOption) *Generator {
	g := &Generator{
		casing:    Lower,
		dict:      NewDictionary(),
		delimiter: "-",
		rand:      rand.New(rand.NewSource(time.Now().UnixNano())),
		size:      2,
	}
	for _, opt := range opts {
		opt(g)
	}
	return g
}

// Generate generates a random name.
func (g *Generator) Generate() (string, error) {
	// Keep generating adjective and noun pairs until they are not the same.
	var adjective, noun string
	for adjective == noun {
		adjective = g.dict.adjectives[g.rand.Intn(g.dict.LengthAdjective())]
		noun = g.dict.nouns[g.rand.Intn(g.dict.LengthNoun())]
	}

	words := []string{adjective, noun}

	switch g.size {
	case 2:
		// do nothing
	case 3:
		verb := g.dict.verbs[g.rand.Intn(g.dict.LengthVerb())]
		words = append(words, verb)
	case 4:
		verb := g.dict.verbs[g.rand.Intn(g.dict.LengthVerb())]
		words = append(words, verb)
		adverb := g.dict.adverbs[g.rand.Intn(g.dict.LengthAdverb())]
		words = append(words, adverb)
	default:
		return "", fmt.Errorf("invalid size: %d", g.size)
	}
	return strings.Join(g.applyCasing(words...), g.delimiter), nil
}

// ParseCasing parses a string into a casing.
func ParseCasing(casing string) (Casing, error) {
	switch casing {
	case "lower":
		return Lower, nil
	case "upper":
		return Upper, nil
	case "title":
		return Title, nil
	default:
		return "", fmt.Errorf("invalid casing: %s", casing)
	}
}

var titleCaser = cases.Title(language.English)

var casingMap = map[Casing]func(string) string{
	Lower: strings.ToLower,
	Upper: strings.ToUpper,
	Title: titleCaser.String,
}

func (g *Generator) applyCasing(words ...string) []string {
	if fn, ok := casingMap[g.casing]; ok {
		for i, word := range words {
			words[i] = fn(word)
		}
	}
	return words
}
