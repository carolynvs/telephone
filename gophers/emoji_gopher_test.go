package gophers

import "testing"

func TestEmojiGopher(t *testing.T) {
	g := EmojiGopher{}

	want := "I 💖 💩y diapers"
	got := g.TransformMessage("I love poopy diapers")
	if got != want {
		t.Fatalf("Expected '%s', Got '%s'", want, got)
	}
}
