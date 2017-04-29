package gophers

import (
	"log"
	"net"
	"time"
)

type DefaultGopher struct{}

func (g DefaultGopher) SendMessages(msg string, done chan bool) {
	// Connect to a friend
	conn, err := net.Dial("tcp", FriendAddress)
	if err != nil {
		time.Sleep(5 * time.Second)
		g.SendMessages(msg, done)
		return
	}
	defer conn.Close()

	// Make a message and convert it to bytes
	msgb := []byte(msg)

	log.Printf("Sending message: %s", msg)
	_, err = conn.Write(msgb)
	if err != nil {
		log.Fatalf("Unable to send the message: %v", err)
	}
	done <- true
}
