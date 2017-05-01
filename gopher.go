package main

// Gopher is the common interface for message passing gophers
type Gopher interface {
	// Name of the gopher
	Name() string

	// HandleMessage performs your gopher's custom actions on the message
	// and sends it to the next gopher in the chain
	HandleMessage(msg Message)
}
