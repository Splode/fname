![fname banner](.github/media/fname_banner_1.1_1280x640.png)

# fname

Generate random, human-friendly names, like `determined-pancake` or `sinister discovery`. `fname` is like a grammatically aware diceware generator for unique names or identifiers.

`fname` isn't meant to provide a secure, globally unique identifier, but with over 500 billion possible combinations, it's good enough for most non-critical use cases.

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
  
Generate a single, random name:

```sh
$ fname
extinct-green
```

Generate multiple names, passing the number of names as an argument:

```sh
$ fname --quantity 3
influential-length
direct-ear
cultural-storage
```

Generate a random name phrase with a custom delimiter:

```sh
$ fname --delimiter "__"
glaring__perception
```

Generate a random name phrase with a custom delimiter and quantity:

```sh
$ fname --size 3
vengeful-toy-identified

$ fname --size 4
ambiguous-anticipation-ignored-keenly
```

Note: the minimum phrase size is 2 (default), and the maximum is 4.

Specify the seed for generating names:

```sh
$ fname --seed 123
foundational-spot

$ fname --seed 123
foundational-spot
```

### Library

#### Install

```sh
go get github.com/splode/fname
```

#### Example

```go
package main

import (
  "fmt"

  "github.com/splode/fname"
)

func main() {
  rng := fname.NewRandomNameGenerator()
  phrase, err := rng.Generate()
  fmt.Println(phrase)
  // => "influential-length"
}
```

## License

[MIT License](./LICENSE)

## Related Projects

- [go-diceware](https://github.com/sethvargo/go-diceware)
- [wordnet-random-name](https://github.com/kohsuke/wordnet-random-name)
