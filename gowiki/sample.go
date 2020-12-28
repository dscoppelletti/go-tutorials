// +build ignore

package main

import (
    "fmt"
    "log"
    "net/http"
)

// The function handler is of the type http.HandlerFunc: it takes a
// ResponseWriter and a Request as its arguments.
// A ResponseWriter value assembles the HTTP server's response; by writing to
// it, we send data to the HTTP client.
// A Request is a data structure that represents the client HTTP request;
// r.URL.Path is the path component of the request URL. The trailing [1:] means
// "create a sub-slice of Path from the 1st character to the end.", so drops the
// leading "/" from the path name.
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// HandleFunc tells the http package to handle all requests to the web root
	// ("/") with handler.
	http.HandleFunc("/", handler)
	// ListenAndServe specifies that it should listen on port 8080 on any
	// interface (":8080"); this function will block until the program is
	// terminated.
	// ListenAndServe always returns an error, since it only returns when an
	// unexpected error occurs. In order to log that error we wrap the function
	// call with log.Fatal.
    log.Fatal(http.ListenAndServe(":8080", nil))
}
