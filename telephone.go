package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Telephone is the game event loop. It listens for incoming messages and has your gopher reply back.
type Telephone struct {
	gopher  Gopher
	address *net.TCPAddr
	friends map[string]Friend
}

func newTelephone(gopher Gopher, address *net.TCPAddr) *Telephone {
	return &Telephone{
		gopher:  gopher,
		address: address,
		friends: make(map[string]Friend),
	}
}

func (t *Telephone) Start() {
	log.Printf("Hi, my name is %s and I am listening on %s\n", t.gopher.Name(), t.address)

	go t.listenForFriends()
	go t.register()
	go t.listenForMessages()
}

func (t *Telephone) Send(value string) {
	msg := Message{
		From: t.gopher.Name(),
		Body: value,
	}

	for _, friend := range t.friends {
		msg.Recipients = append(msg.Recipients, friend)
	}

	msg.Send()
}

func (t *Telephone) listenForMessages() {
	// Start listening for messages
	log.Printf("Listening for messages on tcp %s\n", t.address)
	listen, err := net.ListenTCP("tcp4", t.address)
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
		go t.gopher.HandleMessage(msg)
	}
}

func (t *Telephone) listenForFriends() {
	addr, _ := net.ResolveUDPAddr("udp", registrationAddress)

	// Start listening for broadcasts from friends
	log.Printf("Listening for friends on udp %s\n", registrationAddress)
	listen, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Unable to listen for registration broadcasts: %v", err)
	}
	defer listen.Close()

	// Handle incoming friend registrations
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
		if friend.Name == t.gopher.Name() {
			// ignore myself
			continue
		}
		_, exists := t.friends[friend.Name]
		if !exists {
			log.Printf("\nI found a new friend %s at %s\n", friend.Name, friend.Address)
		}
		t.friends[friend.Name] = *friend
	}
}

func (t *Telephone) register() {
	addr, _ := net.ResolveUDPAddr("udp", registrationAddress)

	// Register as a friendly gopher
	log.Printf("Registering my gopher on %s\n", registrationAddress)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("Unable to connect to registration port: %v", err)
	}
	defer conn.Close()

	self := fmt.Sprintf("%s:%s", t.gopher.Name(), t.address)

	// Broadcast that where I can be reached at regular intervals
	for {
		_, err = conn.Write([]byte(self))
		if err != nil {
			log.Fatalf("Unable to register: %v", err)
		}
		time.Sleep(time.Second)
	}
}
