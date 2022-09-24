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

			adj, err := d.Adjective(0)
			if err != nil {
				t.Error("\t\tShould be able to get an adjective from the dictionary.")
			}
			t.Log("\t\tShould be able to get an adjective from the dictionary.")
			if adj != "able" {
				t.Error("\t\tShould be able to get the first adjective from the dictionary.")
			}
			t.Log("\t\tShould be able to get the first adjective from the dictionary.")

			_, err = d.Adjective(-1)
			if err == nil {
				t.Error("\t\tShould not be able to get an adjective from the dictionary with a negative index.")
			}
			t.Log("\t\tShould not be able to get an adjective from the dictionary with a negative index.")

			adv, err := d.Adverb(0)
			if err != nil {
				t.Error("\t\tShould be able to get an adverb from the dictionary.")
			}
			t.Log("\t\tShould be able to get an adverb from the dictionary.")
			if adv != "abnormally" {
				t.Error("\t\tShould be able to get the first adverb from the dictionary.")
			}
			t.Log("\t\tShould be able to get the first adverb from the dictionary.")

			_, err = d.Adverb(-1)
			if err == nil {
				t.Error("\t\tShould not be able to get an adverb from the dictionary with a negative index.")
			}
			t.Log("\t\tShould not be able to get an adverb from the dictionary with a negative index.")

			noun, err := d.Noun(0)
			if err != nil {
				t.Error("\t\tShould be able to get a noun from the dictionary.")
			}
			t.Log("\t\tShould be able to get a noun from the dictionary.")

			if noun != "aardvark" {
				t.Error("\t\tShould be able to get the first noun from the dictionary.")
			}
			t.Log("\t\tShould be able to get the first noun from the dictionary.")

			_, err = d.Noun(-1)
			if err == nil {
				t.Error("\t\tShould not be able to get a noun from the dictionary with a negative index.")
			}
			t.Log("\t\tShould not be able to get a noun from the dictionary with a negative index.")

			verb, err := d.Verb(0)
			if err != nil {
				t.Error("\t\tShould be able to get a verb from the dictionary.")
			}
			t.Log("\t\tShould be able to get a verb from the dictionary.")

			if verb != "abandoned" {
				t.Error("\t\tShould be able to get the first verb from the dictionary.")
			}
			t.Log("\t\tShould be able to get the first verb from the dictionary.")

			_, err = d.Verb(-1)
			if err == nil {
				t.Error("\t\tShould not be able to get a verb from the dictionary with a negative index.")
			}
			t.Log("\t\tShould not be able to get a verb from the dictionary with a negative index.")

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
