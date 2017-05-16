package gophers

// DefaultGopher is a boring gopher who simply repeats the received message
type DefaultGopher struct{}

func (g DefaultGopher) TransformMessage(msg string) string {
	return msg
}
