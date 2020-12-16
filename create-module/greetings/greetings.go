// Declare a greetings package to collect related functions.
package greetings

import (
    "errors"
	"fmt"
	"math/rand"
    "time"
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
	// Create a message using a random format.
	// Sprintf substitutes the name parameter's value for the %v format verb.
    message := fmt.Sprintf(randomFormat(), name)
	// The second returned value "nil" means no error.
    return message, nil
}

// Changing the Hello function's parameter from a single name to a set of names
// would change the function signature. If you had already published the
// greetings module and users had already written code calling Hello, that
// change would break their programs. In this situation, a better choice is to
// give new functionality a new name.

// Hellos returns a map that associates each of the named people with a greeting
// message.
// The parameter is a slice of names rather than a single name. Also, the type
// of its return values is a map so you can return names mapped to greeting
// messages.
//
func Hellos(names []string) (map[string]string, error) {
    // Initiliaze map to associate names with messages.
    messages := make(map[string]string)
	// Loop through the received slice of names, calling the Hello function to
	// get a message for each name.
	// In the "for" loop, "range" returns two values: the index of the current
	// item in the loop and a copy of the item's value. You don't need the
	// index, so you use the Go blank identifier (an underscore) to ignore it.
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        // In the map, associate the retrieved message with the name.
        messages[name] = message
    }
    return messages, nil
}

// init sets initial values for variables used in the function.
// Go executes init functions automatically at program startup, after global
// variables have been initialized.
func init() {
	// Seed the rand package with the current time.
    rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned message
// is selected at random.
// The name of this function starts with a lowercase letter, making it
// accessible only to code in its own package.
func randomFormat() string {
	// A slice of message formats.
	// A slice is like an array, except that it's dynamically sized as you add
	// and remove items.
	// When declaring a slice, you omit its size in the brackets, so that the
	// array underlying a slice can be dynamically sized.
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }

	// Return a randomly selected message format by specifying a random index
	// for the slice of formats.
	// the math/rand package to generates a random number
    return formats[rand.Intn(len(formats))]
}
