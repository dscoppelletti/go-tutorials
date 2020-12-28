package main

import (
	"fmt"
	"io/ioutil"
)

//  A wiki consists of a series of interconnected pages, each of which has a
// title and a body (the page content). Here, we define Page as a struct with
// two fields representing the title and body.
type Page struct {
	Title string
	Body []byte // This is a slice rather than string because that is the type
				// expected by the io libraries we will use.
}

// This method will save the Page's Body to a text file. For simplicity, we will
// use the Title as the file name.
// The save method returns the error value, to let the application handle it
// should anything go wrong while writing the file. If all goes well,
// Page.save() will return nil (the zero-value for pointers, interfaces, and
// some other types).
func (p *Page) save() error {
	filename := p.Title + ".txt"
	// The save method returns an error value because that is the return type of
	// WriteFile (a standard library function that writes a byte slice to a
	// file).
	// The octal integer literal 0600, passed as the third parameter to
	// WriteFile, indicates that the file should be created with read-write
	// permissions for the current user only.
    return ioutil.WriteFile(filename, p.Body, 0600)
}

// The function loadPage constructs the file name from the title parameter,
// reads the file's contents into a new variable body, and returns a pointer to
// a Page literal constructed with the proper title and body values.
// Functions can return multiple values. Callers of this function can check the
// error returned; if it is nil then it has successfully loaded a Page; if not,
// it will be an error that can be handled by the caller.
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	// The standard library function io.ReadFile returns []byte and error.
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
    return &Page{Title: title, Body: body}, nil
}

func main() {
    p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}
