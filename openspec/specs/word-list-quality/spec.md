### Requirement: Verb list uses consistent present-tense forms
All entries in the verb word list SHALL use 3rd-person singular present tense (e.g., "walks", "runs", "discovers"). Past tense, past participle, and bare infinitive forms SHALL be removed or converted.

#### Scenario: Generated 3-word name uses a present-tense verb
- **WHEN** a user generates a size-3 name phrase multiple times
- **THEN** the verb component consistently reads as a present-tense action (ending in -s or -es for regular verbs)

#### Scenario: No past-participle verbs appear in output
- **WHEN** a user generates a large batch of size-3 names
- **THEN** no verb component ends in "-ed" in a way that reads as past tense (e.g., "abandoned", "admired" are absent)

### Requirement: No word appears in both the adjective and noun lists
Words that are dual-category (e.g., "blue", "dark", "cold") SHALL appear in at most one list. For each overlap word, the appropriate category SHALL be chosen based on how it reads in context as part of a generated name.

#### Scenario: Adjective list contains no noun-list entries
- **WHEN** the adjective and noun data files are compared
- **THEN** there are zero words appearing in both files

#### Scenario: Name quality is unaffected by overlap removal
- **WHEN** overlap words are removed from one list
- **THEN** the total combination space remains above 4 million for 2-word names
