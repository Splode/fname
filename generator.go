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

// ParseCasing parses a casing string and returns the corresponding Casing value.
func ParseCasing(casing string) (Casing, error) {
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

// Deprecated: Use ParseCasing instead.
func CasingFromString(casing string) (Casing, error) {
	return ParseCasing(casing)
}

// Generator generates random name phrases. It is not safe for concurrent use
// from multiple goroutines; create a separate Generator per goroutine instead.
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

// WithDictionary sets a custom Dictionary on the Generator.
// If d is nil, the default embedded Dictionary is used.
func WithDictionary(d *Dictionary) GeneratorOption {
	return func(g *Generator) {
		if d != nil {
			g.dict = d
		}
	}
}

// WithSeed sets the seed used to generate random numbers.
func WithSeed(seed int64) GeneratorOption {
	return func(g *Generator) {
		g.rand = rand.New(rand.NewSource(seed))
	}
}

// WithSize sets the number of words in the generated name.
// Returns an error if size is outside the valid range [2, 4].
func WithSize(size uint) (GeneratorOption, error) {
	if size < 2 || size > 4 {
		return nil, fmt.Errorf("invalid size: %d", size)
	}
	return func(g *Generator) {
		g.size = size
	}, nil
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
	words := make([]string, 0, g.size)
	adjectiveIndex := g.rand.Intn(g.dict.LengthAdjective())
	nounIndex := g.rand.Intn(g.dict.LengthNoun())

	words = append(words, g.dict.adjectives[adjectiveIndex], g.dict.nouns[nounIndex])

	if g.size >= 3 {
		words = append(words, g.dict.verbs[g.rand.Intn(g.dict.LengthVerb())])
	}

	if g.size == 4 {
		words = append(words, g.dict.adverbs[g.rand.Intn(g.dict.LengthAdverb())])
	}

	return strings.Join(g.applyCasing(words), g.delimiter), nil
}

var titleCaser = cases.Title(language.English)

func (g *Generator) applyCasing(words []string) []string {
	for i, word := range words {
		switch g.casing {
		case Lower:
			words[i] = strings.ToLower(word)
		case Upper:
			words[i] = strings.ToUpper(word)
		case Title:
			words[i] = titleCaser.String(word)
		}
	}
	return words
}
