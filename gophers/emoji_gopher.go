package gophers

// EmojiGopher replaces well-known words with their emoji equivalent.
// For example, it replaces the word `love` with ❤️.
type EmojiGopher struct{}

func (g EmojiGopher) TransformMessage(msg string) string {
	// TODO: Pick a set of words and replace them in the message with their emoji equivalent

	// Helpful links:
	// * https://golang.org/pkg/strings/#Replace

	return "I ❤ emoji!"
}
