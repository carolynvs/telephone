package main

// Gopher is the common interface for message passing gophers
type Gopher interface {
	// TransformMessage performs your gopher's custom actions on the message
	TransformMessage(msg string) string
}
