package fname

import (
	"strings"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	t.Log("Given the need to test the NewGenerator function")
	{
		t.Log("\tWhen creating a new Generator with default values")
		{
			g := NewGenerator()
			if g == nil {
				t.Fatal("\t\tShould be able to create a Generator instance.")
			}
			t.Log("\t\tShould be able to create a Generator instance.")

			if g.dict == nil {
				t.Fatal("\t\tShould be able to create a Dictionary instance.")
			}
			t.Log("\t\tShould be able to create a Dictionary instance.")
		}

		t.Log("\tWhen creating a new Generator with custom values")
		{
			g := NewGenerator(WithCasing(Title), WithDelimiter("_"), WithSize(3), WithSeed(12345))
			if g == nil {
				t.Fatal("\t\tShould be able to create a Generator instance.")
			}
			t.Log("\t\tShould be able to create a Generator instance.")

			if g.casing != Title {
				t.Error("\t\tShould be able to set the casing.")
			}

			if g.size != 3 {
				t.Fatal("\t\tShould be able to set the size of the phrase.")
			}
			t.Log("\t\tShould be able to set the size of the phrase.")

			if g.delimiter != "_" {
				t.Fatal("\t\tShould be able to set the delimiter of the phrase.")
			}
			t.Log("\t\tShould be able to set the delimiter of the phrase.")
		}
	}
}

func TestGenerate(t *testing.T) {
	t.Log("Given the need to test the Generate function")
	{
		t.Log("\tWhen generating a phrase")
		{
			g := NewGenerator()
			phrase, err := g.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			if len(phrase) == 0 {
				t.Fatal("\t\tShould be able to generate a phrase of non-zero length.")
			}
			t.Log("\t\tShould be able to generate a phrase of non-zero length.")

			parts := strings.Split(phrase, "-")
			if len(parts) != 2 {
				t.Fatal("\t\tShould be able to generate a phrase with 2 parts.")
			}
			t.Log("\t\tShould be able to generate a phrase with 2 parts.")
		}

		t.Log("\tWhen generating a phrase with a custom case")
		{
			g := NewGenerator(WithCasing(Title))
			phrase, err := g.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			c := phrase[0]
			if c < 'A' || c > 'Z' {
				t.Fatal("\t\tShould be able to generate a phrase with a title case.")
			}
		}

		t.Log("\tWhen generating a phrase with a custom delimiter")
		{
			g := NewGenerator(WithDelimiter("_"))
			phrase, err := g.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			if len(phrase) == 0 {
				t.Fatal("\t\tShould be able to generate a phrase of non-zero length.")
			}
			t.Log("\t\tShould be able to generate a phrase of non-zero length.")

			parts := strings.Split(phrase, "_")
			if len(parts) != 2 {
				t.Fatal("\t\tShould be able to generate a phrase with 2 parts.")
			}
			t.Log("\t\tShould be able to generate a phrase with 2 parts.")

			if !strings.Contains(phrase, "_") {
				t.Fatal("\t\tShould be able to generate a phrase with the custom delimiter.")
			}
			t.Log("\t\tShould be able to generate a phrase with the custom delimiter.")
		}

		t.Log("\tWhen generating a phrase with a custom size")
		{
			g3 := NewGenerator(WithSize(3))
			phrase, err := g3.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			if len(phrase) == 0 {
				t.Fatal("\t\tShould be able to generate a phrase of non-zero length.")
			}
			t.Log("\t\tShould be able to generate a phrase of non-zero length.")

			parts := strings.Split(phrase, "-")
			if len(parts) != 3 {
				t.Fatal("\t\tShould be able to generate a phrase with 3 parts.")
			}
			t.Log("\t\tShould be able to generate a phrase with 3 parts.")

			g4 := NewGenerator(WithSize(4))
			phrase, err = g4.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			if len(phrase) == 0 {
				t.Fatal("\t\tShould be able to generate a phrase of non-zero length.")
			}
			t.Log("\t\tShould be able to generate a phrase of non-zero length.")

			parts = strings.Split(phrase, "-")
			if len(parts) != 4 {
				t.Fatal("\t\tShould be able to generate a phrase with 4 parts.")
			}
			t.Log("\t\tShould be able to generate a phrase with 4 parts.")
		}

		t.Log("\tWhen generating a phrase with a custom seed")
		{
			g1 := NewGenerator(WithSeed(12345))
			phrase1, err := g1.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			g2 := NewGenerator(WithSeed(12345))
			phrase2, err := g2.Generate()
			if err != nil {
				t.Fatal("\t\tShould be able to generate a phrase without error.")
			}
			t.Log("\t\tShould be able to generate a phrase without error.")

			if phrase1 != phrase2 {
				t.Fatal("\t\tShould be able to generate the same phrase with the same seed.")
			}
			t.Log("\t\tShould be able to generate the same phrase with the same seed.")
		}

		t.Log("\tWhen generating a phrase with an invalid size")
		{
			g := NewGenerator(WithSize(1))
			_, err := g.Generate()
			if err == nil {
				t.Fatal("\t\tShould not be able to generate a phrase with an invalid size.")
			}
			t.Log("\t\tShould not be able to generate a phrase with an invalid size.")
		}
	}
}

func TestCasingFromString(t *testing.T) {
	t.Log("Given the need to parse casing strings")
	{
		t.Log("\tWhen parsing a valid casing string")
		{
			testCases := []struct {
				name string
				c    Casing
			}{
				{"lower", Lower},
				{"upper", Upper},
				{"title", Title},
			}
			for _, tc := range testCases {
				c, err := CasingFromString(tc.name)
				if err != nil {
					t.Fatalf("\t\tShould be able to parse a valid casing string : %v", err)
				}
				t.Log("\t\tShould be able to parse a valid casing string")

				if c != tc.c {
					t.Fatalf("\t\tShould be able to parse a valid casing string : got %v, want %v", c, tc.c)
				}
				t.Log("\t\tShould be able to parse a valid casing string")
			}
		}

		t.Log("\tWhen parsing an invalid casing string")
		{
			_, err := CasingFromString("invalid")
			if err == nil {
				t.Fatal("\t\tShould not be able to parse an invalid casing string")
			}
			t.Log("\t\tShould not be able to parse an invalid casing string")
		}
	}
}
