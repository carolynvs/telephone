package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pkg/errors"
)

const registrationAddress string = "224.0.0.1:9000"

var gopher DefaultGopher
var friends map[string]Friend = make(map[string]Friend)

func main() {
	var opts struct {
		Name            string
		InitialMessage  string
		MessagesAddress string
	}
	flag.StringVar(&opts.MessagesAddress, "me", "localhost:8081", "My listener address, defaults to localhost:8081")
	flag.StringVar(&opts.Name, "name", "gopher1", "Your gopher's name")
	flag.Parse()

	gopher.Name = opts.Name
	messageAddr, err := net.ResolveTCPAddr("tcp", opts.MessagesAddress)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Hi, my name is %s and I am listening on TCP %s\n", opts.Name, messageAddr)

	go listenForFriends()
	go register(opts.Name, messageAddr)
	go listenForMessages(messageAddr)

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter message to send: ")
		body, err := reader.ReadString('\n')
		if err != nil {
			err = errors.Wrap(err, "Unable to read message from user")
			log.Printf("%+v", err)
			continue
		}
		msg := buildMessage(body)
		fmt.Print(msg)
		msg.Send()
	}
}

func buildMessage(body string) Message {
	msg := Message{Body: body, From: gopher.Name}
	for _, friend := range friends {
		msg.Recipients = append(msg.Recipients, friend)
	}
	return msg
}

func listenForMessages(messagesAddress *net.TCPAddr) {
	// Start listening for messages
	log.Printf("Listening for messages on tcp %s\n", messagesAddress)
	listen, err := net.ListenTCP("tcp4", messagesAddress)
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

		msg, err := readMessage(conn)
		if err != nil {
			log.Printf("Unable to read incoming message: %+v", err)
			continue
		}

		log.Printf("Received:%s", msg)
		go gopher.HandleMessage(msg)
	}
}

func listenForFriends() {
	addr, _ := net.ResolveUDPAddr("udp", registrationAddress)

	// Start listening for broadcasts from friends
	log.Printf("Listening for friends on udp %s\n", registrationAddress)
	listen, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Unable to listen for registration broadcasts: %v", err)
	}
	defer listen.Close()

	// Handle incoming messages
	for {
		buf := make([]byte, 1024)
		n, _, err := listen.ReadFromUDP(buf)
		if err != nil {
			log.Fatalln(err)
			continue
		}
		msg := string(buf[:n])
		friend, err := parseFriend(msg)
		if err != nil {
			log.Printf("%+v", err)
			continue
		}
		if friend.Name == gopher.Name {
			// ignore myself
			continue
		}
		log.Printf("\nI found my friend %s at %s\n", friend.Name, friend.Address)
		friends[friend.Name] = *friend
	}
}

func register(name string, msgAddr *net.TCPAddr) {
	addr, _ := net.ResolveUDPAddr("udp", registrationAddress)

	// Register as a friendly gopher
	log.Printf("Registering my gopher on %s\n", registrationAddress)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Unable to connect to registration port: %v", err)
	}
	defer conn.Close()

	self := fmt.Sprintf("%s:%s", name, msgAddr)
	_, err = conn.Write([]byte(self))
	if err != nil {
		log.Fatalf("Unable to register: %v", err)
	}
}
