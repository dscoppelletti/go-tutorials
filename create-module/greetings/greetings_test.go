// Ending a file's name with "_test.go" tells the "go test" command that this
// file contains test functions.

// Implement test functions in the same package as the code you're testing.
package greetings

import (
    "testing"
    "regexp"
)

// Test function names have the form "Test"<Name>, where <Name> is specific to
// the test. Also, test functions take a pointer to the testing package's
// "testing.T" as a parameter. You use this parameter's methods for reporting
// and logging from your test.

// TestHelloName calls greetings.Hello with a name, checking for a valid return
// value.
func TestHelloName(t *testing.T) {
    name := "Gladys"
    want := regexp.MustCompile(`\b`+name+`\b`)
    msg, err := Hello("Gladys")
    if !want.MatchString(msg) || err != nil {
		// Print a message to the console and end execution.
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err,
			want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string, checking for an
// error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

// The "go test" command executes test functions (whose names begin with Test)
// in test files (whose names end with _test.go). You can add the -v flag to get
// verbose output that lists all of the tests and their results.
