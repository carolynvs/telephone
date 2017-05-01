package main

import (
	"log"
	"net"
	"regexp"
)

type Friend struct {
	Name    string       `json:"name"`
	Address *net.TCPAddr `json:"address"`
}

func parseFriend(msg string) (*Friend, error) {
	r := regexp.MustCompile(`(.+):(.+:\d+)`)
	matches := r.FindStringSubmatch(msg)
	if len(matches) != 3 {
		log.Fatalf("Invalid friend message '%s'", msg)
	}

	addr, err := net.ResolveTCPAddr("tcp", matches[2])
	if err != nil {
		log.Fatalf("Invalid friend address '%s': %v", matches[2], err)
	}

	f := &Friend{
		Name:    matches[1],
		Address: addr,
	}
	return f, nil
}
