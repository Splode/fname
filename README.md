![fname banner](.github/media/fname_banner_1.1_1280x640.png)

# fname

Generate random, human-friendly names, like `determined-pancake` or `sinister discovery`. fname is like a grammatically aware diceware generator for unique names or identifiers.

fname isn't meant to provide a secure, globally unique identifier, but with over 500 billion possible combinations, it's good enough for most non-critical use cases.

## Table of Contents

- [fname](#fname)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
    - [Download](#download)
    - [Go](#go)
    - [Source](#source)
  - [Usage](#usage)
    - [CLI](#cli)
    - [Library](#library)
      - [Install](#install)
      - [Basic Usage](#basic-usage)
      - [Customization](#customization)
  - [Disclaimers](#disclaimers)
  - [Contributing](#contributing)
    - [Reporting Issues](#reporting-issues)
    - [Suggesting Improvements](#suggesting-improvements)
  - [License](#license)
  - [Related Projects](#related-projects)

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

Generate a name phrase with a specific casing:

```sh
$ fname --casing upper
TRAGIC-MOUNTAIN

$ fname --casing title
Whimsical-Party
```

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

## Contributing

We welcome contributions to the fname project! Whether it's reporting bugs, suggesting improvements, or submitting new features, your input is valuable to us. Here's how you can get started:

1. Fork the repository on GitHub.
2. Clone your fork and create a new branch for your changes.
3. Make your changes and commit them to your branch.
4. Create a pull request, and provide a clear description of your changes.

Before submitting a pull request, please make sure your changes are well-tested and adhere to the code style used throughout the project. If you are unsure how to proceed or need help, feel free to open an issue or ask a question in the [discussions](https://github.com/Splode/fname/discussions) section.

### Reporting Issues

If you encounter a bug or any issue, please [open an issue](https://github.com/Splode/fname/issues) on GitHub. When reporting a bug, try to include as much information as possible, such as the steps to reproduce the issue, the expected behavior, and the actual behavior. This will help us diagnose and fix the issue more efficiently.

### Suggesting Improvements

We are always looking for ways to improve fname. If you have a suggestion for a new feature or an enhancement to an existing feature, please [open an issue](https://github.com/Splode/fname/issues) or start a discussion in the [discussions](https://github.com/Splode/fname/discussions) section. Be sure to explain your idea in detail, and if possible, provide examples or use cases.

Thank you for your interest in contributing to fname!



## License

[MIT License](./LICENSE)

## Related Projects

- [go-diceware](https://github.com/sethvargo/go-diceware)
- [wordnet-random-name](https://github.com/kohsuke/wordnet-random-name)
