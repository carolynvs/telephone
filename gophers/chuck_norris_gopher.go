package gophers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// ChuckNorrisGopher replies with Chuck Norris quotes instead of relaying the original message.
type ChuckNorrisGopher struct {
	endpoint string
}

func NewChuckNorrisGopher() ChuckNorrisGopher {
	return ChuckNorrisGopher{
		endpoint: "http://api.icndb.com/jokes/random",
	}
}

type quoteResult struct {
	Value quoteWrapper `json:"value"`
}
type quoteWrapper struct {
	Joke string `json:"joke"`
}

func (g ChuckNorrisGopher) TransformMessage(msg string) string {
	resp, err := http.Get(g.endpoint)
	if err != nil {
		log.Printf("%+v", errors.Wrapf(err, "Unable to retrieve a quote from %s", g.endpoint))
		return "Chuck Norris failed!"
	}

	result := quoteResult{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Printf("%+v", errors.Wrapf(err, "Unable to unmarshal a quote from %s", g.endpoint))
		return "Chuck Norris failed!"
	}
	return result.Value.Joke
}
