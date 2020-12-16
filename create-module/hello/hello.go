// Code executed as an application must go in a main package.
package main

// Import two packages: example.com/greetings and fmt. This gives your code
// access to functions in those packages. Importing example.com/greetings (the
// package contained in the module you created earlier) gives you access to the
// Hello function. You also import fmt, with functions for handling input and
// output text (such as printing text to the console).
import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
