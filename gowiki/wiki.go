package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
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

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// The html/template package is part of the Go standard library: we can use
	// html/template to keep the HTML in a separate file, allowing us to change
	// the layout of our edit page without modifying the underlying Go code.
	// The function ParseFiles reads the contents of edit.html and return a
	// *Template.
	t, _ := template.ParseFiles(tmpl + ".html")
	// The method Execute executes the template, writing the generated HTML to
	// the ResponseWriter
    t.Execute(w, p)
}

// The function viewHandler allow users to view a wiki page; it will handle URLs
// prefixed with "/view/".
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the page title from r.URL.Path, the path component of the request
	// URL. The Path is re-sliced with [len("/view/"):] to drop the leading
	// "/view/" component of the request path; this is because the path will
	// invariably begin with "/view/", which is not part of the page's title.
	title := r.URL.Path[len("/view/"):]
	// Load the page data.
	p, _ := loadPage(title)
	renderTemplate(w, "view", p)
}

// The function editHandler loads the page (or, if it doesn't exist, create an
// empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/edit/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
