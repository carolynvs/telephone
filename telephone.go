package main

import (
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// Telephone is the game event loop. It listens for incoming messages and has your gopher reply back.
type Telephone struct {
	gopher  Gopher
	me      Friend
	friends map[string]Friend
}

// newTelephone initializes an instance of the telephone game
func NewTelephone(gopher Gopher, name string) *Telephone {
	t := &Telephone{
		gopher:  gopher,
		friends: make(map[string]Friend),
	}

	// Default the name to the host name, if unspecified
	if name == "" {
		name = t.getHostName()
	}
	t.me.Name = name

	return t
}

// Start listening for messages from your gopher friends
func (t *Telephone) Start() {
	log.Printf("Hi, my name is %s. Let's play telephone!\n\n", t.me.Name)

	listener := t.listenOnFreePort()

	go t.listenForFriends()
	go t.register()
	go t.listenForMessages(listener)
}

// Send a message to your gopher friends
func (t *Telephone) Send(value string) {
	if len(t.friends) == 0 {
		log.Println("There is no one online to send the message to. Tell your friends to pick up the phone!")
		return
	}

	// Send to all our friends in a semi-random ordering
	msg := Message{
		From: t.me,
		Body: value,
	}
	for _, friend := range t.friends {
		if msg.To.Name == "" {
			msg.To = friend
		} else {
			msg.CC = append(msg.CC, friend)
		}
	}

	msg.Send()
}

// getLocalLANIP returns the local LAN IP address
// For example 192.168.0.100
// We don't want the public IP as we are playing only with other gophers on the same network
// We don't want 127.0.0.1 or localhost, so that we can tell others how to reach us
func getLocalLANIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func (t *Telephone) listenOnFreePort() net.Listener {
	myIP := getLocalLANIP()

	// port = 0 means give me the next available port
	freePortRequest, _ := net.ResolveTCPAddr("tcp4", myIP+":0")

	// Start listening on a free/open port
	listener, err := net.ListenTCP("tcp", freePortRequest)
	if err != nil {
		log.Fatalf("Unable to listen for incoming message: %v", err)
	}

	// Remember which port was assigned
	addr := listener.Addr().(*net.TCPAddr)
	t.me.Number = PhoneNumber{addr}
	log.Printf("Listening for messages on %s\n", t.me.Number)

	return listener
}

func (t *Telephone) listenForMessages(listener net.Listener) {
	defer listener.Close()

	// Handle incoming messages
	for {
		conn, err := listener.Accept()
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
		go t.handleMessage(msg)
	}
}

func (t *Telephone) handleMessage(msg Message) {
	result := t.gopher.TransformMessage(msg)
	msg.Forward(result)
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
		f := string(buf[:n])
		friend, err := ParseFriend(f)
		if err != nil {
			log.Printf("%+v", err)
			continue
		}

		if friend.Equals(t.me) {
			// ignore my own messages
			continue
		}
		_, exists := t.friends[friend.Name]
		if !exists {
			log.Printf("I found a new friend %s at %s\n", friend.Name, friend.Number)
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

	// Broadcast that where I can be reached at regular intervals
	for {
		_, err = conn.Write([]byte(t.me.String()))
		if err != nil {
			log.Fatalf("Unable to register: %v", err)
		}
		time.Sleep(time.Second)
	}
}

func (t *Telephone) getHostName() string {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal("Unable to determine your hostname, so you must specify it explicityly with the -name flag")
	}

	// Return the first part of the hostname, ignoring the domain
	// e.g. mycomputer.local returns mycomputer
	return strings.SplitN(name, ".", 2)[0]
}
