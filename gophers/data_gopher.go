package gophers

// CSIGopher can read the files on your computer, searching for clues to solve the case.
// Pick words from the message received, and use it to lookup relevant text from the sample
// files in this repository.
type DataScienceGopher struct{}

func (g DataScienceGopher) TransformMessage(msg string) string {
	// TODO: Scan all the words in the message and see if you can find it
	// in the films table. Return the title and description.

	// Helpful links:
	// https://dev.mysql.com/doc/sakila/en/sakila-structure.html
	// https://github.com/mattn/go-sqlite3
	// https://www.thepolyglotdeveloper.com/2017/04/using-sqlite-database-golang-application/

	return "Anything you can say, I can say better with movie quotes."
}
