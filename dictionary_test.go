package fname

import "testing"

func TestNewDictionary(t *testing.T) {
	t.Log("Given the need to test the NewDictionary function")
	{
		t.Log("\tWhen creating a new Dictionary with default values")
		{
			d := NewDictionary()
			if d == nil {
				t.Fatal("\t\tShould be able to create a Dictionary instance.")
			}
			t.Log("\t\tShould be able to create a Dictionary instance.")

			if len(d.adectives) == 0 {
				t.Error("\t\tShould be able to load the adjective file.")
			}
			t.Log("\t\tShould be able to load the adjective file.")

			if len(d.adverbs) == 0 {
				t.Error("\t\tShould be able to load the adverb file.")
			}
			t.Log("\t\tShould be able to load the adverb file.")

			if len(d.nouns) == 0 {
				t.Error("\t\tShould be able to load the noun file.")
			}
			t.Log("\t\tShould be able to load the noun file.")

			if len(d.verbs) == 0 {
				t.Error("\t\tShould be able to load the verb file.")
			}

			if len(d.adectives) != d.LengthAdjective() {
				t.Error("\t\tShould be able to get the length of the adjective file.")
			}
			t.Log("\t\tShould be able to get the length of the adjective file.")

			if len(d.adverbs) != d.LengthAdverb() {
				t.Error("\t\tShould be able to get the length of the adverb file.")
			}
			t.Log("\t\tShould be able to get the length of the adverb file.")

			if len(d.nouns) != d.LengthNoun() {
				t.Error("\t\tShould be able to get the length of the noun file.")
			}
			t.Log("\t\tShould be able to get the length of the noun file.")

			if len(d.verbs) != d.LengthVerb() {
				t.Error("\t\tShould be able to get the length of the verb file.")
			}
			t.Log("\t\tShould be able to get the length of the verb file.")
		}
	}
}
