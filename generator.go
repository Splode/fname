package fname

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Casing int

const (
	Lower Casing = iota
	Upper
	Title
)

func (c Casing) String() string {
	switch c {
	case Lower:
		return "lower"
	case Upper:
		return "upper"
	case Title:
		return "title"
	default:
		return "unknown"
	}
}

func CasingFromString(casing string) (Casing, error) {
	switch strings.ToLower(casing) {
	case Lower.String():
		return Lower, nil
	case Upper.String():
		return Upper, nil
	case Title.String():
		return Title, nil
	default:
		return -1, fmt.Errorf("invalid casing: %s", casing)
	}
}

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
		g.rand = rand.New(rand.NewSource(seed))
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
	if g.size < 2 || g.size > 4 {
		return "", fmt.Errorf("invalid size: %d", g.size)
	}

	words := make([]string, 0, g.size)
	adjectiveIndex := g.rand.Intn(g.dict.LengthAdjective())
	nounIndex := g.rand.Intn(g.dict.LengthNoun())
	for adjectiveIndex == nounIndex {
		nounIndex = g.rand.Intn(g.dict.LengthNoun())
	}

	words = append(words, g.dict.adjectives[adjectiveIndex], g.dict.nouns[nounIndex])

	if g.size >= 3 {
		words = append(words, g.dict.verbs[g.rand.Intn(g.dict.LengthVerb())])
	}

	if g.size == 4 {
		words = append(words, g.dict.adverbs[g.rand.Intn(g.dict.LengthAdverb())])
	}

	return strings.Join(g.applyCasing(words), g.delimiter), nil
}

func (g *Generator) applyCasing(words []string) []string {
	if fn, ok := casingMap[g.casing]; ok {
		for i, word := range words {
			words[i] = fn(word)
		}
	}
	return words
}

var titleCaser = cases.Title(language.English)

var casingMap = map[Casing]func(string) string{
	Lower: strings.ToLower,
	Upper: strings.ToUpper,
	Title: titleCaser.String,
}
