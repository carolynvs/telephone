package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	var opts struct {
		Name            string
		Gopher          string
		MessagesAddress string
	}
	flag.StringVar(&opts.Gopher, "gopher", "", "The type of gopher to use")
	flag.StringVar(&opts.Name, "name", "", "Your gopher's name")
	flag.Parse()

	// Pick a gopher to play with defaulting to a boring gopher who gets the job done
	var gopher Gopher
	switch opts.Gopher {
	default:
		gopher = DefaultGopher{}
	}

	phone := NewTelephone(gopher, opts.Name)
	phone.Start()

	fmt.Println()
	fmt.Println("#############################################################")
	fmt.Println("     Type a message at any time, and press ENTER to send     ")
	fmt.Println("#############################################################")
	fmt.Println()
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
