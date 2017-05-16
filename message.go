package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
)

type Message struct {
	Id   string   `json:"id"`
	To   Friend   `json:"to"`
	From Friend   `json:"from"`
	CC   []Friend `json:"cc"`
	Body string   `json:"body"`
}

func (msg Message) String() string {
	return fmt.Sprintf("\n\tFrom: %s\n\tTo: %s\n\tCC: %s\n\tBody: %s\n\n", msg.From, msg.To, msg.CC, msg.Body)
}

func (msg *Message) generateId() {
	if msg.Id != "" {
		return
	}

	b := make([]byte, 16)
	rand.Read(b)
	msg.Id = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (msg *Message) Send() {
	if msg.Body == "" {
		log.Println("Ignoring empty message")
		return
	}

	// Connect to my friend
	log.Printf("Connecting to my friend %s(%s)\n", msg.To.Name, msg.To.Number)
	conn, err := net.DialTCP("tcp", nil, msg.To.Number.TCPAddr)
	if err != nil {
		err = errors.Wrapf(err, "Unable to connect to %s(%s)", msg.To.Name, msg.To.Number)
		log.Printf("%+v", err)
		return
	}

	msg.transmit(conn)
}

func (msg *Message) Broadcast() {
	addr, err := net.ResolveUDPAddr("udp", resultsAddress)
	if err != nil {
		log.Fatalf("%+v", errors.Wrapf(err, "Unable to resolve upd %s", resultsAddress))
	}

	// Connect to everyone playing
	log.Printf("Connecting to my friends on %s\n", addr)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "Unable to connect to results port"))
	}

	// Broadcast the results of a telephone chain
	msg.transmit(conn)
}

func (msg *Message) transmit(conn net.Conn) {
	defer conn.Close()

	msg.generateId()

	// Marshal the message to bytes
	msgb, err := json.Marshal(msg)
	if err != nil {
		err = errors.Wrapf(err, "Unable to marshal the message %s", msg.To.Name, msg.To.Number)
		log.Printf("%+v", err)
		return
	}

	log.Printf("Sending message:\n%s\n", msg)
	_, err = conn.Write(msgb)
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "Unable to send the message"))
	}
}

func (msg *Message) Forward(body string) {
	if len(msg.CC) == 0 {
		// Nothing to do
		return
	}

	reply := Message{
		Id:   msg.Id,
		From: msg.To,
		To:   msg.CC[0],
		CC:   msg.CC[1:len(msg.CC)],
		Body: body,
	}

	reply.Send()
}

func readMessage(reader io.Reader) (Message, error) {
	log.Println("Parsing incoming message...")

	msg := Message{}
	err := json.NewDecoder(reader).Decode(&msg)
	return msg, errors.Wrap(err, "Unable to parse incomming message")
}
