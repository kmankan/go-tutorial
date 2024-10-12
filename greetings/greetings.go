package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person. the input parameter is type string and the return value is type string
func Hello(name string) (string, error) {
	// If no name is given, return an error with a message
	if name == "" {
		return "", errors.New("empty name")
	}

	// If name is given, return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
