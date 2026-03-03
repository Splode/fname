### Requirement: Size validation occurs at option construction time
`WithSize` SHALL return an error immediately if the provided size is outside the valid range (2–4), so invalid generators cannot be constructed.

#### Scenario: Invalid size is rejected at construction
- **WHEN** a library caller passes `WithSize(1)` to `NewGenerator`
- **THEN** `WithSize` returns a non-nil error before `NewGenerator` is called

#### Scenario: Invalid size 5 is rejected at construction
- **WHEN** a library caller passes `WithSize(5)` to `NewGenerator`
- **THEN** `WithSize` returns a non-nil error

#### Scenario: Valid sizes 2, 3, 4 are accepted
- **WHEN** a library caller passes `WithSize(2)`, `WithSize(3)`, or `WithSize(4)`
- **THEN** `WithSize` returns a nil error and the option applies successfully

### Requirement: CLI rejects zero or negative quantity
The CLI SHALL return an error and non-zero exit code when `--quantity` is zero or negative, rather than producing no output silently.

#### Scenario: Quantity zero prints an error
- **WHEN** a user runs `fname --quantity 0`
- **THEN** the CLI prints a descriptive error message and exits with a non-zero status

#### Scenario: Negative quantity prints an error
- **WHEN** a user runs `fname --quantity -5`
- **THEN** the CLI prints a descriptive error message and exits with a non-zero status

#### Scenario: Positive quantity works normally
- **WHEN** a user runs `fname --quantity 3`
- **THEN** the CLI prints 3 name phrases and exits with status 0
