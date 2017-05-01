package main

type DefaultGopher struct{}

func (g DefaultGopher) TransformMessage(msg Message) string {
	return msg.Body
}
