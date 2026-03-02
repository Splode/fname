## ADDED Requirements

### Requirement: All int64 values are valid seeds
The generator seed option SHALL accept all int64 values, including negative values. No integer value SHALL be treated as a sentinel meaning "no seed."

#### Scenario: Negative seed produces deterministic output
- **WHEN** a user provides `--seed -1`
- **THEN** the generator uses `-1` as the seed and produces deterministic output

#### Scenario: Same negative seed produces same names
- **WHEN** two generators are created with the same negative seed value
- **THEN** both generators produce identical name sequences

#### Scenario: Seed zero is valid
- **WHEN** a user provides `--seed 0`
- **THEN** the generator uses `0` as the seed and produces deterministic output

### Requirement: Omitting seed produces random output
When no seed is provided, the generator SHALL use a time-based random seed, producing non-deterministic output across invocations.

#### Scenario: No seed flag means random generation
- **WHEN** a user runs `fname` without `--seed`
- **THEN** repeated invocations produce different name sequences

#### Scenario: WithSeed option is not applied when seed is absent
- **WHEN** a library caller creates a Generator without `WithSeed`
- **THEN** the generator behaves as if seeded randomly
