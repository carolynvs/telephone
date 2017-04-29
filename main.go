package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/ladygogo/telephone/gophers"
)

var gopher gophers.DefaultGopher

var friend string
var initialMessage string

func main() {
	flag.StringVar(&gophers.ListenerAddress, "me", "localhost:8081", "My listener address, defaults to localhost:8081")
	flag.StringVar(&friend, "friend", "gopher2:localhost:8082", "My friend's name:address:port")
	flag.StringVar(&initialMessage, "msg", "Hello my good friend!", "Initial message to send")
	flag.StringVar(&gopher.Name, "name", "gopher1", "Your gopher's name")

	flag.Parse()

	friendArgs := strings.SplitN(friend, ":", 2)
	gopher.FriendName = friendArgs[0]
	gopher.FriendAddress = friendArgs[1]
	log.Printf("Hi my name is %s, and I am playing telephone with my friend %s", gopher.Name, gopher.FriendName)

	messageReceived := make(chan bool)
	go listenForMessages(messageReceived)

	// The gopher with the lower name starts the game
	if gopher.Name < gopher.FriendName {
		go gopher.Send(initialMessage)
	}

	// Wait until we receive a message, then quit
	<-messageReceived
}

func listenForMessages(messageReceived chan bool) {
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
		gopher.HandleMessage(conn)
		messageReceived <- true
	}
}
