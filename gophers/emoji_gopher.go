package gophers

import "strings"

// EmojiGopher replaces well-known words with their emoji equivalent.
// For example, it replaces the word `love` with ❤️.
type EmojiGopher struct{}

func (g EmojiGopher) TransformMessage(msg string) string {
	// Solution 2 is more efficient
	return g.solution2(msg)
}

func (g EmojiGopher) solution1(msg string) string {
	// Solution 1: Keep a mapping of replacements and apply them one at a time
	// https://golang.org/pkg/strings/#Replace
	emojis := map[string]string{
		"poop": "💩",
		"love": "💖",
	}
	for txt, emj := range emojis {
		msg = strings.Replace(msg, txt, emj, -1)
	}
	return msg
}

func (g EmojiGopher) solution2(msg string) string {
	// Solution 2: Use a Replacer, takes a list of replacements and applies them in a single loop
	// https://golang.org/pkg/strings/#Replacer
	replacer := strings.NewReplacer("poop", "💩", "love", "💖")
	return replacer.Replace(msg)
}
