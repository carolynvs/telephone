package gophers

import "testing"

func TestCSIGopher(t *testing.T) {
	g := EmojiGopher{}

	want := "It is a truth universally acknowledged, that a single man in possession of a good fortune, must be in want of a wife."
	got := g.TransformMessage("truth")
	if got != want {
		t.Fatalf("Expected '%s', Got '%s'", want, got)
	}
}
