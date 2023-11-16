package translator

import "testing"

func TestNewEnglish(t *testing.T) {
	language := "english"
	l := New(language)

	if l == nil {
		t.Errorf("received nil translator for %s, should not be nil\n", language)
	}
}

func TestNewFrench(t *testing.T) {
	language := "french"
	l := New(language)

	if l == nil {
		t.Errorf("received nil translator for %s, should not be nil\n", language)
	}
}

func TestNewNonExistant(t *testing.T) {
	language := "mandarin"
	l := New(language)

	if l != nil {
		t.Errorf("received translator for %s, should be nil\n", language)
	}
}
