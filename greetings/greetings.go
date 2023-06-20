package greetings

import (
	"errors"
	"fmt"
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