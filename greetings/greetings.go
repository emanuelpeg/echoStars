package greetings

import "fmt"

type Greeter interface {
	Hello(name string) (message string)
}

type GreeterImp struct {
}

// Hello returns a greeting for the named person.
func (g GreeterImp) Hello(name string) (message string) {
	// Return a greeting that embeds the name in a message.
	if name == "" {
		message = "Hi!"
	} else {
		message = fmt.Sprintf("Hi, %v. Welcome!", name)
	}
	return message
}
