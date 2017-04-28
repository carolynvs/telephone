package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

const port int = 8081

var addr string = fmt.Sprintf("localhost:%d", port)

func main() {
	done := make(chan bool)

	go listenForMessages(done)
	sendMessages(done)

	<-done
	<-done
}

func listenForMessages(done chan bool) {
	// Start listening for messages
	log.Printf("Listening on %s\n", addr)
	listen, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatalf("Unable to listen for incoming message: %v", err)
	}
	defer listen.Close()

	// Handle incoming messages
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go readMessage(conn, done)
	}
}

func readMessage(conn net.Conn, done chan bool) {
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
	done <- true
}

func sendMessages(done chan bool) {
	// Connect to the listener
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalf("Unable to connect to the listener: %v", err)
	}
	defer conn.Close()

	// Make a message and convert it to bytes
	msg := "Hello World!\r\n\r\n"
	msgb := []byte(msg)

	log.Printf("Sending message: %s", msg)
	_, err = conn.Write(msgb)
	if err != nil {
		log.Fatalf("Unable to send the message: %v", err)
	}
	done <- true
}
