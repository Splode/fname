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
  fname --delimiter "." --number 10

Options:`

	contact = `
Author: Christopher Murphy <flyweight@pm.me>
Repo: https://github.com/splode/fname`
)

var (
	version = ""
)

//go:embed banner
var banner []byte

func main() {
	pflag.Usage = generateUsage

	var (
		delimiter string = ""
		help      bool   = false
		number    int    = 1
		seed      int64  = -1
		ver       bool   = false
		// TODO: add option to use custom dictionary
	)

	pflag.StringVarP(&delimiter, "delimiter", "d", delimiter, "delimiter to use between words")
	pflag.BoolVarP(&help, "help", "h", help, "show fname usage")
	pflag.IntVarP(&number, "number", "n", number, "number of names to generate")
	pflag.Int64VarP(&seed, "seed", "s", seed, "random generator seed")
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
