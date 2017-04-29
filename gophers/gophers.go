package gophers

import (
	"io/ioutil"
	"log"
	"net"
)

var ListenerAddress string

func readMessageString(conn net.Conn) string {
	defer conn.Close()

	// Read all of the message bytes
	log.Println("Reading incoming message...")
	msgb, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatalf("Unable to read incoming message: %v", err)
	}

	// Assume the message is a string
	msg := string(msgb)

	// Print the message
	log.Printf("\n\n\t%s\n\n", msg)

	return msg
}
