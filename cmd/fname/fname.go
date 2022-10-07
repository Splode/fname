package main

import (
	_ "embed"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/pflag"
	"github.com/splode/fname"
)

const (
	usage = `fname generates random, human-friendly names, such as
"awful-fossil" or "constant process".

Usage: 
  fname [options]

Examples:
  # generate a single name phrase using default options
  fname

  # generate 10 names using a custom delimiter
  fname --delimiter "." --quantity 10

Options:`

	contact = `
Author: Christopher Murphy <flyweight@pm.me>
Source: https://github.com/splode/fname`
)

var (
	version = ""
)

//go:embed banner
var banner []byte

func main() {
	pflag.Usage = generateUsage

	var (
		casing    string = "lower"
		delimiter string
		help      bool
		ver       bool
		quantity  int   = 1
		size      uint  = 2
		seed      int64 = -1
		// TODO: add option to use custom dictionary
	)

	pflag.StringVarP(&casing, "casing", "c", casing, "case of generated names: lower, upper, or title")
	pflag.StringVarP(&delimiter, "delimiter", "d", delimiter, "delimiter to use between words")
	pflag.IntVarP(&quantity, "quantity", "q", quantity, "number of name phrases to generate")
	pflag.UintVarP(&size, "size", "z", size, "number of words per phrase (minimum 2, maximum 4)")
	pflag.Int64VarP(&seed, "seed", "s", seed, "random generator seed")
	pflag.BoolVarP(&help, "help", "h", help, "show fname usage")
	pflag.BoolVarP(&ver, "version", "v", ver, "show fname version")
	pflag.Parse()

	if help {
		pflag.Usage()
		os.Exit(0)
	}

	if ver {
		fmt.Println(getVersion())
		os.Exit(0)
	}

	opts := []fname.GeneratorOption{}

	c, err := fname.ParseCasing(casing)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	opts = append(opts, fname.WithCasing(c))

	if delimiter != "" {
		opts = append(opts, fname.WithDelimiter(delimiter))
	}
	if seed != -1 {
		opts = append(opts, fname.WithSeed(seed))
	}
	if size != 2 {
		opts = append(opts, fname.WithSize(size))
	}

	rng := fname.NewGenerator(opts...)

	for i := 0; i < quantity; i++ {
		name, err := rng.Generate()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v", err)
			os.Exit(1)
		}
		fmt.Println(name)
	}
}

func generateUsage() {
	fmt.Println(string(banner))
	fmt.Println(usage)
	pflag.PrintDefaults()
	fmt.Println(contact)
}

func getVersion() string {
	if version != "" {
		return version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "unknown"
	}

	version = info.Main.Version
	if info.Main.Sum != "" {
		version += fmt.Sprintf(" (%s)", info.Main.Sum)
	}

	return version
}
