package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/splode/fname"
)

const (
	usage = `fname generates random, human-friendly names, such as "awful-fossil"
or "zesty zebra".

Usage: 
  fname [options]

Examples:
  # generate a single name phrase using default options
  fname

  # generate 10 names using a custom delimiter
  fname --delimiter "." --number 10

Options:`

	contact = `
Author: Christopher Murphy <flyweight@pm.me>
Repo: https://github.com/splode/fname`
)

//go:embed banner
var banner []byte

func main() {
	pflag.Usage = func() {
		fmt.Println(string(banner))
		fmt.Println(usage)
		pflag.PrintDefaults()
		fmt.Println(contact)
	}

	var (
		delimiter string = ""
		help      bool   = false
		number    int    = 1
		seed      int64  = -1
	)

	pflag.StringVarP(&delimiter, "delimiter", "d", delimiter, "delimiter to use between words")
	pflag.BoolVarP(&help, "help", "h", help, "Show fname usage")
	pflag.IntVarP(&number, "number", "n", number, "number of names to generate")
	pflag.Int64VarP(&seed, "seed", "s", seed, "random generator seed")
	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

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
