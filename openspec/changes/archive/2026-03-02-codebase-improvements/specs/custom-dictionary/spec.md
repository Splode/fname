## ADDED Requirements

### Requirement: Generator accepts a custom Dictionary
The `WithDictionary` option SHALL allow a caller to supply a `*Dictionary` instance, replacing the default embedded word lists.

#### Scenario: Custom adjectives are used in generated names
- **WHEN** a caller provides a Dictionary with a custom adjective list
- **THEN** generated names only use words from that custom adjective list

#### Scenario: Custom noun list is respected
- **WHEN** a caller provides a Dictionary with a custom noun list
- **THEN** generated names only use words from that custom noun list

#### Scenario: Nil dictionary falls back to default
- **WHEN** a caller passes `nil` as the Dictionary to `WithDictionary`
- **THEN** the generator uses the default embedded Dictionary

### Requirement: Dictionary can be constructed with custom word lists
The `NewDictionary` constructor (or an alternative constructor) SHALL accept optional word lists so callers can build a Dictionary without embedding data files.

#### Scenario: Caller-provided word slices are used
- **WHEN** a caller constructs a Dictionary with custom adjective and noun slices
- **THEN** the Dictionary reports the correct lengths for those word categories

#### Scenario: Empty word list for unused category is valid
- **WHEN** a caller constructs a 2-word Generator with a Dictionary that has an empty verb list
- **THEN** generation succeeds because verbs are not used for size-2 names
