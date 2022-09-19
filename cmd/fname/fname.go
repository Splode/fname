package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/splode/fname"
)

func main() {
	var (
		delimiter string = ""
		number    int    = 1
		seed      int64  = -1
	)

	pflag.StringVarP(&delimiter, "delimiter", "d", delimiter, "delimiter to use between words")
	pflag.IntVarP(&number, "number", "n", number, "number of names to generate")
	pflag.Int64VarP(&seed, "seed", "s", seed, "random generator seed")
	pflag.Parse()

	opts := []fname.RandomNameGeneratorOption{}
	if delimiter != "" {
		opts = append(opts, fname.WithDelimiter(delimiter))
	}
	if seed != -1 {
		opts = append(opts, fname.WithSeed(seed))
	}

	rng := fname.NewRandomNameGenerator(opts...)

	for i := 0; i < number; i++ {
		fmt.Println(rng.Generate())
	}
}
