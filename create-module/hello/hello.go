// Code executed as an application must go in a main package.
package main

// Import two packages: example.com/greetings and fmt. This gives your code
// access to functions in those packages. Importing example.com/greetings (the
// package contained in the module you created earlier) gives you access to the
// Hello function. You also import fmt, with functions for handling input and
// output text (such as printing text to the console).
import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    // Set properties of the predefined Logger, including the log entry prefix
    // and a flag to disable printing the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Request a greeting message.
    message, err := greetings.Hello("")
    // If an error was returned, print it to the console and exit the program.
    if err != nil {
        // Print the error and stop the program.
        log.Fatal(err)
    }

    // If no error was returned, print the returned message to the console.
    fmt.Println(message)
}
