package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// return an error if there is no name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	// if a name was recieved, return a value that embeds the name
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// return one of a set of greeting messages
func randomFormat() string {
	// a slice of formats
	formats := []string {
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v, Well met!",
	}
	// return a randomly selected message format
	return formats[rand.Intn(len(formats))]
}