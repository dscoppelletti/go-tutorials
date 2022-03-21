package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
//
// Struct tags such as json:"artist" specify what a field’s name should be when
// the struct’s contents are serialized into JSON. Without them, the JSON would
// use the struct’s capitalized field names – a style not as common in JSON.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
// To keep things simple for the tutorial, you’ll store data in memory. A more
// typical API would interact with a database.
//
// Note that storing data in memory means that the set of albums will be lost
// each time you stop the server, then recreated when you start it.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
//
// Note that you could have given this function any name – neither Gin nor Go
// require a particular function name format.
//
// gin.Context is the most important part of Gin. It carries request details,
// validates and serializes JSON, and more.
//
// Context.IndentedJSON serializes the struct into JSON and add it to the
// response.
// The function’s first argument is the HTTP status code you want to send to the
// client.
//
// Note that you can replace Context.IndentedJSON with a call to Context.JSON to
// send more compact JSON. In practice, the indented form is much easier to work
// with when debugging and the size difference is usually small.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)

    // Add a 201 status code to the response, along with JSON representing the
    // album you added
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    // Use Context.Param to retrieve the id path parameter from the URL. When
    // you map this handler to a path, you’ll include a placeholder for the
    // parameter in the path.
    id := c.Param("id")

    // Loop over the album structs in the slice, looking for one whose ID field
    // value matches the id parameter value. If it’s found, you serialize that
    // album struct to JSON and return it as a response with a 200 OK HTTP code.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }

    // Return an HTTP 404 error with http.StatusNotFound if the album isn’t
    // found.
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	// Initialize a Gin router.
    router := gin.Default()

	// Associate the GET HTTP method and /albums path with a getAlbums function.
    router.GET("/albums", getAlbums)

    // Associate the POST method at the /albums path with the postAlbums
    // function.
    // With Gin, you can associate a handler with an HTTP method-and-path
    // combination. In this way, you can separately route requests sent to a
    // single path based on the method the client is using.
    router.POST("/albums", postAlbums)

    // Associate the /albums/:id path with the getAlbumByID function. In Gin,
    // the colon preceding an item in the path signifies that the item is a path
    // parameter.
    router.GET("/albums/:id", getAlbumByID)

	// Attach the router to an http.Server and start the server.
    router.Run("localhost:8080")
}
