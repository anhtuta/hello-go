package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// This is your database handle
// Making db a global variable simplifies this example. In production, you’d avoid the global variable
var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// Query for multiple rows:
// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	const errorMessage = "albumsByArtist %q: %v"

	// An albums slice to hold data from returned rows.
	var albums []Album

	// By separating the SQL statement from parameter values (rather than concatenating them with, say,
	// fmt.Sprintf), you enable the database/sql package to send the values separate from the SQL text,
	// removing any SQL injection risk
	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf(errorMessage, name, err)
	}

	// Defer closing rows so that any resources it holds will be released when the function exits
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album

		// Scan takes a list of pointers to Go values, where the column values will be written. Here,
		// you pass pointers to fields in the alb variable, created using the & operator. Scan writes
		// through the pointers to update the struct fields
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf(errorMessage, name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf(errorMessage, name, err)
	}
	return albums, nil
}

// Query for a single row
// albumByID queries for the album with the specified ID.
func albumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	// QueryRow doesn’t return an error. Instead, it arranges to return any query error
	// (such as sql.ErrNoRows) from Rows.Scan later
	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// Add data
// addAlbum adds the specified album to the database,
// returning the album ID of the new entry
func addAlbum(alb Album) (int64, error) {
	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"), // need to set OS environment variables: export DB_USER=root
		Passwd: os.Getenv("DB_PASS"), // need to set OS environment variables: export DB_PASS=password
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "go_recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Test connection
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// Query multiple rows
	albums, err := albumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	// Query single row: Hard-code ID 2 here to test the query.
	alb, err := albumByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", alb)

	// Add new row
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
