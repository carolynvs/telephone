package main

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
)

type PhoneNumber struct {
	*net.TCPAddr
}

func ParsePhoneNumber(value string) (*PhoneNumber, error) {
	addr, err := net.ResolveTCPAddr("tcp", value)
	return &PhoneNumber{addr}, errors.Wrapf(err, "Unable to parse %s as a TCP IPv4 address")
}

func (n PhoneNumber) String() string {
	return fmt.Sprintf("%s:%d", n.IP, n.Port)
}
