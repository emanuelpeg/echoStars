package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) (message string) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		message = "Hi!"
	} else {
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	}
	return message
}
