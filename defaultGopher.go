package main

type DefaultGopher struct {
	Name string
}

func (g DefaultGopher) HandleMessage(msg Message) {
	msg.From = g.Name

	// Pass along the message to the next gopher
	msg.Send()
}
