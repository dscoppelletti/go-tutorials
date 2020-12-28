package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

// The function Must is a convenience wrapper that panics when passed a non-nil
// error value, and otherwise returns the *Template unaltered.
// The ParseFiles function takes any number of string arguments that identify
// our template files, and parses those files into templates that are named
// after the base file name.
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	// The method Execute executes the template, writing the generated HTML to
	// the ResponseWriter
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Validation expression for the title.
// The function MustCompile parses and compile the regular expression, and
// return a Regexp. MustCompile is distinct from Compile in that it will panic
// if the expression compilation fails, while Compile returns an error as a
// second parameter.
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// The function getTitle uses the validPath expression to validate path and
// extract the page title.
// If the title is valid, it will be returned along with a nil error value. If
// the title is invalid, the function will write a "404 Not Found" error to the
// HTTP connection, and return an error to the handler.
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
    m := validPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return "", errors.New("invalid Page Title")
    }
    return m[2], nil // The title is the second subexpression.
}

// The function viewHandler allow users to view a wiki page; it will handle URLs
// prefixed with "/view/".
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	// Load the page data.
	p, err := loadPage(title)
	if err != nil {
		// If the requested Page doesn't exist, it redirects the client to the
		// edit Page so the content may be created.
		// The Redirect function adds an HTTP status code 302 and a Location
		// header to the HTTP response.
		http.Redirect(w, r, "/edit/" + title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

// The function editHandler loads the page (or, if it doesn't exist, create an
// empty Page struct), and displays an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// The function saveHandler handles the submission of forms located on the edit
// pages.
func saveHandler(w http.ResponseWriter, r *http.Request) {
	// The page title (provided in the URL) and the form's only field, Body, are
	// stored in a new Page.
	title, err := getTitle(w, r)
    if err != nil {
        return
    }
	body := r.FormValue("body")
	// The value returned by FormValue is of type string, so we must convert
	// that value to []byte before it will fit into the Page struct.
	p := &Page{Title: title, Body: []byte(body)}
	// The save() method writes the data to a file
	err = p.save()
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	// The client is redirected to the /view/ page.
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
