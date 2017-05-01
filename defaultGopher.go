package main

type DefaultGopher struct {
	name string
}

func (g DefaultGopher) Name() string {
	return g.name
}

func (g DefaultGopher) HandleMessage(msg Message) {
	msg.From = g.name

	// Pass along the message to the next gopher
	msg.Send()
}
