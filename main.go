package main

import (
	"flag"
	"log"
	"net"
	"strings"

	"github.com/ladygogo/telephone/gophers"
)

var listenerAddress string
var friend string
var initialMessage string
var gopher gophers.DefaultGopher

func main() {
	flag.StringVar(&listenerAddress, "me", "localhost:8081", "My listener address, defaults to localhost:8081")
	flag.StringVar(&friend, "friend", "", "My friend's name:address:port")
	flag.StringVar(&initialMessage, "msg", "", "Initial message to send")
	flag.StringVar(&gopher.Name, "name", "gopher1", "Your gopher's name")

	flag.Parse()

	if friend != "" {
		friendArgs := strings.SplitN(friend, ":", 2)
		gopher.FriendName = friendArgs[0]
		gopher.FriendAddress = friendArgs[1]
		log.Printf("Hi, my name is %s, and I am playing telephone with my friend %s", gopher.Name, gopher.FriendName)
	} else {
		log.Printf("Hi, my name is %s, and I am the last gopher in the telephone chain.", gopher.Name)
	}

	if initialMessage != "" {
		// I am the first gopher, start the telephone chain, and then quit
		gopher.SendMessage(initialMessage)
	} else {
		// Wait until we receive a message, and then quit
		messageReceived := make(chan bool)
		go listenForMessages(messageReceived)
		<-messageReceived
	}

	log.Println("Game over! ☎️")
}

func listenForMessages(messageReceived chan bool) {
	// Start listening for messages
	log.Printf("Listening on %s\n", listenerAddress)
	listen, err := net.Listen("tcp4", listenerAddress)
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
		close(messageReceived)
	}
}
