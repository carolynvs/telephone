# Let's Play Gopher Telephone!

This is the classic game of Telephone, played out by helpful gophers.

## Rules
* All gophers get in a line, sorted by their name.
* The first gopher thinks up a really special message and tells it to the next gopher.
* Each gopher repeats the message to the next in line.
* When the last gopher receives the message, she tells everyone the message.

If every gopher faithfully repeats the message, then the game isn't terribly fun. However, if each gopher has a slighty different personality, altering the message in unique ways, then the final message can be quite different from the original. Haha, things that are different from what we expect are funny! Get it?! 😎

## Gophers

* **Default**: This gopher is quite boring and repeats the message exactly as received. Everyone starts out with this gopher so that they have a working game to play with.
* **Emoji Gopher**: This gopher ❤️ emoji, and replaces well-known words with their emoji equivalent. For example, it replaces the word `love` with ❤️.
* **Chuck Norris Gopher**: This gopher is obsessed with Chuck Norris, and instead of relaying the message received, sends [Chuck Norris quotes](norris).
* **Data Science Gopher**: Data is awesome! This gopher queries the [SQLite](sqlite) database provided in this repository and sends interesting gopher facts.
* **CSI Cyber Gopher**: Lookout criminals, CSI gopher can read the files on your computer, searching for clues to solve the case. Pick words from the message received, and use it to lookup relevant text from the sample files in this repository.

[norris]: http://api.icndb.com/jokes/random
[sqlite]: https://github.com/mattn/go-sqlite3
[aciitext]: http://artii.herokuapp.com/make?text=gophers

# Run the Game

```
go get -u github.com/ladygogo/telephone
cd $GOPATH/github.com/ladygogo/telephone
go build
# open a terminal and run the following
./telephone -name gopher1
# open another terminal in the same directory and run the following
./telephone -name gopher2
```