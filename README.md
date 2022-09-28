![fname banner](.github/media/fname_banner_1.1_1280x640.png)

# fname

Generate random, human-friendly names, like `determined-pancake` or `sinister discovery`. fname is like a grammatically aware diceware generator for unique names or identifiers.

fname isn't meant to provide a secure, globally unique identifier, but with over 500 billion possible combinations, it's good enough for most non-critical use cases.

## Installation

### Download

Download the latest release from the [releases page](https://github.com/Splode/fname/releases).

### Go

```sh
go install github.com/splode/fname/cmd/fname@latest
```

### Source

```sh
git clone https://github.com/splode/fname.git
cd fname
go install ./cmd/fname
```

## Usage

### CLI
  
Generate a single, random name phrase:

```sh
$ fname
extinct-green
```

Generate multiple name phrases, passing the number of names as an argument:

```sh
$ fname --quantity 3
influential-length
direct-ear
cultural-storage
```

Generate a name phrase with a custom delimiter:

```sh
$ fname --delimiter "__"
glaring__perception
```

Generate a name phrase with more words:

```sh
$ fname --size 3
vengeful-toy-identified

$ fname --size 4
spellbinding-project-presented-fully
```

Note: the minimum phrase size is 2 (default), and the maximum is 4.

Specify the seed for generating names:

```sh
$ fname --seed 123 --quantity 2
pleasant-joy
eligible-tenant

$ fname --seed 123 --quantity 2
pleasant-joy
eligible-tenant
```

### Library

#### Install

```sh
go get github.com/splode/fname
```

#### Basic Usage

```go
package main

import (
  "fmt"

  "github.com/splode/fname"
)

func main() {
  rng := fname.NewGenerator()
  phrase, err := rng.Generate()
  fmt.Println(phrase)
  // => "influential-length"
}
```

#### Customization

```go
package main

import (
  "fmt"

  "github.com/splode/fname"
)

func main() {
  rng := fname.NewGenerator(fname.WithDelimiter("__"), fname.WithSize(3))
  phrase, err := rng.Generate()
  fmt.Println(phrase)
  // => "established__shark__destroyed"
}
```

## Disclaimers

fname is not cryptographically secure, and should not be used for anything that requires a truly unique identifier. It is meant to be a fun, human-friendly alternative to UUIDs.

fname's dictionary is curated to exclude words that are offensive, or could be considered offensive, either alone or when generated in a phrase. Nevertheless, all cases are not and cannot be covered. If you find a word that you think should be removed, please [open an issue](https://github.com/Splode/fname/issues).


## License

[MIT License](./LICENSE)

## Related Projects

- [go-diceware](https://github.com/sethvargo/go-diceware)
- [wordnet-random-name](https://github.com/kohsuke/wordnet-random-name)
