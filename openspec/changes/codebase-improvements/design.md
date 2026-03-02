## Context

`fname` is a small Go library and CLI for generating human-friendly random name phrases. The codebase is compact (~150 LOC across 3 Go files) but has accumulated several bugs and rough edges. The public library API is used externally (`go get github.com/splode/fname`), so breaking changes require care.

Current pain points:
- The seed `-1` sentinel makes a valid int64 input silently unusable
- An index-based collision loop guards against a problem that doesn't exist
- Size validation happens too late (at generate time, not construction time)
- The word-list parser uses streaming I/O on in-memory strings
- Verb tenses are inconsistent across the 448-word verb list
- 72 words appear in both adjective and noun lists

## Goals / Non-Goals

**Goals:**
- Fix the seed sentinel bug with a clean API that accepts all int64 values
- Remove dead/incorrect logic (collision loop)
- Validate generator options eagerly at construction time
- Improve word list quality (verb tense consistency, cross-list overlap)
- Minor performance improvements (split, casing switch)
- Add `WithDictionary()` to fulfill the existing TODO
- Add `--format` output flag to the CLI

**Non-Goals:**
- Rewriting the generator architecture
- Changing the default name format or word order
- Expanding the size range beyond 2–4
- Adding cryptographic randomness

## Decisions

### D1: Seed flag uses `*int64` instead of sentinel `-1`

**Decision**: Change the `seed` variable in `main()` from `int64 = -1` to `*int64` (nil = unset). Pass `fname.WithSeed(*seed)` only when non-nil.

**Alternatives considered**:
- Separate `--no-seed` bool flag: adds a flag just to undo another flag, confusing
- Use `0` as sentinel: same problem, `0` is a valid seed
- `*int64` nil pointer: idiomatic Go for "optional value", clean, no reserved values

**Impact**: CLI-only change. The library's `WithSeed(int64)` signature is unchanged.

### D2: Remove collision-avoidance loop, no replacement

**Decision**: Delete the `for adjectiveIndex == nounIndex` loop entirely. No replacement needed.

**Rationale**: The loop compares an index into the adjective list (0..1745) against an index into the noun list (0..2663). Equal integers do not mean equal words — `adjective[5]="absolute"`, `noun[5]="absence"`. The loop solves a phantom problem and could theoretically spin indefinitely on equal-length lists.

**Alternatives considered**:
- Fix the loop to compare word strings: adds a real-collision check, but two-word names from a 4.6M combination space make same-word collisions negligible (~0.07% for largest overlap category)
- Keep loop as-is: incorrect semantics, wasted cycles

### D3: Eager size validation in `WithSize`

**Decision**: Return an error from `WithSize()`, changing its signature to `func WithSize(size uint) (GeneratorOption, error)`. Callers get immediate feedback on invalid sizes.

**Alternatives considered**:
- Validate in `NewGenerator()` and return `(*Generator, error)`: larger API change, but cleaner; deferred for a future refactor
- Keep deferred validation in `Generate()`: current state, poor library ergonomics
- Panic in `WithSize()`: not idiomatic Go for user input validation

**Note**: This is a **BREAKING** change to the `WithSize` function signature.

### D4: Replace `bufio.Scanner` with `strings.Split`

**Decision**: Replace the `split()` function body with `strings.Split(strings.TrimRight(s, "\n"), "\n")`. The embedded data is already in memory; a streaming scanner adds unnecessary overhead.

**Rationale**: Simpler, faster, no reader allocation. The only edge case is a trailing newline on the embedded string, which `strings.TrimRight` handles.

### D5: Replace `casingMap` with a `switch` in `applyCasing`

**Decision**: Remove `casingMap` and replace the map lookup in `applyCasing` with a direct `switch` on `g.casing`.

**Rationale**: Three cases don't benefit from a map. A switch is zero-allocation, branch-predictor friendly, and more readable.

### D6: Document (not fix) goroutine safety

**Decision**: Add a doc comment to `Generator` stating it is not safe for concurrent use. Creating a new `Generator` per goroutine is the idiomatic solution.

**Alternatives considered**:
- Add a `sync.Mutex` around rand calls: adds lock contention overhead for the common single-goroutine case
- Switch to `math/rand/v2` global rand: changes minimum Go version requirement and behavior

### D7: `WithDictionary()` takes a `*Dictionary`

**Decision**: Add `WithDictionary(d *Dictionary)` as a new `GeneratorOption`. `NewDictionary()` remains the default; callers who want custom words construct their own `Dictionary` and pass it in.

**Rationale**: Minimal API surface, composable with existing options, fulfills the existing TODO with no breakage.

### D8: `CasingFromString` → `ParseCasing`

**Decision**: Rename `CasingFromString` to `ParseCasing`. Keep `CasingFromString` as a deprecated alias for one release cycle.

**Rationale**: `ParseX` is the idiomatic Go convention for string-to-value parsing (cf. `strconv.ParseInt`, `time.Parse`).

### D9: `--format` flag with `plain` (default) and `json`

**Decision**: Add `--format` / `-f` flag accepting `plain` (default, current behavior) or `json`. JSON output is an array of name strings: `["name1","name2"]`.

**Rationale**: Enables scripting without `xargs` / `sed` gymnastics. An array is more useful than newline-delimited JSON objects for batch generation.

## Risks / Trade-offs

- **`WithSize` signature change is breaking** → Mitigated by it being a small, focused library; version bump to communicate the change
- **Verb list edits are manual and subjective** → Mitigated by focusing on clearly wrong tenses (past participles like "abandoned" in a verb slot that reads as present action); a full linguistic audit is out of scope
- **Adj/noun overlap cleanup could remove intentionally dual-use words** → Mitigated by only removing from the list where the word is clearly stronger in one category (e.g., "blue" as adjective, not noun)

## Open Questions

- Should `NewGenerator` be changed to return `(*Generator, error)` now, or deferred to a future version? (Current proposal keeps it non-error-returning for minimal breakage)
- Should JSON output for `--format json` include metadata (seed used, size, count)? Or just the name array?
