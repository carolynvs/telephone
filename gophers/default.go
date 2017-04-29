package gophers

import (
	"log"
	"net"
	"time"
)

type DefaultGopher struct {
	Name          string
	FriendName    string
	FriendAddress string
}

func (g DefaultGopher) Send(msg string) {
	// Connect to a friend
	conn, err := net.Dial("tcp", g.FriendAddress)
	if err != nil {
		log.Println("My friend isn't ready yet to receive my message. Waiting...")
		time.Sleep(5 * time.Second)
		g.Send(msg)
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
}

func (g DefaultGopher) HandleMessage(conn net.Conn) {
	msg := readMessageString(conn)

	if g.FriendAddress != "" {
		// I am not the last gopher in the chain
		// So keep passing the message along
		g.Send(msg)
	} else {
		log.Println("Game over! ☎️")
	}
}
