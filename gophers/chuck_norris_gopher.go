package gophers

// ChuckNorrisGopher replies with Chuck Norris quotes instead of relaying the original message.
type ChuckNorrisGopher struct {
	endpoint string
}

func NewChuckNorrisGopher() ChuckNorrisGopher {
	return ChuckNorrisGopher{
		endpoint: "http://api.icndb.com/jokes/random",
	}
}

func (g ChuckNorrisGopher) TransformMessage(msg string) string {
	// TODO: Lookup quotes from an HTTP API and return them
	// Feel free to use a different quote API if Chuck isn't to your liking

	// Helpful links:
	// * https://golang.org/pkg/net/http/#example_Get
	// * https://golang.org/pkg/encoding/json/#example_Decoder_Decode_stream

	return "Chuck Norris quotes are great!"
}
