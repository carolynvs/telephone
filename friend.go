package main

import (
	"fmt"
	"log"
	"regexp"

	"github.com/pkg/errors"
)

type Friend struct {
	Name   string      `json:"name"`
	Number PhoneNumber `json:"number"`
}

func ParseFriend(msg string) (*Friend, error) {
	r := regexp.MustCompile(`(.+)\((.+:\d+)\)`)
	matches := r.FindStringSubmatch(msg)
	if len(matches) != 3 {
		log.Fatalf("%+v", errors.Errorf("Invalid friend message '%s'", msg))
	}

	number, err := ParsePhoneNumber(matches[2])
	if err != nil {
		log.Fatalf("%+v", errors.Wrapf(err, "Invalid friend phone number '%s'", matches[2]))
	}

	f := &Friend{
		Name:   matches[1],
		Number: *number,
	}
	return f, nil
}

func (f Friend) String() string {
	return fmt.Sprintf("%s(%s)", f.Name, f.Number)
}

func (f Friend) Equals(o Friend) bool {
	return f.Name == o.Name &&
		f.Number.IP.String() == o.Number.IP.String() &&
		f.Number.Port == o.Number.Port
}
