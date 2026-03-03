## ADDED Requirements

### Requirement: CLI supports structured JSON output
The CLI SHALL accept a `--format` flag that controls output format. The default value is `plain` (current behavior). When `json` is specified, output SHALL be a JSON array of name strings.

#### Scenario: Default plain format is unchanged
- **WHEN** a user runs `fname --quantity 3` without `--format`
- **THEN** output is three names, one per line, as before

#### Scenario: JSON format produces a valid JSON array
- **WHEN** a user runs `fname --format json --quantity 3`
- **THEN** output is a single JSON array, e.g. `["name1","name2","name3"]`

#### Scenario: JSON format with quantity 1 produces a single-element array
- **WHEN** a user runs `fname --format json`
- **THEN** output is `["name"]` (an array, not a bare string)

#### Scenario: Invalid format value produces an error
- **WHEN** a user runs `fname --format csv`
- **THEN** the CLI prints a descriptive error and exits with a non-zero status

### Requirement: Short flag `-f` is the alias for `--format`
The `--format` flag SHALL have `-f` as its short-form alias.

#### Scenario: Short flag produces same output as long flag
- **WHEN** a user runs `fname -f json`
- **THEN** output is identical to `fname --format json`
