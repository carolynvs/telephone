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

func main() {
	var opts struct {
		Name            string
		Gopher          string
		MessagesAddress string
	}
	flag.StringVar(&opts.MessagesAddress, "me", "localhost:8081", "My listener address, defaults to localhost:8081")
	flag.StringVar(&opts.Gopher, "gopher", "", "The type of gopher to use")
	flag.StringVar(&opts.Name, "name", "gopher1", "Your gopher's name")
	flag.Parse()

	// TODO: pick a random unused port
	addr, err := net.ResolveTCPAddr("tcp", opts.MessagesAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Pick a gopher to play with defaulting to a boring gopher who gets the job done
	var gopher Gopher
	switch opts.Gopher {
	default:
		gopher = DefaultGopher{name: opts.Name}
	}

	phone := newTelephone(gopher, addr)
	phone.Start()

	fmt.Println("### Type a message to send at any time, and press ENTER to send ###")
	reader := bufio.NewReader(os.Stdin)
	for {
		value, err := reader.ReadString('\n')
		if err != nil {
			err = errors.Wrap(err, "Unable to read message from user")
			log.Printf("%+v", err)
			continue
		}

		phone.Send(value)
	}
}
