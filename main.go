package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/ladygogo/telephone/gophers"
)

var gopher gophers.DefaultGopher

func main() {
	flag.StringVar(&gophers.ListenerAddress, "me", "localhost:8081", "My listener address, defaults to localhost:8081")
	flag.StringVar(&gophers.FriendAddress, "friend", "localhost:8082", "My friend's address, defaults to localhost:8082")
	flag.StringVar(&gophers.Message, "msg", "Hello there, my friend!", "Initial message to send")
	flag.Parse()

	done := make(chan bool)

	go listenForMessages(done)
	go gopher.SendMessages(gophers.Message, done)
	<-done
	<-done
}

func listenForMessages(done chan bool) {
	// Start listening for messages
	log.Printf("Listening on %s\n", gophers.ListenerAddress)
	listen, err := net.Listen("tcp4", gophers.ListenerAddress)
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

	gopher.SendMessages(msg, done)
}
