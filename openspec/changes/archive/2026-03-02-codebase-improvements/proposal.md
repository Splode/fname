## Why

A codebase audit identified a set of bugs, UX rough edges, performance inefficiencies, and library API gaps that have accumulated over time. Addressing them systematically will improve correctness, usability, and the quality of the library as a dependency. The seed sentinel bug, in particular, is a functional defect where a valid input value is silently discarded.

## What Changes

- Fix seed sentinel bug: `-1` is currently treated as "no seed provided," making it impossible to use `-1` as a seed value
- Remove false collision-avoidance loop that compares indices across differently-sized arrays (solving a non-existent problem)
- Move size validation to construction time so invalid generators fail eagerly
- Add validation for `--quantity` flag to reject zero or negative values with a clear error
- Replace `bufio.Scanner` with `strings.Split` for in-memory word list parsing
- Replace `casingMap` map lookup with a direct `switch` in `applyCasing`
- Add goroutine safety to `Generator` (or document that it is not safe for concurrent use)
- Normalize verb tense across the verb word list (consistent 3rd-person singular present tense)
- Audit and clean up adjective/noun word overlap (72 words appear in both lists)
- Add `--format` flag to CLI for structured output (e.g., JSON, newline-delimited)
- Implement `WithDictionary()` option to fulfill the existing `TODO` in `dictionary.go`
- Rename `CasingFromString` to `ParseCasing` for idiomatic Go style

## Capabilities

### New Capabilities

- `seed-handling`: Correct, unambiguous seed input — use `*int64` or a separate flag rather than a sentinel value
- `generator-validation`: Eager validation of generator options at construction time, including size and quantity
- `custom-dictionary`: `WithDictionary()` option allowing callers to supply their own word lists
- `output-formatting`: `--format` CLI flag supporting structured output (JSON, plain)
- `word-list-quality`: Normalized verb tenses and cleaned adjective/noun overlap in data files

### Modified Capabilities

## Impact

- `generator.go`: seed logic, collision loop removal, size validation, `applyCasing` switch, concurrency docs
- `cmd/fname/fname.go`: seed flag type change, quantity validation, `--format` flag
- `dictionary.go`: `split()` implementation, `WithDictionary()` option, `ParseCasing` rename
- `data/verb`, `data/adjective`, `data/noun`: word list data file edits
- Library public API: `CasingFromString` → `ParseCasing` is a **BREAKING** rename; `WithDictionary` is additive
