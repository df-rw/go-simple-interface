package languages

import (
	"testing"
)

func TestFrenchHello(t *testing.T) {
	e := NewFrench()

	got := e.Hello()
	want := "Bon jour"

	if got != want {
		t.Errorf("mismatch (got %q, want %q)\n", got, want)
	}
}

func TestFrenchGoodbye(t *testing.T) {
	e := NewFrench()

	got := e.Goodbye()
	want := "Au revoir"

	if got != want {
		t.Errorf("mismatch (got %q, want %q)\n", got, want)
	}
}
