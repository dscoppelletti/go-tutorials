package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// Declare a database handler
// Making db a global variable simplifies this example. In production, you’d
// avoid the global variable, such as by passing the variable to functions that
// need it or by wrapping it in a struct.
var db *sql.DB

type Album struct {
    ID     int64
    Title  string
    Artist string
    Price  float32
}

func main() {
    // Capture connection properties into a DNS.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }

    // Get a database handle.
    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
		// To simplify the code, you’re calling log.Fatal to end execution and
		// print the error to the console. In production code, you’ll want to
		// handle errors in a more graceful way.
        log.Fatal(err)
    }

	// Confirm that connecting to the database works. At run time, sql.Open
	// might not immediately connect, depending on the driver. You’re using Ping
	// here to confirm that the database/sql package can connect when it needs
	// to.
    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")

	// Query for multiple rows
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Printf("Albums found: %v\n", albums)

	// Query for a single row.
	// Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
    	log.Fatal(err)
	}

	fmt.Printf("Album found: %v\n", alb)

	// Add data
	albID, err := addAlbum(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ID of added album: %v\n", albID)
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
    // An albums slice to hold data from returned rows.
    var albums []Album

	// Executes a SELECT statement.
    rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
    if err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

	// Defer closing rows so that any resources it holds will be released when
	// the function exits.
    defer rows.Close()

    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var alb Album
        if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price);
			err != nil {
            return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
        }
        albums = append(albums, alb)
    }

	// Checking for an error here is the only way to find out that the results
	// are incomplete.
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
    }

    return albums, nil
}

// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
    // An album to hold data from the returned row.
    var alb Album

	// Execute a SELECT statement.
	// To simplify the calling code, QueryRow doesn’t return an error. Instead,
	// it arranges to return any query error (such as sql.ErrNoRows) from
	// Rows.Scan later.
    row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
    if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price);
		err != nil {
        if err == sql.ErrNoRows {
			// Typically this error is worth replacing with more specific text,
			// such as “no such album” here.
            return alb, fmt.Errorf("albumsById %d: no such album", id)
        }

        return alb, fmt.Errorf("albumsById %d: %v", id, err)
    }

    return alb, nil
}

// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {

	// Execute an INSERT statement.
    result, err := db.Exec(
		"INSERT INTO album (title, artist, price) VALUES (?, ?, ?)",
		alb.Title, alb.Artist, alb.Price)
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }

	// Retrieve the ID of the inserted database row.
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addAlbum: %v", err)
    }

    return id, nil
}
