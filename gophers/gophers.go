package gophers

import (
	"io/ioutil"
	"log"
	"net"
	"time"
)

func sendMessage(msg string, destination string) {
	if destination == "" {
		// Nothing to do
		// I am the final gopher in the telephone chain
		return
	}

	// Connect to my friend
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		log.Println("My friend isn't ready yet to receive my message. Waiting...")
		time.Sleep(5 * time.Second)
		sendMessage(msg, destination)
		return
	}
	defer conn.Close()

	// Conver the message to bytes
	msgb := []byte(msg)

	log.Printf("Sending message: %s", msg)
	_, err = conn.Write(msgb)
	if err != nil {
		log.Fatalf("Unable to send the message: %v", err)
	}
}

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
