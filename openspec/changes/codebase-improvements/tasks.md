## 1. Seed Bug Fix

- [x] 1.1 Change `seed` variable in `main()` from `int64 = -1` to `*int64` (nil pointer)
- [x] 1.2 Update pflag registration to use `pflag.Int64VarP` with a pointer receiver
- [x] 1.3 Replace the `if seed != -1` guard with a nil-pointer check
- [x] 1.4 Verify that `--seed -1`, `--seed 0`, and omitting `--seed` all behave correctly

## 2. Remove False Collision-Avoidance Loop

- [ ] 2.1 Delete the `for adjectiveIndex == nounIndex` loop in `generator.go`
- [ ] 2.2 Confirm tests still pass after removal

## 3. Generator Validation

- [ ] 3.1 Change `WithSize` signature to return `(GeneratorOption, error)`
- [ ] 3.2 Move the `size < 2 || size > 4` check into `WithSize` and return error there
- [ ] 3.3 Remove the size check from `Generate()`
- [ ] 3.4 Update all `WithSize` call sites (tests, CLI) to handle the returned error
- [ ] 3.5 Add CLI validation for `--quantity`: print error and exit non-zero if ≤ 0

## 4. Performance: Word List Parsing

- [ ] 4.1 Replace `bufio.Scanner` implementation in `split()` with `strings.Split` + `strings.TrimRight`
- [ ] 4.2 Run benchmarks before and after to confirm improvement (optional but recommended)

## 5. Performance: Casing Switch

- [ ] 5.1 Remove the `casingMap` package-level variable from `generator.go`
- [ ] 5.2 Rewrite `applyCasing` to use a `switch` on `g.casing` directly
- [ ] 5.3 Verify casing tests still pass

## 6. Concurrency Documentation

- [ ] 6.1 Add a doc comment to `Generator` stating it is not safe for concurrent use from multiple goroutines

## 7. Custom Dictionary

- [ ] 7.1 Add a `WithDictionary(d *Dictionary) GeneratorOption` function to `generator.go`
- [ ] 7.2 Remove the `TODO: allow for custom dictionary` comment from `dictionary.go`
- [ ] 7.3 Add tests for `WithDictionary` with a custom word list
- [ ] 7.4 Document `WithDictionary` usage in the README library section

## 8. Output Formatting

- [ ] 8.1 Add `--format` / `-f` flag to CLI accepting `plain` (default) and `json`
- [ ] 8.2 Implement JSON output: marshal collected names as a `[]string` JSON array
- [ ] 8.3 Add validation: unrecognized format values print error and exit non-zero
- [ ] 8.4 Add README examples for `--format json`

## 9. API Rename: CasingFromString → ParseCasing

- [ ] 9.1 Add `ParseCasing` as the canonical function (same body as `CasingFromString`)
- [ ] 9.2 Mark `CasingFromString` as deprecated with a doc comment pointing to `ParseCasing`
- [ ] 9.3 Update internal call site in `cmd/fname/fname.go` to use `ParseCasing`

## 10. Word List Quality

- [ ] 10.1 Audit verb list and convert past-tense / past-participle entries to 3rd-person present tense
- [ ] 10.2 Run `task data:dupe` and `task data:spellcheck` after verb edits to verify no duplicates or typos
- [ ] 10.3 Identify the 72 adjective/noun overlap words and remove each from the less-appropriate list
- [ ] 10.4 Re-run combination count to confirm 2-word space stays above 4 million
