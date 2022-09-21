# fname

Command-line utility that generates random, human-friendly name, like `determined-pancake` or `silly-zebra`. `fname` is like a grammatically correct diceware generator for unique names or identifiers.

`fname` isn't meant to provide a secure, unique identifier, but with over 3.4 million possible combinations, it's good enough for most non-critical use cases.

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
  
Generate a single, random name:

```sh
$ fname
extinct-green
```

Generate multiple names, passing the number of names as an argument:

```sh
$ fname --number 3
influential-length
direct-ear
cultural-storage
```

Generate a random name phrase with a custom delimiter:

```sh
$ fname --delimiter _
glaring_perception
```

Specify the seed for generating names:

```sh
$ fname --seed 123
foundational-spot

$ fname --seed 123
foundational-spot
```

## License

## Related Projects
