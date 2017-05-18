package gophers

// CSIGopher can read the files on your computer, searching for clues to solve the case.
// Pick words from the message received, and use it to lookup relevant text from the sample
// files in this repository.
type CSIGopher struct{}

func (g CSIGopher) TransformMessage(msg string) string {
	// TODO: Pick a word from the message, and look for it in the sample books/ directory.
	// Return the sentence that contains the code word.

	// Helpful links:
	// https://golang.org/pkg/os/#Open
	// https://golang.org/pkg/bufio/#example_Scanner_lines

	return "You cannot hide from justice!"
}
