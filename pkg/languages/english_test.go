package languages

import (
	"testing"
)

func TestEnglishHello(t *testing.T) {
	e := NewEnglish()

	got := e.Hello()
	want := "Hello"

	if got != want {
		t.Errorf("mismatch (got %q, want %q)\n", got, want)
	}
}

func TestEnglishGoodbye(t *testing.T) {
	e := NewEnglish()

	got := e.Goodbye()
	want := "Goodbye"

	if got != want {
		t.Errorf("mismatch (got %q, want %q)\n", got, want)
	}
}
