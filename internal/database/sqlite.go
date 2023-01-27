package database

import (
	"database/sql"
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
		INSERT INTO image (id, filename, path)
		VALUES (?, ?, ?)
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
	row, err := db.Query("SELECT * FROM image ORDER BY filename")
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	var (
		id       string
		filename string
		path     string
	)

	for row.Next() {
		row.Scan(&id, &filename, &path)
		fmt.Printf("[%s] filename: %s\tpath: %s\n", id, filename, path)
	}
}
