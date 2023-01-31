package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DATABASE_PATH string = "./internal/database/sqlite_database.db"
)

var db *sql.DB

type Image struct {
	id       string
	filename string
	path     string
}

func NewSQLiteConnection() (err error) {

	db, err = sql.Open("sqlite3", DATABASE_PATH)
	if err != nil {
		return
	}

	err = db.Ping()

	return
}

func CreateTable() {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS image (
			"id" TEXT NOT NULL PRIMARY KEY,
			"name" TEXT NOT NULL,
			"filename" TEXT NOT NULL,
			"path" TEXT NOT NULL
		)
	`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Printf("error: %v\n", err)
	}

	statement.Exec()
	log.Println("Image table was created!")
}

func InsertImage(name, filename, path string) {
	insertImageSQL := `
		INSERT INTO image (id, name, filename, path)
		VALUES (?, ?, ?, ?)
	`

	statement, err := db.Prepare(insertImageSQL)
	if err != nil {
		log.Fatalln(err)
	}

	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)

	_, err = statement.Exec(id, name, filename, path)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Image has inserted successfully!")
}

func InfoImages() {
	row, err := db.Query("SELECT * FROM image ORDER BY name")
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	var (
		id       string
        name     string
		filename string
		path     string
	)

	for row.Next() {
		row.Scan(&id, &name,  &filename, &path)
        fmt.Printf("[%s] name: %s\tfilename: %s\tpath: %s\n", id, name, filename, path)
	}
}

func DelImage(id, name, path string) error {
    deleleImageSQL := `
        DELETE FROM image
        WHERE  image.id = ? 
            OR image.name = ?
            OR image.path = ?

    `
    statement, err := db.Prepare(deleleImageSQL)
    if err != nil {
        return err
    }

    _, err = statement.Exec(id, name, path)
    if err != nil {
        return err
    }

    log.Println("Image has deleted successfully!")
    return nil
}

func SearchPath(id, name string) (string, error) {

    pathSearchQuerySQL := `
        SELECT path FROM image
        WHERE image.id == ? 
        OR image.name == ?
    `

    row, err := db.Query(pathSearchQuerySQL, id, name) 
    if err != nil {
        return "", err
    }

    defer row.Close() 

    var path string

    for row.Next() {
        row.Scan(&path)
    }
    
    if path == "" {
        return "", errors.New("error in scan")
    }

    return path, nil
}
