package gophers

// YodaGopher replies with Yoda transformation of text quotes instead of relaying the original message.
type YodaGopher struct {
	endpoint string
}

func NewYodaGopher() YodaGopher {
	return YodaGopher{
         endpoint: "https://yoda.p.mashape.com/yoda",
	}
}


func (g YodaGopher) TransformMessage(msg string) string {

	return "A yoda gopher I am."
}
