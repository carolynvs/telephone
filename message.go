package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/pkg/errors"
)

type Message struct {
	to         Friend
	From       string   `json:"from"`
	Recipients []Friend `json:"recipients"`
	Body       string   `json:"body"`
}

func (msg *Message) String() string {
	return fmt.Sprintf("\n\tFrom: %s\n\tRecipients: %s\n\tBody: %s\n\n", msg.From, msg.Recipients, msg.Body)
}

func (msg *Message) prepare() (shouldSend bool) {
	if msg.Body == "" {
		log.Println("Ignoring empty message")
		shouldSend = false
		return
	}

	if len(msg.Recipients) == 0 {
		shouldSend = false
		return
	}

	// Identify the next and remaining recipients
	msg.to = msg.Recipients[0]
	msg.Recipients = msg.Recipients[1:len(msg.Recipients)]
	shouldSend = true
	return
}

func (msg *Message) Send() {
	if !msg.prepare() {
		return
	}

	log.Print(msg)
	// Connect to my friend
	log.Printf("Connecting to my friend %s(%s):\n\t%s\n", msg.to.Name, msg.to.Address, msg.Body)
	conn, err := net.DialTCP("tcp", nil, msg.to.Address)
	if err != nil {
		err = errors.Wrapf(err, "Unable to connect to %s(%s)", msg.to.Name, msg.to.Address)
		log.Printf("%+v", err)
		return
	}
	defer conn.Close()

	// Marshal the message to bytes
	msgb, err := json.Marshal(msg)
	if err != nil {
		err = errors.Wrapf(err, "Unable to marshal the message %s", msg.to.Name, msg.to.Address)
		log.Printf("%+v", err)
		return
		log.Printf(": %v", err)
	}

	log.Println("Sending message")
	_, err = conn.Write(msgb)
	if err != nil {
		log.Fatalf("Unable to send the message: %v", err)
	}
}

func readMessage(reader io.ReadCloser) (Message, error) {
	defer reader.Close()

	log.Println("Parsing incoming message...")

	msg := Message{}
	err := json.NewDecoder(reader).Decode(&msg)
	return msg, errors.Wrap(err, "Unable to parse incomming message")
}
