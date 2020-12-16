// Declare a greetings package to collect related functions.
package greetings

import (
    "errors"
    "fmt"
)

// Hello returns a greeting for the named person.
// The name of this function starts with a capital letter because can be called
// by a function not in the same package. This is known in Go as an exported
// name.
// The function returns two values: a string and an error. Your caller will
// check the second value to see if an error occurred.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		// The errors.New function returns an error with your message inside.
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	// In Go, the := operator is a shortcut for declaring and initializing a
	// variable in one line (Go uses the value on the right to determine the
	// variable's type).
	// Sprintf substitutes the name parameter's value for the %v format verb.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// The second returned value "nil" means no error.
    return message, nil
}
