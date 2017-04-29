package gophers

import (
	"net"
)

type DefaultGopher struct {
	Name          string
	FriendName    string
	FriendAddress string
}

func (g DefaultGopher) HandleMessage(conn net.Conn) {
	msg := readMessageString(conn)

	// Pass along the message unchanged to my friend
	g.SendMessage(msg)
}

func (g DefaultGopher) SendMessage(msg string) {
	sendMessage(msg, g.FriendAddress)
}
